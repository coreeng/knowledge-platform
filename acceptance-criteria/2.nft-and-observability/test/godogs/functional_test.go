package godogs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
	"net/url"
	"nft-and-observability/consts"
	metrics "nft-and-observability/metrics"
	"nft-and-observability/structs"
	"strconv"
	"strings"
	"testing"
	"time"
)

var kubernetesClient *kubernetes.Clientset
var owner string
var pods *corev1.PodList
var k6OperatorSystemNamespace *corev1.Namespace
var k6OperatorSystemPods *corev1.PodList
var resourceQuota *corev1.ResourceQuotaList
var k6TestDataResultMetric *structs.K6TestResultMetric
var request *resty.Request
var response resty.Response
var elapsedMilliseconds int64

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

func theBootcampParticipantIsDefined() error {
	deploymentsList, err := kubernetesClient.AppsV1().Deployments(consts.ReferenceServiceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed fetching the deployments for namespace %s, err: %v", consts.ReferenceServiceNamespace, err)
	}
	if len(deploymentsList.Items) == 0 {
		metrics.PushFailureMetric(consts.DeploymentPresent, "")
		return fmt.Errorf("no deployments found in namespace %v", consts.ReferenceServiceNamespace)
	}

	for _, deployment := range deploymentsList.Items {
		if deployment.Name == consts.ReferenceServiceDeployment {
			owner = deployment.Labels[consts.DeploymentOwnerLabel]
			if owner == "" {
				metrics.PushFailureMetric(consts.DeploymentPresentOwner, "")
			}
			return nil
		}
	}
	metrics.PushFailureMetric(consts.DeploymentPresent, owner)
	return fmt.Errorf("could not find a \"%s\" deployment", consts.ReferenceServiceDeployment)
}

func iFetchPodsForTheNamespace() error {
	podList, err := kubernetesClient.CoreV1().Pods(consts.DefaultNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error occured while listing pods in namespace %s: %v", consts.DefaultNamespace, err)
	}

	if len(podList.Items) == 0 {
		metrics.PushFailureMetric(consts.NoPodsRunningInNamespace, owner)
		return fmt.Errorf("no running pods detected in namespace %s", consts.DefaultNamespace)
	}
	pods = podList
	return nil
}

func isPodRunningSuccessfully(podNameLabel string) error {
	for _, pod := range pods.Items {
		if pod.Labels != nil && pod.Labels[consts.PodNameLabel] == podNameLabel {
			if pod.Status.Phase == consts.PodStatusRunning {
				metrics.PushSuccessMetric(fmt.Sprintf(consts.PodRunning, podNameLabel), owner)
				if isAllContainersRunning(&pod) {
					metrics.PushSuccessMetric(fmt.Sprintf(consts.PodContainersRunning, podNameLabel), owner)
					return nil
				} else {
					metrics.PushFailureMetric(fmt.Sprintf(consts.PodContainersRunning, podNameLabel), owner)
					return fmt.Errorf("%s pod has non-running containers", podNameLabel)
				}
			} else {
				metrics.PushFailureMetric(fmt.Sprintf(consts.PodRunning, podNameLabel), owner)
				return fmt.Errorf("%s pod is up but not in a running state", podNameLabel)
			}
		}
	}
	return nil
}

func thereIsARunningPrometheus() error {
	return isPodRunningSuccessfully(consts.PrometheusPodNameLabel)
}

func thereIsARunningAlertmanager() error {
	return isPodRunningSuccessfully(consts.AlertManagerPodNameLabel)
}

func thereIsARunningGrafana() error {
	return isPodRunningSuccessfully(consts.GrafanaPodNameLabel)
}

func isAllContainersRunning(prometheusPod *corev1.Pod) bool {
	for _, containerStatus := range prometheusPod.Status.ContainerStatuses {
		if !containerStatus.Ready {
			return false
		}
	}
	return true
}

func iGetTheK6OperatorSystemNamespace() error {
	fetchedK6OperatorSystemNamespace, err := kubernetesClient.CoreV1().Namespaces().Get(context.TODO(), consts.K6OperatorNamespace, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Not found -- handled in the subsequent step
			return nil
		}
		return fmt.Errorf("error occurred while fetching the %s namespace", consts.K6OperatorNamespace)
	}
	k6OperatorSystemNamespace = fetchedK6OperatorSystemNamespace
	return nil
}

func theNamespaceIsReturned() error {
	if k6OperatorSystemNamespace == nil {
		metrics.PushFailureMetric(consts.K6OperatorNamespaceSetUp, owner)
		return fmt.Errorf("%s namespace not yet set up", consts.K6OperatorNamespace)
	}
	metrics.PushSuccessMetric(consts.K6OperatorNamespaceSetUp, owner)
	return nil
}

func iGetThePodsInTheK6OperatorSystemNamespace() error {
	k6OperatorNsPods, err := kubernetesClient.CoreV1().Pods("k6-operator-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error while listing pods in the %s namespace", consts.K6OperatorNamespace)
	}
	k6OperatorSystemPods = k6OperatorNsPods
	return nil
}

func thereIsARunningK6OperatorPod() error {
	if len(k6OperatorSystemPods.Items) == 0 {
		metrics.PushFailureMetric(consts.K6OperatorPodRunning, owner)
		return fmt.Errorf("k6 operator pod not running")
	}

	for _, pod := range k6OperatorSystemPods.Items {
		if strings.Contains(pod.Name, "k6-operator") {
			metrics.PushSuccessMetric(consts.K6OperatorPodRunning, owner)
			return nil
		}
	}
	metrics.PushFailureMetric(consts.K6OperatorPodRunning, owner)
	return fmt.Errorf("k6 operator pod not running")
}

func theK6TestResultDashboardIsPresent() error {
	req, err := http.NewRequest("GET", consts.GrafanaSearchUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating grafana request, err: %v", err)
	}

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(consts.DefaultGrafanaUsername+":"+consts.DefaultGrafanaPassword))
	req.Header.Set("Authorization", basicAuth)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making grafana request, err: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received a non-200 status code from the grafana request, status code: %d", resp.StatusCode)
	}

	dashboardItems, err := extractDashboardItems(err, resp)

	if err != nil {
		return err
	}
	for _, dashboardItem := range dashboardItems {
		if dashboardItem.Title == consts.DefaultGrafanaK6DashboardTitle {
			metrics.PushSuccessMetric(consts.K6TestRunDashboardSetUp, owner)
			return nil
		}
	}
	metrics.PushFailureMetric(consts.K6TestRunDashboardSetUp, owner)
	return fmt.Errorf("k6 test run dashboard not set up yet")
}

func extractDashboardItems(err error, resp *http.Response) ([]structs.DashboardSearchItem, error) {
	var dashboardItems []structs.DashboardSearchItem
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading grafana response body, err: %v", err)
	}
	err = json.Unmarshal(body, &dashboardItems)
	if err != nil {
		return nil, fmt.Errorf("error decoding grafana response, err: %v", err)
	}
	return dashboardItems, nil
}

func iFetchPrometheusMetrics() error {
	// Noop
	return nil
}

func theFollowingK6MetricsArePresent(expectedK6Metrics *godog.Table) error {
	for _, expectedK6Metric := range expectedK6Metrics.Rows {
		k6Metric := expectedK6Metric.Cells[0].Value
		query := fmt.Sprintf(consts.CountOverTimeQuery, k6Metric)
		queryParam := url.Values{"query": {query}}

		parsedUrl, err := url.Parse(consts.PrometheusQueryUrl)
		if err != nil {
			return fmt.Errorf("error parsing Prometheus URL: %v", err)
		}
		parsedUrl.RawQuery = queryParam.Encode()
		queryResponse, err := http.Get(parsedUrl.String())
		if err != nil {
			return fmt.Errorf("error making Prometheus GET request: %v", err)
		}

		body, err := io.ReadAll(queryResponse.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %v", err)
		}
		var result structs.PrometheusQueryResult
		err = json.Unmarshal(body, &result)
		if err != nil {
			return fmt.Errorf("error unmarshaling Prometheus JSON response: %v", err)
		}
		if len(result.Data.Result) == 0 {
			metrics.PushFailureMetric(fmt.Sprintf(consts.K6MetricPresent, k6Metric), owner)
		} else {
			metrics.PushSuccessMetric(fmt.Sprintf(consts.K6MetricPresent, k6Metric), owner)
		}
	}
	return nil
}

func iFetchResourceQuotaForTheReferenceServiceApp() error {
	resourceQuotaList, err := kubernetesClient.CoreV1().ResourceQuotas(consts.ReferenceServiceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error fetching resource quota for the %s namespace, err: %v", consts.ReferenceServiceNamespace, err)
	}
	resourceQuota = resourceQuotaList
	return nil
}

func memoryAndCPUIsSet() error {
	if len(resourceQuota.Items) == 0 {
		metrics.PushFailureMetric(consts.ResourceQuotaSet, owner)
		return fmt.Errorf("resource quota not yet set")
	}

	memoryRequests := resourceQuota.Items[0].Spec.Hard[consts.MemoryRequests]
	cpuRequests := resourceQuota.Items[0].Spec.Hard[consts.CpuRequests]
	memoryLimits := resourceQuota.Items[0].Spec.Hard[consts.MemoryLimits]

	if &memoryRequests == nil || &cpuRequests == nil {
		metrics.PushFailureMetric(consts.ResourceQuotaSet, owner)
		return fmt.Errorf("memory or cpu requests not yet set")
	}

	if &memoryLimits == nil {
		metrics.PushFailureMetric(consts.ResourceQuotaSet, owner)
		return fmt.Errorf("memory and cpu requests set, but memory limits not yet set")
	}

	metrics.PushSuccessMetric(consts.ResourceQuotaSet, owner)
	return nil
}

func iFetchTheK6TestResultMetric() error {
	err, metric := metrics.FetchK6TestResultMetric()
	if err != nil {
		return err
	}
	k6TestDataResultMetric = metric
	return nil
}

func theK6TestResultMetricIsPresent() error {
	if k6TestDataResultMetric == nil {
		metrics.PushFailureMetric(consts.K6TestResultMetricPresent, owner)
		return fmt.Errorf("k6-test-result metric not yet present")
	}
	metrics.PushSuccessMetric(consts.K6TestResultMetricPresent, owner)
	return nil
}

func theLastTestRunWasSuccessful() error {
	if k6TestDataResultMetric.Status == "success" {
		metrics.PushSuccessMetric(consts.K6TestRunSuccess, owner)
		return nil
	}
	metrics.PushFailureMetric(consts.K6TestRunSuccess, owner)
	return fmt.Errorf("last k6 test run failed")
}

func maxVUSWere500ForTheLastTestRun() error {
	maxVus, err := strconv.Atoi(k6TestDataResultMetric.MaxVUS)
	if err != nil {
		return fmt.Errorf("failed to parse %s into an int", k6TestDataResultMetric.MaxVUS)
	}

	if maxVus >= 500 {
		metrics.PushSuccessMetric(consts.K6TestSupport500VUS, owner)
		return nil
	}

	metrics.PushFailureMetric(consts.K6TestSupport500VUS, owner)
	return fmt.Errorf("support 500 VUS requirement not yet met")
}

func maxVUSWere1500ForTheLastTestRun() error {
	maxVus, err := strconv.Atoi(k6TestDataResultMetric.MaxVUS)
	if err != nil {
		return fmt.Errorf("failed to parse %s into an int", k6TestDataResultMetric.MaxVUS)
	}

	if maxVus >= 1500 {
		metrics.PushSuccessMetric(consts.K6TestSupport1500VUS, owner)
		return nil
	}

	metrics.PushFailureMetric(consts.K6TestSupport1500VUS, owner)
	return fmt.Errorf("support 1500 VUS requirement not yet met")
}

func the95thPercentileWas50MSOrLessForTheLastTestRun() error {
	p95ResponseTime, err := strconv.ParseFloat(k6TestDataResultMetric.P95ResponseTime, 64)
	if err != nil {
		return fmt.Errorf("failed to parse %s into an int", k6TestDataResultMetric.P95ResponseTime)
	}

	if p95ResponseTime <= 50 {
		metrics.PushSuccessMetric(consts.K6TestP95ResponseWithin50ms, owner)
		return nil
	}
	metrics.PushFailureMetric(consts.K6TestP95ResponseWithin50ms, owner)
	return fmt.Errorf("p95 response was greater than 50 ms")
}

func anHttpClientIsSetUp() {
	httpClient := resty.New()
	request = httpClient.R()
}

func iCallTheDownstreamDependencyWithASecondOfResponseDelay() error {
	delayUrlWithSecondOfDelay := fmt.Sprintf(consts.ReferenceServiceDelayUrl, "1")
	beginningOfRequest := time.Now()
	res, err := request.Get(delayUrlWithSecondOfDelay)
	if err != nil {
		return fmt.Errorf("error occurred while calling %s. Err: %v", delayUrlWithSecondOfDelay, err)
	}
	response = *res
	elapsedMilliseconds = time.Since(beginningOfRequest).Milliseconds()

	return nil
}

func iGetAResponseBackWithin600MS() error {
	if elapsedMilliseconds <= 600 {
		metrics.PushSuccessMetric(consts.DownstreamDelayHandling, owner)
		return nil
	}
	metrics.PushFailureMetric(consts.DownstreamDelayHandling, owner)
	return fmt.Errorf("response not received within 600 ms")
}

func theStatusCodeIs503() error {
	if response.StatusCode() == 503 {
		metrics.PushSuccessMetric(consts.DownstreamDelayStatus, owner)
		return nil
	}
	metrics.PushFailureMetric(consts.DownstreamDelayStatus, owner)
	return fmt.Errorf("non 503 status code received")
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the kubernetes client is set up$`, theKubernetesClientIsSetUp)
	ctx.Step(`^the bootcamp participant is defined$`, theBootcampParticipantIsDefined)
	ctx.Step(`^i fetch pods for the namespace$`, iFetchPodsForTheNamespace)
	ctx.Step(`^there is a running prometheus$`, thereIsARunningPrometheus)
	ctx.Step(`^there is a running alertmanager$`, thereIsARunningAlertmanager)
	ctx.Step(`^there is a running grafana$`, thereIsARunningGrafana)
	ctx.Step(`^the k6 test result dashboard is present in grafana$`, theK6TestResultDashboardIsPresent)
	ctx.Step(`^i get the k6-operator-system namespace$`, iGetTheK6OperatorSystemNamespace)
	ctx.Step(`^the namespace is returned$`, theNamespaceIsReturned)
	ctx.Step(`^i get pods in the k6-operator-system namespace$`, iGetThePodsInTheK6OperatorSystemNamespace)
	ctx.Step(`^there is a running k6 operator pod$`, thereIsARunningK6OperatorPod)
	ctx.Step(`^i fetch prometheus metrics$`, iFetchPrometheusMetrics)
	ctx.Step(`^the following k6 metrics are present:$`, theFollowingK6MetricsArePresent)
	ctx.Step(`^i fetch resource quota for the reference-service application$`, iFetchResourceQuotaForTheReferenceServiceApp)
	ctx.Step(`^memory and CPU is set$`, memoryAndCPUIsSet)
	ctx.Step(`^i fetch the k6 test result metric$`, iFetchTheK6TestResultMetric)
	ctx.Step(`^the k6 test result metric is present$`, theK6TestResultMetricIsPresent)
	ctx.Step(`^the last test run was successful`, theLastTestRunWasSuccessful)
	ctx.Step(`^max VUs were 500 during the last test run$`, maxVUSWere500ForTheLastTestRun)
	ctx.Step(`^max VUs were 1500 during the last test run$`, maxVUSWere1500ForTheLastTestRun)
	ctx.Step(`^the response time 95th percentile was 50 ms or less for the last test run$`, the95thPercentileWas50MSOrLessForTheLastTestRun)
	ctx.Step(`^an HTTP client is set up$`, anHttpClientIsSetUp)
	ctx.Step(`^i call the downstream dependency with a second of response delay$`, iCallTheDownstreamDependencyWithASecondOfResponseDelay)
	ctx.Step(`^i get a response back within 600ms$`, iGetAResponseBackWithin600MS)
	ctx.Step(`^the status code is 503$`, theStatusCodeIs503)
}

func TestFeatures(t *testing.T) {
	metrics.Init()

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
	metrics.PushTestSuiteOutcomeMetric(consts.NftAndObsModule, owner)
}
