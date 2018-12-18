package retrieve

import (
    "errors"
    "fmt"
    log "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "regexp"
    "time"
)

const dilbertUrlPrefix = "https:"
const dilbertPageUrl = "https://dilbert.com/strip/%d-%d-%d" // year-month-day
const dilbertRegexpStr = "<img class=\"img-responsive img-comic\" width=\"900\" height=\"280\" alt=\"(?P<title>.+?) - Dilbert by Scott Adams\" src=\"(?P<link>[^\"]+)\" />"
const dilbertComicId = "dilbert"

type dilbertRetrieverType struct{}

var dilbertRegexp *regexp.Regexp

func init() {
    var err error
    dilbertRegexp, err = regexp.Compile(dilbertRegexpStr)
    if err != nil {
        log.Panic("Cannot compile the regexp for the Dilbert comic", err)
    }

    var instance dilbertRetrieverType
    registerRetriever(dilbertComicId, instance)
}

func (dilbertRetrieverType) RetrieveIssue() (*model.Issue, error) {
    // https://dilbert.com/strip/2018-12-14
    // <img class="img-responsive img-comic" width="900" height="280" alt="Cake Is Healthy - Dilbert by Scott Adams" src="//assets.amuniversal.com/f2d7d7f0c8c601366722005056a9545d">
    var currentTime = time.Now()
    var year = currentTime.Year()
    var month = currentTime.Month()
    var day = currentTime.Day()

    var dilbertPageUrlStr = fmt.Sprintf(dilbertPageUrl, year, month, day)

    log.Info("Dilbert: retrieving ", dilbertPageUrlStr)
    httpResp, err := http.Get(dilbertPageUrlStr)
    if err != nil {
        log.Warn("Dilbert: failed to retrieve page for current date", err)
        return nil, err
    }

    defer httpResp.Body.Close()

    if httpResp.StatusCode != 200 {
        log.Warn("Dilbert: got bad status code: ", httpResp.StatusCode)
        return nil, err
    }

    htmlContent, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        log.Warn("Dilbert: failed to parse the HTML content", err)
        return nil, err
    }

    match := dilbertRegexp.FindStringSubmatch(string(htmlContent))
    if match == nil {
        log.Warn("Dilbert: no match found in retrieved HTML content")
        return nil, errors.New("no comic issue data found in Dilbert HTML")
    }

    groups := extractGroupsAsMap(match, dilbertRegexp)

    var issue = Issue{
        Comic: dilbertComicId,
        Time:  currentTime,
        Url:   dilbertUrlPrefix + groups["link"],
        Title: groups["title"],
    }

    return &issue, nil
}
