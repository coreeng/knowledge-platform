Feature: Canary Operator Acceptance Criteria

  Background:
    Given the kubernetes client is setup

  Scenario: Canary operator namespace exists
    When I get the "canary-operator-system" namespace
    Then the namespace "canary-operator-system" exists

  Scenario: Canary operator controller manager is running
    When I get the pods in the "canary-operator-system" namespace
    Then there is a running controller manager pod

  Scenario: A custom resource definition for "CanariedApp" exists
    Given the kubernetes api extension client exists
    When I get the custom resource definition "canariedapps.canary.cecg.com"
    Then the custom resource definition "canariedapps.canary.cecg.com" exists

  Scenario: When I update the version of my app deployed using the CR "CanariedApp", a canary deployment is created
    Given I have a namespace "canary-operator-autograding"
    And I have the following CR:
      | Kind  | CanariedApp             |
      | name  | canariedapp-autograding |
      | image | cecg/minimal-ref-app:v1 |
    When I update the CR with:
      | Kind  | CanariedApp             |
      | name  | canariedapp-autograding |
      | image | cecg/minimal-ref-app:v2 |
    Then the canary deployment "canariedapp-autograding-canary" is created

#  @wip
#  Scenario: When the new canary deployment is stable, traffic through ingress goes through both deployments
#    Given I have a namespace "canary-operator-autograding"
#    And I have the following CR:
#      | Kind  | CanariedApp             |
#      | name  | canariedapp-autograding |
#      | image | cecg/minimal-ref-app:v1 |
#    And I update the CR with:
#      | Kind  | CanariedApp             |
#      | name  | canariedapp-autograding |
#      | image | cecg/minimal-ref-app:v2 |
#    When my smoke tests via ingress pass
#    Then traffic goes through both canary and non-canary deployments
#
#  @wip
#  Scenario: When the new canary deployment is not stable, no traffic is going through the canary deployment
#    Given I have a namespace "canary-operator-autograding"
#    And I have the following CR:
#      | Kind  | CanariedApp             |
#      | name  | canariedapp-autograding |
#      | image | cecg/minimal-ref-app:v1 |
#    And I update the CR with:
#      | Kind  | CanariedApp             |
#      | name  | canariedapp-autograding |
#      | image | cecg/minimal-ref-app:v2 |
#    When my smoke tests produce a "CanaryHealthCheck" alert
#    Then traffic goes through both canary and non-canary deployments
