package io.cecg.referenceapplication.api.dtos;

import lombok.Builder;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@Builder
@RequiredArgsConstructor
public class StatusResponse {
    private final String status;
}
