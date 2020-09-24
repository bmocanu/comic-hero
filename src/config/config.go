package config

import (
    "comic-hero/model"
    log "github.com/sirupsen/logrus"
    "os"
)

const AppVersion = "1.6"
const AppReleaseDate = "Sep/24/2020"

var Server model.ServerConfig
var Store model.StoreConfig
var Retrieve model.RetrieveConfig

var comics = make(map[string]model.ComicConfig)
var appDir string
var configDir string
var configFile string

func init() {
    log.SetLevel(log.InfoLevel)
    log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
    logAppBanner()

    parseCommandLineArgs()

    if !fileExists(configFile) {
        log.Fatal("Cannot find configuration file: ", configFile)
        os.Exit(model.ExitConfigFileNotFound)
    }

    var err = loadConfigFromFile(configFile)
    if err != nil {
        log.Fatal("Cannot load configuration file: ", configFile, ": ", err)
        os.Exit(model.ExitConfigFileNotLoaded)
    }

    logConfigState()
}

func IsComicEnabled(name string) bool {
    for key, comicConfig := range comics {
        if name == key {
            return comicConfig.Enabled
        }
    }
    return false
}

func IsComicImageProxyEnabled(name string) bool {
    for key, comicConfig := range comics {
        if name == key {
            return comicConfig.ProxyImage
        }
    }
    return false
}
