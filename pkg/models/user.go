package models

type User struct {
	ChatID int64  `json:"chat_id" gorm:"primary_key"`
	State  string `json:"state"`
}
