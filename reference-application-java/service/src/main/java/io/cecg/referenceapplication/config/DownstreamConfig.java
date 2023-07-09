package io.cecg.referenceapplication.config;

import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;

@ConfigurationProperties(prefix = "downstream")
@Configuration
@Getter
@Setter
public class DownstreamConfig {
    private String url;
    private Integer port;
    private Long readTimeoutMs;
    private Long connectTimeoutMs;

    public String getFullUrl() {
        if (port == null)
            return url;
        else
            return String.format("%s:%d", url, port);
    }

    public String getFullUrlWithPath(String path) {
        return String.format("%s/%s", getFullUrl(), path);
    }
}
