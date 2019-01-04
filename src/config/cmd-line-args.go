package config

import (
    "comic-hero/model"
    "flag"
    log "github.com/sirupsen/logrus"
    "os"
    "path/filepath"
    "strings"
)

func parseCommandLineArgs() {
    var configDirFlag = flag.String("config-dir", "config", "The folder containing the configuration of comic-hero")
    var configFileFlag = flag.String("config-file", "config.json", "The main configuration file of the application")
    flag.Parse()

    appDir = getAppStartDir()
    configDir = *configDirFlag
    configFile = *configFileFlag

    if !strings.HasPrefix(configDir, "/") {
        configDir = filepath.Join(appDir, configDir)
    }
    configFile = filepath.Join(configDir, configFile)

    log.Info("Start params: App dir: ", appDir)
    log.Info("Start params: Config dir: ", configDir)
    log.Info("Start params: Config file: ", configFile)
}

func getAppStartDir() string {
    ex, err := os.Executable()
    if err != nil {
        log.Fatal("Internal system error. Cannot get app startup dir: ", err)
        os.Exit(model.ExitInternalSystemError)
    }
    return filepath.Dir(ex)
}
