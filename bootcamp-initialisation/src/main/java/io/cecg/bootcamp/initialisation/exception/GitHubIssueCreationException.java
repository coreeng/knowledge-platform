package io.cecg.bootcamp.initialisation.exception;

import java.io.IOException;

public class GitHubIssueCreationException extends RuntimeException {
    public GitHubIssueCreationException(String message, IOException exception) {
        super(message, exception);
    }
}
