package main


import (
	"os"
	"fmt"
	"log"
	"strings"
	"strconv"
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
	log.SetPrefix("[TWITTER-BOT]$: ")

	UserId, err := GetUser(*api)
	if err != nil{
		log.Fatalf("cannot get user %v", err)
	}

	FollowedUsersId := make([]string, 0)
	c := api.GetFriendsIdsAll(url.Values{
		"user_id": []string{fmt.Sprintf("%d", UserId)},
	})

	for fpid := range c {
		for _, id := range fpid.Ids{
			FollowedUsersId = append(FollowedUsersId, strconv.FormatInt(id, 10))
		}
	}
	ids := strings.Join(FollowedUsersId, ",")

	fmt.Printf("got %d friends\n", len(FollowedUsersId))

	stream :=  api.PublicStreamFilter(url.Values{
		"follow": []string{ids},
	})

	defer stream.Stop()
	for v := range stream.C {
		tw, ok := v.(anaconda.Tweet)
		if !ok {
			log.Printf("cannot conveted to anaconda.Tweet, received: %T", v)
			continue
		}

		log.Printf("[reciving tweet: %s from %s. Time: %s, Following: %t, userid: %d]\n", tw.Text, tw.User.Name, tw.CreatedAt, tw.User.Following, tw.User.Id)
	}


}

func GetUser(api anaconda.TwitterApi) (int64, error) {
	username := "dbmak2"
	users, err := api.GetUsersLookup(username,nil)
	if err != nil {
		return 0, err
	}
	for _, user := range users {
		fmt.Printf("Found id %d - %s - %s - %s", user.Id, user.URL, user.Description, user.Location)
	}
	return users[0].Id, nil
}
func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}