package structure

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

func main() {
	// Read the YAML file
	yamlData, err := ioutil.ReadFile("/Users/nzi02/tempdir/cecg-bootcamp/structure/content_structure.yaml")
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
