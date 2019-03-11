package entity

type Team struct {
	DbAwareEntity
	Name     string    `json:"name"`
	Owner    *User      `json:"owner"`
	Members  *[]User    `json:"members"`
	Projects []Project `json:"projects"`
}
