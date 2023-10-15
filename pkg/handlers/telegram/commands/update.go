package commands

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/errors"
	"Telegram/pkg/repo/room"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdateCom(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, userRepo user.Repository, roomRepo room.Repository) {
	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	state, err := userRepo.GetUserState(update.Message.Chat.ID)
	if err != nil {
		errors.LogError(errors.ErrorGettingUserState, err)
		return
	}

	switch update.Message.Text {
	case cst.Search:
		responseText, err = HandleSearchCom(userRepo, update.Message.Chat.ID, "search")
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard()
	case cst.Info:
		responseText, err = HandleSearchCom(userRepo, update.Message.Chat.ID, "info")
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard()
	case cst.Back:
		msg, err := HandleBackCom(userRepo, update.Message.Chat.ID, state)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		update.Message.Text = msg
		HandleUpdateCom(botApi, update, userRepo, roomRepo)
	case cst.Menu:
		responseText, err = HandleMenuCom(userRepo, update.Message.Chat.ID)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
	default:
		switch state {
		case "info":
			// wait for building name
			// change state to info_number
			replyMarkup = keyboard.CreateMiniKeyboard(cst.Back)
		case "info_number":
			// wait for room number
			// change state to finish
		case "search":
			// wait for building name
			// change state to search_date
		case "search_date":
			// wait for date
			// change state to search_time
		case "search_time":
			// wait for date
			// change state to search_time
		case "finish":
			responseText = cst.CantUnderstand
			replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
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
