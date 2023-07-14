package io.cecg.referenceapplication.api.controllers;

import io.cecg.referenceapplication.api.dtos.StatusResponse;
import io.cecg.referenceapplication.api.exceptions.ApiException;
import io.cecg.referenceapplication.domain.http.DownstreamConnector;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController(value = "downstream")
@ResponseBody
@RequiredArgsConstructor(onConstructor = @__({@Autowired}))
public class DownstreamController {

    private final StatusResponse OK_RESPONSE = StatusResponse.builder().status("OK").build();
    private final DownstreamConnector downstreamConnector;

    @GetMapping("/delay/{delay}")
    public StatusResponse getDelay(@PathVariable(value = "delay") Integer delay) throws ApiException {
        downstreamConnector.getDelay(delay);
        return OK_RESPONSE;
    }

    @GetMapping("/status/{status}")
    public StatusResponse getStatus(@PathVariable(value = "status") Integer status) throws ApiException {
        downstreamConnector.getStatus(status);
        return OK_RESPONSE;
    }

}
