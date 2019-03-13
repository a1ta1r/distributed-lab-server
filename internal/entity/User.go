package entity

type User struct {
	DbAwareEntity
	Name           string    `json:"name"`
	Position       string    `json:"position"`
	Status         string    `json:"status"`
	Team           []Team    `gorm:"many2many:user_teams;" json:"team_id"`
	TeamsOwned     []Team    `gorm:"foreignkey:Owner"`
	Projects       []Project `gorm:"many2many:project_members;"`
	Tasks          []Task    `gorm:"many2many:user_tasks;"`
	TasksInitiated []Task    `gorm:"foreignkey:Initiator"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
}
