package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func init() {
	fmt.Println("Inside config package initialization")
}

/*
Config is the structure used for loading application configuration from an external JSON file
*/
type Config struct {
	Server ServerConfig  `json:"server"`
	Comics []ComicConfig `json:"comics"`
}

/*
ServerConfig is the structure used for loading the server runtime configuration.
This struct is aggregated into Config.
*/
type ServerConfig struct {
	Address     string `json:"address"`
	Port        int    `json:"port"`
	ContextPath string `json:"contextPath"`
}

/*
ComicConfig is the structure used for loading the configuration of each comic that is supported by
this application. This struct is aggregated into Config.
*/
type ComicConfig struct {
	Name           string `json:"name"`
	Enabled        bool   `json:"enabled"`
	ProxyImage     bool   `json:"proxyImage"`
	CheckFrequency string `json:"checkFrequency"`
}

/*
LoadConfiguration unmarshalls the given JSON file content and produces a completely populated Config
structure, containing all the application configuration.
*/
func LoadConfiguration(file string) Config {
	var config Config
	configFile, error := os.Open(file)
	defer configFile.Close()
	if error != nil {
		fmt.Println(error.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
