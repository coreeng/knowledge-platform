package godogs

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
	"strings"
	"testing"
)

var kubernetesClient *kubernetes.Clientset
var impersonateAccountClient *kubernetes.Clientset
var accountToImpersonate string
var namespaces map[string]*corev1.Namespace
var podPerNamespace map[string]*corev1.PodList

func theKubernetesClientIsSetUp() error {
	if kubernetesClient == nil {
		logrus.Info("Attempting to fetch in-cluster config..")
		config, err := rest.InClusterConfig()
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

func iCanAccessTheFollowingNamespaces(namespaces string) error {
	//TODO - add namespace here from the input
	pods, err := impersonateAccountClient.CoreV1().Pods("team-a").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		println("----Error while fetching pods----")
		println(err.Error())
		return fmt.Errorf("could not fetch namespaces")
	}
	for _, pod := range pods.Items {
		println("***Pods***")
		println(pod.Name)
	}
	return nil
}

func iCannotAccessTheFollowingNamespaces(arg1 string) error {
	return nil
}

func aServiceAccountForNamespaceAlreadyExists(arg1 string) error {
	return nil
}

func iImpersonateTheServiceAccount() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return fmt.Errorf("failed fetching in cluster config, err: %v", err)
	}

	//TODO - add account to impersonate from input
	accountToImpersonate = fmt.Sprintf("system:serviceaccount:%s", "team-a:team-a-sa")
	println("---Account to impersonate")
	println(accountToImpersonate)
	config.Impersonate = rest.ImpersonationConfig{UserName: accountToImpersonate}
	logrus.Info("Attempting to create a kubernetes client after impersonation..")
	impersonateAccountClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error occured while impersonating service account")
	}
	logrus.Info("Impersonation was successful..")
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I get the "([^"]*)" namespace$`, iGetTheNamespace)
	ctx.Step(`^the kubernetes client is setup$`, theKubernetesClientIsSetUp)
	ctx.Step(`^the namespace "([^"]*)" is returned$`, theNamespaceIsReturned)
	ctx.Step(`^I get the pods in the "([^"]*)" namespace$`, iGetThePodsInTheNamespace)
	ctx.Step(`^there is a running "([^"]*)" pod in the namespace "([^"]*)"$`, thereIsARunningPod)
	ctx.Step(`^the "([^"]*)" namespace has the subnamespaces "([^"]*)"$`, theNamespaceHasTheSubnamespaces)
	ctx.Step(`^I can access the following namespaces: "([^"]*)"$`, iCanAccessTheFollowingNamespaces)
	ctx.Step(`^I cannot access the following namespaces: "([^"]*)"$`, iCannotAccessTheFollowingNamespaces)
	ctx.Step(`^a service account for "([^"]*)" already exists$`, aServiceAccountForNamespaceAlreadyExists)
	ctx.Step(`^I impersonate the service account$`, iImpersonateTheServiceAccount)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
