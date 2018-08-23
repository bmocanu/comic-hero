package comichero.web;

import comichero.model.ComicIssue;
import comichero.retrievers.SinfestRetriever;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDate;

@RestController
public class RootController {

    private final SinfestRetriever sinfestRetriever;

    @Autowired
    public RootController(SinfestRetriever sinfestRetriever) {
        this.sinfestRetriever = sinfestRetriever;
    }

    @RequestMapping(value = "/", method = RequestMethod.GET)
    @SuppressWarnings("unused")
    String homepage(Model model) {
        return "comics-list";
    }

    @RequestMapping("/sinfest")
    @SuppressWarnings("unused")
    String sinfest() {
        String PREFIX = "<html><body><img src=\"";
        String SUFFIX = "\"></body</html>";
        ComicIssue issue = sinfestRetriever.retrieveIssue(LocalDate.of(2018, 8, 19));
        return PREFIX + issue.getImageUrl() + SUFFIX;
    }


}
