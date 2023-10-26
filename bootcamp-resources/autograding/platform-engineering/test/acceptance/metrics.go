package acceptance

import (
	"context"
	"fmt"

	"time"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
)

const (
	PushGatewayUrl = "http://pushgateway.pushgateway-autograding:9091"
	promURL        = "kube-prometheus-stack-prometheus.platform-engineering-autograding.svc.cluster.local"
)

type PrometheusResponse struct {
	Data struct {
		Result []struct {
			Value []interface{}
		}
	}
}

var scenarioOutcomeMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "acceptance_criteria_scenario",
		Help: "The acceptance criteria scenario that either passed or failed",
	},
	[]string{"scenario", "outcome"})

func PushScenarioMetric(scenarioName string, scenarioSuccess bool) {
	var scenarioOutcomeLabel string
	var metricValue float64

	if scenarioSuccess == true {
		scenarioOutcomeLabel = "success"
		metricValue = 1
	} else {
		scenarioOutcomeLabel = "failure"
		metricValue = 0
	}
	pusher := push.New(PushGatewayUrl, "acceptance-criteria-platform-engineering-scenario-outcome")
	scenarioOutcomeMetric.WithLabelValues(scenarioName, scenarioOutcomeLabel).Set(metricValue)
	pusher.Collector(scenarioOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to pushgateway, err: %v", err)
	}
}

func getMetrics(ctx context.Context, query string) (string, error) {
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("%s:9090", promURL),
	})
	if err != nil {
		return "", fmt.Errorf("create prometheus client: %w", err)
	}

	v1api := v1.NewAPI(client)

	res, warnings, err := v1api.Query(ctx, query, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		return "", fmt.Errorf("querying Prometheus: %w", err)
	}

	if len(warnings) > 0 {
		return "", fmt.Errorf("prometheus warnings: %v", warnings)
	}
	result := res.(model.Vector)
	var out string

	for _, v := range result {
		out = v.Value.String()
	}

	return out, nil
}

func getMetricsRate(ctx context.Context, query string, interval time.Time) (int, error) {
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("%s:9090", promURL),
	})
	if err != nil {
		return 0, fmt.Errorf("create prometheus client: %w", err)
	}

	v1api := v1.NewAPI(client)
	r := v1.Range{
		Start: interval,
		End:   time.Now(),
		Step:  time.Minute,
	}

	res, warnings, err := v1api.QueryRange(ctx, query, r, v1.WithTimeout(5*time.Second))
	if err != nil {
		return 0, fmt.Errorf("querying Prometheus: %w", err)
	}

	if len(warnings) > 0 {
		return 0, fmt.Errorf("prometheus warnings: %v", warnings)
	}
	result := res.(model.Matrix)
	var out int
	for _, v := range result {
		for _, val := range v.Values {
			out += int(val.Value)
		}
	}

	return out, nil
}
