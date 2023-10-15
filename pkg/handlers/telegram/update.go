package telegram

import (
	"Telegram/pkg/bot/keyboard"
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleUpdate(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository, roomRepo room.Repository) {

	if update.Message == nil {
		return
	}

	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	switch update.Message.Text {
	case "Поиск свободного кабинета":
		responseText = "Here's the info for free rooms..."
		replyMarkup = keyboard.CreateMiniKeyboard("Назад")
	case "Информация о кабинете":
		responseText = "Here's the room information..."
		replyMarkup = keyboard.CreateMiniKeyboard("Назад")
	case "Назад":
		responseText = "Choose an option:"
		replyMarkup = keyboard.CreateMainKeyboard()
	default:
		responseText = "Choose an option:"
		replyMarkup = keyboard.CreateMainKeyboard()
	}

	replyMarkup.ResizeKeyboard = true
	replyMarkup.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	msg.ReplyMarkup = replyMarkup

	_, err := botApi.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
