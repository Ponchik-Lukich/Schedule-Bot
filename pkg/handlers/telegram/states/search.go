package states

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
	"Telegram/pkg/utils"
	"time"
)

func HandleSearchDateState(chatID int64, date, state string, repos repo.Repositories) (string, bool, error) {
	targetDay := ""

	switch date {
	case cst.Today:
		today := time.Now().Add(3 * time.Hour)
		targetDay = today.Format(cst.DateLayout)
	case cst.Tomorrow:
		tomorrow := time.Now().Add(27 * time.Hour)
		targetDay = tomorrow.Format(cst.DateLayout)
	default:
		if !utils.IsDate(date) {
			return cst.WrongDateFormat, false, nil
		}
	}

	updates := map[string]any{
		"state":      state,
		"saved_date": targetDay,
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return cst.EnterTime, true, nil
}

func HandleSearchTimeState(chatID int64, searchTime string, repos repo.Repositories) (string, bool, error) {
	targetTime := ""

	if searchTime == cst.Now {
		now := time.Now()
		targetTime = now.Format(cst.TimeLayout)
	} else {
		if _, err := time.Parse(cst.TimeLayout, searchTime); err != nil {
			return cst.WrongTimeFormat, false, nil
		}
		targetTime = searchTime
	}

	user, err := repos.GetUserRepo().GetUser(chatID)
	if err != nil {
		return "", false, err
	}

	building, date := user.SavedBuilding, user.SavedDate

	res, err := repos.GetRoomRepo().GetFreeRooms(building, date, targetTime)
	if err != nil {
		return "", false, err
	}

	updates := map[string]any{
		"state": "finish",
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", false, err
	}

	return res, true, nil
}
