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
	Lessons      []LessonInfoDto
}

func (r RoomInfoDto) String() string {
	var b strings.Builder
	available, projector := cst.No, cst.No
	if r.IsAvailable {
		available = cst.Yes
	}
	if r.HasProjector {
		projector = cst.Yes
	}

	b.WriteString(fmt.Sprintf("%s\n", cst.RoomInfo))
	b.WriteString(fmt.Sprintf("%s\n", r.RoomName))
	b.WriteString(fmt.Sprintf("%s %v\n", cst.IsAvailable, available))
	b.WriteString(fmt.Sprintf("%s %v\n", cst.Projector, projector))
	b.WriteString(fmt.Sprintf("%s\n", cst.Schedule))
	for _, lesson := range r.Lessons {
		b.WriteString(fmt.Sprintf("%s\n", lesson.String()))
	}

	return b.String()
}
