package comichero.app;

import comichero.api.ComicIssue;
import comichero.retrievers.SinfestRetriever;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDate;

@RestController
@EnableAutoConfiguration
@ComponentScan(basePackages = {
        "comichero.retrievers",
})
public class MainClass {

    private static final String PREFIX = "<html><body><img src=\"";
    private static final String SUFFIX = "\"></body</html>";

    private final SinfestRetriever sinfestRetriever;

    @Autowired
    public MainClass(SinfestRetriever sinfestRetriever) {
        this.sinfestRetriever = sinfestRetriever;
    }

    @RequestMapping("/")
    @SuppressWarnings("unused")
    String home() {
        ComicIssue issue = sinfestRetriever.retrieveIssue(LocalDate.of(2018,8,19));
        return PREFIX + issue.getImageUrl() + SUFFIX;
    }

    public static void main(String[] args) throws Exception {
        SpringApplication.run(MainClass.class, args);
    }

}