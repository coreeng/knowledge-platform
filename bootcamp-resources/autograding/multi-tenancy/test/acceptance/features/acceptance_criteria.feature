Feature: Multi tenant kubernetes Acceptance Criteria

  Background:
    Given the kubernetes client is setup

  Scenario: The HNC Controller namespace should exist
    When I get the "hnc-system" namespace
    Then the namespace "hnc-system" is returned

  Scenario: The HNC Controller should be running in the cluster
    When I get the pods in the "hnc-system" namespace
    Then there is a running "hnc-controller-manager" pod in the namespace "hnc-system"

  Scenario Outline: The hierarchical namespaces should exist and they should be setup correctly
    When I get the "<parent-namespace>" namespace
    Then the namespace "<parent-namespace>" is returned
    And the "<parent-namespace>" namespace has the subnamespaces "<subnamespaces>"
    Examples:
      | parent-namespace | subnamespaces                 |
      | cecg             | team-a,team-b                 |
      | team-a           | app-1,app-2,team-a-monitoring |
      | team-b           | app-3,team-b-monitoring       |

  Scenario Outline: Each tenant's namespaces should be isolated from each other, one tenant should not be able to affect other tenant's resources
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

  Scenario Outline: The default deny is applied for all inbound traffic. Inter namespace communication is not allowed
    Given I have a destination pod "destination-pod" in the namespace "<dest-namespace>"
    And the "destination-pod" has a service called "destination-service" in the namespace "<dest-namespace>"
    And I have a source pod "source-pod" in the namespace "<source-namespace>"
    When I try to connect from "source-pod" to "destination-pod"
    Then the access is denied
    Examples:
      | source-namespace | dest-namespace    |
      | team-a           | app-1             |
      | team-a           | app-2             |
      | team-a           | team-a-monitoring |
      | team-b           | app-3             |
      | team-b           | team-b-monitoring |

  Scenario Outline: The default deny is applied for all inbound traffic. Intra namespace communication is allowed
    Given I have a destination pod "destination-pod" in the namespace "<namespace>"
    And the "destination-pod" has a service called "destination-service" in the namespace "<namespace>"
    And I have a source pod "source-pod" in the namespace "<namespace>"
    When I try to connect from "source-pod" to "destination-pod"
    Then the access is allowed
    Examples:
      | namespace         |
      | team-a            |
#      | app-1             |
#      | app-2             |
#      | team-a-monitoring |
#      | team-b            |
#      | app-3             |
#      | team-b-monitoring |
