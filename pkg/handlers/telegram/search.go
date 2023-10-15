package telegram

import (
	"Telegram/pkg/bot/keyboard"
	"Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleSearch(userRepo user.Repository, chatID int64, userState string) (string, tgbotapi.ReplyKeyboardMarkup, error) {
	err := userRepo.SetUserState(chatID, userState)
	if err != nil {
		return "", tgbotapi.ReplyKeyboardMarkup{}, err
	}
	reply := constants.BuildingChoice
	buttons := keyboard.CreateBuildingKeyboard()

	return reply, buttons, nil
}
