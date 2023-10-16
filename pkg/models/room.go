package models

import (
	cst "Telegram/pkg/constants"
	"fmt"
	"strings"
)

type Room struct {
	Model
	ID           string `json:"id" gorm:"primary_key"`
	Name         string `json:"name" gorm:"index"`
	Building     string `json:"building" gorm:"index"`
	IsLaboratory bool   `json:"is_laboratory" gorm:"index"`
	HasComputer  bool   `json:"has_computer" gorm:"index"`
	HasProjector bool   `json:"has_projector" gorm:"index"`
	IsAvailable  bool   `json:"is_available" gorm:"index"`
	HasDot       bool   `json:"has_dot" gorm:"index"`
	IsTemporary  bool   `json:"is_temporary" gorm:"index"`
}

type RoomInfoDto struct {
	RoomName     string
	IsAvailable  bool
	HasProjector bool
	Lessons      []Lesson
}

func (r RoomInfoDto) String() string {
	var b strings.Builder

	sortLessons(r.Lessons)

	available, projector := cst.Emoji["No"], cst.Emoji["No"]
	if r.IsAvailable {
		available = cst.Emoji["Yes"]
	}
	if r.HasProjector {
		projector = cst.Emoji["Yes"]
	}

	b.WriteString(fmt.Sprintf("%s %s: %s\n\n", cst.Emoji["Room"], cst.Info, r.RoomName))
	b.WriteString(fmt.Sprintf("%s: %s     %s: %s\n", cst.Emoji["Ava"], available, cst.Emoji["Proj"], projector))

	currentDay := -1
	for _, lesson := range r.Lessons {
		if lesson.WeekDay != currentDay {
			b.WriteString(fmt.Sprintf("%s\n\n", cst.Days[lesson.WeekDay]))
			currentDay = lesson.WeekDay
		}
		b.WriteString(fmt.Sprintf("%s\n", lesson.String()))
	}

	return b.String()
}
