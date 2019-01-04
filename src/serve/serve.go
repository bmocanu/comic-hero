package serve

import (
    "comic-hero/config"
    _ "github.com/gorilla/feeds"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
    "strings"
)

var httpHandler http.Handler

func init() {
    var contextPath = config.Server.ContextPath
    var localHandler = mux.NewRouter()
    localHandler.HandleFunc(concat(contextPath, ""), getFeedList).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/"), getFeedList).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/css"), getCss).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/css/"), getCss).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/feed/rss/{id}"), getRss20Feed).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/feed/rss/{id}/"), getRss20Feed).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/feed/atom/{id}"), getAtomFeed).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/feed/atom/{id}/"), getAtomFeed).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/get/{id}/{hash}"), getProxyImage).Methods("GET")
    localHandler.HandleFunc(concat(contextPath, "/get/{id}/{hash}/"), getProxyImage).Methods("GET")
    httpHandler = localHandler
}

func StartServing() {
    var addressAndPort = config.Server.ListenAddress + ":" + strconv.Itoa(config.Server.ListenPort)
    log.Info("Listening for HTTP requests on: ", addressAndPort, ": contextPath: ", config.Server.ContextPath)
    log.Fatal(http.ListenAndServe(addressAndPort, httpHandler))
}

func concat(part1 string, part2 string) string {
    var path string
    if part2 == "" {
        path = part1
    } else if part1 == "" {
        path = part2
    } else if !strings.HasSuffix(part1, "/") && !strings.HasPrefix(part2, "/") {
        path = part1 + "/" + part2
    } else {
        path = part1 + part2
    }
    return strings.Replace(path, "//", "/", -1)
}

