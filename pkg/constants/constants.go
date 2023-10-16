package constants

const (
	Ydb                 = "ydb"
	Greeting            = "Привет! Я бот, который поможет тебе найти свободный кабинет в университете."
	BuildingChoice      = "Выберите корпус: "
	Choice              = "Выберите функцию:"
	Info                = "Информация о кабинете"
	Search              = "Поиск свободного кабинета"
	Menu                = "Вернуться в главное меню"
	Back                = "Назад"
	CantUnderstand      = "Извини, я не понимаю тебя."
	EnterRoom           = "Введите номер кабинета:"
	BuildingDoesntExist = "Корпуса с таким названием не существует."
	MapPath             = "assets/images/map.png"
	RoomDoesntExist     = "Кабинета с таким номером не существует."
	RoomPattern         = `^(|.*\D)число(|\D.*)$`
	RoomsFound          = "Найдено несколько кабинетов! \nПожалуйста, уточните номер кабинета:"
	IsAvailable         = "Возможность занять:"
	No                  = "Нет"
	Yes                 = "Да"
	Schedule            = "Расписание:"
	Projector           = "Наличие проектора:"
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
	Buildings = map[string]struct{}{
		"Корпус А":   {},
		"Корпус Б":   {},
		"Корпус В":   {},
		"Корпус Г":   {},
		"Корпус Д":   {},
		"Корпус И":   {},
		"Корпус К":   {},
		"Корпус Т":   {},
		"Корпус Э":   {},
		"Корпус НЛК": {},
		"Корпус 31":  {},
		"Корпус 33":  {},
		"Корпус 45":  {},
		"Корпус 46":  {},
		"Корпус 47":  {},
		"Корпус 64":  {},
		"Корпус МПК": {},
		"Всё равно":  {},
	}
)
