package serve

import (
    "comic-hero/config"
    "fmt"
    log "github.com/sirupsen/logrus"
    "html"
    "io"
    "math/rand"
    "net/http"
    "sort"
    "strconv"
)

var pagePrefix = `
<html>
<head>
    <title>Comic Hero</title>
    <link rel="shortcut icon" type="image/png" href="%s"/>
    <link rel="icon" type="image/png" href="%s">
    <link rel="stylesheet" href="%s"/>
</head>
<body>
<a href="https://github.com/bmocanu/comic-hero"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_orange_ff7600.png" alt="Fork me on GitHub"/></a>
<h1>comic-hero <span class="subH1">rss streamer</span></h1>
<div class="comicsTable">
`

var feedDiv = `
    <div class="row">
        <div class="title"><a href="%s">%s</a></div>
        <div class="description">%s</div>
        <div class="rss20FeedLink"><a href="%s">RSS 2.0 feed</a></div>
        <div class="atomFeedLink"><a href="%s">Atom feed</a></div>
    </div>
`

var pageSuffix = `
</div>
<div id="versionContainer">
    <a href="https://github.com/bmocanu/comic-hero">comic-hero</a>, <span class="version">version %s, released at %s</span>
</div>
</body>
</html>
`

func getFeedList(response http.ResponseWriter, request *http.Request) {
    log.Info("HTTP Get for feed list HTML page: ", request.RequestURI)
    var contextPath = config.Server.ContextPath
    var cssRandomUrl = urlConcat(contextPath, fmt.Sprintf("/css?rnd=%d", rand.Intn(1000000)))
    var favIconUrl = urlConcat(contextPath, "/favicon")
    var pageContent = fmt.Sprintf(pagePrefix, favIconUrl, favIconUrl, cssRandomUrl)

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
            html.EscapeString(comicDef.Name),
            html.EscapeString(comicDef.Description),
            urlConcat(contextPath, "/feed/rss/"+strconv.Itoa(comicDef.Id)),
            urlConcat(contextPath, "/feed/atom/"+strconv.Itoa(comicDef.Id)))
    }

    pageContent += fmt.Sprintf(pageSuffix,
        html.EscapeString(config.AppVersion),
        html.EscapeString(config.AppReleaseDate))

    setContentTypeHeader(response, "text/html")
    _, err := io.WriteString(response, pageContent)
    if err != nil {
        log.Error("Failed to write feed page HTML content to HTTP response: ", err)
        response.WriteHeader(http.StatusInternalServerError)
    }
}
