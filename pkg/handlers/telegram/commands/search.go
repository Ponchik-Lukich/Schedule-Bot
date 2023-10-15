package commands

import (
	"Telegram/pkg/constants"
	"Telegram/pkg/repo"
)

func HandleSearchCom(chatID int64, userState string, repos repo.Repositories) (string, error) {
	updates := map[string]any{
		"state": userState,
	}
	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", err
	}
	reply := constants.BuildingChoice

	return reply, nil
}
