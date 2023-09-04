package io.cecg.referenceapplication;

import io.cecg.referenceapplication.api.filters.ConnectionDrainingFilter;
import lombok.RequiredArgsConstructor;
import org.eclipse.jetty.server.handler.HandlerWrapper;
import org.eclipse.jetty.server.handler.StatisticsHandler;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.embedded.jetty.JettyServletWebServerFactory;
import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.stereotype.Component;

@Component
@RequiredArgsConstructor(onConstructor = @__({@Autowired}))
public class GracefulShutdownHandler implements WebServerFactoryCustomizer<JettyServletWebServerFactory> {

    private final ConnectionDrainingFilter connectionDrainingFilter;

    @Override
    public void customize(JettyServletWebServerFactory factory) {
        factory.addServerCustomizers(server -> {
            server.addLifeCycleListener(connectionDrainingFilter);

            HandlerWrapper wrapperStatistics = new StatisticsHandler(); //metrics
            wrapperStatistics.setServer(server);
            wrapperStatistics.setHandler(server.getHandler());

            server.setHandler(wrapperStatistics);
            server.setStopAtShutdown(false);
        });
    }
}
