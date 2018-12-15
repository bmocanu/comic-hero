package retrieve

import (
	"comic-hero/src/model"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

const sinfestUrlPrefix = "http://sinfest.net/"
const sinfestPageUrl = "http://sinfest.net/view.php?date=%d-%d-%d" // year-month-day
const sinfestRegexpStr = "<img src=\"(?P<link>[^\"]+)\" alt=\"(?P<title>[^\"]+)\">"
const sinfestComicId = "sinfest"

var sinfestRegexp *regexp.Regexp

func init() {
	var err error
	sinfestRegexp, err = regexp.Compile(sinfestRegexpStr)
	if err != nil {
		log.Panic("Cannot compile the regexp for Sinfest comic", err)
	}
}

func SinfestIssue() (*model.Issue, error) {
	// http://sinfest.net/view.php?date=2018-12-14
	// <img src="btphp/comics/2018-12-14.gif" alt="MMXVIII 29">
	var currentTime = time.Now()
	var year = currentTime.Year()
	var month = currentTime.Month()
	var day = currentTime.Day()

	var sinfestPageUrlStr = fmt.Sprintf(sinfestPageUrl, year, month, day)

	log.Info("Sinfest: retrieving ", sinfestPageUrlStr)
	httpResp, err := http.Get(sinfestPageUrlStr)
	if err != nil {
		log.Warn("Sinfest: failed to retrieve HTML page for current date", err)
		return nil, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode == 200 {
		log.Info("Sinfest: OK")
	} else {
		log.Warn("Sinfest: Got bad status code: ", httpResp.StatusCode)
		return nil, err
	}

	htmlContent, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Warn("Sinfest: failed to parse the HTML content")
		return nil, err
	}

	match := sinfestRegexp.FindStringSubmatch(string(htmlContent))
	if match == nil {
		log.Warn("Sinfest: no match found in retrieved HTML content")
		return nil, errors.New("no comic issue data found in Sinfest HTML")
	}

	groups := extractGroupsAsMap(match, dilbertRegexp)
	var issue = model.Issue{
		Comic: sinfestComicId,
		Time:  currentTime,
		Url:   sinfestUrlPrefix + groups["link"],
		Title: groups["title"],
	}

	return &issue, nil
}
