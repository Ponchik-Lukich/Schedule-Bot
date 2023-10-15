package telegram

import (
	"Telegram/pkg/bot/keyboard"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleStart(botAPI *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository) {
	chatID := update.Message.Chat.ID
	err := userRepo.CreateUser(chatID)
	if err != nil {
		log.Println(err)
	}

	msg := tgbotapi.NewMessage(chatID, "Привет! Я бот, который поможет тебе найти свободный кабинет в университете.")
	msg.ReplyMarkup = keyboard.CreateMainKeyboard()

	_, err = botAPI.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
