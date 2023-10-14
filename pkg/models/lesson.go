package models

import (
	"time"
)

type Lesson struct {
	Model
	ID       string     `json:"id" gorm:"primary_key"`
	Type     string     `json:"type"`
	Subject  string     `json:"subject"`
	Week     int        `json:"week"`
	WeekDay  int        `json:"week_day"`
	TimeFrom time.Time  `json:"time_from"`
	TimeTo   time.Time  `json:"time_to"`
	DateFrom *time.Time `json:"date_from"`
	DateTo   *time.Time `json:"date_to"`
	Date     *time.Time `json:"date"`
	Groups   []Group    `json:"groups" gorm:"many2many:lessons_groups;constraint:OnDelete:CASCADE"`
	Semester int        `json:"semester"`
	Tutors   []Tutor    `json:"tutors" gorm:"many2many:lessons_tutors;constraint:OnDelete:CASCADE"`
	Room     Room       `json:"room"`
	RoomID   string
}
