package acceptance

import (
	"context"
	"embed"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/rdumont/assistdog"
	"github.com/sirupsen/logrus"
	"html/template"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/client-go/util/retry"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

type nsCtxKey struct{}
type controllerPods struct{}
type crdCtxKey struct{}
type autogradingNsCtxKey struct{}
type manifestCrKey struct{}

var kubernetesClient *kubernetes.Clientset
var k8ApiExtensionClient *apiextensionsclientset.Clientset

//go:embed features/*

var features embed.FS

func getConfig() (*rest.Config, error) {
	logrus.Info("Attempting to fetch in-cluster config.")
	var (
		config *rest.Config
		err    error
	)
	if os.Getenv("RUN_OUTSIDE_CLUSTER") == "true" {
		config, err = clientcmd.BuildConfigFromFlags(
			"",
			filepath.Join(homedir.HomeDir(), ".kube", "config"),
		)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return config, fmt.Errorf("failed fetching in cluster config, err: %v", err)
	}
	return config, err
}
func iGetTheNamespace(ctx context.Context, namespaceName string) (context.Context, error) {
	var namespace, err = kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Not found -- handled in the subsequent step
			return ctx, nil
		}
		return ctx, fmt.Errorf("error occurred while fetching the %s namespace", namespace)
	}
	return context.WithValue(ctx, nsCtxKey{}, namespace), nil
}

func iGetThePodsInTheNamespace(ctx context.Context, namespaceName string) (context.Context, error) {
	pods, err := kubernetesClient.CoreV1().Pods(namespaceName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return ctx, fmt.Errorf("error while listing pods in the %s namespace", namespaceName)
	}
	return context.WithValue(ctx, controllerPods{}, pods), nil
}

func theKubernetesClientIsSetup(ctx context.Context) (context.Context, error) {
	config, err := getConfig()
	if err != nil {
		return ctx, fmt.Errorf("failed fetching in cluster config, err: %v", err)
	}
	logrus.Info("Attempting to create a kubernetes client.")
	kubernetesClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return ctx, fmt.Errorf("error creating a kubernetes client, err: %v", err)
	}
	logrus.Info("Kubernetes client creates. Proceeding with the test.")
	return ctx, nil
}

func theNamespaceExists(ctx context.Context, namespaceName string) (context.Context, error) {
	if ctx.Value(nsCtxKey{}) != nil && ctx.Value(nsCtxKey{}).(*corev1.Namespace).Name == namespaceName {
		return ctx, nil
	} else {
		return ctx, fmt.Errorf("controller namespace %s not found", namespaceName)
	}

}

func thereIsARunningControllerManagerPod(ctx context.Context) (context.Context, error) {
	if ctx.Value(controllerPods{}) == nil {

		return ctx, fmt.Errorf("no controller manager pods are running")
	}

	for _, pod := range ctx.Value(controllerPods{}).(*corev1.PodList).Items {
		if pod.GetLabels()["control-plane"] == "controller-manager" {
			return ctx, nil
		}
	}
	return ctx, fmt.Errorf("no controller manager pods are running")
}

func theKubernetesApiExtensionClientExists() error {
	config, err := getConfig()

	if err != nil {
		return fmt.Errorf("failed fetching in cluster config, err: %v", err)
	}
	logrus.Info("Attempting to create a kubernetes client..")

	k8ApiExtensionClient, err = apiextensionsclientset.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed creating a kubernetes api extension client, err: %v", err)
	}
	return nil

}

func iGetTheCustomResourceDefinition(ctx context.Context, resourceDefinitionName string) (context.Context, error) {
	customResource, err := k8ApiExtensionClient.
		ApiextensionsV1().
		CustomResourceDefinitions().
		Get(context.TODO(), resourceDefinitionName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Not found -- handled in the subsequent step
			return ctx, nil
		}
		return ctx, fmt.Errorf("the custom resource definition %s does not exist", resourceDefinitionName)
	}
	return context.WithValue(ctx, crdCtxKey{}, customResource), nil
}

func theCustomResourceDefinitionExists(ctx context.Context, resourceDefinitionName string) (context.Context, error) {
	if ctx.Value(crdCtxKey{}) != nil && ctx.Value(crdCtxKey{}).(*apiextensionsv1.CustomResourceDefinition).Name == resourceDefinitionName {
		return ctx, nil
	} else {
		return ctx, fmt.Errorf("crd %s not found", resourceDefinitionName)
	}
}

func theCanaryDeploymentIsCreated(ctx context.Context, deploymentName string) (context.Context, error) {
	if err := retry.OnError(
		wait.Backoff{Duration: 5 * time.Second, Steps: 5},
		func(err error) bool { return true },
		func() error {
			_, err := kubernetesClient.
				AppsV1().
				Deployments(ctx.Value(autogradingNsCtxKey{}).(string)).Get(context.TODO(), deploymentName, metav1.GetOptions{})
			return err
		},
	); err != nil {
		return ctx, fmt.Errorf("could not find any canary deployment")
	}

	return ctx, nil
}

func iHaveANamespace(ctx context.Context, namespaceName string) (context.Context, error) {
	_, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		return ctx, fmt.Errorf("error retrieve %s", namespaceName)
	}
	return context.WithValue(ctx, autogradingNsCtxKey{}, namespaceName), nil
}

func iHaveTheFollowingCR(ctx context.Context, data *godog.Table) (context.Context, error) {
	pwd, err := os.Getwd()
	workingDirPath, err := filepath.Abs(pwd)
	manifestsPath := filepath.Join(workingDirPath, "test/acceptance/manifests/")
	templatePath := filepath.Join(manifestsPath, "cr-template.yaml")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		logrus.Info(err.Error())
		return ctx, fmt.Errorf("error parsing the CR template")
	}
	substitutionValues, err := assistdog.NewDefault().ParseMap(data)
	if err != nil {
		return ctx, fmt.Errorf("error parsing the data table from the step")
	}
	crManifestFile, _ := os.Create(filepath.Join(manifestsPath, "cr.yaml"))

	tmpl.Execute(crManifestFile, substitutionValues)

	_, err = exec.Command("kubectl", "apply", "-n", ctx.Value(autogradingNsCtxKey{}).(string), "-f", crManifestFile.Name()).Output()
	if err != nil {
		logrus.Info(err.Error())
		return ctx, fmt.Errorf("error creating the custom resource")
	}
	return context.WithValue(ctx, manifestCrKey{}, crManifestFile.Name()), nil
}

func iUpdateTheCRWith(ctx context.Context, data *godog.Table) (context.Context, error) {
	return iHaveTheFollowingCR(ctx, data)
}

// cleanup functions

func cleanupCrs(ctx context.Context) error {
	if ctx.Value(autogradingNsCtxKey{}) != nil {
		output, err := exec.Command("kubectl", "-n", ctx.Value(autogradingNsCtxKey{}).(string), "delete", "canariedapps", "--all").Output()
		logrus.Info(output)
		return err
	}
	return nil

}

func cleanupManifests(ctx context.Context) error {
	if ctx.Value(manifestCrKey{}) != nil {
		err := os.Remove(ctx.Value(manifestCrKey{}).(string))
		if err != nil {
			return fmt.Errorf("error deleting cr manifest")
		}
	}
	return nil
}
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		e := cleanupCrs(ctx)
		if e != nil {
			return ctx, e
		}

		e = cleanupManifests(ctx)
		if e != nil {
			return ctx, e
		}
		var scenarioSuccess = err == nil
		PushScenarioMetric(sc.Name, scenarioSuccess)
		return ctx, err
	})
	ctx.Step(`^I get the "([^"]*)" namespace$`, iGetTheNamespace)
	ctx.Step(`^I get the pods in the "([^"]*)" namespace$`, iGetThePodsInTheNamespace)
	ctx.Step(`^the kubernetes client is setup$`, theKubernetesClientIsSetup)
	ctx.Step(`^the namespace "([^"]*)" exists$`, theNamespaceExists)
	ctx.Step(`^there is a running controller manager pod$`, thereIsARunningControllerManagerPod)
	ctx.Step(`^I get the custom resource definition "([^"]*)"$`, iGetTheCustomResourceDefinition)
	ctx.Step(`^the custom resource definition "([^"]*)" exists$`, theCustomResourceDefinitionExists)
	ctx.Step(`^the kubernetes api extension client exists$`, theKubernetesApiExtensionClientExists)
	ctx.Step(`^the canary deployment "([^"]*)" is created$`, theCanaryDeploymentIsCreated)
	ctx.Step(`^I have a namespace "([^"]*)"$`, iHaveANamespace)
	ctx.Step(`^I have the following CR:$`, iHaveTheFollowingCR)
	ctx.Step(`^I update the CR with:$`, iUpdateTheCRWith)
}

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
