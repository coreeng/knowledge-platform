package nft_and_obs_metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"nft-and-observability/consts"
	"nft-and-observability/structs"
	"strings"
)

const K6TestDataMetric = "k6_test_data"

var stepOutcomeMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "acceptance_criteria_step",
		Help: "The acceptance criteria step that either passed or failed",
	},
	[]string{"is_success", "stepName", "participant"})

var acceptanceCriteriaSuccessMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "p2p_fast_feedback_acceptance_criteria_success",
		Help: "The acceptance criteria for the p2p-fast-feedback module are met",
	}, []string{"module", "participant"},
)

func Init() {
	prometheus.MustRegister(stepOutcomeMetric, acceptanceCriteriaSuccessMetric)
}

func FetchK6TestResultMetric() (error, *structs.K6TestResultMetric) {
	var k6TestDataMetricLine string
	var metric structs.K6TestResultMetric

	resp, err := http.Get(consts.PushGatewayUrl + "/metrics")
	if err != nil {
		return fmt.Errorf("error occurred while fetching metrics from push gateway. Err: %v", err), nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error occurred while reading push gateway metrics response. Err: %v", err), nil
	}

	metricsData := string(body)

	for _, line := range strings.Split(metricsData, "\n") {
		if strings.HasPrefix(line, K6TestDataMetric) {
			k6TestDataMetricLine = line
			break
		}
	}

	if k6TestDataMetricLine != "" {
		parts := strings.Split(k6TestDataMetricLine, " ")
		labelsAndValues := parts[0]
		labelValuePairs := strings.Split(labelsAndValues, ",")
		for _, pair := range labelValuePairs {
			labelValue := strings.Split(pair, "=")
			sanitisedLabelValue := sanitiseLabelValue(labelValue[1])
			if len(labelValue) == 2 {
				switch labelValue[0] {
				case "status":
					metric.Status = sanitisedLabelValue
				case "p95_response_time":
					metric.P95ResponseTime = sanitisedLabelValue
				case "vus_max":
					metric.MaxVUS = sanitisedLabelValue
				case "timestamp":
					metric.Timestamp = sanitisedLabelValue
				default:
				}
			}
		}
	}

	return nil, &metric
}

func sanitiseLabelValue(labelValue string) string {
	removedRightCurly := strings.ReplaceAll(labelValue, "}", "")
	removedLeftCurly := strings.ReplaceAll(removedRightCurly, "{", "")
	trimmedValue := strings.Trim(removedLeftCurly, `"`)
	return trimmedValue
}

func PushSuccessMetric(stepName string, participant string) {
	pusher := push.New(consts.PushGatewayUrl, "acceptance-criteria-nft-and-obs-step-outcome")
	stepOutcomeMetric.WithLabelValues("1", stepName, participant).Set(1)
	pusher.Collector(stepOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to push gateway, err: %v", err)
	}
}

func PushFailureMetric(stepName string, participant string) {
	pusher := push.New(consts.PushGatewayUrl, "acceptance-criteria-nft-and-obs-step-outcome")
	stepOutcomeMetric.WithLabelValues("0", stepName, participant).Set(0)
	pusher.Collector(stepOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to push gateway, err: %v", err)
	}
}

func PushTestSuiteOutcomeMetric(module string, participant string) {
	pusher := push.New(consts.PushGatewayUrl, "acceptance-criteria-nft-and-obs-satisfied")
	acceptanceCriteriaSuccessMetric.WithLabelValues(module, participant).Set(1)
	pusher.Collector(acceptanceCriteriaSuccessMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Error("Error while pushing acceptance criteria met to push gateway, err: %v", err)
	}
}
