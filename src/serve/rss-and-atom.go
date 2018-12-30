package serve

import (
    "comic-hero/config"
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

const imageHtmlContent = `<img src="%s" title="%s" alt="%s" />`

func getRss20Feed(w http.ResponseWriter, r *http.Request) {
    var feed, err = generateFeedObject(w, r)
    if err != nil {
        log.Warn("HTTP request for RSS 2.0 feed failed: ", err)
        return
    }

    xmlContent, err := feed.ToRss()
    if err != nil {
        log.Warn("HTTP request for RSS 2.0 feed failed: ", err)
    }

    w.Header().Set("Content-Type", "text/xml")
    w.WriteHeader(http.StatusOK)
    _, err = io.WriteString(w, xmlContent)
    if err != nil {
        log.Error("Failed to write XML feed content to HTTP response: ", err)
    }
}

func getAtomFeed(w http.ResponseWriter, r *http.Request) {
    var feed, err = generateFeedObject(w, r)
    if err != nil {
        log.Warn("HTTP request for Atom feed failed: ", err)
        return
    }

    xmlContent, err := feed.ToAtom()
    if err != nil {
        log.Warn("HTTP request for Atom feed failed: ", err)
    }

    w.Header().Set("Content-Type", "text/xml")
    w.WriteHeader(http.StatusOK)
    _, err = io.WriteString(w, xmlContent)
    if err != nil {
        log.Error("Failed to write XML feed content to HTTP response: ", err)
    }
}

func generateFeedObject(w http.ResponseWriter, r *http.Request) (*feeds.Feed, error) {
    var params = mux.Vars(r)
    var reqIdStr = params["id"]
    log.Info("HTTP Get for feed: id=[", reqIdStr, "]: ", r.RequestURI)

    if reqIdStr == "" {
        w.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for feed, without any ID")
    }

    var reqId, err = strconv.Atoi(reqIdStr)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for feed, without any ID")
    }

    var comicDef, found = config.GetComicDefForId(reqId)
    if !found {
        w.WriteHeader(http.StatusNotFound)
        return nil, fmt.Errorf("HTTP request for RSS 2.0 feed, with unknown ID")
    }

    var issueLink = store.GetIssuesForComicId(reqId)

    feed := &feeds.Feed{
        Title:       comicDef.Name,
        Link:        &feeds.Link{Href: comicDef.Url},
        Description: comicDef.Description,
        Author:      &feeds.Author{Name: comicDef.Author},
        Created:     time.Now(),
    }

    if issueLink != nil {
        feed.Items = make([]*feeds.Item, issueLink.IssueCount)
        var idx = 0
        for linkCursor := issueLink; linkCursor != nil; linkCursor = linkCursor.NextLink {
            var newItem feeds.Item
            newItem.Id = linkCursor.Hash
            newItem.Title = linkCursor.Issue.Title
            newItem.Link = &feeds.Link{Href: linkCursor.Issue.IssueUrl}
            newItem.Description = fmt.Sprintf(imageHtmlContent, linkCursor.Issue.ImageUrl, linkCursor.Issue.Title, linkCursor.Issue.Title)
            newItem.Created = linkCursor.Issue.Time
            feed.Items[idx] = &newItem
            idx++
        }
    }

    return feed, nil
}
