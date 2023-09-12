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
	"strings"
	"testing"
)

type multiTenancyFeature struct {
	kubernetesClient *kubernetes.Clientset
	namespaces       map[string]*corev1.Namespace
	podPerNamespace  map[string]*corev1.PodList
}

func (feature *multiTenancyFeature) theKubernetesClientIsSetUp() error {
	if feature.kubernetesClient == nil {
		logrus.Info("Attempting to fetch in-cluster config..")
		config, err := rest.InClusterConfig()
		if err != nil {
			return fmt.Errorf("failed fetching in cluster config, err: %v", err)
		}
		logrus.Info("Attempting to create a kubernetes client..")
		feature.kubernetesClient, err = kubernetes.NewForConfig(config)
		if err != nil {
			return fmt.Errorf("failed creating a kubernetes client, err: %v", err)
		}
		logrus.Info("Proceeding with the acceptance criteria test..")
		return nil
	}
	return nil
}

func (feature *multiTenancyFeature) iGetTheNamespace(namespaceName string) error {
	namespace, err := feature.kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Not found -- handled in the subsequent step
			return nil
		}
		return fmt.Errorf("error occurred while fetching the %s namespace", namespaceName)
	}
	if feature.namespaces == nil {
		feature.namespaces = make(map[string]*corev1.Namespace)
	}
	feature.namespaces[namespaceName] = namespace
	return nil
}

func (feature *multiTenancyFeature) theNamespaceIsReturned(namespaceName string) error {
	if feature.namespaces[namespaceName] == nil {
		return fmt.Errorf("%s namespace not yet set up", namespaceName)
	}
	return nil
}

func (feature *multiTenancyFeature) iGetThePodsInTheNamespace(namespace string) error {
	pods, err := feature.kubernetesClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error while listing pods in the %s namespace", namespace)
	}
	if feature.podPerNamespace == nil {
		feature.podPerNamespace = make(map[string]*corev1.PodList)
	}
	feature.podPerNamespace[namespace] = pods
	return nil
}

func (feature *multiTenancyFeature) thereIsARunningPod(podName string, namespace string) error {
	if len(feature.podPerNamespace[namespace].Items) == 0 {
		return fmt.Errorf("no pods are running in the %s namespace", namespace)
	}

	for _, pod := range feature.podPerNamespace[namespace].Items {
		if strings.Contains(pod.Name, podName) {
			return nil
		}
	}
	return fmt.Errorf("the pod %s is not running in the namespace %s", podName, namespace)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	feature := &multiTenancyFeature{}
	ctx.Step(`^I get the "([^"]*)" namespace$`, feature.iGetTheNamespace)
	ctx.Step(`^the kubernetes client is setup$`, feature.theKubernetesClientIsSetUp)
	ctx.Step(`^the namespace "([^"]*)" is returned$`, feature.theNamespaceIsReturned)
	ctx.Step(`^I get the pods in the "([^"]*)" namespace$`, feature.iGetThePodsInTheNamespace)
	ctx.Step(`^there is a running "([^"]*)" pod in the namespace "([^"]*)"$`, feature.thereIsARunningPod)
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
