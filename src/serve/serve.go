package serve

import (
    _ "github.com/gorilla/feeds"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
)

var httpHandler http.Handler

func init() {
    var localHandler = mux.NewRouter()
    localHandler.HandleFunc("/", getFeedList).Methods("GET")
    localHandler.HandleFunc("/css", getCss).Methods("GET")
    localHandler.HandleFunc("/feed/rss/{id}", getRss20Feed).Methods("GET")
    localHandler.HandleFunc("/feed/atom/{id}", getAtomFeed).Methods("GET")
    httpHandler = localHandler
}

func StartServing() {
    log.Fatal(http.ListenAndServe(":8080", httpHandler))
}
