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

// <div id="comic-wrap".+?<div id="comic">.+?<img src="(?P<link>[^"]+)".+?<h2 class="post-title"><a href="[^"]+">(?P<title>.+?)</a></h2>
// (?P<title>.+?)
// (?P<link>[^"]+)

//<div id="comic-wrap" class="comic-id-1438">
//<div id="comic-head"></div>
//<div class="comic-table">
//<div id="comic">
//<img src="http://twoleafclover.net/wp-content/uploads/2018/10/2018-10-26-HeadinaHole.png" alt="" title="" scale="0">				</div>
//</div>
//<div id="comic-foot">
//<table id="comic-nav-wrapper">
//<tbody><tr class="comic-nav-container">
//<td class="comic-nav"><a href="http://twoleafclover.net/?comic=healthcaretlc" class="comic-nav-base comic-nav-first">‹‹ First</a></td>
//<td class="comic-nav"><a href="http://twoleafclover.net/?comic=hi-fred" class="comic-nav-base comic-nav-previous">‹ Prev</a></td>
//<td class="comic-nav"><a href="http://twoleafclover.net/?comic=head-in-a-hole#comments" class="comic-nav-comments" title="Head in a Hole">Comments(<span class="comic-nav-comment-count">0</span>)</a></td>
//<td class="comic-nav"><a href="http://twoleafclover.net?random&amp;nocache=1" class="comic-nav-random" title="Random Comic">Random</a></td>
//<td class="comic-nav"><span class="comic-nav-base comic-nav-next comic-nav-void ">Next ›</span></td>
//<td class="comic-nav"><span class="comic-nav-base comic-nav-last comic-nav-void ">Last ››</span></td>
//</tr>
//
//</tbody></table>
//</div>
//<div class="clear"></div>
//</div>
//<div id="content" class="narrowcolumn">
//<div id="post-1438" class="post-1438 comic type-comic status-publish hentry uentry postonpage-1 odd post-author-admin">
//<div class="post-content">
//<div class="post-info">
//<h2 class="post-title"><a href="http://twoleafclover.net/?comic=head-in-a-hole">Head in a Hole</a></h2>

const twoLeafCloverPageUrlStr = "http://twoleafclover.net/"
const twoLeafCloverRegexpStr = `(?Us)<div id="comic-wrap".+?<div id="comic">.+?<img src="(?P<imageUrl>[^"]+)".+?<h2 class="post-title"><a href="(?P<issueUrl>[^"]+)">(?P<title>.*?)</a></h2>`
const twoLeafCloverComicName = "two-leaf-clover"

type twoLeafCloverRetrieverType struct{}

var twoLeafCloverRegexp *regexp.Regexp

func init() {
    var err error
    twoLeafCloverRegexp, err = regexp.Compile(twoLeafCloverRegexpStr)
    if err != nil {
        log.Panic("Cannot compile the regexp for the TwoLeafClover comic: ", err)
    }

    registerRetriever(twoLeafCloverComicName, &twoLeafCloverRetrieverType{})
}

func (twoLeafCloverRetrieverType) RetrieveIssue() (*model.Issue, error) {
    var currentTime = time.Now()

    log.Info("TwoLeafClover: retrieving ", twoLeafCloverPageUrlStr)
    httpResp, err := http.Get(twoLeafCloverPageUrlStr)
    if err != nil {
        log.Warn("TwoLeafClover: failed to retrieve page for current date: ", err)
        return nil, err
    }

    defer httpResp.Body.Close()

    if httpResp.StatusCode != 200 {
        log.Warn("TwoLeafClover: got bad status code: ", httpResp.StatusCode)
        return nil, err
    }

    htmlContent, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        log.Warn("TwoLeafClover: failed to parse the HTML content: ", err)
        return nil, err
    }

    match := twoLeafCloverRegexp.FindStringSubmatch(string(htmlContent))
    if match == nil {
        log.Warn("TwoLeafClover: no match found in retrieved HTML content")
        return nil, errors.New("no comic issue data found in TwoLeafClover HTML")
    }

    groups := extractGroupsAsMap(match, twoLeafCloverRegexp)

    var issue = model.Issue{
        ComicName: twoLeafCloverComicName,
        Time:      currentTime,
        IssueUrl:  groups["issueUrl"],
        ImageUrl:  groups["imageUrl"],
        Title:     groups["title"],
    }

    return &issue, nil
}
