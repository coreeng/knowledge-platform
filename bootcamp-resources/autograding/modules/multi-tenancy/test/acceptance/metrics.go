package acceptance

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/sirupsen/logrus"
)

const PushGatewayUrl = "http://pushgateway.pushgateway-autograding:9091"

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
	pusher := push.New(PushGatewayUrl, "acceptance-criteria-multi-tenancy-scenario-outcome")
	scenarioOutcomeMetric.WithLabelValues(scenarioName, scenarioOutcomeLabel).Set(metricValue)
	pusher.Collector(scenarioOutcomeMetric)

	err := pusher.Push()
	if err != nil {
		logrus.Errorf("Error while pushing metrics to pushgateway, err: %v", err)
	}
}
