package entity

type Project struct {
	DbAwareEntity
	Team    Team   `gorm:"foreignkey:ProjectId"`
	Owner   User   `gorm:"foreignkey:OwnerId"`
	OwnerId uint   `json:"owner_id"`
	Members []User `gorm:"many2many:project_members;"`
	Tasks   []Task `gorm:"foreignkey:Project"`
}
