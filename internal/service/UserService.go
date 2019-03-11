package service

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	userStorage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return UserService{userStorage: userStorage}
}

func (us UserService) AddUser(user entity.User) entity.User {
	user, err := us.userStorage.AddUser(user)
	if err != nil {
		panic(err)
	}
	return user
}

func (us UserService) GetUsers() []entity.User {
	users, err := us.userStorage.GetUsers()
	if err != nil {
		panic(err)
	}

	return users
}

func (us UserService) GetUser(user entity.User) (entity.User, error) {
	user, err := us.userStorage.GetUser(user)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return user, err
}

func (us UserService) UpdateUser(user entity.User) (entity.User, error) {
	user, err := us.userStorage.UpdateUser(user)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return user, err
}

func (us UserService) DeleteUser(user entity.User) {
	err := us.userStorage.DeleteUser(user)
	if err != nil {
		panic(err)
	}
}
