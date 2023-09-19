package acceptance

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/sirupsen/logrus"
	"io"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
	"strings"
	"testing"
	"time"
)

var kubernetesClient *kubernetes.Clientset
var impersonateAccountClient *kubernetes.Clientset
var namespaces map[string]*corev1.Namespace
var podPerNamespace map[string]*corev1.PodList
var propagatedRoleBindingPerTenant map[string]rbacv1.RoleBinding
var serviceAccountToImpersonate *corev1.ServiceAccount
var tenantParentNamespace string
var destinationPod *corev1.Pod
var sourcePod *corev1.Pod
var destinationNamespaceName string
var sourceNamespaceName string
var destinationService *corev1.Service

func theKubernetesClientIsSetUp() error {
	if kubernetesClient == nil {
		logrus.Info("Attempting to fetch in-cluster config..")
		var config *rest.Config
		var err error
		if os.Getenv("RUN_OUTSIDE_CLUSTER") == "true" {
			config, err = clientcmd.BuildConfigFromFlags(
				"",
				filepath.Join(homedir.HomeDir(), ".kube", "config"),
			)
		} else {
			config, err = rest.InClusterConfig()
		}

		if err != nil {
			return fmt.Errorf("failed fetching in cluster config, err: %v", err)
		}
		logrus.Info("Attempting to create a kubernetes client..")
		kubernetesClient, err = kubernetes.NewForConfig(config)
		if err != nil {
			return fmt.Errorf("failed creating a kubernetes client, err: %v", err)
		}
		logrus.Info("Proceeding with the acceptance criteria test..")
		return nil
	}
	return nil
}

func iGetTheNamespace(namespaceName string) error {
	namespace, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Not found -- handled in the subsequent step
			return nil
		}
		return fmt.Errorf("error occurred while fetching the %s namespace", namespaceName)
	}
	if namespaces == nil {
		namespaces = make(map[string]*corev1.Namespace)
	}
	namespaces[namespaceName] = namespace
	return nil
}

func theNamespaceIsReturned(namespaceName string) error {
	if namespaces[namespaceName] == nil {
		return fmt.Errorf("%s namespace not yet set up", namespaceName)
	}
	return nil
}

func iGetThePodsInTheNamespace(namespace string) error {
	pods, err := kubernetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error while listing pods in the %s namespace", namespace)
	}
	if podPerNamespace == nil {
		podPerNamespace = make(map[string]*corev1.PodList)
	}
	podPerNamespace[namespace] = pods
	return nil
}

func thereIsARunningPod(podName string, namespace string) error {
	if len(podPerNamespace[namespace].Items) == 0 {
		return fmt.Errorf("no pods are running in the %s namespace", namespace)
	}

	for _, pod := range podPerNamespace[namespace].Items {
		if strings.Contains(pod.Name, podName) {
			return nil
		}
	}
	return fmt.Errorf("the pod %s is not running in the namespace %s", podName, namespace)
}

func theNamespaceHasTheSubnamespaces(namespaceName string, subnamespacesNames string) error {
	// at this point we know that the namespace already exists
	for _, nsName := range strings.Split(subnamespacesNames, ",") {
		childNamespace, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), nsName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("namespae %s is not found", nsName)
			}
			return fmt.Errorf("error occurred while fetching the %s namespace", nsName)
		}
		subnamespaceAnnotation := childNamespace.Annotations[v1alpha2.SubnamespaceOf]
		if subnamespaceAnnotation != namespaceName {
			return fmt.Errorf("the namespace %s is not setup correctly", nsName)
		}
	}
	return nil

}

func aRoleBindingExistsInAllMyNamespaces(namespaceNames string) error {
	if propagatedRoleBindingPerTenant == nil {
		propagatedRoleBindingPerTenant = make(map[string]rbacv1.RoleBinding)
	}
	for _, nsName := range strings.Split(namespaceNames, ",") {
		_, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), nsName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("namespace %s is not found", nsName)
			}
			return fmt.Errorf("error occurred while fetching the %s namespace", nsName)
		}
		roleBindings, err := kubernetesClient.RbacV1().RoleBindings(nsName).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error while fetching the role bindings in the namespace %s", nsName)
		}
		if propagatedRoleBindingPerTenant[tenantParentNamespace].Name == "" {
			// pick the first role binding in the namespace - in the future
			propagatedRoleBindingPerTenant[tenantParentNamespace] = roleBindings.Items[0]
		} else {
			// if the role binding was already found, check that the rest of the namespaces have it as well
			// in this scenario all the namespaces that belong to a tenant should have the role binding
			_, err := kubernetesClient.RbacV1().RoleBindings(nsName).
				Get(context.TODO(), propagatedRoleBindingPerTenant[tenantParentNamespace].Name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("the role bindings are not defined in the namespace %s", nsName)
			}
		}
	}
	return nil
}

func iAmATenantCalled(tenantName string) error {
	// check if a namespace with this tenant's name exists
	_, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), tenantName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("the tenant %s has no namespace associated with it", tenantName)
		}
		return fmt.Errorf("error occurred while fetching the %s namespace", tenantName)
	}
	tenantParentNamespace = tenantName
	return nil
}

func theRoleBindingIsAssociatedWithAServiceAccount() error {
	serviceAccount, err := kubernetesClient.CoreV1().
		ServiceAccounts(tenantParentNamespace).
		Get(context.TODO(), propagatedRoleBindingPerTenant[tenantParentNamespace].Subjects[0].Name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error trying to fetch a service account for the tenant %s", tenantParentNamespace)
	}
	serviceAccountToImpersonate = serviceAccount
	return nil
}

func iImpersonateTheServiceAccount() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("failed fetching in cluster config, err: %v", err)
	}
	account := fmt.Sprintf("system:serviceaccount:%s:%s", tenantParentNamespace, serviceAccountToImpersonate.Name)
	config.Impersonate = rest.ImpersonationConfig{UserName: account}
	logrus.Info("Attempting to create a kubernetes client after impersonation of account ..")
	impersonateAccountClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error occured while impersonating service account")
	}
	logrus.Info("Impersonation was successful..")
	return nil
}

func iCanAccessAllMyNamespaces(namespaceNames string) error {
	for _, ns := range strings.Split(namespaceNames, ",") {
		_, err := impersonateAccountClient.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("cannot access the pods in the namespace %s", ns)
		}
	}
	return nil
}

func iCannotAccessOtherTenantsNamespaces(namespaceNames string) error {
	for _, ns := range strings.Split(namespaceNames, ",") {
		_, err := impersonateAccountClient.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
		if err == nil {
			return fmt.Errorf("I can access another tenants's pods in the namespace %s", ns)
		}
	}
	return nil
}

func iHaveADestinationPodInTheNamespace(destinationPodName, nsName string) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   destinationPodName,
			Labels: map[string]string{"app": destinationPodName},
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:    "main",
					Image:   "busybox",
					Command: []string{"sh"},
					Args:    []string{"-c", "echo \"hello world\" > index.html && /bin/httpd -p 9000 -f"},
					Ports: []corev1.ContainerPort{
						{
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 9000,
						},
					},
				},
			},
		},
	}
	pod, err := kubernetesClient.CoreV1().Pods(nsName).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("error creating destination pod %s in the namespace %s", destinationPodName, nsName)
	}
	destinationPod = pod
	destinationNamespaceName = nsName
	return nil
}

func iHaveASourcePodInTheNamespace(sourcePodName, nsName string) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   sourcePodName,
			Labels: map[string]string{"app": sourcePodName},
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Name:    "main",
					Image:   "alpine",
					Command: []string{"sh"},
					Args: []string{"-c", fmt.Sprintf(
						"timeout 3s nc -vz %s.%s 80", destinationService.Name, destinationNamespaceName)},
					Ports: []corev1.ContainerPort{
						{
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 9000,
						},
					},
				},
			},
		},
	}
	pod, err := kubernetesClient.CoreV1().Pods(nsName).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("error creating source pod %s in the namespace %s", sourcePodName, nsName)
	}
	sourcePod = pod
	sourceNamespaceName = nsName
	return nil
}

func iTryToConnectFromTo(sourcePodName, destinationPodName string) error {
	err := wait.PollUntilContextTimeout(context.TODO(), 3*time.Second, 2*time.Minute, true, func(ctx context.Context) (bool, error) {
		p, err := kubernetesClient.CoreV1().Pods(sourceNamespaceName).Get(context.Background(), sourcePod.Name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if len(p.Status.ContainerStatuses) == 0 {
			return false, nil
		}
		state := p.Status.ContainerStatuses[0].State
		if state.Terminated != nil {
			_ = state.Terminated.ExitCode
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return fmt.Errorf("error occured while connecting from %s to %s", sourcePodName, destinationPodName)
	}
	return nil
}

func theAccessIsDenied() error {
	req := kubernetesClient.CoreV1().Pods(sourceNamespaceName).GetLogs(sourcePod.Name, &corev1.PodLogOptions{})
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return fmt.Errorf("error fetching logs from the source pod")
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return fmt.Errorf("error copying logs from the source pod")
	}
	logs := buf.String()

	if strings.Contains(logs, "open") {
		return fmt.Errorf("the connection from namespace %s to namespace %s is open", sourceNamespaceName, destinationNamespaceName)
	}
	return nil
}

func theAccessIsAllowed() error {
	req := kubernetesClient.CoreV1().Pods(sourceNamespaceName).GetLogs(sourcePod.Name, &corev1.PodLogOptions{})
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return fmt.Errorf("error fetching logs from the source pod")
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return fmt.Errorf("error copying logs from the source pod")
	}
	logs := buf.String()

	logrus.Info(fmt.Sprintf("******: logs - %s", logs))

	if !strings.Contains(logs, "open") {
		return fmt.Errorf("the connection between applications in the same namespace is denied")
	}
	return nil
}

func thePodHasAServiceCalledInTheNamespace(podName, serviceName, nsName string) error {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: serviceName},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:       80,
					TargetPort: intstr.FromInt32(9000),
				},
			},
			Selector: map[string]string{"app": podName},
		},
	}
	service, err := kubernetesClient.CoreV1().Services(nsName).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("error creating service %s in the namespace %s", serviceName, nsName)
	}
	destinationService = service
	return nil
}

func cleanupPods(scenario *godog.Scenario) {
	if sourcePod == nil {
		return
	}
	err := kubernetesClient.CoreV1().Pods(sourceNamespaceName).Delete(context.TODO(), sourcePod.Name, metav1.DeleteOptions{})
	if err != nil {
		logrus.Info(fmt.Sprintf("no source pod to delete in namespace %s", sourceNamespaceName))
	} else {
		waitUntilPodTerminated(sourcePod.Name, sourceNamespaceName)
	}
	err = kubernetesClient.CoreV1().Pods(destinationNamespaceName).Delete(context.TODO(), destinationPod.Name, metav1.DeleteOptions{})
	if err != nil {
		logrus.Info(fmt.Sprintf("no destination pod to delete in namespace %s", destinationNamespaceName))
	} else {
		waitUntilPodTerminated(destinationPod.Name, destinationNamespaceName)
	}
	err = kubernetesClient.CoreV1().Services(destinationNamespaceName).Delete(context.TODO(), destinationService.Name, metav1.DeleteOptions{})
	if err != nil {
		logrus.Info(fmt.Sprintf("no destination service to delete in namespace %s", destinationNamespaceName))
	}

}

func waitUntilPodTerminated(podName string, nsName string) error {
	err := wait.PollUntilContextTimeout(context.TODO(), 3*time.Second, 2*time.Minute, true, func(ctx context.Context) (bool, error) {
		p, err := kubernetesClient.CoreV1().Pods(nsName).Get(context.Background(), podName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		if len(p.Status.ContainerStatuses) == 0 {
			return false, nil
		}
		state := p.Status.ContainerStatuses[0].State
		if state.Terminated != nil {
			_ = state.Terminated.ExitCode
			return true, nil
		}
		return false, nil
	})
	return err

}
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		cleanupPods(sc)
		return ctx, nil
	})

	ctx.Step(`^I get the "([^"]*)" namespace$`, iGetTheNamespace)
	ctx.Step(`^the kubernetes client is setup$`, theKubernetesClientIsSetUp)
	ctx.Step(`^the namespace "([^"]*)" is returned$`, theNamespaceIsReturned)
	ctx.Step(`^I get the pods in the "([^"]*)" namespace$`, iGetThePodsInTheNamespace)
	ctx.Step(`^there is a running "([^"]*)" pod in the namespace "([^"]*)"$`, thereIsARunningPod)
	ctx.Step(`^the "([^"]*)" namespace has the subnamespaces "([^"]*)"$`, theNamespaceHasTheSubnamespaces)
	ctx.Step(`^a roleBinding exists in all my namespaces: "([^"]*)"$`, aRoleBindingExistsInAllMyNamespaces)
	ctx.Step(`^I am a tenant called "([^"]*)"$`, iAmATenantCalled)
	ctx.Step(`^I can access all my namespaces: "([^"]*)"$`, iCanAccessAllMyNamespaces)
	ctx.Step(`^I cannot access other tenant\'s namespaces: "([^"]*)"$`, iCannotAccessOtherTenantsNamespaces)
	ctx.Step(`^the roleBinding is associated with a serviceAccount$`, theRoleBindingIsAssociatedWithAServiceAccount)
	ctx.Step(`^I impersonate the service account$`, iImpersonateTheServiceAccount)
	ctx.Step(`^I can access all my namespaces: "([^"]*)"$`, iCanAccessAllMyNamespaces)
	ctx.Step(`^I cannot access other tenant's namespaces: "([^"]*)"$`, iCannotAccessOtherTenantsNamespaces)
	ctx.Step(`^I have a destination pod "([^"]*)" in the namespace "([^"]*)"$`, iHaveADestinationPodInTheNamespace)
	ctx.Step(`^I have a source pod "([^"]*)" in the namespace "([^"]*)"$`, iHaveASourcePodInTheNamespace)
	ctx.Step(`^I try to connect from "([^"]*)" to "([^"]*)"$`, iTryToConnectFromTo)
	ctx.Step(`^the access is denied$`, theAccessIsDenied)
	ctx.Step(`^the access is allowed$`, theAccessIsAllowed)
	ctx.Step(`^the "([^"]*)" has a service called "([^"]*)" in the namespace "([^"]*)"$`, thePodHasAServiceCalledInTheNamespace)
}

//go:embed features/*
var features embed.FS

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
			FS:       features,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
