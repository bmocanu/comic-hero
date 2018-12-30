package serve

import (
    "comic-hero/config"
    "fmt"
    log "github.com/sirupsen/logrus"
    "io"
    "net/http"
    "sort"
    "strconv"
)

var pagePrefix = `
<html>
    <head>
        <title>Comic Hero</title>
        <link rel="stylesheet" href="%s"/>
    </head>
    <body>
        <h1>Comic Hero - RSS comic streamer</h1>
        <div id="comicsContainer">
`

var feedDiv = `
            <div class="row">
                <div class="title"><a href="%s">%s</a></div>
                <div class="link"><a href="%s">Atom feed</a></div>
                <div class="link"><a href="%s">RSS 2.0 feed</a></div>
            </div>
`

var pageSuffix = `
        </div>
    </body>
</html>
`

func getFeedList(w http.ResponseWriter, r *http.Request) {
    log.Info("HTTP Get for feed list HTML page: ", r.RequestURI)
    var contextPath = config.Server.ContextPath
    var pageContent = fmt.Sprintf(pagePrefix, concat(contextPath, "/css"))

    var sortedComicNames = make([]string, len(config.ComicDefs))
    var index = 0
    for comicName := range config.ComicDefs {
        sortedComicNames[index] = comicName
        index++
    }
    sort.Strings(sortedComicNames)

    for _, comicName := range sortedComicNames {
        var comicDef = config.ComicDefs[comicName]
        pageContent += fmt.Sprintf(feedDiv,
            comicDef.Url,
            comicDef.Name,
            concat(contextPath, "/feed/atom/"+strconv.Itoa(comicDef.Id)),
            concat(contextPath, "/feed/rss/"+strconv.Itoa(comicDef.Id)))
    }

    pageContent += pageSuffix

    w.Header().Set("Content-Type", "text/html")
    _, err := io.WriteString(w, pageContent)
    if err != nil {
        log.Error("Failed to write feed page HTML content to HTTP response: ", err)
    }
}
