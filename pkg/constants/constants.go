package constants

const (
	Ydb                 = "ydb"
	Postgres            = "postgres"
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
	RoomPattern         = `^(|.*\D)number(|\D.*)$`
	RoomsFound          = "Найдено несколько кабинетов! \nПожалуйста, уточните номер кабинета:"
	FreeRooms           = "Свободные кабинеты:"
	IsAvailable         = "Возможность занять"
	No                  = "Нет"
	Yes                 = "Да"
	Schedule            = "Расписание:"
	Projector           = "Наличие проектора"
	FullDateLayout      = "2006-01-02"
	DateLayout          = "02-01"
	EnterDate           = "Введите дату (День-Месяц)"
	Today               = "Сегодня"
	Tomorrow            = "Завтра"
	Now                 = "Текущее время"
	EnterTime           = "Введите время (Часы:Минуты)"
	WrongDateFormat     = "Неверный формат даты. Пожалуйста, введите дату в формате День-Месяц"
	TimeLayout          = "15:04"
	WrongTimeFormat     = "Неверный формат времени. Пожалуйста, введите время в формате Часы:Минуты"
	NoRoomsFound        = "К сожалению, свободных кабинетов не найдено. Попробуйте выбрать другое время."
	Building            = "AND building = ?"
	Any                 = "Всё равно"
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
		"Корпус 5":   {},
		"Корпус МПК": {},
		"Всё равно":  {},
	}
	Emoji = map[string]string{
		"Room": "🚪",
		"Proj": "📽",
		"Ava":  "🔓",
		"Sch":  "📅",
		"Time": "🕒",
		"Tut":  "👨‍🎓",
		"No":   "❌",
		"Yes":  "✅",
		"Lec":  "📚",
		"Pra":  "📝",
		"Lab":  "🔬",
		"Ayd":  "🎯",
		"Rez":  "🚫",
	}
	Days = map[int]string{
		0: "Понедельник",
		1: "Вторник",
		2: "Среда",
		3: "Четверг",
		4: "Пятница",
		5: "Суббота",
		6: "Воскресенье",
	}
)
