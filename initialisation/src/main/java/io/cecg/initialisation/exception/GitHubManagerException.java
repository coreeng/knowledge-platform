package io.cecg.initialisation.exception;

import java.io.IOException;

public class GitHubManagerException extends RuntimeException {
    public GitHubManagerException(String message, IOException ioException) {
        super(message, ioException);
    }
}
