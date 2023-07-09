package io.cecg.referenceapplication.api.filters;

import com.google.common.util.concurrent.Uninterruptibles;
import lombok.extern.slf4j.Slf4j;
import org.eclipse.jetty.util.component.LifeCycle;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;

@Slf4j
@Component
public class ConnectionDrainingFilter extends OncePerRequestFilter implements LifeCycle.Listener {
    private final AtomicBoolean shuttingDown = new AtomicBoolean(false);
    private final Integer drainingMs;

    public ConnectionDrainingFilter(@Value("${server.drainingMs}") Integer drainingMs) {
        this.drainingMs = drainingMs;
    }


    @Override
    public void lifeCycleStarting(LifeCycle event) {

    }

    @Override
    public void lifeCycleStarted(LifeCycle event) {

    }

    @Override
    public void lifeCycleFailure(LifeCycle event, Throwable cause) {

    }

    @Override
    public void lifeCycleStopping(LifeCycle event) {
        log.info("Draining connections");
        shuttingDown.set(true);
        Uninterruptibles.sleepUninterruptibly(drainingMs, TimeUnit.MILLISECONDS);
        log.info("Finished draining connections");

    }

    @Override
    public void lifeCycleStopped(LifeCycle event) {

    }

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        if(shuttingDown.get()) {
            response.addHeader("Connection", "close");
        }
        filterChain.doFilter(request, response);
    }
}
