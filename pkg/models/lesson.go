package models

import (
	cst "Telegram/pkg/constants"
	"fmt"
	"time"
)

type Lesson struct {
	Model
	ID       string     `json:"id" gorm:"primary_key"`
	Type     string     `json:"type"`
	Subject  string     `json:"subject"`
	Week     int        `json:"week" gorm:"index"`
	WeekDay  int        `json:"week_day" gorm:"index"`
	TimeFrom time.Time  `json:"time_from" gorm:"index:time_interval;index"`
	TimeTo   time.Time  `json:"time_to" gorm:"index:time_interval;index"`
	DateFrom *time.Time `json:"date_from" gorm:"index:date_interval"`
	DateTo   *time.Time `json:"date_to" gorm:"index:date_interval"`
	Date     *time.Time `json:"date" gorm:"index"`
	Groups   []Group    `json:"groups" gorm:"many2many:lessons_groups;constraint:OnDelete:CASCADE"`
	Semester int        `json:"semester"`
	Tutors   []Tutor    `json:"tutors" gorm:"many2many:lessons_tutors;constraint:OnDelete:CASCADE"`
	Room     Room       `json:"room"`
	RoomID   string     `json:"room_id" gorm:"index"`
}

func (l Lesson) String() string {
	timeFrom := l.TimeFrom.UTC().Add(3 * time.Hour).Format("15:04")
	timeTo := l.TimeTo.UTC().Add(3 * time.Hour).Format("15:04")
	lessonType, res := "", ""

	switch l.Type {
	case "Лек":
		lessonType = cst.Emoji["Lec"]
	case "Лаб":
		lessonType = cst.Emoji["Lab"]
	case "Пр":
		lessonType = cst.Emoji["Pra"]
	case "Ауд":
		lessonType = cst.Emoji["Ayd"]
	case "Резерв":
		lessonType = cst.Emoji["Rez"]
	}

	lessonDetails := fmt.Sprintf("%s %s - %s  %s <b>%s</b> %s", cst.Emoji["Time"], timeFrom, timeTo, lessonType, l.Type, l.Subject)

	var tutorsString string
	for _, tutor := range l.Tutors {
		tutorsString += tutor.ShortName + " "
	}

	var dateRangeString string
	if l.DateFrom != nil && l.DateTo != nil {
		dateRangeString = fmt.Sprintf("%s (%s - %s)\n", cst.Emoji["Sch"], l.DateFrom.Format("02.01.2006"),
			l.DateTo.Format("02.01.2006"))
	}

	if len(l.Tutors) != 0 {
		res = fmt.Sprintf("%s\n%s %s\n%s", lessonDetails, cst.Emoji["Tut"], tutorsString, dateRangeString)
	} else {
		res = fmt.Sprintf("%s\n %s", lessonDetails, dateRangeString)
	}

	return res
}
