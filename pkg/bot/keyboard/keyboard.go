package keyboard

import (
	cst "Telegram/pkg/constants"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func CreateDateKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cst.Today),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cst.Tomorrow),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cst.Back),
		),
	)

	return keyboard
}

func CreateTimeKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cst.Now),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(cst.Back),
		),
	)

	return keyboard
}

func CreateBuildingKeyboard(expanded bool) tgbotapi.ReplyKeyboardMarkup {
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
			tgbotapi.NewKeyboardButton("Корпус 5"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Корпус 64"),
			tgbotapi.NewKeyboardButton("Корпус МПК"),
		),
	)
	if expanded {
		keyboard.Keyboard = append(keyboard.Keyboard, tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Всё равно"),
		))
	}
	keyboard.Keyboard = append(keyboard.Keyboard, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Назад"),
	))

	return keyboard
}
