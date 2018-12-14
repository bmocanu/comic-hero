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
LoadConfiguration unmarshalls the given JSON file content and produces a completely populated Config
structure, containing all the application configuration.
*/
func LoadConfiguration(file string) (*Config, error) {
	var config Config

	configFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
