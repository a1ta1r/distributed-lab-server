package entity

import "github.com/jinzhu/gorm"

type Team struct {
	gorm.Model
	Name     string
	Owner    User
	Members  []User
	Projects []Project
}
