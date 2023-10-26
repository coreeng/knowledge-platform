package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

var promURL = "http://kube-prometheus-stack-prometheus.platform-engineering-autograding.svc.cluster.local"

func AlertFiring(ctx context.Context, alertname string) (bool, error) {
	query := fmt.Sprintf("ALERTS{handler='/canaryTest',alertname='%s',alertstate='firing'}", alertname)
	resp, err := getMetrics(ctx, query)
	if err != nil {
		return false, fmt.Errorf("getting ingress metric: %w", err)
	}

	// Not firing
	if resp == "" {
		return false, nil
	}

	alertFiring, err := strconv.Atoi(resp)
	if err != nil {
		return false, fmt.Errorf("convert string to integer: %w", err)
	}

	if alertFiring == 0 {
		return false, nil
	}

	return true, nil
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
