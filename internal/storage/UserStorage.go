package storage

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
)

type UserStorage struct {
	db gorm.DB
}

func NewUserStorage(db gorm.DB) UserStorage {
	return UserStorage{db: db}
}

func (us UserStorage) AddUser(user entity.User) (entity.User, error) {
	err := us.db.Save(user).Error
	return user, err
}

func (us UserStorage) GetUser(id uint) (entity.User, error) {
	var user entity.User

	err := us.db.First(&user, user.ID).Error
	return user, err
}

func (us UserStorage) GetUsers() ([]entity.User, error) {
	var users []entity.User

	err := us.db.Find(&users).Error
	return users, err
}

func (us UserStorage) UpdateUser(user entity.User) (entity.User, error) {
	err := us.db.Save(user).Error
	return user, err
}

func (us UserStorage) DeleteUser(user entity.User) error {
	err := us.db.Delete(&user, user.ID).Error
	return err
}
