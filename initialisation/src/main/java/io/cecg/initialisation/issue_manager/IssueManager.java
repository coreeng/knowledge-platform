package io.cecg.initialisation.issue_manager;

import io.cecg.initialisation.exception.IssueManagerException;
import io.cecg.initialisation.model.CmdArguments;
import org.kohsuke.github.GHIssue;
import org.kohsuke.github.GHIssueState;
import org.kohsuke.github.GHRepository;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Map;
import java.util.Objects;
import java.util.Scanner;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class IssueManager {
    public static Map<String, GHIssue> getIssuesByTitle(GHRepository ghRepository) {
        try {
            return ghRepository.getIssues(GHIssueState.ALL).stream()
                    .collect(Collectors.toMap(GHIssue::getTitle, gh -> gh));
        } catch (IOException e) {
            throw new IssueManagerException("Error occurred while getting repository issues", e);
        }
    }

    public static Set<File> createModuleIssues(CmdArguments cmdArguments, GHRepository ghRepository, Set<File> modules, Map<String, GHIssue> issuesByTitle) {
        return modules.stream()
                .filter(module -> cmdArguments.getModules().contains(module.getName()))
                .map(module -> {
                    createLabel(ghRepository, module);
                    Set<File> collect = Stream.of(Objects.requireNonNull(module.listFiles()))
                            .filter(f -> f.getName().endsWith(".md"))
                            .filter(f -> f.getName().startsWith("epic-"))
                            .collect(Collectors.toSet());
                    collect
                            .stream()
                            .map(issue -> {
                                String fileContent = readContentFromFile(issue);
                                Scanner scanner = new Scanner(fileContent);
                                String epicTitle = "<unknown>";

                                if (!scanner.nextLine().equals("+++")) {
                                    throw new IssueManagerException("Expected epic" + issue + " to start with +++");
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
            throw new IssueManagerException("Error occurred while creating epic", e);
        }
    }

    private static String readContentFromFile(File issue) {
        String content;
        try {
            content = Files.readString(Path.of(issue.getPath()));
        } catch (IOException e) {
            throw new IssueManagerException("Failed to read file " + issue.getPath(), e);
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
}
