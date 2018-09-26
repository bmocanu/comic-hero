package comichero.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

/**
 * Configuration bundle for the application. This gather in one place all the configuration of
 * this application.
 *
 * @author Bogdan Mocanu
 */
@Component
@ConfigurationProperties("app")
public class Configuration {

    private ComicsItemNode sinfest = new ComicsItemNode();
    private ComicsItemNode dilbert = new ComicsItemNode();

    public ComicsItemNode getSinfest() {
        return sinfest;
    }

    public ComicsItemNode getDilbert() {
        return dilbert;
    }

    @Override
    public String toString() {
        return "Configuration{" +
                "sinfest=" + sinfest +
                ", dilbert=" + dilbert +
                '}';
    }
}
