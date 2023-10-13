package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			handleStart(bot, update)
		}
	}
}

func handleStart(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func RunBotLocal(botToken string) {
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
		HandleUpdate(bot, &update)
	}
}
