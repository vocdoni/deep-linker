package main

import (
	"io/ioutil"
	"os"
	"strconv"

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
		Fallback string `toml:"fallback"`
		Entities struct {
			Prefix string `toml:"prefix"`
		} `toml:"entities"`
		Processes struct {
			Prefix string `toml:"prefix"`
		} `toml:"processes"`
		Posts struct {
			Prefix string `toml:"prefix"`
		} `toml:"posts"`
		Validation struct {
			Prefix string `toml:"prefix"`
		} `toml:"validation"`
		Recovery struct {
			Prefix string `toml:"prefix"`
		} `toml:"recovery"`
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

	// Read config.toml
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	if _, err := toml.Decode(string(data), &conf); err != nil {
		return conf, err
	}

	// Environment vars
	if os.Getenv("HTTP_PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
		if err != nil {
			return conf, err
		}
		conf.HTTP.Port = port
	}
	if os.Getenv("REDIRECT_DOMAIN") != "" {
		conf.Redirect.Domain = os.Getenv("REDIRECT_DOMAIN")
	}
	if os.Getenv("LINK_FALLBACK") != "" {
		conf.Link.Fallback = os.Getenv("LINK_FALLBACK")
	}
	if os.Getenv("ANDROID_PACKAGE") != "" {
		conf.Android.Package = os.Getenv("ANDROID_PACKAGE")
	}
	if os.Getenv("IOS_BUNDLE_ID") != "" {
		conf.Ios.BundleID = os.Getenv("IOS_BUNDLE_ID")
	}
	if os.Getenv("SOCIAL_TITLE") != "" {
		conf.Social.Title = os.Getenv("SOCIAL_TITLE")
	}
	if os.Getenv("SOCIAL_DESCRIPTION") != "" {
		conf.Social.Description = os.Getenv("SOCIAL_DESCRIPTION")
	}

	return conf, nil
}
