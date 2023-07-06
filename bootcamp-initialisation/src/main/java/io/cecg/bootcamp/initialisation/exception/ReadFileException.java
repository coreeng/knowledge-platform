package io.cecg.bootcamp.initialisation.exception;

import java.io.IOException;

public class ReadFileException extends RuntimeException {
    public ReadFileException(String message, IOException exception) {
        super(message, exception);
    }
}
