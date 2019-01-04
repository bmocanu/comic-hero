package serve

import (
    log "github.com/sirupsen/logrus"
    "io"
    "net/http"
)

var cssContent = `
body {
    font-family: Arial, sans-serif;
    font-size: 14pt;
    text-align: center;
}

h1 {
    font-size: 22pt;
}

#comicsContainer {
    background-color: #DDD;
    display: table;
    border: solid 1px black;
    margin-left: auto;
    margin-right: auto;
}

#comicsContainer .row {
    display: table-row;
}

#comicsContainer .row .title {
    padding: 5px;
    width: 400px;
    display: table-cell;
    text-align: left;
}

#comicsContainer .row .link {
    padding: 5px;
    width: 120px;
    display: table-cell;
}

#comicsContainer .row .link a {
    color: brown;
    text-decoration: none;
}

#comicsContainer .row .link a:hover {
    text-decoration: underline;
}

.enabled {
    background-color: #f0ffe5;
}

.disabled {
    background-color: #ffe5e5;
}
`

func getCss(response http.ResponseWriter, request *http.Request) {
    log.Info("HTTP Get for CSS page: ", request.RequestURI)
    response.Header().Set("Content-Type", "text/css")
    _, err := io.WriteString(response, cssContent)
    if err != nil {
        log.Error("Failed to write CSS content to HTTP response: ", err)
    }
}
