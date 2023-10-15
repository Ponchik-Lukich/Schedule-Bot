package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func CreateMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Поиск свободного кабинета"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Информация о кабинете"),
		),
	)

	return keyboard
}

func CreateMiniKeyboard(direction string) tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(direction),
		),
	)

	return keyboard
}

func CreateBuildingKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус А"),
			tgbotapi.NewKeyboardButton("Корпус Б"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус В"),
			tgbotapi.NewKeyboardButton("Корпус Г"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус Д"),
			tgbotapi.NewKeyboardButton("Корпус И"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус К"),
			tgbotapi.NewKeyboardButton("Корпус Т"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус Э"),
			tgbotapi.NewKeyboardButton("Корпус НЛК"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус 31"),
			tgbotapi.NewKeyboardButton("Корпус 33"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус 45"),
			tgbotapi.NewKeyboardButton("Корпус 46"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус 47"),
			tgbotapi.NewKeyboardButton("Корпус 5 (Технопарк)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус 64"),
			tgbotapi.NewKeyboardButton("Корпус МПК"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Всё равно"),
		),
	)

	return keyboard
}
