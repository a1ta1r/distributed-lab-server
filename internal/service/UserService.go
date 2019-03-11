package service

import "github.com/a1ta1r/distributed-lab-server/internal/storage"

type UserService struct {
	userStorage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return UserService{userStorage:userStorage}
}
