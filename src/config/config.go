package config

import (
    "comic-hero/model"
    log "github.com/sirupsen/logrus"
    "os"
    "strconv"
    "strings"
)

var Server model.ServerConfig

func init() {
    Server.ListenAddress = getStringVar("LISTEN_ADDRESS", true, "")
    Server.ListenPort = getIntVar("LISTEN_PORT", true)
    Server.BaseUrl = getStringVar("BASE_URL", true, "")
    Server.ContextPath = getStringVar("CONTEXT_PATH", false, "")
    // SINFEST=enabled;no-proxy
    // DILBERT=enabled;no-proxy
    // OGLAF=enabled;proxy

    if strings.HasSuffix(Server.BaseUrl, "/") {
        Server.BaseUrl = Server.BaseUrl[:len(Server.BaseUrl)-1]
    }

    log.Info("Config log: ListenAddress=", Server.ListenAddress)
    log.Info("Config log: ListenPort=", Server.ListenPort)
    log.Info("Config log: BaseUrl=", Server.BaseUrl)
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
