package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config holds the general parameters of the service
type Config struct {
	HTTP struct {
		Port int `toml:"port"`
	} `toml:"http"`
	Redirect struct {
		Domain string `toml:"domain"`
	} `toml:"redirect"`
	Link struct {
		Prefix   string `toml:"prefix"`
		Fallback string `toml:"fallback"`
	} `toml:"link"`
	Android struct {
		Package  string `toml:"package"`
		Scheme   string `toml:"scheme"`
		MinBuild uint   `toml:"defaultMinBuild"`
	} `toml:"android"`
	Ios struct {
		BundleID   string `toml:"bundleId"`
		Scheme     string `toml:"scheme"`
		AppStoreID string `toml:"appStoreId"`
		MinVersion string `toml:"defaultMinVersion"`
	} `toml:"ios"`
	Social struct {
		Title       string `toml:"title"`
		Description string `toml:"description"`
		Image       string `toml:"image"`
	} `toml:"social"`
	Misc struct {
		Verbose bool `toml:"verbose"`
	} `toml:"misc"`
}

// ReadConfig parses `config.toml` and returns a struct with the desired config
func ReadConfig(path string) (Config, error) {
	conf := Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	if _, err := toml.Decode(string(data), &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
