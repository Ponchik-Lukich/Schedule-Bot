package commands

import (
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
)

func HandleMenuCom(chatID int64, repos repo.Repositories) (string, error) {
	updates := map[string]any{
		"state": "wait",
	}
	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
		return "", err
	}
	reply := cst.Choice

	return reply, nil
}

func HandleBackCom(chatID int64, userState string, repos repo.Repositories) (string, error) {
	if userState == "finish" {
		return cst.Menu, nil
	}

	if newState, ok := cst.NumberStates[cst.StatesNumber[userState]-1]; !ok {
		userState = newState
	} else {
		userState = "wait"
	}

	updates := map[string]any{
		"state": userState,
	}

	if err := repos.GetUserRepo().UpdateUser(chatID, updates); err != nil {
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
