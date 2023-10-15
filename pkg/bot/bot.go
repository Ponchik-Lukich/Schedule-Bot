package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	var responseText string
	var replyMarkup tgbotapi.ReplyKeyboardMarkup

	switch update.Message.Text {
	case "Поиск свободного кабинета":
		responseText = "Here's the info for free rooms..."
		replyMarkup = CreateMiniKeyboard("Назад")
	case "Информация о кабинете":
		responseText = "Here's the room information..."
		replyMarkup = CreateMiniKeyboard("Назад")
	case "Назад":
		responseText = "Choose an option:"
		replyMarkup = CreateMainKeyboard()
	default:
		responseText = "Choose an option:"
		replyMarkup = CreateMainKeyboard()
	}

	replyMarkup.ResizeKeyboard = true
	replyMarkup.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseText)
	msg.ReplyMarkup = replyMarkup

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
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
