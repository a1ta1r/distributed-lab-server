package entity

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name    string
	Team    Team
	Owner   User
	Members []User
	Tasks   []Task
}
