package config

import (
    "comic-hero/model"
    log "github.com/sirupsen/logrus"
    "os"
    "strconv"
)

var Server model.ServerConfig

const AppVersion = "1.1"

func init() {
    Server.ListenAddress = getStringVar("LISTEN_ADDRESS", true, "")
    Server.ListenPort = getIntVar("LISTEN_PORT", true)
    Server.ContextPath = getStringVar("CONTEXT_PATH", true, "")

    // SINFEST=enabled;no-proxy
    // DILBERT=enabled;no-proxy
    // OGLAF=enabled;proxy

    log.Info("--------------------------------------------------")
    log.Info("comic-hero RSS streamer | ver. ", AppVersion)
    log.Info("--------------------------------------------------")
    log.Info("Config log: ListenAddress=", Server.ListenAddress)
    log.Info("Config log: ListenPort=", Server.ListenPort)
    log.Info("Config log: ContextPath=", Server.ContextPath)
}

func getStringVar(name string, mandatory bool, defaultValue string) string {
    var value = os.Getenv(name)
    if mandatory && value == "" {
        log.Fatal("The " + name + " variable is not set")
        os.Exit(1)
    }

    if value == "" {
        return defaultValue
    } else {
        return value
    }
}

func getIntVar(name string, mandatory bool) int {
    var value = os.Getenv(name)
    if mandatory && value == "" {
        log.Fatal("The " + name + " variable is not set")
        os.Exit(1)
    }

    var intValue, err = strconv.Atoi(value)
    if err != nil {
        log.Fatal("The " + name + " variable is not a number")
        os.Exit(1)
    }

    return intValue
}
