package states

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
	"strings"
)

func HandleInfoState(chatID int64, building string, repos repo.Repositories) (string, bool, error) {
	if _, ok := cst.Buildings[building]; !ok {
		return cst.BuildingDoesntExist, false, nil
	}

	updates := map[string]any{
		"state":          "info_number",
		"saved_building": strings.Split(building, " ")[1],
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return cst.EnterRoom, true, nil
}

func HandleInfoNumberState(chatID int64, number string, repos repo.Repositories) (string, bool, error) {
	user, err := repos.GetUserRepo().GetUser(chatID)
	if err != nil {
		return "", false, err
	}

	building := user.SavedBuilding

	res, err := repos.GetRoomRepo().GetRoomInfo(*building, number)
	if err != nil {
		return "", false, err
	}

	updates := map[string]any{
		"state":          "finish",
		"saved_building": "",
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return res, true, nil
}
