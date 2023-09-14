Feature: Multi tenant kubernetes Acceptance Criteria

  Background:
    Given the kubernetes client is setup

  Scenario: HNC Controller namespace exists
    When I get the "hnc-system" namespace
    Then the namespace "hnc-system" is returned

  Scenario: HNC controller installed
    When I get the pods in the "hnc-system" namespace
    Then there is a running "hnc-controller-manager" pod in the namespace "hnc-system"

  Scenario Outline: Hierarchical namespaces exist and are created correctly
    When I get the "<parent-namespace>" namespace
    Then the namespace "<parent-namespace>" is returned
    And the "<parent-namespace>" namespace has the subnamespaces "<subnamespaces>"
    Examples:
      | parent-namespace | subnamespaces                 |
      | cecg             | team-a,team-b                 |
      | team-a           | app-1,app-2,team-a-monitoring |
      | team-b           | app-3,team-b-monitoring       |

  Scenario Outline: The access to cluster namespaces is isolated per tenant - RBAC rules are applied correctly
    Given a service account for "<team-name>" already exists
    When I impersonate the service account
    Then I can access the following namespaces: "<allowed_namespaces>"
    And I cannot access the following namespaces: "<not_allowed_namespaces>"
    Examples:
      | team-name | allowed_namespaces                   | not_allowed_namespaces               |
      | team-a    | team-a,app-1,app-2,team-a-monitoring | team-b,app-3,team-b-monitoring       |
      | team-b    | team-b,app-3,team-b-monitoring       | team-a,app-1,app-2,team-a-monitoring |
