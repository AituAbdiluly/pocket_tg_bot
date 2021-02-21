package main

import (
	"log"

	"github.com/AituAbdiluly/pocket_tg_bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	// bot init
	bot, err := tgbotapi.NewBotAPI("1603407543:AAHmGyqZcQaGMmKIabQ38qbIi63JMF694xU")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// pocket client init
	pocketClient, err := pocket.NewClient("95913-040972201e73766f6dd110a4")
	if err != nil {
		log.Fatal(err)
	}

		telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost/")
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
}