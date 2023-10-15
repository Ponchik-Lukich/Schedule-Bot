package constants

const (
	Ydb            = "ydb"
	Greeting       = "Привет! Я бот, который поможет тебе найти свободный кабинет в университете."
	BuildingChoice = "Выбери корпус: "
	Choice         = "Выбери функцию:"
	Info           = "Информация о кабинете"
	Search         = "Поиск свободного кабинета"
	Menu           = "Вернуться в главное меню"
	Back           = "Назад"
	CantUnderstand = "Извини, я не понимаю тебя."
)

var (
	NumberStates = map[int64]string{
		1:   "wait",
		2:   "search",
		3:   "search_date",
		4:   "search_time",
		5:   "finish",
		100: "info",
		101: "info_number",
	}
	StatesNumber = map[string]int64{
		"wait":        1,
		"search":      2,
		"search_day":  3,
		"search_time": 4,
		"finish":      5,
		"info":        100,
		"info_number": 101,
	}
)
