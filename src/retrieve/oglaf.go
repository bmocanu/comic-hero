package retrieve

import (
    "comic-hero/model"
    "errors"
    log "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "regexp"
    "time"
)

const oglafPageUrl = "https://www.oglaf.com/"
const oglafRegexpStr = `<img id="strip"\s+src="(?P<link>[^"]+)"\s+.+?\s+title="(?P<title>[^"]+)"\s+/>`
const oglafComicName = "oglaf"

type oglafRetrieverType struct{}

var oglafRegexp *regexp.Regexp

func init() {
    var err error
    oglafRegexp, err = regexp.Compile(oglafRegexpStr)
    if err != nil {
        log.Panic("Cannot compile the regexp for Oglaf comic: ", err)
    }

    registerRetriever(oglafComicName, &oglafRetrieverType{})
}

func (oglafRetrieverType) RetrieveIssue() (*model.Issue, error) {
    // HTML sample
    // <img id="strip" src="https://media.oglaf.com/comic/rectitude.jpg" alt="maybe just get a bigger funnel?" title="the inaccurate conception">

    log.Info("Oglaf: retrieving ", oglafPageUrl)
    httpResp, err := http.Get(oglafPageUrl)
    if err != nil {
        log.Warn("Oglaf: failed to retrieve HTML page: ", err)
        return nil, err
    }

    defer httpResp.Body.Close()

    if httpResp.StatusCode != 200 {
        log.Warn("Oglaf: got bad status code: ", httpResp.StatusCode)
        return nil, err
    }

    htmlContent, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        log.Warn("Oglaf: failed to parse the HTML content")
        return nil, err
    }

    match := oglafRegexp.FindStringSubmatch(string(htmlContent))
    if match == nil {
        log.Warn("Oglaf: no match found in retrieved HTML content")
        return nil, errors.New("no comic issue data found in Oglaf HTML")
    }

    groups := extractGroupsAsMap(match, oglafRegexp)
    var issue = model.Issue{
        ComicName: oglafComicName,
        Time:      time.Now(),
        IssueUrl:  oglafPageUrl,
        ImageUrl:  groups["link"],
        Title:     groups["title"],
    }

    return &issue, nil
}
