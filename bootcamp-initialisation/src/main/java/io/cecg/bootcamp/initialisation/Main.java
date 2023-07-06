package io.cecg.bootcamp.initialisation;

import io.cecg.bootcamp.initialisation.exception.GitHubIssueCreationException;
import io.cecg.bootcamp.initialisation.exception.MalformedEpicException;
import io.cecg.bootcamp.initialisation.exception.ModuleUnavailableException;
import io.cecg.bootcamp.initialisation.exception.ReadFileException;
import io.cecg.bootcamp.initialisation.model.CmdArguments;
import org.apache.commons.cli.*;
import org.kohsuke.github.*;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.*;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Main {

    private static final String ORG = "org";
    private static final String BOOTCAMPEE_REPO = "bootcampee-repo";
    private static final String GIT_TOKEN = "git-token";
    private static final String MODULES = "modules";
    private static final String CECG_BOOTCAMP_MODULES_LOCATION = "../bootcamp-content/content/bootcamp/modules";
    private static final String GITHUB_BASE_URL = "https://github.com/";

    public static void main(String[] args) throws Exception {
        Options options = buildCmdOptions();

        Set<File> availableModules = getAvailableModules();
        CmdArguments cmdArguments = extractCmdArguments(args, options, availableModules);

        if (cmdArguments.isHelpRequest() || cmdArguments.isModulesRequest()) {
            return;
        }

        String token = cmdArguments.getGitToken();
        GitHub github = new GitHubBuilder().withOAuthToken(token).build();
        GHOrganization organization = github.getOrganization(cmdArguments.getOrg());
        GHRepository ghRepository = getOrCreateRepository(cmdArguments, organization);


        Map<String, GHIssue> issuesByTitle = getIssuesByTitle(ghRepository);

        Set<File> issuesCreatedForModules = createModuleIssues(cmdArguments, ghRepository, availableModules, issuesByTitle);

        System.out.println("Modules affected: \n");
        issuesCreatedForModules.forEach(module -> System.out.println(module.getName() + "\n"));

        System.out.println("Github Link: " + GITHUB_BASE_URL + cmdArguments.getOrg() + "/" + cmdArguments.getBootcampeeRepo() + "/issues/");

    }

    private static Map<String, GHIssue> getIssuesByTitle(GHRepository ghRepository) throws IOException {
        return ghRepository.getIssues(GHIssueState.ALL).stream()
                .collect(Collectors.toMap(GHIssue::getTitle, gh -> gh)
                );
    }

    private static Set<File> getAvailableModules() {
        File[] moduleFiles = new File(CECG_BOOTCAMP_MODULES_LOCATION).listFiles();
        if (moduleFiles != null) {
            return Stream.of(moduleFiles)
                    .filter(File::isDirectory)
                    .collect(Collectors.toSet());
        }
        return Collections.emptySet();
    }

    private static Set<File> createModuleIssues(CmdArguments cmdArguments, GHRepository ghRepository, Set<File> modules, Map<String, GHIssue> issuesByTitle) {
        return modules.stream()
                .filter(module -> cmdArguments.getModules().contains(module.getName()))
                .map(module -> {
                    createLabel(ghRepository, module);
                    Stream.of(Objects.requireNonNull(module.listFiles()))
                            .filter(f -> f.getName().endsWith(".md"))
                            .filter(f -> f.getName().startsWith("epic-"))
                            .collect(Collectors.toSet())
                            .stream()
                            .map(issue -> {
                                String fileContent = readContentFromFile(issue);
                                Scanner scanner = new Scanner(fileContent);
                                String epicTitle = "<unknown>";

                                if (!scanner.nextLine().equals("+++")) {
                                    throw new MalformedEpicException("Expected epic" + issue + " to start with +++");
                                }
                                String nextLine = scanner.nextLine();
                                while (!nextLine.equals("+++")) {
                                    if (nextLine.startsWith("title")) {
                                        epicTitle = nextLine.split(" = ")[1].replace("\"", "").trim();
                                    }
                                    nextLine = scanner.nextLine();
                                }

                                if (!issuesByTitle.containsKey(epicTitle)) {
                                    StringBuilder epicContent = buildEpicContent(scanner);
                                    createGithubIssue(ghRepository, module, epicTitle, epicContent);
                                } else {
                                    System.out.println("Not updating issue" + issue + " as it already exists\n");
                                }
                                return module;
                            }).collect(Collectors.toSet());
                    return module;
                }).collect(Collectors.toSet());
    }

    private static StringBuilder buildEpicContent(Scanner scanner) {
        StringBuilder epicContent = new StringBuilder();
        while (scanner.hasNext()) {
            epicContent.append(scanner.nextLine());
            epicContent.append("\n");
        }
        return epicContent;
    }

    private static void createGithubIssue(GHRepository ghRepository, File module, String epicTitle, StringBuilder epicContent) {
        System.out.println("Creating epic: " + epicTitle + "\n");
        try {
            ghRepository.createIssue(epicTitle)
                    .body(epicContent.toString())
                    .label(module.getName())
                    .label("epic")
                    .create();
        } catch (IOException e) {
            throw new GitHubIssueCreationException("Error occurred while creating epic", e);
        }
    }

    private static String readContentFromFile(File issue) {
        String content;
        try {
            content = Files.readString(Path.of(issue.getPath()));
        } catch (IOException e) {
            throw new ReadFileException("Failed to read file " + issue.getPath(), e);
        }
        return content;
    }

    private static void createLabel(GHRepository ghRepository, File module) {
        try {
            ghRepository.createLabel(module.getName(), "3238a8");
        } catch (IOException e) {
            System.out.println("Label for module " + module.getName() + " already exists\n");
        }
    }

    private static GHRepository getOrCreateRepository(CmdArguments cmdArguments, GHOrganization organization) throws IOException {
        GHRepository ghRepository = organization.getRepository(cmdArguments.getBootcampeeRepo());

        if (ghRepository == null) {
            ghRepository = organization.createRepository(cmdArguments.getBootcampeeRepo())
                    .owner(cmdArguments.getOrg())
                    .private_(true)
                    .create();
        }
        return ghRepository;
    }

    private static CmdArguments extractCmdArguments(String[] args, Options options, Set<File> availableModules) {
        CommandLineParser parser = new DefaultParser();
        CmdArguments.Builder cmdArgumentBuilder = new CmdArguments.Builder();
        boolean isHelpRequest = checkIfHelpRequest(parser, args);
        boolean isModuleRequest = checkIfModuleRequest(parser, args);

        if (!isHelpRequest && !isModuleRequest) {
            try {
                CommandLine cmd = parser.parse(options, args);

                String gitToken = cmd.getOptionValue(GIT_TOKEN);
                String org = cmd.getOptionValue(ORG);
                String repo = cmd.getOptionValue(BOOTCAMPEE_REPO);
                String modules = cmd.getOptionValue(MODULES);

                Set<String> providedModuleNames = Arrays.stream(modules.trim().split(",")).collect(Collectors.toSet());

                ensureModuleIsAvailable(availableModules, providedModuleNames);

                cmdArgumentBuilder
                        .setGitToken(gitToken)
                        .setOrg(org)
                        .setBootcampeeRepo(repo)
                        .setModules(providedModuleNames);

            } catch (ParseException e) {
                System.err.println("Error parsing command-line arguments: " + e.getMessage());
                System.exit(1);
            }
        }

        cmdArgumentBuilder.setHelpRequest(isHelpRequest)
                .setModulesRequest(isModuleRequest);

        return cmdArgumentBuilder.build();
    }

    private static void ensureModuleIsAvailable(Set<File> availableModules, Set<String> providedModuleNames) {
        providedModuleNames.forEach(providedModuleName -> {
            Set<String> availableModuleNames = availableModules.stream().map(File::getName).collect(Collectors.toSet());
            if (!availableModuleNames.contains(providedModuleName)) {
                throw new ModuleUnavailableException("Module " + providedModuleName + " does not exist. Available Modules: " + availableModuleNames);
            }
        });
    }

    private static boolean checkIfModuleRequest(CommandLineParser parser, String[] args) {
        Options options = new Options();
        options.addOption(Option.builder().longOpt("available-modules").hasArg(false).desc("Show available modules").build());

        try {
            CommandLine cmd = parser.parse(options, args);
            if (cmd.hasOption("available-modules")) {

                Set<File> availableModules = getAvailableModules();
                System.out.println("\nAvailable modules: \n");
                availableModules.forEach(module -> System.out.println(module.getName()));
                System.out.println("\n");
            }
        } catch (ParseException e) {
            // Not an available modules request
            return false;
        }
        return true;
    }

    private static boolean checkIfHelpRequest(CommandLineParser parser, String[] args) {
        Options options = new Options();
        options.addOption(Option.builder().longOpt("help").hasArg(false).desc("Show tool usage").build());

        try {
            CommandLine cmd = parser.parse(options, args);
            if (cmd.hasOption("help")) {
                String mandatoryParams = "\nMandatory parameters: \n\n" +
                        "--git-token            ---- Needs to be a valid token with access to create repos/issues in the provided organisation.\n" +
                        "--org                  ---- The github organisation.\n" +
                        "--modules              ---- A comma separated list of modules you want to create the issues for(E.g: p2p-fast-feedback,platform-engineering).\n" +
                        "--bootcampee-repo      ---- The repository of the bootcampee. If it doesn't exist, it will be created.";

                String optionalParams = "\nHelp options: \n\n" +
                        "--help                 ---- Displays help text\n" +
                        "--available-modules    ---- Displays available modules for referencing in the --modules parameter\n";

                String commandUsage = "\n\n./bootcamp-initialize --git-token=<yourToken> --org=<yourOrg> --modules=<yourCommaSeparatedModulesList> --bootcampee-repo=<bootcampeeRepo>";

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

    private static Options buildCmdOptions() {
        Options options = new Options();

        options.addOption(Option.builder().longOpt(GIT_TOKEN).hasArg().required().desc("Git Token").build());
        options.addOption(Option.builder().longOpt(ORG).hasArg().required().desc("Organization").build());
        options.addOption(Option.builder().longOpt(BOOTCAMPEE_REPO).hasArg().required().desc("Repository").build());
        options.addOption(Option.builder().longOpt(MODULES).hasArg().required().desc("Modules").build());

        return options;
    }
}
