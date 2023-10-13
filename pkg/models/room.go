package models

type Room struct {
	Id           string `json:"room_id"`
	Name         string `json:"name"`
	Building     string `json:"building"`
	Laboratory   bool   `json:"is_laboratory"`
	Computer     bool   `json:"has_computer"`
	Projector    bool   `json:"has_projector"`
	Availability bool   `json:"is_available"`
	Dot          bool   `json:"has_dot"`
	Temporary    bool   `json:"is_temporary"`
}
