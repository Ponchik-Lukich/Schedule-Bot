package commands

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/errors"
	"Telegram/pkg/handlers/telegram/states"
	"Telegram/pkg/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleUpdateCom(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, repos repo.Repositories) {
	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	next, photo := false, false

	state, err := repos.GetUserRepo().GetUserState(update.Message.Chat.ID)
	if err != nil {
		errors.LogError(errors.ErrorGettingUserState, err)
		return
	}

	switch update.Message.Text {
	case cst.Search:
		responseText, err = HandleSearchCom(update.Message.Chat.ID, "search", repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard()
		photo = true
	case cst.Info:
		responseText, err = HandleSearchCom(update.Message.Chat.ID, "info", repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard()
		photo = true
	case cst.Back:
		msg, err := HandleBackCom(update.Message.Chat.ID, state, repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		update.Message.Text = msg
		HandleUpdateCom(botApi, update, repos)
	case cst.Menu:
		responseText, err = HandleMenuCom(update.Message.Chat.ID, repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
	default:
		switch state {
		case "info":
			responseText, next, err = states.HandleInfoState(update.Message.Chat.ID, update.Message.Text, repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateBuildingKeyboard()
			} else {
				replyMarkup = keyboard.CreateMiniKeyboard(cst.Back)
			}
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

	if photo {
		if err = functions.SendPhoto(botApi, update.Message.Chat.ID, cst.MapPath); err != nil {
			errors.LogError(errors.ErrorSendingMessage, err)
			return
		}
	}

	if err = functions.SendMessage(botApi, replyMarkup, update.Message.Chat.ID, responseText); err != nil {
		errors.LogError(errors.ErrorSendingMessage, err)
		return
	}
}
