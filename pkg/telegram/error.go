package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	errInvalidURL   = errors.New("URL is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

// msg.Text = "You are not authorized! Use command /start to authorize."
// msg.Text = "This is not valid link!"
// msg.Text = "Oops, failed to add a the link. Try again later."

func (b *Bot) handleError(chatID int64, err error) {
	msg := tgbotapi.NewMessage(chatID, "Something went wrong, try again.")

	switch err {
	case errInvalidURL:
		msg.Text = "This is not valid link!"
		b.bot.Send(msg)
	case errUnauthorized:
		msg.Text = "You are not authorized! Use command /start to authorize."
		b.bot.Send(msg)
	case errUnableToSave:
		msg.Text = "Oops, failed to add a the link. Try again later."
		b.bot.Send(msg)
	}
}
