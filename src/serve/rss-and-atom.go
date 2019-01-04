package serve

import (
    "comic-hero/config"
    "comic-hero/model"
    "comic-hero/store"
    "fmt"
    "github.com/gorilla/feeds"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "io"
    "net/http"
    "strconv"
    "time"
)

func getRss20Feed(response http.ResponseWriter, request *http.Request) {
    var feed, err = generateFeedObject(response, request)
    if err != nil {
        log.Warn("HTTP request for RSS 2.0 feed failed: ", err)
        return
    }

    xmlContent, err := feed.ToRss()
    if err != nil {
        log.Warn("HTTP request for RSS 2.0 feed failed: ", err)
    }

    response.Header().Set("Content-Type", "text/xml")
    response.WriteHeader(http.StatusOK)
    _, err = io.WriteString(response, xmlContent)
    if err != nil {
        log.Error("Failed to write XML feed content to HTTP response: ", err)
    }
}

func getAtomFeed(response http.ResponseWriter, request *http.Request) {
    var feed, err = generateFeedObject(response, request)
    if err != nil {
        log.Warn("HTTP request for Atom feed failed: ", err)
        return
    }

    xmlContent, err := feed.ToAtom()
    if err != nil {
        log.Warn("HTTP request for Atom feed failed: ", err)
    }

    response.Header().Set("Content-Type", "text/xml")
    response.WriteHeader(http.StatusOK)
    _, err = io.WriteString(response, xmlContent)
    if err != nil {
        log.Error("Failed to write XML feed content to HTTP response: ", err)
    }
}

func generateFeedObject(response http.ResponseWriter, request *http.Request) (*feeds.Feed, error) {
    var params = mux.Vars(request)
    var reqIdStr = params["id"]
    log.Info("HTTP Get for feed: id=[", reqIdStr, "]: ", request.RequestURI)

    if reqIdStr == "" {
        response.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for feed, without any ID")
    }

    var reqId, err = strconv.Atoi(reqIdStr)
    if err != nil {
        response.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for feed, with invalid ID")
    }

    var comicDef, found = config.GetComicDefForId(reqId)
    if !found {
        response.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for RSS 2.0 feed, with unknown ID")
    }

    var issueList = store.GetIssueListForComicId(reqId)

    feed := &feeds.Feed{
        Title:       comicDef.Name,
        Link:        &feeds.Link{Href: comicDef.Url},
        Description: comicDef.Description,
        Author:      &feeds.Author{Name: comicDef.Author},
        Created:     time.Now(),
    }

    if issueList != nil {
        feed.Items = make([]*feeds.Item, issueList.LinkCount)
        var idx = 0
        for linkCursor := issueList.FirstLink; linkCursor != nil; linkCursor = linkCursor.NextLink {
            var newItem feeds.Item
            newItem.Id = linkCursor.Hash
            newItem.Title = linkCursor.Issue.Title
            newItem.Link = &feeds.Link{Href: linkCursor.Issue.IssueUrl}
            newItem.Description = calculateDescriptionForFeedItem(linkCursor)
            newItem.Created = linkCursor.Issue.Time
            feed.Items[idx] = &newItem
            idx++
        }
    }

    return feed, nil
}

func calculateDescriptionForFeedItem(issueLink *model.IssueLink) string {
    var imageUrl string
    if issueLink.ProxyImage {
        imageUrl = issueLink.ProxyImageUrl
    } else {
        imageUrl = issueLink.Issue.ImageUrl
    }

    var imageHtmlContent = `<a href="%s" title="%s"><img src="%s" title="%s" alt="%s"/></a>`
    return fmt.Sprintf(imageHtmlContent, imageUrl, issueLink.Issue.Title, imageUrl, issueLink.Issue.Title, issueLink.Issue.Title)
}
