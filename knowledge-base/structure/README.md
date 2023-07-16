# Weight Generator

The Weight Generator is a Go program that generates weights for Hugo modules and pages based on a specified structure in a YAML file. The program operates in conjunction with Hugo, a popular static site generator.

## Usage

To use the Weight Generator, follow these steps:

1. Make sure the Weight Generator (`weight_generator.go`) and the structure YAML file (`content_structure.yaml`) are in the same directory as your Hugo project, at the same level as the `bootcamp-content` folder.

2. Modify the structure YAML file (`content_structure.yaml`) to define the desired structure of your Hugo modules and pages. The structure should follow the following syntax:

3. The placeholder `cecg_bootcamp_module` must exist in structure yaml file.
```yaml
- path: <string>
  pages:
    - path: <string>
      pages:
        - path: <string>
          pages:
            - path: <string>
              pages:
                - path: <string>
                - path: <string>
                - path: <string>
```

3. Run the Weight Generator by executing the `weight_generator.go` file. The program will read the structure YAML file, replace the placeholder (`cecg_bootcamp_module`) with the desired structure, generate weights for the pages, and update the Markdown files accordingly.

4. Add/replace weight for each markdown file specified in the `content_structure.yaml`

5. Check the console output to view the pages with their assigned weights.

## Examples

Here are some examples to illustrate the usage and functionality of the Weight Generator:


### Example 1: Structure YAML File with Placeholder Only

The structure YAML file (`content_structure.yaml`) contains only the placeholder (`cecg_bootcamp_module`):

```yaml
cecg_bootcamp_module
```

The Weight Generator will replace the placeholder with the desired structure and assign weights to the pages based on the structure. The resulting structure will be:

```yaml
- path: bootcamp-content/content/_index.md # -> weight: 1
  pages:
    - path: bootcamp-content/content/bootcamp/_index.md # -> weight: 1
      pages:
        - path: bootcamp-content/content/bootcamp/modules/_index.md # -> weight: 1
          pages:
            - path: bootcamp-content/content/bootcamp/modules/cloud-iac/_index.md # -> weight: 1
              pages:
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/background.md # -> weight: 1
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-core-platform.md # -> weight: 2
                - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-iac-setup.md # -> weight: 3
```

### Example 2: Structure YAML File with Content and Placeholder

The structure YAML file (`content_structure.yaml`) contains the desired structure along with the placeholder:

```yaml
- path: <string>
  pages:
    - path: <string>
      pages:
        - path: <string>
          pages:
              cecg_bootcamp_module
                - path: <string>
                pages:
                  - path: <string>
                  - path: <string>
                  - path: <string>
```

The Weight Generator will replace the placeholder with the cecg bootcamp module and  will assign weights to the pages based on the provided structure and update the Markdown files accordingly.

The result will be : 

```yaml
- path: <string> # -> weight: 1
  pages:
    - path: <string> # -> weight: 1
      pages:
        - path: <string> # -> weight: 1
          pages:
            - path: bootcamp-content/content/_index.md # -> weight: 1
              pages:
                - path: bootcamp-content/content/bootcamp/_index.md # -> weight: 1
                  pages:
                    - path: bootcamp-content/content/bootcamp/modules/_index.md # -> weight: 1
                      pages:
                        - path: bootcamp-content/content/bootcamp/modules/cloud-iac/_index.md # -> weight: 1
                          pages:
                            - path: bootcamp-content/content/bootcamp/modules/cloud-iac/background.md # -> weight: 1
                            - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-core-platform.md # -> weight: 2
                            - path: bootcamp-content/content/bootcamp/modules/cloud-iac/epic-iac-setup.md # -> weight: 3
                            - path: <string> # -> weight: 4
                              pages:
                                - path: <string> # -> weight: 1
                                - path: <string> # -> weight: 2
                                - path: <string> # -> weight: 3
```
Please note that the examples assume the YAML file is correctly formatted and aligned.

## Error Handling

The Weight Generator includes basic error handling. If any error occurs during the execution, an error message will be logged, and the program will terminate. It is recommended to check the console output for any error messages and address them accordingly.

Please ensure that the Weight Generator and the structure YAML file are located in the correct directory and that the required dependencies (such as the `gopkg.in/yaml.v2` package) are properly installed.

Feel free to modify the Weight Generator code and the structure YAML file according to your specific needs.