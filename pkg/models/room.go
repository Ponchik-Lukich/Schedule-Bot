package models

import (
	"fmt"
	"strings"
)

type Room struct {
	Model
	ID           string `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	Building     string `json:"building"`
	IsLaboratory bool   `json:"is_laboratory"`
	HasComputer  bool   `json:"has_computer"`
	HasProjector bool   `json:"has_projector"`
	IsAvailable  bool   `json:"is_available"`
	HasDot       bool   `json:"has_dot"`
	IsTemporary  bool   `json:"is_temporary"`
}

func (r Room) String() string {
	var b strings.Builder

	b.WriteString("Room Details:\n")
	b.WriteString(fmt.Sprintf("ID: %s\n", r.ID))
	b.WriteString(fmt.Sprintf("Name: %s\n", r.Name))
	b.WriteString(fmt.Sprintf("Building: %s\n", r.Building))
	b.WriteString(fmt.Sprintf("Is a Laboratory: %v\n", r.IsLaboratory))
	b.WriteString(fmt.Sprintf("Has Computer: %v\n", r.HasComputer))
	b.WriteString(fmt.Sprintf("Has Projector: %v\n", r.HasProjector))
	b.WriteString(fmt.Sprintf("Is Available: %v\n", r.IsAvailable))
	b.WriteString(fmt.Sprintf("Has Dot: %v\n", r.HasDot))
	b.WriteString(fmt.Sprintf("Is Temporary: %v\n", r.IsTemporary))

	return b.String()
}

type RoomInfoDto struct {
	RoomName       string
	IsAvailability bool
	HasProjector   bool
	Lessons        []LessonInfoDto
}

func (r RoomInfoDto) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("Room: %s\n", r.RoomName))
	b.WriteString(fmt.Sprintf("Is Available: %v\n", r.IsAvailability))
	b.WriteString(fmt.Sprintf("Has Projector: %v\n", r.HasProjector))
	b.WriteString("Lessons:\n")
	for _, lesson := range r.Lessons {
		b.WriteString(fmt.Sprintf("%s\n", lesson.String()))
	}

	return b.String()
}
