package models

import "time"

type Lesson struct {
	Id       string    `json:"lesson_id"`
	Subject  string    `json:"subject"`
	Week     int       `json:"week"`
	WeekDay  int       `json:"week_day"`
	Date     string    `json:"date"`
	TimeFrom string    `json:"time_from"`
	TimeTo   string    `json:"time_to"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
	Groups   []Group   `json:"groups"`
	Semester int       `json:"semester"`
	Tutors   []Tutor   `json:"tutors"`
	Room     Room      `json:"room"`
}
