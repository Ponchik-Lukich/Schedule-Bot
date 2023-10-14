package models

type Tutor struct {
	Model
	ID        string `json:"id" gorm:"primary_key"`
	ShortName string `json:"short_name"`
}
