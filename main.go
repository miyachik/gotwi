package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// For User Authentication
type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

var config Config

func main() {
	data, err := Asset("settings.local.toml")
	if err != nil {
		panic(err)
	}

	_, err = toml.Decode(string(data), &config)
	if err != nil {
		fmt.Println(err)
	}

	if config.ConsumerKey == "" || config.ConsumerSecret == "" || config.AccessToken == "" || config.AccessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Update (POST!) Tweet (uncomment to run)
	flag.Parse()
	client.Statuses.Update(flag.Args()[0], nil)
}
