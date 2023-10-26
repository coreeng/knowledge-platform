package structs

type PrometheusQueryResult struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type DashboardSearchItem struct {
	Title string `json:"title"`
	UID   string `json:"uid"`
}

type K6TestResultMetric struct {
	MaxVUS          string `json:"vus_max"`
	P95ResponseTime string `json:"p(95)"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}
