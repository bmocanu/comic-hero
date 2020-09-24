package retrieve

import (
    "comic-hero/model"
    "errors"
    "fmt"
    log "github.com/sirupsen/logrus"
    "io/ioutil"
    "net/http"
    "regexp"
    "time"
)

const gocomicsPageUrl = "https://www.gocomics.com/%s/%d/%d/%d" // comicName-year-month-day
const gocomicsRegexpStr = `(?Us)<div class="comic container js-comic[^"]+".*?data-title="(?P<title>.*?) \| GoComics\.com".*?data-image="(?P<link>[^"]+)"`

var gocomicsNames = [...]string{
    "bc", "calvinandhobbes", "fowl-language", "realitycheck", "pearlsbeforeswine", "moderately-confused",
    "speedbump", "garfield", "dicktracy", "freerange", "brevity", "deflocked", "birdbrains", "peanuts",
    "duplex", "jerryholbert", "life-on-earth", "whyatt-cartoons", "pot-shots", "waynovision", "dogeatdoug",
    "luann", "mikeluckovich", "nonsequitur", "bignate", "inthebleachers", "imaginethis", "bloomcounty",
    "wizardofid", "strangebrew", "wumo", "gasolinealley", "getfuzzy", "roseisrose", "robrogers", "offthemark",
    "libertymeadows", "heathcliff", "herbandjamaal", "harley", "9to5", "herman", "pickles",
}

var gocomicsRegexp *regexp.Regexp

type gocomicsRetrieverType struct {
    comicName string
}

func init() {
    var err error
    gocomicsRegexp, err = regexp.Compile(gocomicsRegexpStr)
    if err != nil {
        log.Panic("Cannot compile the regexp for GoComics: ", err)
    }

    for _, comicName := range gocomicsNames {
        registerRetriever(comicName, &gocomicsRetrieverType{comicName: comicName})
    }
}

func (retriever gocomicsRetrieverType) RetrieveIssue() (*model.Issue, error) {
    //<div class="comic container js-comic-2663988 js-item-init js-item-share js-comic-swipe bg-white border rounded" data-shareable-model="FeatureItem"
    //data-shareable-id="2663988"
    //data-transcript=""
    //data-id="2663988"
    //data-feature-id="193"
    //data-feature-name="B.C."
    //data-feature-code="crbc"
    //data-feature-type="comic"
    //data-feature-format="print"
    //data-date="2018-12-29"
    //data-formatted-date="December 29, 2018"
    //data-url="https://www.gocomics.com/bc/2018/12/29"
    //data-creator="Mastroianni and Hart"
    //data-title="B.C. for December 29, 2018 | GoComics.com"
    //data-tags=""
    //data-description="For December 29, 2018"
    //data-image="https://assets.amuniversal.com/181fcc10dac601366dcc005056a9545d"
    //itemtype="http://schema.org/CreativeWork"
    //accountableperson="Andrews McMeel Universal"
    //creator="Mastroianni and Hart">

    //<div class="comic container js-comic[^"]+"[.\s]+?data-title="(?P<title>.+?) \| GoComics\.com"[.\s]+?data-image="(?P<link>[^"]+)"

    var comicName = retriever.comicName
    var currentTime = time.Now()
    var year = currentTime.Year()
    var month = currentTime.Month()
    var day = currentTime.Day()

    var gocomicsPageUrlStr = fmt.Sprintf(gocomicsPageUrl, comicName, year, month, day)
    log.Info("GoComics: ", comicName, ": ", gocomicsPageUrlStr)

    httpResp, err := http.Get(gocomicsPageUrlStr)
    if err != nil {
        log.Warn("GoComics: ", comicName, ": failed to retrieve HTML page for current date: ", err)
        return nil, err
    }

    defer httpResp.Body.Close()

    if httpResp.StatusCode != 200 {
        log.Warn("GoComics: ", comicName, ": Got bad status code: ", httpResp.StatusCode)
        return nil, err
    }

    htmlContent, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        log.Warn("GoComics: ", comicName, ": failed to parse the HTML content")
        return nil, err
    }

    match := gocomicsRegexp.FindStringSubmatch(string(htmlContent))
    if match == nil {
        log.Warn("GoComics: ", comicName, ": no match found in retrieved HTML content")
        return nil, errors.New("no comic issue data found in GoComics-" + comicName + " HTML")
    }

    groups := extractGroupsAsMap(match, gocomicsRegexp)
    var issue = model.Issue{
        ComicName: comicName,
        Time:      currentTime,
        IssueUrl:  gocomicsPageUrlStr,
        ImageUrl:  groups["link"],
        Title:     groups["title"],
    }

    return &issue, nil
}
