package entity

import "time"

type Project struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Name    string `json:"name"`
	Team    Team `json:"team"`
	Owner   User `json:"owner"`
	Members []User `json:"members"`
	Tasks   []Task `json:"tasks"`
}
