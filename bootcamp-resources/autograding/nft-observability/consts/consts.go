package consts

const (
	//Metrics

	DeploymentPresent           = "deployment-present"
	DeploymentPresentOwner      = "deployment-present-owner"
	NoPodsRunningInNamespace    = "no-pods-running-in-namespace"
	PodRunning                  = "%s-pod-running"
	PodContainersRunning        = "%s-containers-running"
	K6OperatorNamespaceSetUp    = "k6-operator-system-namespace-set-up"
	K6OperatorPodRunning        = "k6-operator-pod-running"
	K6TestRunDashboardSetUp     = "k6-test-run-dashboard-set-up"
	K6MetricPresent             = "%s-metric-present"
	K6TestResultMetricPresent   = "k6-test-result-metric-present"
	K6TestRunSuccess            = "k6-last-test-run-success"
	K6TestSupport500VUS         = "support-500-vus"
	K6TestSupport1500VUS        = "support-1500-vus"
	K6TestP95ResponseWithin50ms = "k6-test-p95-response-within-50ms"
	ResourceQuotaSet            = "resource-quota-set"
	DownstreamDelayHandling     = "downstream-delay-handling"
	DownstreamDelayStatus       = "downstream-delay-status"
	NftAndObsModule             = "nft-and-obs"

	// Cluster values

	DefaultNamespace           = "default"
	K6OperatorNamespace        = "k6-operator-system"
	ReferenceServiceDeployment = "reference-service"
	ReferenceServiceNamespace  = "reference-service-showcase"
	PrometheusQueryUrl         = "http://prom-kube-prometheus-stack-prometheus.default.svc.cluster.local:9090/api/v1/query"
	ReferenceServiceDelayUrl   = "http://reference-service.reference-service-showcase/delay/%s"
	GrafanaSearchUrl           = "http://prom-grafana.default/api/search"
	PushGatewayUrl             = "http://pushgateway.pushgateway-autograding:9091"
	DeploymentOwnerLabel       = "owner"
	PodStatusRunning           = "Running"
	PodNameLabel               = "app.kubernetes.io/name"
	PrometheusPodNameLabel     = "prometheus"
	AlertManagerPodNameLabel   = "alertmanager"
	GrafanaPodNameLabel        = "grafana"
	MemoryRequests             = "requests.memory"
	CpuRequests                = "requests.cpu"
	MemoryLimits               = "limits.memory"

	// Grafana

	DefaultGrafanaUsername         = "admin"
	DefaultGrafanaPassword         = "prom-operator"
	DefaultGrafanaK6DashboardTitle = "K6 Test Result"

	// Prom Query

	CountOverTimeQuery = "count_over_time({__name__=\"%s\"}[48h])"
)
