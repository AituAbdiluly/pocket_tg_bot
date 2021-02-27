package main

import (
	"log"

	"github.com/AituAbdiluly/pocket_tg_bot/pkg/config"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/repository"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/repository/boltdb"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/server"
	"github.com/AituAbdiluly/pocket_tg_bot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	// bot init
	bot, err := tgbotapi.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// pocket client init
	pocketClient, err := pocket.NewClient(config.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, config.RedirectURL)

	authServer := server.NewAuthorizationServer(pocketClient, tokenRepository, config.TelegramBotURL)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB(config *Config) (*bolt.DB, error) {
	db, err := bolt.Open(config.DBPath, 0600, nil)
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
