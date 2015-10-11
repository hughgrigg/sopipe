package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	anaconda.SetConsumerKey(viper.GetString("accounts.twitter.consumerKey"))
	anaconda.SetConsumerSecret(viper.GetString("accounts.twitter.secretKey"))
	twitter := anaconda.NewTwitterApi(
		viper.GetString("accounts.twitter.accessToken"),
		viper.GetString("accounts.twitter.accessTokenSecret"),
	)

	tweet, _ := twitter.PostTweet("", nil)
	fmt.Println(tweet.IdStr)
}
