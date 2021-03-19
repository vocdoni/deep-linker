package main

import (
	"fmt"
	"net/url"
)

// MakeLink generates a Dynamic link containing the parameters for all the supported use cases
func MakeLink(mainURL string, config Config) string {
	link := url.PathEscape(mainURL)
	title := url.PathEscape(config.Social.Title)
	description := url.PathEscape(config.Social.Description)
	image := url.PathEscape(config.Social.Image)

	return fmt.Sprintf("https://%s/?link=%s&apn=%s&amv=%d&ibi=%s&ius=%s&isi=%s&ivm=%s&st=%s&sd=%s&si=%s",
		config.Redirect.Domain,
		link,
		config.Android.Package,
		config.Android.MinBuild,
		config.Ios.BundleID,
		config.Ios.Scheme,
		config.Ios.AppStoreID,
		config.Ios.MinVersion,
		title,
		description,
		image,
	)
}

// MakeEntityLink returns a Dynamic link pointing to a Vocdoni entity, targeting the mobile app and fallback web browsers
func MakeEntityLink(entityID string, config Config) string {
	webURL := fmt.Sprintf("%s/entities/#/%s", config.Link.Entities.Prefix, entityID)
	return MakeLink(webURL, config)
}

// MakeProcessLink returns a Dynamic link pointing to a Vocdoni Process, targeting the mobile app and fallback web browsers
func MakeProcessLink(entityID, processID string, config Config) string {
	webURL := fmt.Sprintf("%s/processes/#/%s/%s", config.Link.Processes.Prefix, entityID, processID)
	return MakeLink(webURL, config)
}

// MakePostLink returns a Dynamic link pointing to an entity's post, targeting the mobile app and fallback web browsers
func MakePostLink(entityID, postIdx string, config Config) string {
	webURL := fmt.Sprintf("%s/posts/#/%s/%s", config.Link.Posts.Prefix, entityID, postIdx)
	return MakeLink(webURL, config)
}

// MakeRegistryValidationLink returns a Dynamic link targeting the app to prompt for a registration confirmation
func MakeRegistryValidationLink(entityID, validationToken string, config Config) string {
	webURL := fmt.Sprintf("%s/validation/#/%s/%s", config.Link.Validation.Prefix, entityID, validationToken)
	return MakeLink(webURL, config)
}

// MakeRecoveryLink returns a Dynamic link targeting the manager or the app to recover an account
func MakeRecoveryLink(name, date, payload string, config Config) string {
	webURL := fmt.Sprintf("%s/recovery/#/%s/%s/%s", config.Link.Recovery.Prefix, name, date, payload)
	return MakeLink(webURL, config)
}
