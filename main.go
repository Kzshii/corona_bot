package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

func getToken() string {
	viper.SetConfigName("secret")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}

	return viper.Get("BOT_TOKEN").(string)
}

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  getToken(),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Authorized on account %s", bot.Me.Username)

	initBot(bot)
}
