package entity

type Project struct {
	DbAwareEntity
	Team    Team `gorm:"foreignkey:TeamId"`
	TeamId  uint `json:"team_id"`
	Owner   User `gorm:"foreignkey:OwnerId"`
	OwnerId uint `json:"owner_id"`
}
