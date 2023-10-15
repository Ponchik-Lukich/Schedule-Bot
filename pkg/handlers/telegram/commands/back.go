package commands

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo/user"
)

func HandleMenuCom(userRepo user.Repository, chatID int64) (string, error) {
	updates := map[string]interface{}{
		"state": "wait",
	}
	if err := userRepo.UpdateUser(chatID, updates); err != nil {
		return "", err
	}
	reply := cst.Choice

	return reply, nil
}

func HandleBackCom(userRepo user.Repository, chatID int64, userState string) (string, error) {
	if userState == "finish" {
		return cst.Menu, nil
	}

	if newState, ok := cst.NumberStates[cst.StatesNumber[userState]-1]; !ok {
		userState = newState
	} else {
		userState = "wait"
	}

	updates := map[string]interface{}{
		"state": userState,
	}

	if err := userRepo.UpdateUser(chatID, updates); err != nil {
		return "", err
	}

	switch userState {
	case "wait":
		return cst.Menu, nil
	case "search":
		return cst.Search, nil
	case "info":
		return cst.Info, nil
	}

	return "", nil
}
