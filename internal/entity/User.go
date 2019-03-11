package entity

type User struct {
	DbAwareEntity
	Name     string `json:"name"`
	Position string `json:"position"`
	Status   string `json:"status"`
	Team     Team   `json:"team"`
}
