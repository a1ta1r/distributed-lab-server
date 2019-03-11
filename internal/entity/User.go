package entity

import "time"

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Name      string     `json:"name"`
	Position  string     `json:"position"`
	Status    string     `json:"status"`
	Team      Team       `json:"team"`
}
