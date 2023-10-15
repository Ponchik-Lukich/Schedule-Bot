package commands

import (
	"Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
)

func HandleSearchCom(userRepo user.Repository, chatID int64, userState string) (string, error) {
	updates := map[string]interface{}{
		"state": "search_building",
	}
	if err := userRepo.UpdateUser(chatID, updates); err != nil {
		return "", err
	}
	reply := constants.BuildingChoice

	return reply, nil
}
