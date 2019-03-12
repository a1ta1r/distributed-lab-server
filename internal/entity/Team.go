package entity

type Team struct {
	DbAwareEntity
	Name  string `json:"name"`
	Owner User
	Users []User `gorm:"many2many:user_teams;"`
}
