package entity

type Project struct {
	DbAwareEntity
	Name    string `json:"name"`
	Team    Team   `json:"team"`
	Owner   User   `json:"owner"`
	Members []User `json:"members"`
	Tasks   []Task `json:"tasks"`
}
