Feature: NFT and Observability Acceptance Criteria

  Scenario: Monitoring stack setup
    Given the kubernetes client is set up
    And the bootcamp participant is defined
    And i fetch pods for the namespace
    Then there is a running prometheus
    And  there is a running alertmanager
    And there is a running grafana
    And the k6 test result dashboard is present in grafana

  Scenario: Deployed load testing
    Given the kubernetes client is set up
    When i get the k6-operator-system namespace
    Then the namespace is returned
    When i get pods in the k6-operator-system namespace
    Then there is a running k6 operator pod

  Scenario: Deployed load testing metrics in monitoring stack
    When i fetch prometheus metrics
    Then the following k6 metrics are present:
      | k6_http_req_duration_p95        |
      | k6_http_reqs                    |
      | k6_vus                          |
      | k6_data_received                |
      | k6_data_sent                    |
      | k6_http_req_blocked_avg         |
      | k6_http_req_waiting_avg         |
      | k6_http_req_sending_avg         |
      | k6_http_req_receiving_avg       |
      | k6_http_req_tls_handshaking_avg |

  Scenario: Resource quota are set
    Given the kubernetes client is set up
    When i fetch resource quota for the reference-service application
    Then memory and CPU is set

  Scenario: Support 500VUs with P95 being 50 ms or less
    When i fetch the k6 test result metric
    Then the k6 test result metric is present
    And the last test run was successful
    And max VUs were 500 during the last test run
    And the response time 95th percentile was 50 ms or less for the last test run

  Scenario: Support 1500VUs with p95 being 50 ms or less
    When i fetch the k6 test result metric
    Then the k6 test result metric is present
    And the last test run was successful
    And max VUs were 1500 during the last test run
    And the response time 95th percentile was 50 ms or less for the last test run

  Scenario: Fail gracefully when downstream dependencies are slow
    Given an HTTP client is set up
    When i call the downstream dependency with a second of response delay
    Then i get a response back within 600ms
    And the status code is 503