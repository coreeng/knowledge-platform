Feature: P2P Fast Feedback Acceptance Criteria

  Scenario: Reference service has a deployment and 1 service replica
    Given the kubernetes client is set up
    When I list deployments for the reference-service-showcase namespace
    Then a reference-service deployment is present
    And a single replica of the reference-service is running

  Scenario: Counter endpoints are available
    Given An HTTP Client is set up
    When I poll the DELETE counter endpoint
    Then I receive a successful response
