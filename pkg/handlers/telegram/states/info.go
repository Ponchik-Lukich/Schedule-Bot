package states

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
)

func HandleInfoState(userRepo user.Repository, chatID int64, building string) (string, bool, error) {
	if _, ok := cst.Buildings[building]; !ok {
		return cst.BuildingDoesntExist, false, nil
	}

	updates := map[string]interface{}{
		"state":          "info_number",
		"saved_building": building,
	}

	if err := userRepo.UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return cst.EnterRoom, true, nil
}
