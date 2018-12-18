package main

import (
    log "github.com/sirupsen/logrus"
    "time"
)

func main() {
    log.SetLevel(log.InfoLevel)

    // retrieve.init() 	=> sets up the cron tab and schedules the periodical comic scanning
    // store.init() 	=> sets up the issue storage
    // serve.init()		=> sets up the HTTP interface + controllers

    time.Sleep(10 * time.Minute)
}
