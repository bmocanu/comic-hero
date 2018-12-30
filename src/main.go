package main

import (
    _ "comic-hero/config"
    "comic-hero/retrieve"
    _ "comic-hero/retrieve"
    "comic-hero/serve"
    _ "comic-hero/store"
)

func main() {
    // config.init() 	=> sets up the configuration based on env vars
    // retrieve.init() 	=> sets up the cron tab and schedules the periodical comic scanning
    // store.init() 	=> sets up the issue storage
    // serve.init()		=> sets up the HTTP interface + controllers

    // initial retrieving upon startup
    retrieve.FetchNewIssues()

    // start listening for incoming HTTP requests
    serve.StartServing()
}
