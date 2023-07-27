package godogs

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"p2p-fast-feedback/test/godogs/pushgateway"
	"testing"
)

const Namespace = "reference-service-showcase"
const ReferenceServiceCounterUrl = "http://reference-service.reference-service-showcase/counter"

var kubernetesClient *kubernetes.Clientset
var deployments *v1.DeploymentList
var referenceServiceReplicas int
var request *resty.Request
var response resty.Response
var owner string

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

func iListDeploymentsForTheReferenceServiceShowcaseNamespace() error {
	deploymentsList, err := kubernetesClient.AppsV1().Deployments(Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed fetching the deployments for namespace %s, err: %v", Namespace, err)
	}
	deployments = deploymentsList
	return nil
}

func aDeploymentIsPresent() error {
	if len(deployments.Items) == 0 {
		pushgateway.PushFailureMetric("deployment-present", "")
		return fmt.Errorf("no deployments found in namespace %v", Namespace)
	}

	var deploymentNames []string
	for _, deployment := range deployments.Items {
		if deployment.Name == "reference-service" {
			referenceServiceReplicas = int(*deployment.Spec.Replicas)
			owner = deployment.Labels["owner"]
			if owner == "" {
				pushgateway.PushFailureMetric("deployment-present-owner", "")
			}
			pushgateway.PushSuccessMetric("deployment-present", owner)
			return nil
		}
		deploymentNames = append(deploymentNames, deployment.Name)
	}
	pushgateway.PushFailureMetric("deployment-present", owner)
	return fmt.Errorf("could not find a \"reference-service\" deployment, deployment names present: %v", deploymentNames)
}

func aSingleReplicaOfTheReferenceServiceIsRunning() error {
	if referenceServiceReplicas == 1 {
		pushgateway.PushSuccessMetric("single-replica-running", owner)
		return nil
	}
	pushgateway.PushFailureMetric("single-replica-running", owner)
	return fmt.Errorf("more than one replica of the reference-service found")
}

func anHttpClientIsSetUp() {
	httpClient := resty.New()
	request = httpClient.R()
}

func iReceiveASuccessfulResponse() error {
	if response.IsSuccess() == true {
		pushgateway.PushSuccessMetric("delete-counter-present", owner)
		return nil
	}
	pushgateway.PushFailureMetric("delete-counter-present", owner)
	return fmt.Errorf("unexpected response received. Request URL: %s, Request Method: %s, Response code: %d, error: %v", response.Request.URL, response.Request.Method, response.StatusCode(), response.Error())
}
func iPollTheDeleteCounterEndpoint() error {
	logrus.Infof("Hitting DELETE counter endpoint")
	httpResponse, err := request.Delete(ReferenceServiceCounterUrl + "/p2p-fast-feedback-ac")

	if err != nil {
		return fmt.Errorf("error occurred while calling DELETE counter, err: %v", err)
	}

	response = *httpResponse
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the kubernetes client is set up$`, theKubernetesClientIsSetUp)
	ctx.Step(`^I list deployments for the reference-service-showcase namespace$`, iListDeploymentsForTheReferenceServiceShowcaseNamespace)
	ctx.Step(`^a reference-service deployment is present$`, aDeploymentIsPresent)
	ctx.Step(`^a single replica of the reference-service is running$`, aSingleReplicaOfTheReferenceServiceIsRunning)
	ctx.Step(`^An HTTP Client is set up$`, anHttpClientIsSetUp)
	ctx.Step(`^I receive a successful response$`, iReceiveASuccessfulResponse)
	ctx.Step(`^I poll the DELETE counter endpoint$`, iPollTheDeleteCounterEndpoint)
}

func TestFeatures(t *testing.T) {
	pushgateway.Init()

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
	pushgateway.PushTestSuiteOutcomeMetric("p2p-fast-feedback", owner)
}
