package models

type User struct {
	Model
	ChatID        int64  `json:"chat_id" gorm:"primary_key"`
	State         string `json:"state"`
	SavedBuilding string `json:"saved_building"`
	SavedRoom     string `json:"saved_room"`
	SavedDate     string `json:"saved_date"`
	SavedFilter   bool   `json:"saved_filter"`
}
