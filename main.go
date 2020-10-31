package main


import (
	"os"
	"fmt"
	"net/url"
	"github.com/ChimeraCoder/anaconda"
)


var (
	TWITTER_API_SECRET=getenv("TWITTER_API_SECRET")
	TWITTER_API_SECRET_KEY=getenv("TWITTER_API_SECRET_KEY")
	TWITTER_API_BEARER_TOKEN=getenv("TWITTER_API_BEARER_TOKEN")
	TWITTER_API_ACCESS_TOKEN=getenv("TWITTER_API_ACCESS_TOKEN")
	TWITTER_API_ACCESS_TOKEN_SECRET=getenv("TWITTER_API_ACCESS_TOKEN_SECRET")
)

func main() {
	anaconda.SetConsumerKey(TWITTER_API_SECRET)
	anaconda.SetConsumerSecret(TWITTER_API_SECRET_KEY)
	api := anaconda.NewTwitterApi(TWITTER_API_ACCESS_TOKEN, TWITTER_API_ACCESS_TOKEN_SECRET)
 
	stream :=  api.PublicStreamFilter(url.Values{
		"track": []string{"#love"},
	})

	defer stream.Stop()

	for v := range stream.C {
		tw, ok := v.(anaconda.Tweet)
		if !ok {
			fmt.Printf("cannot conveted to anaconda.Tweet: %v, received type: %T", ok, v)
			continue
		}
		fmt.Printf("Reciving id %d tweet: %s from %s\n", tw.Id, tw.Text, tw.User.Name)
	}
}


func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}