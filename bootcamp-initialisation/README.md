# Bootcamp Initialization Tool

The Bootcamp Initialization Tool is a command-line utility written in Java that simplifies the process
of setting up a GitHub repository for bootcamp participants. It automates the creation of GitHub issues based
on predefined modules and their associated content.

## Building the Binary

To build the Bootcamp Initialization Tool binary, follow these steps:

Navigate to the project directory:

`cd bootcamp-initialization` 


Run the build script:

`./buildCliTool.sh`

The build script performs the following actions:

1) Cleans the project and builds it using Gradle.
2) Generates the bootcamp-initialize executable, which serves as a convenient means of executing the tool.
3) Gives execute permissions to the `bootcamp-initialize` executable.

## Running the Tool

To run the Bootcamp Initialization Tool, use the bootcamp-initialize executable created by the `buildCliTool.sh` script.

Here's an example command to run the Bootcamp Initialization Tool:

`./bootcamp-initialize --git-token=<yourToken> --org=<yourOrganization> --modules=<module1,module2..> --bootcampee-repo=<repo>`

## Command-line Arguments

The Bootcamp Initialization Tool accepts the following mandatory command-line arguments:

`--git-token`: Specifies the GitHub token with access to create repositories and issues in the provided organization.

`--org`: Specifies the GitHub organization.

`--modules`: Specifies a comma-separated list of modules for which to create issues.

`--bootcampee-repo`: Specifies the repository for the bootcamp participant. If it doesn't exist, it will be created.


## Help

To display the help text and usage instructions, use the `--help` option:

`./bootcamp-initialize --help`

## Available modules

To display modules available, to utilize in the `--modules` parameter, use the `--available-modules` option:

`./bootcamp-initialize --available-modules`