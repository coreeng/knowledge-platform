package main

import (
	_ "fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"path/filepath"
	"testing"
)

func TestWeightGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WeightGenerator Suite")
}

var _ = Describe("WeightGenerator", func() {
	var (
		yamlFilePath string
		tempDir      string
	)

	BeforeSuite(func() {
		// Create a temporary directory
		tempDir, err := os.MkdirTemp("", "weight_generator_test")
		Expect(err).ToNot(HaveOccurred())

		// Create a temporary YAML file
		yamlFilePath = filepath.Join(tempDir, "content_structure.yaml")
		err = os.WriteFile(yamlFilePath, []byte(`
path: cecg-bootcamp/structure/content/_index.md
pages:
  - path: cecg-bootcamp/structure/content/bootcamp/_index.md
    pages:
      - path: cecg-bootcamp/structure/content/bootcamp/modules/_index.md
        pages:
          - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/_index.md
            pages:
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/background.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-core-platform.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-iac-setup.md
`), os.ModePerm)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterSuite(func() {
		// Clean up temporary directory
		err := os.RemoveAll(tempDir)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("assignWeights", func() {
		Context("assign weights to pages - Depth 1", func() {
			It("should assign weights to pages", func() {
				page := &Page{
					Path: "/path/to/page",
					Pages: []*Page{
						{Path: "/path/to/page1"},
						{Path: "/path/to/page2"},
					},
				}

				assignWeights(page, 0, true)

				Expect(page.Weight).To(Equal(1))
				Expect(page.Pages[0].Weight).To(Equal(1))
				Expect(page.Pages[1].Weight).To(Equal(2))
			})
		})
		Context("assign weights to pages - Depth 3", func() {
			It("should assign weights to pages", func() {
				page := &Page{
					Path: "/path/to/page",
					Pages: []*Page{
						{
							Path: "/path/to/page/page1",
							Pages: []*Page{
								{Path: "/path/to/page/page1/page1"},
								{Path: "/path/to/page/page1/page2"},
								{Path: "/path/to/page/page1/page3",
									Pages: []*Page{
										{
											Path: "/path/to/page2-1",
											Pages: []*Page{
												{Path: "/path/to/page2-1-1"},
												{Path: "/path/to/page2-1-2"},
											},
										},
									},
								},
							},
						},
					},
				}

				assignWeights(page, 0, true)

				Expect(page.Weight).To(Equal(1))
				Expect(page.Pages[0].Weight).To(Equal(1))
				Expect(page.Pages[0].Pages[0].Weight).To(Equal(1))
				Expect(page.Pages[0].Pages[1].Weight).To(Equal(2))
				Expect(page.Pages[0].Pages[2].Weight).To(Equal(3))
				Expect(page.Pages[0].Pages[2].Pages[0].Weight).To(Equal(1))
				Expect(page.Pages[0].Pages[2].Pages[0].Pages[0].Weight).To(Equal(1))
				Expect(page.Pages[0].Pages[2].Pages[0].Pages[1].Weight).To(Equal(2))
			})
		})
	})

	Describe("replacePlaceholder", func() {
		Context("when the placeholder exists", func() {
			It("should replace the placeholder with the desired structure", func() {
				yamlData := []byte(`
path: cecg-bootcamp/structure/content/_index.md
pages:
  - path: cecg-bootcamp/structure/content/bootcamp/_index.md
    pages:
      - path: cecg-bootcamp/structure/content/bootcamp/modules/_index.md
        pages:
            cecg_bootcamp_module
          - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/_index.md
            pages:
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/background.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-core-platform.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-iac-setup.md
`)

				modifiedYAMLData, err := replacePlaceholder(yamlData)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(modifiedYAMLData)).To(ContainSubstring("cecg-bootcamp/structure/content/bootcamp/_index.md"))
				Expect(string(modifiedYAMLData)).To(ContainSubstring("path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/background.md"))
				Expect(string(modifiedYAMLData)).NotTo(ContainSubstring("cecg_bootcamp_module"))
			})
		})

		Context("when the placeholder does not exist", func() {
			It("should return an error", func() {
				yamlData := []byte(`
path: cecg-bootcamp/structure/content/_index.md
pages:
  - path: cecg-bootcamp/structure/content/bootcamp/_index.md
    pages:
      - path: cecg-bootcamp/structure/content/bootcamp/modules/_index.md
        pages:
          - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/_index.md
            pages:
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/background.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-core-platform.md
              - path: cecg-bootcamp/structure/content/bootcamp/modules/cloud-iac/epic-iac-setup.md
`)

				_, err := replacePlaceholder(yamlData)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("cecg_bootcamp_module must exist in the YAML data"))
			})
		})
	})

	Describe("getLeadingWhitespace", func() {
		It("should extract the leading whitespace from a line until the position of the placeholder", func() {
			line := "    cecg_bootcamp_module"
			placeholder := "cecg_bootcamp_module"
			leadingWhitespace := getLeadingWhitespace(line, placeholder)

			Expect(leadingWhitespace).To(Equal("    "))
		})
	})
})
