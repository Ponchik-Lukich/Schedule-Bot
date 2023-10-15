package telegram

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/errors"
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdate(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository, roomRepo room.Repository) {
	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	state, err := userRepo.GetUserState(update.Message.Chat.ID)
	if err != nil {
		errors.LogError(errors.ErrorGettingUserState, err)
		return
	}

	switch update.Message.Text {
	case cst.Search:
		responseText, replyMarkup, err = HandleSearch(userRepo, update.Message.Chat.ID, "search")
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
	case cst.Info:
		responseText, replyMarkup, err = HandleSearch(userRepo, update.Message.Chat.ID, "info")
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
	case cst.Back:
		msg, err := HandleBack(userRepo, update.Message.Chat.ID, state)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		update.Message.Text = msg
		HandleUpdate(botApi, update, userRepo, roomRepo)
	case cst.Menu:
		responseText, replyMarkup, err = HandleBackToMenu(userRepo, update.Message.Chat.ID)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
	default:
		switch state {
		case "info":
		case "info_number":
		case "search":
		case "search_date":
		case "search_time":
		case "finish":
		default:
			responseText = cst.CantUnderstand
			replyMarkup = keyboard.CreateMainKeyboard()
		}
	}

	err = functions.SendMessage(botApi, replyMarkup, update.Message.Chat.ID, responseText)
	if err != nil {
		errors.LogError(errors.ErrorSendingMessage, err)
		return
	}
}
