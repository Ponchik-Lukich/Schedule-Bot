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
		"saved_building": building,
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return cst.EnterRoom, true, nil
}

func HandleInfoNumberState(chatID int64, number string, repos repo.Repositories) (string, bool, error) {
	room, err := repos.GetRoomRepo().GetRoomInfo(number)
	if err != nil {
		return "", false, err
	}

	if room == nil {
		return cst.RoomDoesntExist, false, nil
	}

	var res strings.Builder
	res.WriteString(cst.RoomInfo + "\n")
	for _, r := range room {
		res.WriteString(r.String() + "\n")
	}

	updates := map[string]any{
		"state": "finish",
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return res.String(), true, nil
}
