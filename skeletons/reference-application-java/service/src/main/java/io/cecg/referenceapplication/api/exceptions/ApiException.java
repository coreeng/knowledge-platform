package io.cecg.referenceapplication.api.exceptions;

import lombok.Getter;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;

@RequiredArgsConstructor
@Getter
public abstract class ApiException extends Exception {
    private final HttpStatus statusCode;

    public ApiException(String message, HttpStatus statusCode) {
        super(message);
        this.statusCode = statusCode;
    }

    public ApiException(Exception exception, HttpStatus statusCode) {
        super(exception);
        this.statusCode = statusCode;
    }
}
