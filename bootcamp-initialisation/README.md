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

`./bootcamp-initialize --git-token=<yourToken> --org=<yourOrganization> --module-location=<directoryWithModules> --modules=<module1,module2..> --bootcampee-repo=<repo>`

## Command-line Arguments

The Bootcamp Initialization Tool accepts the following mandatory command-line arguments:

`--git-token`: Specifies the GitHub token with access to create repositories and issues in the provided organization.

`--org`: Specifies the GitHub organization.

`--module-location`: This expects the directory location for your modules. Each module is expected to have its own dedicated directory, containing the corresponding .md files.

`--modules`: Specifies a comma-separated list of modules for which to create issues.

`--bootcampee-repo`: Specifies the repository for the bootcamp participant. If it doesn't exist, it will be created.


## Help

To display the help text and usage instructions, use the `--help` option:

`./bootcamp-initialize --help`


## Important information

To use the tool effectively, you need to specify a directory where your modules are located using the `--module-location`
parameter. Let's say you choose to provide the directory `/Users/Projects/custom/modules`. Inside this directory, you should have separate
folders for each module. These folders represent your modules and should include all relevant files, such as an
`_index.md` file for that module, any epics and background info.

As an example, if you have a directory structure like so:

```
Your-Project
  content
    modules
      my-module-1
        ...
        .md files
        ...
      my-module-2
        ...
        .md files
        ...
```
you will need to provide `Your-Project/content/modules` as the parameter to the `--module-location` parameter.

If you want a file to be considered for issue creation, it should have a specific naming format. The file needs to be
prefixed with `epic-` and suffixed with `.md.` For example, `epic-deployed-testing.md` and `epic-resiliency.md` are eligible files
that will be recognized by the tool.

However, there are some files that are not eligible for issue creation. Examples include files like `monitoring.md`, `_index.md`,
and `background.md`. The tool will actively filter out these ineligible files.