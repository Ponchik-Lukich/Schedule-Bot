package functions

import (
	"Telegram/pkg/errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessage(botApi *tgbotapi.BotAPI, replyMarkup tgbotapi.ReplyKeyboardMarkup, chatID int64, text string) error {
	replyMarkup.ResizeKeyboard = true
	replyMarkup.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = replyMarkup

	_, err := botApi.Send(msg)
	if err != nil {
		return fmt.Errorf("%s: %v", errors.ErrorSendingMessage, err)
	}
	return nil
}
