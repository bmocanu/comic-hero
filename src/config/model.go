package config

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
    Id             int    `json:"id"`
    Name           string `json:"name"`
    Enabled        bool   `json:"enabled"`
    ProxyImage     bool   `json:"proxyImage"`
}
