package telegram

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	"Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleStart(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository) {
	chatID := update.Message.Chat.ID
	err := userRepo.CreateUser(chatID)
	if err != nil {
		log.Println(err)
	}

	reply := keyboard.CreateMainKeyboard()
	err = functions.SendMessage(botApi, reply, update.Message.Chat.ID, constants.Greeting)
	if err != nil {
		log.Println(err)
	}
}
