package comichero.retrievers;

import comichero.api.ComicIssue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import java.io.IOException;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.text.MessageFormat;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Comic issue retriever for <a href="http://sinfest.net">Sinfest</a>.
 *
 * @author Bogdan Mocanu
 * @see Retriever
 * @see ComicIssue
 */
@Component
public class SinfestRetriever implements Retriever {

    private static final Logger log = LoggerFactory.getLogger(SinfestRetriever.class);

    private static final String TITLE = "sinfest";
    private static final String ID = "{0}-{1}";
    private static final String DATE = "yyyy-MM-dd";
    private static final String WEB_PAGE = "http://sinfest.net/view.php?date={0}";
    private static final String REGEX = "<img\\s+src=\"([^\"]+)\"\\s+alt=\"([^\"]+)\"";
    private static final Pattern REGEX_PATTERN = Pattern.compile(REGEX);

    // Format is:
    // Web page: http://sinfest.net/view.php?date=2018-04-29
    // Target content: <img src="btphp/comics/2018-04-29.gif" alt="some title">

    public ComicIssue retrieveIssue(LocalDate date) {
        String dateString = DateTimeFormatter.ofPattern(DATE).format(date);
        log.info("Sinfest: preparing to load web page for date {}", dateString);
        String issueId = MessageFormat.format(ID, TITLE, dateString);
        try {
            String pageContent = getWebPageAsString(MessageFormat.format(WEB_PAGE, dateString));
            log.info("Sinfest: page loaded, total {} chars", pageContent.length());
            Matcher regexMatcher = REGEX_PATTERN.matcher(pageContent);
            if (regexMatcher.matches()) {
                log.info("Sinfest: target content successfully found, comic issue generated");
                String comicImageUrl = regexMatcher.group(1);
                String comicTitle = regexMatcher.group(2);
                return new ComicIssue(issueId, comicTitle, comicImageUrl);
            }
            log.warn("Cannot retrieve Sinfest web page for date:[" + dateString + "]. Reason: cannot find pattern in page");
        } catch (IOException exception) {
            log.warn("Cannot retrieve Sinfest web page for date [" + dateString + "]", exception);
        }
        return null;
    }

    private String getWebPageAsString(String url) throws IOException {
        try (Scanner scanner = new Scanner(new URL(url).openStream(),
                StandardCharsets.UTF_8.toString())) {
            scanner.useDelimiter("\\A");
            return scanner.hasNext() ? scanner.next() : "";
        }
    }

}
