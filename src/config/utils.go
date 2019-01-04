package config

import (
    log "github.com/sirupsen/logrus"
    "os"
)

func fileExists(file string) bool {
    if _, err := os.Stat(file); !os.IsNotExist(err) {
        return true
    } else {
        return false
    }
}

func logAppBanner() {
    log.Info("--------------------------------------------------")
    log.Info("comic-hero RSS streamer | ver. ", AppVersion)
    log.Info("--------------------------------------------------")
}

func logConfigState() {
    log.Info("Config log: ListenAddress=", Server.ListenAddress)
    log.Info("Config log: ListenPort=", Server.ListenPort)
    log.Info("Config log: ContextPath=", Server.ContextPath)
    log.Info("Config log: BaseUrl=", Server.BaseUrl)

    var enabledComics string
    var proxyImageComics string
    for key, comicConfig := range comics {
        if comicConfig.Enabled {
            enabledComics = enabledComics + ", " + key
        }
        if comicConfig.ProxyImage {
            proxyImageComics = proxyImageComics + ", " + key
        }
    }

    if enabledComics == "" {
        log.Warn("No comics are enabled in the config file")
    } else {
        log.Info("The following comics are enabled and WILL be retrieved periodically: ", enabledComics[2:])
    }

    if proxyImageComics == "" {
        log.Info("No comics are set for proxy image")
    } else {
        log.Info("The following comics are set to proxy image: ", proxyImageComics[2:])
    }

    log.Info("Config log: Retrieve.IssuesFetchingCronJobConfig: ", Retrieve.IssuesFetchingCronJobConfig)
}
