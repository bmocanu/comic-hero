package main

import (
    _ "comic-hero/config"
    _ "comic-hero/retrieve"
    "comic-hero/serve"
    _ "comic-hero/store"
    log "github.com/sirupsen/logrus"
)

func main() {
    log.SetLevel(log.InfoLevel)
    // config.init() 	=> sets up the configuration based on env vars
    // retrieve.init() 	=> sets up the cron tab and schedules the periodical comic scanning
    // store.init() 	=> sets up the issue storage
    // serve.init()		=> sets up the HTTP interface + controllers
    serve.StartServing()
}
