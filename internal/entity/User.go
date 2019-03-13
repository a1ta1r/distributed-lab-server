package entity

type User struct {
	DbAwareEntity
	Position string `json:"position"`
	Status   string `json:"status"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserIds(users []User) []uint {
	var userIds []uint
	for _, user := range users {
		userIds = append(userIds, user.ID)
	}
	return userIds
}
