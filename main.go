package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	_ "./statik"
	"github.com/BurntSushi/toml"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/rakyll/statik/fs"
)

// Config For User Authentication
type Config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

var config Config
var regxNewline = regexp.MustCompile(`\r\n|\r|\n`)

func main() {
	fsStatik, _ := fs.New()
	f, err := fsStatik.Open("/settings.toml")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	tokens := strings.TrimSpace(string(buf))
	_, err = toml.Decode(tokens, &config)
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
