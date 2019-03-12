package entity

type Project struct {
	DbAwareEntity
	Name    string `json:"name"`
	Team    Team   `json:"team"`
	Owner   User   `json:"owner"`
	Members []User `gorm:"many2many:project_members;"`
	Tasks   []Task `gorm:"foreignkey:Project"`
}
