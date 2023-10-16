package models

import (
	"fmt"
	"strings"
	"time"
)

type Lesson struct {
	Model
	ID       string     `json:"id" gorm:"primary_key"`
	Type     string     `json:"type"`
	Subject  string     `json:"subject"`
	Week     int        `json:"week" gorm:"index"`
	WeekDay  int        `json:"week_day" gorm:"index"`
	TimeFrom time.Time  `json:"time_from" gorm:"index:time_interval"`
	TimeTo   time.Time  `json:"time_to" gorm:"index:time_interval"`
	DateFrom *time.Time `json:"date_from" gorm:"index:date_interval"`
	DateTo   *time.Time `json:"date_to" gorm:"index:date_interval"`
	Date     *time.Time `json:"date" gorm:"index"`
	Groups   []Group    `json:"groups" gorm:"many2many:lessons_groups;constraint:OnDelete:CASCADE"`
	Semester int        `json:"semester"`
	Tutors   []Tutor    `json:"tutors" gorm:"many2many:lessons_tutors;constraint:OnDelete:CASCADE"`
	Room     Room       `json:"room"`
	RoomID   string     `json:"room_id" gorm:"index"`
}

type LessonInfoDto struct {
	TimeFrom   time.Time
	TimeTo     time.Time
	TutorNames []string
	GroupNames []string
}

func (l LessonInfoDto) String() string {
	return fmt.Sprintf("Время: %s - %s\nПреподаватели: %s\nГруппы: %s\n", l.TimeFrom.Format("15:04"),
		l.TimeTo.Format("15:04"), strings.Join(l.TutorNames, ", "), strings.Join(l.GroupNames, ", "))
}
