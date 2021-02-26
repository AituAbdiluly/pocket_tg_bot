package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart           = "start"
	replyStartTemplate     = "Hey! To save links to your Pocket account, first you need to give an access to it. To do this, follow the link:\n%s"
	replyAlreadyAuthorized = "You have already been authorized. Now you can use Pocket Bot. Send some links :)"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
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
