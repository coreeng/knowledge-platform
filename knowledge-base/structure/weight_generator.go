package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	_ "path/filepath"
	"regexp"
	"strings"
)

type Page struct {
	Path   string  `yaml:"path"`
	Pages  []*Page `yaml:"pages"`
	Weight int     `yaml:"weight,omitempty"`
}

func main() {

	// Read the YAML file
	yamlFilePath := "content_structure.yaml"
	//Read the structure YAML file
	yamlData, err := readYamlFile(yamlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Replace the placeholder with the desired structure
	yamlData, err = replacePlaceholder(yamlData)
	if err != nil {
		log.Fatal(err)
	}

	// Save the modified YAML data back to the file
	err = writeYamlFile(yamlFilePath, yamlData)
	if err != nil {
		log.Fatal(err)
	}

	var pages []*Page
	// Unmarshal the YAML data into the page structure
	err = yaml.Unmarshal(yamlData, &pages)
	if err != nil {
		log.Fatal(err)
	}

	// Assign weights to the pages
	assignWeights(pages[0], 0, true)

	// Print the pages with their assigned weights
	printPages(pages, "")

	// Edit the markdown files
	err = editMarkdownFiles(pages[0])
	if err != nil {
		log.Fatal(err)
	}
}

func printPages(pages []*Page, indent string) {
	for _, page := range pages {
		fmt.Printf("%s%s -> weight: %d\n", indent, page.Path, page.Weight)
		if len(page.Pages) > 0 {
			printPages(page.Pages, indent+"  ")
		}
	}
}

func editMarkdownFiles(page *Page) error {
	if page == nil {
		return nil
	}

	err := editMarkdownFile(page.Path, page.Weight)
	if err != nil {
		return err
	}

	for _, child := range page.Pages {
		err := editMarkdownFiles(child)
		if err != nil {
			return err
		}
	}

	return nil
}

func editMarkdownFile(filePath string, weight int) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Find the weight in the front matter
	weightPattern := regexp.MustCompile(`(?ms)(^(\s*\+\+\+)\s*.*?^\s*weight\s*=\s*)(?P<Weight>\d+)(.*?^\s*\+\+\+\s*)`)
	match := weightPattern.FindStringSubmatch(string(content))

	if len(match) > 0 {
		// Weight found, replace the existing weight value with the new weight value
		newContent := weightPattern.ReplaceAllString(string(content), fmt.Sprintf("${1}%d${4}", weight))
		err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		// Weight not found, append it after the date
		titlePattern := regexp.MustCompile(`(?ms)(^(\s*\+\+\+)\s*.*?^\s*title\s*=\s*".*?".*?^\s*\+\+\+\s*)`)
		match := titlePattern.FindStringSubmatch(string(content))

		if len(match) > 0 {
			newContent := strings.Replace(string(content), match[0], fmt.Sprintf("%s\nweight = %d\n%s", match[0], weight, match[0]), 1)
			err = os.WriteFile(filePath, []byte(newContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			// Date not found, unable to append weight
			return fmt.Errorf("unable to find date in file %s", filePath)
		}
	}

	return nil
}

func readYamlFile(yamlFilePath string) ([]byte, error) {
	yamlData, err := os.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("%s not found. Must be in the same directory with the weight_generator.go", yamlFilePath)
	}

	return yamlData, nil
}

func writeYamlFile(yamlFilePath string, yamlData []byte) error {

	// Save the modified YAML data back to the file
	err := os.WriteFile(yamlFilePath, yamlData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot write yamlData to %s: %w", yamlFilePath, err)
	}
	return nil

}
func replacePlaceholder(yamlData []byte) ([]byte, error) {
	placeholder := "cecg_bootcamp_module"
	// Check if the placeholder exists in the YAML data
	if !strings.Contains(string(yamlData), placeholder) {
		return nil, fmt.Errorf("%s must exist in the YAML data", placeholder)
	}
	// Define the desired structure to replace the placeholder
	desiredStructure := `
- path: ../content/_index.md
  pages:
    - path: ../content/bootcamp/_index.md
      pages:
        - path: ../content/bootcamp/ways-of-working/_index.md
          pages:
            - path: ../content/bootcamp/ways-of-working/roles.md
            - path: ../content/bootcamp/ways-of-working/rituals-artfifacts.md
            - path: ../content/bootcamp/ways-of-working/practices.md
        - path: ../content/bootcamp/setup-and-access/_index.md
          pages:
            - path: ../content/bootcamp/setup-and-access/workstation-setup.md
            - path: ../content/bootcamp/setup-and-access/epic-setup.md
            - path: ../content/bootcamp/setup-and-access/epic-katas.md
            - path: ../content/bootcamp/setup-and-access/epic-tekton.md
        - path: ../content/bootcamp/modules/_index.md
          pages:
            - path: ../content/bootcamp/modules/p2p-fast-feedback/_index.md
              pages:
                - path: ../content/bootcamp/modules/p2p-fast-feedback/background.md
                - path: ../content/bootcamp/modules/p2p-fast-feedback/epic-p2p-fastfeedback-learning-plan.md
                - path: ../content/bootcamp/modules/p2p-fast-feedback/epic-reference-application.md
                - path: ../content/bootcamp/modules/p2p-fast-feedback/epic-p2p.md
                - path: ../content/bootcamp/modules/p2p-fast-feedback/epic-counter-resets.md
            - path: ../content/bootcamp/modules/nft-and-observability/_index.md
              pages:
                - path: ../content/bootcamp/modules/nft-and-observability/background.md
                - path: ../content/bootcamp/modules/nft-and-observability/epic-p2p-nft-learning-plan.md
                - path: ../content/bootcamp/modules/nft-and-observability/epic-monitoring-stack.md
                - path: ../content/bootcamp/modules/nft-and-observability/epic-deployed-testing.md
                - path: ../content/bootcamp/modules/nft-and-observability/epic-resiliency.md
            - path: ../content/bootcamp/modules/multitenant-kubernetes/_index.md
              pages:
                - path: ../content/bootcamp/modules/multitenant-kubernetes/background.md
                - path: ../content/bootcamp/modules/multitenant-kubernetes/epic-kubernetes-learning.md
                - path: ../content/bootcamp/modules/multitenant-kubernetes/epic-canary-deployments.md
                - path: ../content/bootcamp/modules/multitenant-kubernetes/epic-multi-tenancy-rbac.md
                - path: ../content/bootcamp/modules/multitenant-kubernetes/epic-multi-tenancy-network-isolation.md
            - path: ../content/bootcamp/modules/platform-engineering/_index.md
              pages:
                - path: ../content/bootcamp/modules/platform-engineering/epic-platform-engineering-learning-plan.md
                - path: ../content/bootcamp/modules/platform-engineering/epic-canary-operator.md
            - path:  ../content/bootcamp/modules/debugging-containers/_index.md
              pages:
                - path: ../content/bootcamp/modules/debugging-containers/background.md
                - path: ../content/bootcamp/modules/debugging-containers/epic-troubleshooting.md
            - path: ../content/bootcamp/modules/cloud-iac/_index.md
              pages:
                - path: ../content/bootcamp/modules/cloud-iac/background.md
                - path: ../content/bootcamp/modules/cloud-iac/epic-iac-learning-plan.md
                - path: ../content/bootcamp/modules/cloud-iac/epic-iac-setup.md
                - path: ../content/bootcamp/modules/cloud-iac/epic-core-platform.md
    - path: ../content/core-platform/_index.md
      pages:
        - path: ../content/core-platform/building-a-core-platform.md
        - path: ../content/core-platform/faqs.md
        - path: ../content/core-platform/features/_index.md
          pages:
              - path: ../content/core-platform/features/sprint-0/_index.md
                pages:
                 - path: ../content/core-platform/features/sprint-0/feature-version-control-access-control.md
                 - path: ../content/core-platform/features/sprint-0/feature-cloud-accounts.md
                 - path: ../content/core-platform/features/sprint-0/feature-adr-process.md
                 - path: ../content/core-platform/features/sprint-0/feature-base-deployment-pipeline.md
                 - path: ../content/core-platform/features/sprint-0/feature-platform-teams-ways-of-working.md
                 - path: ../content/core-platform/features/sprint-0/feature-platform-testing-strategy.md
                 - path: ../content/core-platform/features/sprint-0/feature-platform-ci-developer-infra-setup.md
                 - path: ../content/core-platform/features/sprint-0/feature-golang-dev-environment.md
              - path: ../content/core-platform/features/platform-path-to-prod/_index.md
                pages:
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-core-platform-environments.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-monolithic-deployment.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-decoupled-platform-deployment.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-basic-promotion.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-base-infrastructure.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-platform-services-provisioning.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-provisioning-single-plane-of-glass.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-cluster-config-management.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-platform-e2e-testing.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-continuous-e2e-testing.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-regular-full-rebuild.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-segregated-sandbox.md
                 - path: ../content/core-platform/features/platform-path-to-prod/feature-automatic-promotion.md
              - path: ../content/core-platform/features/connected-kubernetes/_index.md
                pages:
                  - path: ../content/core-platform/features/connected-kubernetes/feature-basic-cluster.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-sso-integration.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-base-networking-design.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-node-pools.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-cloud-provider-registries.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-network-connectivity.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-block-storage.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-basic-ingress.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-egressless-bootstrap.md
                  - path: ../content/core-platform/features/connected-kubernetes/feature-autoscaling.md
              - path: ../content/core-platform/features/multi-tenant-access/_index.md
                pages:
                  - path: ../content/core-platform/features/multi-tenant-access/feature-tenant-kubernetes-access.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-cloud-based-groups.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-corporate-ad-based-rbac.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-registries.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-cluster-resource-creation.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-production-access-model.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-quality-of-service.md
                  - path: ../content/core-platform/features/multi-tenant-access/feature-rbac-audit.md
              - path: ../content/core-platform/features/kubernetes-network/_index.md
                pages:
                  - path: ../content/core-platform/features/kubernetes-network/feature-cni.md
                  - path: ../content/core-platform/features/kubernetes-network/feature-service-mesh.md
                  - path: ../content/core-platform/features/kubernetes-network/_index.md
                  - path: ../content/core-platform/features/kubernetes-network/feature-default-deny.md
                  - path: ../content/core-platform/features/kubernetes-network/feature-ipv6.md
              - path: ../content/core-platform/features/ingress/_index.md
                pages:
                  - path: ../content/core-platform/features/ingress/feature-platform-ingress.md
                  - path: ../content/core-platform/features/ingress/feature-cross-region-loadbalancing.md
                  - path: ../content/core-platform/features/ingress/feature-edge-integration.md
                  - path: ../content/core-platform/features/ingress/feature-segregated-ingress.md
                  - path: ../content/core-platform/features/ingress/feature-corporate-dns-reliability.md
                  - path: ../content/core-platform/features/ingress/feature-cluster-dns.md
                  - path: ../content/core-platform/features/ingress/feature-tenant-dns.md
                  - path: ../content/core-platform/features/ingress/feature-http-ingress.md
                  - path: ../content/core-platform/features/ingress/feature-tcp-ingress.md
                  - path: ../content/core-platform/features/ingress/feature-udp-ingress.md
                  - path: ../content/core-platform/features/ingress/feature-ingress-scalability-testing.md
              - path: ../content/core-platform/features/connectivity/_index.md
                pages:
                  - path: ../content/core-platform/features/connectivity/feature-per-tenant-egress-firewall.md
                  - path: ../content/core-platform/features/connectivity/feature-same-cloud-provider-connectivity.md
                  - path: ../content/core-platform/features/connectivity/feature-alternative-cloud-provider-connectivity.md
                  - path: ../content/core-platform/features/connectivity/feature-private-third-party-connectivity.md
                  - path: ../content/core-platform/features/connectivity/feature-advanced-source-nating.md
                  - path: ../content/core-platform/features/connectivity/feature-inbound-workstation-access.md
              - path: ../content/core-platform/features/platform-observability/_index.md
                pages:
                  - path: ../content/core-platform/features/platform-observability/feature-aggregated-platform-logs.md
                  - path: ../content/core-platform/features/platform-observability/feature-aggregated-platform-metrics.md
                  - path: ../content/core-platform/features/platform-observability/feature-platform-service-monitoring.md
                  - path: ../content/core-platform/features/platform-observability/feature-platform-single-pane-of-glass.md
                  - path: ../content/core-platform/features/platform-observability/feature-incident-management-integration.md
              - path: ../content/core-platform/features/governance/_index.md
                pages:
                  - path: ../content/core-platform/features/governance/feature-policy-controller.md
                  - path: ../content/core-platform/features/governance/feature-security-policy-enforcement.md
                  - path: ../content/core-platform/features/governance/feature-kubernetes-best-practices.md
                  - path: ../content/core-platform/features/governance/feature-per-tenant-policy.md
                  - path: ../content/core-platform/features/governance/feature-retrospective-policy-validation.md
              - path: ../content/core-platform/features/tenant-observability/_index.md
                pages:
                  - path: ../content/core-platform/features/tenant-observability/feature-export-aggregated-metrics.md
                  - path: ../content/core-platform/features/tenant-observability/feature-export-aggregated-logs.md
                  - path: ../content/core-platform/features/tenant-observability/feature-external-dashboarding-and-incident-management-integration.md
                  - path: ../content/core-platform/features/tenant-observability/feature-aggregated-metrics-store.md
                  - path: ../content/core-platform/features/tenant-observability/feature-aggregated-log-store.md
                  - path: ../content/core-platform/features/tenant-observability/feature-dashboarding.md
                  - path: ../content/core-platform/features/tenant-observability/feature-tracing.md
                  - path: ../content/core-platform/features/tenant-observability/feature-incident-management-integration.md
                  - path: ../content/core-platform/features/tenant-observability/feature-service-interaction.md
              - path: ../content/core-platform/features/platform-security/_index.md
                pages:
                  - path: ../content/core-platform/features/platform-security/feature-build-deployment-time-protection.md
                  - path: ../content/core-platform/features/platform-security/feature-runtime-endpoint-protection.md
                  - path: ../content/core-platform/features/platform-security/feature-runtime-misconfiguration.md
                  - path: ../content/core-platform/features/platform-security/feature-runtime-vulnerabilities.md
                  - path: ../content/core-platform/features/platform-security/feature-host-protection.md
                  - path: ../content/core-platform/features/platform-security/feature-network-segregation.md
                  - path: ../content/core-platform/features/platform-security/feature-workload-segregation.md
                  - path: ../content/core-platform/features/platform-security/feature-corporate-integrations.md
                  - path: ../content/core-platform/features/platform-security/feature-encryption-in-transit.md
                  - path: ../content/core-platform/features/platform-security/feature-encryption-at-rest.md
                  - path: ../content/core-platform/features/platform-security/feature-secure-kubernetes-setup.md
                  - path:
              - path: ../content/core-platform/features/providers-locations-and-dr/_index.md
                pages:
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-single-cloud-region-multi-az.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-regular-dr-rebuild-in-new-region.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-multiple-cloud-regions.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-multi-cloud-provider.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-infrastructure-autoscaling.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-on-prem.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-hybrid.md
                  - path: ../content/core-platform/features/providers-locations-and-dr/feature-multi-provider-registries.md
              - path: ../content/core-platform/features/secrets-management/_index.md
              - path: ../content/core-platform/features/continuous-load/_index.md
              - path: ../content/core-platform/features/developer-portal/_index.md
                pages:
                  - path: ../content/core-platform/features/developer-portal/feature-developer-portal.md
                  - path: ../content/core-platform/features/developer-portal/feature-service-catalog.md
                  - path: ../content/core-platform/features/developer-portal/feature-application-blueprints.md
              - path: ../content/core-platform/features/persistence/_index.md
                pages: 
                  - path: ../content/core-platform/features/persistence/feature-paired-cloud-account.md
                  - path: ../content/core-platform/features/persistence/feature-managed-service-blueprints.md
        - path: ../content/core-platform/kubernetes-upgrade.md
     - path: ../content/core-p2p/_index.md
      pages:
    - path: ../content/core-p2p/_index.md
      pages:
       - path: ../content/core-p2p/p2p-consolidation.md
       - path: ../content/core-p2p/design-a-p2p.md
       - path: ../content/core-p2p/service-encapsulation.md
       - path: ../content/core-p2p/testing-strategy/_index.md
         pages:
           - path: ../content/core-p2p/testing-strategy/unit-testing.md
           - path: ../content/core-p2p/testing-strategy/functional-testing.md
           - path: ../content/core-p2p/testing-strategy/non-functional-testing.md
       - path: ../content/core-p2p/promotion.md
       - path: ../content/core-p2p/application-architecture.md
       - path: ../content/core-p2p/deployment-models.md
       - path: ../content/core-p2p/features/_index.md
         pages:
           - path: ../content/core-p2p/features/reference-application/_index.md
             pages:
               - path: ../content/core-p2p/features/reference-application/feature-basic-reference-app.md
           - path: ../content/core-p2p/features/developer-advocacy/_index.md
             pages:
               - path: ../content/core-p2p/features/developer-advocacy/feature-seed-community.md
               - path: ../content/core-p2p/features/developer-advocacy/feature-chat-community.md
               - path: ../content/core-p2p/features/developer-advocacy/feature-wiki.md
           - path: ../content/core-p2p/features/pipeline-testing/_index.md
           - path: ../content/core-p2p/features/p2p-monitoring/_index.md
             pages:
               - path: ../content/core-p2p/features/p2p-monitoring/feature-tenant-logs-forwarding.md
               - path: ../content/core-p2p/features/p2p-monitoring/feature-aggregated-logs.md
               - path: ../content/core-p2p/features/p2p-monitoring/feature-observability-integrations.md
           - path: ../content/core-p2p/features/continous-integration/_index.md
             pages:
               - path: ../content/core-p2p/features/continous-integration/feature-tenant-onboarding.md
               - path: ../content/core-p2p/features/continous-integration/feature-tenant-definition.md
               - path: ../content/core-p2p/features/continous-integration/feature-pipeline-generation.md
           - path: ../content/core-p2p/features/pipeline-security/_index.md
             pages:
               - path: ../content/core-p2p/features/pipeline-security/feature-vulnerability-scanning.md
           - path: ../content/core-p2p/features/continuous-deployment/_index.md
             pages:
               - path: ../content/core-p2p/features/continuous-deployment/feature-change-management.md
               - path: ../content/core-p2p/features/continuous-deployment/feature-deployment-access.md
               - path: ../content/core-p2p/features/continuous-deployment/feature-manual-gates.md
               - path: ../content/core-p2p/features/continuous-deployment/feature-promotion-mechanism.md
               - path: ../content/core-p2p/features/continuous-deployment/feature-scheduled-releases.md
           - path: ../content/core-p2p/features/artifacts/_index.md
             pages:
               - path: ../content/core-p2p/features/artifacts/feature-build-agent.md
               - path: ../content/core-p2p/features/artifacts/feature-build-service.md
               - path: ../content/core-p2p/features/artifacts/feature-immutable-artifacts.md
               - path: ../content/core-p2p/features/artifacts/feature-tenant-provided-build-agent.md
               - path: ../content/core-p2p/features/artifacts/feature-versioning.md
           - path: ../content/core-p2p/features/deployment-models/_index.md
             pages:
               - path: ../content/core-p2p/features/deployment-models/feature-deployment-mechanism.md
           - path: ../content/core-p2p/features/database-lifecycle/_index.md
           - path: ../content/core-p2p/features/p2p-infra/_index.md
             pages:
               - path: ../content/core-p2p/features/p2p-infra/feature-managed-tool.md
               - path: ../content/core-p2p/features/p2p-infra/feature-artifact-registry.md
               - path: ../content/core-p2p/features/p2p-infra/feature-source-control.md
    - path: ../content/core-engineer/_index.md
      pages:
        - path: ../content/core-engineer/engineering-principles.md
        - path: ../content/core-engineer/adrs.md
        - path: ../content/core-engineer/techniques.md
    - path: ../content/delivery/_index.md
      pages:
        - path: ../content/delivery/delivery-principles.md`
	lines := strings.Split(string(yamlData), "\n")

	// Find the leading whitespace from the placeholder lines
	leadingWhitespaces := make([]string, 0)
	for _, line := range lines {
		if strings.Contains(line, placeholder) {
			leadingWhitespaces = append(leadingWhitespaces, getLeadingWhitespace(line, placeholder))
		}
	}

	// Replace the placeholder with the modified structure and prepend the leading whitespace to each line
	replacementLines := strings.Split(desiredStructure, "\n")
	for i, line := range replacementLines {
		replacementLines[i] = leadingWhitespaces[0] + line
	}

	// Insert the modified lines at the position of the placeholder lines
	for i, line := range lines {
		if strings.Contains(line, placeholder) {
			lines = append(lines[:i], append(replacementLines, lines[i+1:]...)...)
			break
		}
	}

	// Join the modified lines and convert back to bytes
	modifiedYAMLData := []byte(strings.Join(lines, "\n"))
	return modifiedYAMLData, nil
}

// Helper function to extract the leading whitespace from a line until the position of the placeholder
func getLeadingWhitespace(line, placeholder string) string {
	index := strings.Index(line, placeholder)
	if index == -1 {
		return ""
	}
	return line[:index]
}

func assignWeights(page *Page, prevSiblingWeight int, isParent bool) int {
	if page == nil {
		return prevSiblingWeight
	}

	weight := prevSiblingWeight

	if isParent {
		weight = prevSiblingWeight + 1
	} else if prevSiblingWeight > 0 {
		weight = prevSiblingWeight + 1
	}

	page.Weight = weight

	childWeight := 0
	for _, child := range page.Pages {
		childWeight = assignWeights(child, childWeight, true)
	}

	return weight
}
