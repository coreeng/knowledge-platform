package io.cecg.referenceapplication.api.filters;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.springframework.util.StopWatch;
import org.springframework.util.StringUtils;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

@Slf4j
@Component
public class LogFilter extends GenericFilter {

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain) throws IOException, ServletException {
        StopWatch watch = new StopWatch();
        watch.start();
        String path = null;

        if (request instanceof HttpServletRequest) {
            path = ((HttpServletRequest) request).getRequestURI();
        }
        int status = 0;

        chain.doFilter(request, response);
        watch.stop();
        if (response instanceof HttpServletResponse) {
            status = ((HttpServletResponse) response).getStatus();
        }
        long time = watch.getLastTaskTimeMillis();

        if (!StringUtils.isEmpty(path)) {
            log.info("Request for {} took {} ms. Had status {}.", path, time, status);
        }
    }
}
