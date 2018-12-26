package store

import (
    "comic-hero/config"
    "comic-hero/model"
    log "github.com/sirupsen/logrus"
    "hash/fnv"
)

var comicStore = make(map[int]*model.IssueLink)

func NewIssue(issue *model.Issue) {
    log.Info("New issue to be stored: ", issue.Comic, " - ", issue.Title, " - ", issue.Url)
    var comicId, _ = config.GetIdForComicName(issue.Comic)
    var issueHash = calculateHashForIssue(issue)
    log.Info("New issue data: comicId ", comicId, ", issueHash ", issueHash)
    var link, _ = comicStore[comicId]

    // list sanitization
    var newIssueExists = false
    var totalIssueCount = 0
    var linkBeforeLast *model.IssueLink

    for linkCursor := link; linkCursor != nil; linkCursor = linkCursor.NextLink {
        totalIssueCount++
        if linkCursor.Hash == issueHash {
            newIssueExists = true
        }
        if linkCursor.NextLink != nil {
            linkBeforeLast = linkCursor
        }
    }

    if newIssueExists {
        log.Info("New issue already stored, dropping it")
        return
    }

    if totalIssueCount >= 10 {
        linkBeforeLast.NextLink = nil
    }

    var newLink model.IssueLink
    newLink.Issue = issue
    newLink.Hash = issueHash
    newLink.NextLink = link
    newLink.IssueCount = totalIssueCount + 1

    comicStore[comicId] = &newLink
    log.Info("New issue stored, total issues currently: ", totalIssueCount+1)
}

func GetIssuesForComicId(id int) *model.IssueLink {
    var issueLink, _ = comicStore[id]
    return issueLink
}

func calculateHashForIssue(issue *model.Issue) uint32 {
    hashTarget := issue.Url
    h := fnv.New32a()
    _, err := h.Write([]byte(hashTarget))
    if err != nil {
        log.Error("Failed to calculate hash for string ", hashTarget, err)
    }
    return h.Sum32()
}
