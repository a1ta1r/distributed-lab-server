package entity

type Team struct {
	DbAwareEntity
	Name     string    `json:"name"`
	OwnerID  uint      `json:"owner_id"`
	Users    []User    `json:"users"`
	Projects []Project `json:"projects"`
}
