package functions

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessage(botApi *tgbotapi.BotAPI, buttons tgbotapi.ReplyKeyboardMarkup, chatID int64, text string) error {
	buttons.ResizeKeyboard = true
	buttons.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = buttons
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := botApi.Send(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SendPhoto(botApi *tgbotapi.BotAPI, chatID int64, imagePath string) error {
	photo := tgbotapi.NewPhotoUpload(chatID, imagePath)
	_, err := botApi.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
