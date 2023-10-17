package states

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
)

func HandleInfoState(chatID int64, building, state string, repos repo.Repositories) (string, bool, error) {
	if _, ok := cst.Buildings[building]; !ok {
		return cst.BuildingDoesntExist, false, nil
	}

	updates := map[string]any{
		"state":          state,
		"saved_building": building,
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	if state == "search_date" {
		return cst.EnterDate, true, nil
	}

	return cst.EnterRoom, true, nil
}

func HandleInfoNumberState(chatID int64, number string, repos repo.Repositories) (string, bool, error) {
	user, err := repos.GetUserRepo().GetUser(chatID)
	if err != nil {
		return "", false, err
	}

	building := user.SavedBuilding

	res, next, err := repos.GetRoomRepo().GetRoomInfo(building, number)
	if err != nil {
		return "", false, err
	}
	if !next {
		return res, false, nil
	}

	updates := map[string]any{
		"state": "finish",
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return res, true, nil
}
