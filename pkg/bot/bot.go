package bot

import (
	"Telegram/pkg/handlers/telegram"
	"Telegram/pkg/repo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func RunBotLocal(botToken string, repos repo.Repositories) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			telegram.HandleStart(bot, &update, repos.GetUserRepo())
			continue
		}
		telegram.HandleUpdate(bot, &update, repos.GetUserRepo(), repos.GetRoomRepo())
	}
}
