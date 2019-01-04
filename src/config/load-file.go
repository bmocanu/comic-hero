package config

import (
    "comic-hero/model"
    "encoding/json"
    log "github.com/sirupsen/logrus"
    "os"
)

type jsonModel struct {
    Server struct {
        Address     string `json:"address"`
        Port        int    `json:"port"`
        ContextPath string `json:"contextPath"`
        BaseUrl     string `json:"baseUrl"`
    } `json:"server"`
    Store struct {
        IssuesStoredPerComic int `json:"issuesStoredPerComic"`
    } `json:"store"`
    Retrieve struct {
        IssuesFetchingCronJobConfig string `json:"issuesFetchingCronJobConfig"`
    } `json:"retrieve"`
    Comics []struct {
        Name       string `json:"name"`
        Enabled    bool   `json:"enabled"`
        ProxyImage bool   `json:"proxyImage"`
    } `json:"comics"`
}

func loadConfigFromFile(configFileStr string) error {
    log.Info("Loading config from file: ", configFileStr)
    var config jsonModel

    configFile, err := os.Open(configFileStr)
    if err != nil {
        return err
    }

    defer configFile.Close()

    jsonParser := json.NewDecoder(configFile)
    err = jsonParser.Decode(&config)
    if err != nil {
        return err
    }

    Server.ListenAddress = config.Server.Address
    Server.ListenPort = config.Server.Port
    Server.ContextPath = config.Server.ContextPath
    Server.BaseUrl = config.Server.BaseUrl

    Store.IssuesStoredPerComic = config.Store.IssuesStoredPerComic
    Retrieve.IssuesFetchingCronJobConfig = config.Retrieve.IssuesFetchingCronJobConfig

    for _, comicModel := range config.Comics {
        _, found := GetIdForComicName(comicModel.Name)
        if !found {
            log.Fatal("Invalid comic name in configuration file: ", comicModel.Name)
            os.Exit(model.ExitInvalidComicsConfig)
        }

        var newComicConfig model.ComicConfig
        newComicConfig.Enabled = comicModel.Enabled
        newComicConfig.ProxyImage = comicModel.ProxyImage
        comics[comicModel.Name] = newComicConfig
    }

    return nil
}
