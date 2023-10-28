package commands

import (
	"Telegram/pkg/bot/functions"
	"Telegram/pkg/bot/keyboard"
	cst "Telegram/pkg/constants"
	"Telegram/pkg/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleStartCom(botApi *tgbotapi.BotAPI, update *tgbotapi.Update, repos repo.Repositories) {
	chatID := update.Message.Chat.ID
	err := repos.GetUserRepo().CreateUser(chatID)
	if err != nil {
		log.Println(err)
	}

	reply := keyboard.CreateMainKeyboard()
	err = functions.SendMessage(botApi, reply, update.Message.Chat.ID, cst.Greeting+"\n"+cst.Commands)
	if err != nil {
		log.Println(err)
	}
}
