package models

type Group struct {
	Model
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
