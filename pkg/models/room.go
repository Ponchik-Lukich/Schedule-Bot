package models

type Room struct {
	Model
	ID             string `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Building       string `json:"building"`
	IsLaboratory   bool   `json:"is_laboratory"`
	HasComputer    bool   `json:"has_computer"`
	HasProjector   bool   `json:"has_projector"`
	IsAvailability bool   `json:"is_available"`
	HasDot         bool   `json:"has_dot"`
	IsTemporary    bool   `json:"is_temporary"`
}
