package main

import (
	"testing"
)

func TestMakeLink(t *testing.T) {
	// Default config
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		t.Error(err)
	}

	url := "https://vocdoni.io"
	expectedLink := "https://vocdoni.page.link/?link=https:%2F%2Fvocdoni.io&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink := MakeLink(url, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	url = "https://testing.org"
	expectedLink = "https://vocdoni.page.link/?link=https:%2F%2Ftesting.org&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink = MakeLink(url, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	// Custom config
	cfg = makeConfig1()

	url = "https://vocdoni.io"
	expectedLink = "https://testing.page.link/?link=https:%2F%2Fvocdoni.io&apn=org.testing.app&amv=1&ibi=com.testing.app&ius=testing&isi=1505234624&ivm=0.8&st=Testing&sd=Description%20here&si=https:%2F%2Fdomain%2Fimage.png"
	actualLink = MakeLink(url, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}
}

func TestMakeEntityLink(t *testing.T) {
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		t.Error(err)
	}

	entityID := "0x1234"
	expectedLink := "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fentities%2F%23%2F0x1234&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink := MakeEntityLink(entityID, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	entityID = "0x23456"
	expectedLink = "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fentities%2F%23%2F0x23456&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink = MakeEntityLink(entityID, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}
}

func TestMakeProcessLink(t *testing.T) {
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		t.Error(err)
	}

	entityID := "0x1234"
	processID := "0x5678"
	expectedLink := "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fprocesses%2F%23%2F0x1234%2F0x5678&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink := MakeProcessLink(entityID, processID, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	entityID = "0x23456"
	processID = "0x6789"
	expectedLink = "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fprocesses%2F%23%2F0x23456%2F0x6789&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink = MakeProcessLink(entityID, processID, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}
}

func TestMakePostLink(t *testing.T) {
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		t.Error(err)
	}

	entityID := "0x1234"
	postIdx := "100"
	expectedLink := "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fposts%2F%23%2F0x1234%2F100&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink := MakePostLink(entityID, postIdx, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	entityID = "0x23456"
	postIdx = "200"
	expectedLink = "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fposts%2F%23%2F0x23456%2F200&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink = MakePostLink(entityID, postIdx, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}
}

func TestMakeRegistryValidationLink(t *testing.T) {
	cfg, err := ReadConfig("./config.toml")
	if err != nil {
		t.Error(err)
	}

	entityID := "0x1234"
	validationToken := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expectedLink := "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fvalidation%2F%23%2F0x1234%2FABCDEFGHIJKLMNOPQRSTUVWXYZ&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink := MakeRegistryValidationLink(entityID, validationToken, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}

	entityID = "0x23456"
	validationToken = "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	expectedLink = "https://vocdoni.page.link/?link=https:%2F%2Fapp.vocdoni.net%2Fvalidation%2F%23%2F0x23456%2FZYXWVUTSRQPONMLKJIHGFEDCBA&apn=org.vocdoni.app&amv=1&ibi=com.vocdoni.app&ius=vocdoni&isi=1505234624&ivm=0.8&st=Vocdoni&sd=Make%20your%20community%20thrive%20with%20the%20universally%20verifiable%20voting%20system&si=https:%2F%2Fblog.vocdoni.io%2Fcontent%2Fimages%2F2020%2F04%2FSquare_512_alpha-1.png"
	actualLink = MakeRegistryValidationLink(entityID, validationToken, cfg)

	if actualLink != expectedLink {
		t.Errorf("Incorrect link\nWanted: %s\nGot: %s", expectedLink, actualLink)
	}
}

// Helpers

func makeConfig1() Config {
	cfg := Config{}
	cfg.HTTP.Port = 8080
	cfg.Redirect.Domain = "testing.page.link"
	cfg.Link.Prefix = "https://app.testing.net"
	cfg.Link.Fallback = "https://testing.io"

	cfg.Android.Package = "org.testing.app"
	cfg.Android.Scheme = "testing"
	cfg.Android.MinBuild = 1

	cfg.Ios.BundleID = "com.testing.app"
	cfg.Ios.Scheme = "testing"
	cfg.Ios.AppStoreID = "1505234624"
	cfg.Ios.MinVersion = "0.8"

	cfg.Social.Title = "Testing"
	cfg.Social.Description = "Description here"
	cfg.Social.Image = "https://domain/image.png"

	return cfg
}
