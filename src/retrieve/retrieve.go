package retrieve

import (
    "comic-hero/model"
    "comic-hero/store"
    "github.com/mileusna/crontab"
    log "github.com/sirupsen/logrus"
    "regexp"
)

//const cronOncePerHour = "0 * * * *"
const cronOncePerMinute = "* * * * *"

var retrievers = make(map[string]Retriever)

func init() {
    crontab.New().MustAddJob(cronOncePerMinute, FetchNewIssues)
}

// Retriever defines the behaviour of types that are capable of retrieving comic content (image, title, date) for a
// particular comic website
type Retriever interface {
    RetrieveIssue() (*model.Issue, error)
}

// registerRetriever adds a new retriever in the map of known retrievers. The map is later on used for activating
// each retriever and getting a new issue for each comic
func registerRetriever(name string, retriever Retriever) {
    if _, retrieverExists := retrievers[name]; retrieverExists {
        log.Error("Cannot register a new retriever, as there is already one with the name ", name)
        return
    }
    retrievers[name] = retriever
    log.Info("New comic retriever registered: ", name)
}

func FetchNewIssues() {
    log.Info("Calling all retrievers for a new round of fetching comic issues")
    for _, retriever := range retrievers {
        issue, err := retriever.RetrieveIssue()
        if err == nil {
            store.NewIssue(issue)
        }
    }
}

// Utility functions ----------------------------------------------------------------------------------------

// ExtractGroupsAsMap creates a map out of the regex groups found in the given regular expression, putting the
// group name as key and the actual matched content as a value
func extractGroupsAsMap(matches []string, regexp *regexp.Regexp) map[string]string {
    groups := make(map[string]string)
    matchNames := regexp.SubexpNames()
    for index, groupName := range matchNames {
        if index != 0 && groupName != "" {
            groups[groupName] = matches[index]
        }
    }
    return groups
}
