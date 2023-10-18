package commands

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/errors"
	"Telegram/pkg/handlers/telegram/states"
	"Telegram/pkg/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleUpdateCom(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, repos repo.Repositories) {
	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup
	chatID, msgText := update.Message.Chat.ID, update.Message.Text
	next, photo := false, false

	user, err := repos.GetUserRepo().GetUser(chatID)
	if err != nil {
		errors.LogError(errors.ErrorGettingUserState, err)
		return
	}
	state := user.State
	log.Println(state, msgText)

	switch msgText {
	case cst.Search:
		responseText, err = HandleSearchCom(chatID, "search", repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard(true)
		photo = true
	case cst.Info:
		responseText, err = HandleSearchCom(chatID, "info", repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateBuildingKeyboard(false)
		photo = true
	case cst.Back:
		msg, err := HandleBackCom(chatID, state, repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		update.Message.Text = msg
		HandleUpdateCom(botApi, update, repos)
	case cst.Menu:
		responseText, err = HandleMenuCom(chatID, repos)
		if err != nil {
			errors.LogError(errors.ErrorGettingUserState, err)
			return
		}
		replyMarkup = keyboard.CreateMainKeyboard()
	default:
		switch state {
		case "info":
			responseText, next, err = states.HandleInfoState(chatID, msgText, "info_number", repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateBuildingKeyboard(false)
			} else {
				replyMarkup = keyboard.CreateMiniKeyboard(cst.Back)
			}
		case "info_number":
			responseText, next, err = states.HandleInfoNumberState(chatID, msgText, repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateMiniKeyboard(cst.Back)
			} else {
				replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
			}
		case "search":
			responseText, next, err = states.HandleInfoState(chatID, msgText, "search_date", repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateBuildingKeyboard(true)
			} else {
				replyMarkup = keyboard.CreateDateKeyboard()
			}
		case "search_date":
			responseText, next, err = states.HandleSearchDateState(chatID, msgText, "search_time", repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateDateKeyboard()
			} else {
				replyMarkup = keyboard.CreateTimeKeyboard()
			}

		case "search_time":
			responseText, next, err = states.HandleSearchTimeState(chatID, msgText, repos)
			if err != nil {
				errors.LogError(errors.ErrorUpdatingUser, err)
				return
			}
			if !next {
				replyMarkup = keyboard.CreateDateKeyboard()
			} else {
				replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
			}
		case "finish":
			responseText = cst.CantUnderstand
			replyMarkup = keyboard.CreateMiniKeyboard(cst.Menu)
		default:
			responseText = cst.CantUnderstand
			replyMarkup = keyboard.CreateMainKeyboard()
		}
	}

	if photo {
		if err = functions.SendPhoto(botApi, chatID, cst.MapPath); err != nil {
			errors.LogError(errors.ErrorSendingMessage, err)
			return
		}
	}

	if err = functions.SendMessage(botApi, replyMarkup, chatID, responseText); err != nil {
		errors.LogError(errors.ErrorSendingMessage, err)
		return
	}
}
