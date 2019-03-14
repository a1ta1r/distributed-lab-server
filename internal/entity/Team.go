package entity

type Team struct {
	DbAwareEntity
	Owner   User   `gorm:"foreignkey:OwnerId"`
	OwnerId uint   `json:"owner_id"`
	Users   []User `gorm:"many2many:user_teams;"`
}
