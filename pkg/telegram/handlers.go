package telegram

import (
	"context"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	commandStart           = "start"
	replyStartTemplate     = "Hey! To save links to your Pocket account, first you need to give an access to it. To do this, follow the link:\n%s"
	replyAlreadyAuthorized = "You have already been authorized. Now you can use Pocket Bot. Send some links :)"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Link is saved to your Pocket.")

	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		msg.Text = "This is not valid link!"
		_, err = b.bot.Send(msg)
		return err
	}

	accessToken, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		msg.Text = "You are not authorized! Use command /start to authorize."
		_, err = b.bot.Send(msg)
		return err
	}

	if err := b.pocketClient.Add(context.Background(), pocket.AddInput{
		AccessToken: accessToken,
		URL:         message.Text,
	}); err != nil {
		msg.Text = "Oops, failed to add a the link. Try again later."
		_, err = b.bot.Send(msg)
		return err
	}

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	_, err := b.getAccessToken(message.Chat.ID)
	if err != nil {
		return b.initAuthProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, replyAlreadyAuthorized)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "No such command :(")
	_, err := b.bot.Send(msg)
	return err
}
