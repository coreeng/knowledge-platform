package io.cecg.bootcamp.initialisation.exception;

import java.io.IOException;

public class IssueManagerException extends RuntimeException {
    public IssueManagerException(String message, IOException e) {
        super(message, e);
    }

    public IssueManagerException(String message) {
        super(message);
    }
}
