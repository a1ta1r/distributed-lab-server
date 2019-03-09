package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Position string
	Status   string
	Team     Team
}
