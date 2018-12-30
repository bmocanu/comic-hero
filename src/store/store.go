package store

import (
    "comic-hero/config"
    "comic-hero/model"
    "crypto/sha1"
    "fmt"
    log "github.com/sirupsen/logrus"
)

var comicStore = make(map[int]*model.IssueLink)

func NewIssue(issue *model.Issue) {
    var comicId, _ = config.GetIdForComicName(issue.ComicName)
    var issueHash = calculateHashForIssue(issue)
    log.Info("New issue to store: comicName=[", issue.ComicName,
        "], comicId=[", comicId,
        "], title=[", issue.Title,
        "], imageUrl=[", issue.ImageUrl,
        "], hash=[", issueHash, "]")
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

    if totalIssueCount >= config.IssueStoreSize {
        linkBeforeLast.NextLink = nil
    }
    totalIssueCount++

    var newLink model.IssueLink
    newLink.Issue = issue
    newLink.Hash = issueHash
    newLink.NextLink = link
    newLink.IssueCount = totalIssueCount

    comicStore[comicId] = &newLink
    log.Info("New issue stored, total issues currently: ", totalIssueCount)
}

func GetIssuesForComicId(id int) *model.IssueLink {
    var issueLink, _ = comicStore[id]
    return issueLink
}

func calculateHashForIssue(issue *model.Issue) string {
    hashTarget := issue.ImageUrl
    h := sha1.New()
    _, err := h.Write([]byte(hashTarget))
    if err != nil {
        log.Error("Failed to calculate hash for string: ", hashTarget, ": ", err)
    }
    return fmt.Sprintf("%x", h.Sum(nil))
}
