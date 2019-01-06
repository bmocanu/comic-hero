package serve

import (
    "comic-hero/config"
    _ "github.com/gorilla/feeds"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

var httpHandler http.Handler

func init() {
    var contextPath = config.Server.ContextPath
    var localHandler = mux.NewRouter()
    localHandler.HandleFunc(urlConcat(contextPath, ""), getFeedList).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/"), getFeedList).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/css"), getCss).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/favicon"), getFavIcon).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/feed/rss/{id}"), getRss20Feed).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/feed/atom/{id}"), getAtomFeed).Methods("GET")
    localHandler.HandleFunc(urlConcat(contextPath, "/get/{id}/{hash}"), getProxyImage).Methods("GET")
    httpHandler = localHandler
}

func StartServing() {
    var addressAndPort = config.Server.ListenAddress + ":" + strconv.Itoa(config.Server.ListenPort)
    log.Info("Listening for HTTP requests on: ", addressAndPort, ": contextPath: ", config.Server.ContextPath)
    log.Fatal(http.ListenAndServe(addressAndPort, httpHandler))
}
