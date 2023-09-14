Feature: Multi tenant kubernetes Acceptance Criteria

  Background:
    Given the kubernetes client is setup

#  Scenario: HNC Controller namespace exists
#    When I get the "hnc-system" namespace
#    Then the namespace "hnc-system" is returned
#
#  Scenario: HNC controller installed
#    When I get the pods in the "hnc-system" namespace
#    Then there is a running "hnc-controller-manager" pod in the namespace "hnc-system"
#
#  Scenario Outline: Hierarchical namespaces exist and are created correctly
#    When I get the "<parent-namespace>" namespace
#    Then the namespace "<parent-namespace>" is returned
#    And the "<parent-namespace>" namespace has the subnamespaces "<subnamespaces>"
#    Examples:
#      | parent-namespace | subnamespaces                 |
#      | cecg             | team-a,team-b                 |
#      | team-a           | app-1,app-2,team-a-monitoring |
#      | team-b           | app-3,team-b-monitoring       |

  Scenario Outline: The access to cluster namespaces is isolated per tenant - RBAC rules are applied correctly
    Given I am a tenant called "<tenant-name>"
    And a roleBinding exists in all my namespaces: "<my-namespaces>"
    And the roleBinding is associated with a serviceAccount
    When I impersonate the service account
    Then I can access all my namespaces: "<my-namespaces>"
    And I cannot access other tenant's namespaces: "<other-namespaces>"
    Examples:
      | tenant-name | my-namespaces                        | other-namespaces                     |
      | team-a      | team-a,app-1,app-2,team-a-monitoring | team-b,app-3,team-b-monitoring       |
      | team-b      | team-b,app-3,team-b-monitoring       | team-a,app-1,app-2,team-a-monitoring |
