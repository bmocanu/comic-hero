package main

import (
    _ "comic-hero/retrieve"
    "comic-hero/serve"
    _ "comic-hero/store"
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetLevel(log.InfoLevel)

    // retrieve.init() 	=> sets up the cron tab and schedules the periodical comic scanning
    // store.init() 	=> sets up the issue storage
    // serve.init()		=> sets up the HTTP interface + controllers

    //for true {
    //    time.Sleep(10 * time.Second)
    //    retrieve.FetchNewIssues()
    //}

    serve.StartServing()
}
