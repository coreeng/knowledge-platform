package pushgateway

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
)

const PushGatewayUrl = "http://pushgateway.acceptance-criteria-pushgateway:9091"

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
	prometheus.MustRegister(stepOutcomeMetric)
}

func PushSuccessMetric(stepName string, participant string) {
	pusher := push.New(PushGatewayUrl, "acceptance-criteria-p2p-fast-feedback-step-outcome")
	stepOutcomeMetric.WithLabelValues("1", stepName, participant).Set(1)
	pusher.Collector(stepOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to push gateway, err: %v", err)
	}
}

func PushFailureMetric(stepName string, participant string) {
	pusher := push.New(PushGatewayUrl, "acceptance-criteria-p2p-fast-feedback-step-outcome")
	stepOutcomeMetric.WithLabelValues("0", stepName, participant).Set(0)
	pusher.Collector(stepOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to push gateway, err: %v", err)
	}
}

func PushTestSuiteOutcomeMetric(module string, participant string) {
	pusher := push.New(PushGatewayUrl, "acceptance-criteria-p2p-fast-feedback-satisfied")
	acceptanceCriteriaSuccessMetric.WithLabelValues(module, participant).Set(1)
	pusher.Collector(acceptanceCriteriaSuccessMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Error("Error while pushing acceptance criteria met to push gateway, err: %v", err)
	}
}
