package telegram

import (
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleBackToMenu(userRepo user.Repository, chatID int64) (string, tgbotapi.ReplyKeyboardMarkup, error) {
	err := userRepo.SetUserState(chatID, "wait")
	if err != nil {
		return "", tgbotapi.ReplyKeyboardMarkup{}, err
	}
	reply := cst.Choice
	buttons := keyboard.CreateMainKeyboard()

	return reply, buttons, nil
}

func HandleBack(userRepo user.Repository, chatID int64, state string) (string, error) {
	if state == "finish" {
		return cst.Menu, nil
	}
	
	if newState, ok := cst.NumberStates[cst.StatesNumber[state]-1]; !ok {
		state = newState
	} else {
		state = "wait"
	}

	err := userRepo.SetUserState(chatID, state)
	if err != nil {
		return "", err
	}

	switch state {
	case "wait":
		return cst.Menu, nil
	case "search":
		return cst.Search, nil
	case "info":
		return cst.Info, nil
	}

	return "", nil
}
