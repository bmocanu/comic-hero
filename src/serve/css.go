package serve

import (
    log "github.com/sirupsen/logrus"
    "io"
    "net/http"
)

var cssContent = `
/* basics ============================================ */
body {
    font-family: Arial, sans-serif;
    font-size: 16px;
    text-align: left;
    padding-left: 15px;
}

body a {
    text-decoration: none;
    color: orangered;
}

body a:hover {
    text-decoration: underline;
}

h1 {
    font-size: 28px;
    margin: 25px 0 20px 0;
}

h1 .subH1 {
    color: #BBB;
}

/* comics table ============================================ */
.comicsTable {
    display: inline-block;
    margin: 0 15px 0 0;
    max-width: 1200px;
    min-width: 500px;
}

.comicsTable .row {
    display: inline-block;
    float: left;
    width: 100%;
    margin-bottom: 2px;
    padding: 3px;
    background-color: #DDD;
    border-radius: 5px;
    border-top: 1px solid #EEE;
    border-left: 1px solid #EEE;
    border-right: 1px solid #999;
    border-bottom: 1px solid #999;
}

.comicsTable .row * {
    float: left;
    padding: 2px;
    vertical-align: middle;
}

.comicsTable .row .title {
    width: calc(15% - 4px);
}

.comicsTable .row .title a {
    color: blue;
}

.comicsTable .row .description {
    width: calc(80% - 16px - 75px - 90px);
    font-size: 90%;
}

.comicsTable .row .atomFeedLink {
    float: right;
    text-wrap: none;
    white-space: nowrap;
    font-size: 90%;
}

.comicsTable .row .rss20FeedLink {
    float: right;
    text-wrap: none;
    font-size: 90%;
}

/* other elements ============================================ */
#versionContainer {
    margin: 20px 0 0 0;
}
`

func getCss(response http.ResponseWriter, request *http.Request) {
    log.Info("HTTP Get for CSS page: ", request.RequestURI)
    setContentTypeHeader(response, "text/css")
    _, err := io.WriteString(response, cssContent)
    if err != nil {
        log.Error("Failed to write CSS content to HTTP response: ", err)
        response.WriteHeader(http.StatusInternalServerError)
    }
}
