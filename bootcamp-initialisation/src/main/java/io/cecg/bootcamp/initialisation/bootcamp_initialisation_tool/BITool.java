package io.cecg.bootcamp.initialisation.bootcamp_initialisation_tool;

import io.cecg.bootcamp.initialisation.github_manager.GitHubManager;
import io.cecg.bootcamp.initialisation.issue_manager.IssueManager;
import io.cecg.bootcamp.initialisation.model.CmdArguments;
import io.cecg.bootcamp.initialisation.module_manager.ModuleManager;
import org.apache.commons.cli.*;
import org.kohsuke.github.GHIssue;
import org.kohsuke.github.GHRepository;

import java.io.File;
import java.util.Arrays;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

public class BITool {

    private static final String ORG = "org";
    private static final String BOOTCAMPEE_REPO = "bootcampee-repo";
    private static final String GIT_TOKEN = "git-token";
    private static final String MODULES = "modules";
    private static final String MODULE_LOCATION = "module-location";
    private static final String GITHUB_BASE_URL = "https://github.com/";

    public static void run(String[] args) {
        Options options = buildCmdOptions();

        CmdArguments cmdArguments = extractCmdArguments(args, options);

        if (cmdArguments.isHelpRequest()) {
            return;
        }

        Set<File> availableModules = ModuleManager.getAvailableModules(cmdArguments.getModuleLocation());
        ModuleManager.ensureModuleIsAvailable(availableModules, cmdArguments.getModules());
        GHRepository ghRepository = GitHubManager.getOrCreateRepository(cmdArguments);

        Map<String, GHIssue> issuesByTitle = IssueManager.getIssuesByTitle(ghRepository);
        Set<File> issuesCreatedForModules = IssueManager.createModuleIssues(cmdArguments, ghRepository, availableModules, issuesByTitle);

        System.out.println("Modules affected: \n");
        issuesCreatedForModules.forEach(module -> System.out.println(module.getName() + "\n"));

        System.out.println("Github Link: " + GITHUB_BASE_URL + cmdArguments.getOrg() + "/" + cmdArguments.getBootcampeeRepo() + "/issues/");
    }

    private static Options buildCmdOptions() {
        Options options = new Options();

        options.addOption(Option.builder().longOpt(GIT_TOKEN).hasArg().required().desc("Git Token").build());
        options.addOption(Option.builder().longOpt(ORG).hasArg().required().desc("Organization").build());
        options.addOption(Option.builder().longOpt(BOOTCAMPEE_REPO).hasArg().required().desc("Repository").build());
        options.addOption(Option.builder().longOpt(MODULES).hasArg().required().desc("Modules").build());
        options.addOption(Option.builder().longOpt(MODULE_LOCATION).hasArg().required().desc("Module location").build());

        return options;
    }

    private static CmdArguments extractCmdArguments(String[] args, Options options) {
        CommandLineParser parser = new DefaultParser();
        CmdArguments.Builder cmdArgumentBuilder = new CmdArguments.Builder();
        boolean isHelpRequest = checkIfHelpRequest(parser, args);

        if (!isHelpRequest) {
            try {
                CommandLine cmd = parser.parse(options, args);

                String gitToken = cmd.getOptionValue(GIT_TOKEN);
                String org = cmd.getOptionValue(ORG);
                String repo = cmd.getOptionValue(BOOTCAMPEE_REPO);
                String modules = cmd.getOptionValue(MODULES);
                String moduleLocation = cmd.getOptionValue(MODULE_LOCATION);

                Set<String> providedModuleNames = Arrays.stream(modules.trim().split(",")).collect(Collectors.toSet());

                cmdArgumentBuilder
                        .setGitToken(gitToken)
                        .setOrg(org)
                        .setBootcampeeRepo(repo)
                        .setModules(providedModuleNames)
                        .setModuleLocation(moduleLocation);

            } catch (ParseException e) {
                System.err.println("Error parsing command-line arguments: " + e.getMessage());
                System.exit(1);
            }
        }

        cmdArgumentBuilder.setHelpRequest(isHelpRequest);

        return cmdArgumentBuilder.build();
    }

    private static boolean checkIfHelpRequest(CommandLineParser parser, String[] args) {
        Options options = new Options();
        options.addOption(Option.builder().longOpt("help").hasArg(false).desc("Show tool usage").build());

        try {
            CommandLine cmd = parser.parse(options, args);
            if (cmd.hasOption("help")) {
                String mandatoryParams = """

                        Mandatory parameters:\s

                        --git-token            ---- Needs to be a valid token with access to create repos/issues in the provided organisation.
                        --org                  ---- The github organisation.
                        --module-location      ---- The directory location for your modules. Each module is expected to have its own dedicated directory, containing the corresponding .md files.
                        --modules              ---- A comma separated list of modules you want to create the issues for(E.g: p2p-fast-feedback,platform-engineering).
                        --bootcampee-repo      ---- The repository of the bootcampee. If it doesn't exist, it will be created.""";

                String optionalParams = """

                        Help options:\s

                        --help                 ---- Displays help text
                        """;

                String commandUsage = "\n\n./bootcamp-initialize --git-token=<yourToken> --org=<yourOrg> --modules=<yourCommaSeparatedModulesList> --module-location=<Your/Module/Location> --bootcampee-repo=<bootcampeeRepo>";

                System.out.println("\nCommand usage: " + commandUsage);
                System.out.println("\n" + mandatoryParams);
                System.out.println("\n" + optionalParams);
            }
        } catch (ParseException e) {
            // Not a help request
            return false;
        }
        return true;
    }
}
