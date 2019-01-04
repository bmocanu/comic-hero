package store

import (
    "comic-hero/config"
    "comic-hero/model"
    "crypto/sha1"
    "fmt"
    log "github.com/sirupsen/logrus"
    "strings"
)

var comicStore = make(map[int]*model.IssueList)

func NewIssue(issue *model.Issue) {
    var comicId, _ = config.GetIdForComicName(issue.ComicName)
    var issueHash = calculateHashForIssue(issue)

    // issue sanitization
    if strings.TrimSpace(issue.Title) == "" {
        issue.Title = fmt.Sprintf("No title for %d/%d/%d", issue.Time.Day(), issue.Time.Month(), issue.Time.Year())
    }

    log.Info("New issue to store: comicName=[", issue.ComicName,
        "], comicId=[", comicId,
        "], title=[", issue.Title,
        "], issueUrl=[", issue.IssueUrl,
        "], imageUrl=[", issue.ImageUrl,
        "], hash=[", issueHash, "]")
    var list, listFound = comicStore[comicId]

    if !listFound {
        list = &model.IssueList{}
        comicStore[comicId] = list
    }

    // list sanitization
    var newIssueExists = false
    var linkBeforeLast *model.IssueLink

    for linkCursor := list.FirstLink; linkCursor != nil; linkCursor = linkCursor.NextLink {
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

    if list.LinkCount >= config.Store.IssuesStoredPerComic {
        linkBeforeLast.NextLink = nil
        list.LinkCount--
    }

    var proxyImage = config.IsComicImageProxyEnabled(issue.ComicName)
    var proxyUrl string
    if proxyImage {
        proxyUrl = calculateProxyUrlForComic(comicId, issueHash)
    }

    var newLink model.IssueLink
    newLink.Issue = issue
    newLink.Hash = issueHash
    newLink.NextLink = list.FirstLink
    newLink.ProxyImage = proxyImage
    newLink.ProxyImageUrl = proxyUrl

    list.FirstLink = &newLink
    list.LinkCount++

    log.Info("New issue stored, total issues currently: ", list.LinkCount)
}

func GetIssueListForComicId(comicId int) *model.IssueList {
    var list, _ = comicStore[comicId]
    return list
}

func GetIssueLinkByComicIdAndHash(comicId int, hash string) *model.IssueLink {
    var issueList = GetIssueListForComicId(comicId)
    if issueList == nil {
        return nil
    }

    for linkCursor := issueList.FirstLink; linkCursor != nil; linkCursor = linkCursor.NextLink {
        if linkCursor.Hash == hash {
            return linkCursor
        }
    }
    return nil
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

func calculateProxyUrlForComic(comicId int, issueHash string) string {
    return fmt.Sprintf("%s/get/%d/%s", config.Server.BaseUrl, comicId, issueHash)
}
