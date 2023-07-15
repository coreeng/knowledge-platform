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
		err = ioutil.WriteFile(filePath, []byte(newContent), os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		// Weight not found, append it after the date
		titlePattern := regexp.MustCompile(`(?ms)(^(\s*\+\+\+)\s*.*?^\s*title\s*=\s*".*?".*?^\s*\+\+\+\s*)`)
		match := titlePattern.FindStringSubmatch(string(content))

		if len(match) > 0 {
			newContent := strings.Replace(string(content), match[0], fmt.Sprintf("%s\nweight = %d\n%s", match[0], weight, match[0]), 1)
			err = ioutil.WriteFile(filePath, []byte(newContent), os.ModePerm)
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
	desiredStructure :=
		`- path: bootcamp-content/content/_index.md
  pages:
    - path: bootcamp-content/content/bootcamp/_index.md
      pages:
        - path: bootcamp-content/content/bootcamp/modules/_index.md
          pages:
            - path: bootcamp-content/content/bootcamp/modules/cloud-iac/_index.md
              pages:
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/background.md
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-core-platform.md
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-iac-setup.md`
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