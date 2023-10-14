package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func CreateMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поиск свободного кабинета"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Информация о кабинете"),
		),
	)
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true

	return keyboard
}

func CreateMiniKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = true

	return keyboard
}

func HandleUpdate(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	var responseText string
	var replyMarkup interface{}

	switch update.Message.Text {
	case "Поиск свободного кабинета":
		responseText = "Here's the info for free rooms..."
		replyMarkup = CreateMiniKeyboard()
	case "Информация о кабинете":
		responseText = "Here's the room information..."
		replyMarkup = CreateMiniKeyboard()
	case "Назад":
		responseText = "Choose an option:"
		replyMarkup = CreateMainKeyboard()
	default:
		responseText = "Choose an option:"
		replyMarkup = CreateMainKeyboard()
	}

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
