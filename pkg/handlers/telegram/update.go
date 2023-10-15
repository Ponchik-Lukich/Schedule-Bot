package telegram

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	"Telegram/pkg/errors"
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleUpdate(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository, roomRepo room.Repository) {
	if update.Message == nil {
		return
	}

	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	state, err := userRepo.GetUserState(update.Message.Chat.ID)
	if err != nil {
		log.Println(fmt.Sprintf("%s: %v", errors.ErrorGettingUserState, err))
	}

	switch update.Message.Text {
	case "Поиск свободного кабинета":
		// handle search
	case "Информация о кабинете":
		// handle info
	case "Назад":
		// handle back
	case "Вернуться в главное меню":
		responseText = "Выберите функцию:"
		replyMarkup = keyboard.CreateMainKeyboard()
		err := userRepo.SetUserState(update.Message.Chat.ID, "wait")
		if err != nil {
			log.Println(fmt.Sprintf("%s: %v", errors.ErrorSettingUserState, err))
		}
	default:
		switch state {
		case "wait":
			responseText = "Извини, я не понимаю тебя."
		}
	}

	err = functions.SendMessage(botApi, replyMarkup, update.Message.Chat.ID, responseText)
	if err != nil {
		log.Println(err)
	}
}
