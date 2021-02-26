package main

import (
	"log"

	"github.com/AituAbdiluly/pocket_tg_bot/pkg/repository"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/repository/boltdb"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/server"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/telegram"
	"github.com/boltdb/bolt"
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

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")

	authServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/pocket_go_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
