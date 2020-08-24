package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/kelseyhightower/envconfig"
)

type TwitterConfig struct {
	ConsumerKey    string `required:"true" split_words:"true"`
	ConsumerSecret string `required:"true" split_words:"true"`
	AccessToken    string `required:"true" split_words:"true"`
	AccessSecret   string `required:"true" split_words:"true"`
}

func main() {
	var text string
	flag.StringVar(&text, "text", "", "tweet text body")
	flag.Parse()

	if text == "" {
		log.Fatalf("please set tweet body")
	}

	fmt.Printf("%+v", text)

	var keys TwitterConfig
	if err := envconfig.Process("t", &keys); err != nil {
		log.Fatalf("failed to process env")
	}

	if keys.ConsumerKey == "" || keys.ConsumerSecret == "" || keys.AccessToken == "" || keys.AccessSecret == "" {
		log.Fatalf("please set twitter credentials with T_ prefix (ex. T_CONSUMERKEY)")
	}

	config := oauth1.NewConfig(keys.ConsumerKey, keys.ConsumerSecret)
	token := oauth1.NewToken(keys.AccessToken, keys.AccessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	client.Statuses.Update(text, nil)
}
