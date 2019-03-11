package storage

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
)

type StatusStorage struct {
	db gorm.DB
}

func NewStatusStorage(db gorm.DB) StatusStorage {
	return StatusStorage{db: db}
}

func (us StatusStorage) AddStatus(status entity.Status) (entity.Status, error) {
	err := us.db.Save(&status).Error
	return status, err
}

func (us StatusStorage) GetStatus(status entity.Status) (entity.Status, error) {
	err := us.db.First(&status, status.ID).Error
	return status, err
}

func (us StatusStorage) GetStatuses() ([]entity.Status, error) {
	var statuses []entity.Status

	err := us.db.Find(&statuses).Error
	return statuses, err
}

func (us StatusStorage) UpdateStatus(status entity.Status) (entity.Status, error) {
	err := us.db.Save(&status).Error
	return status, err
}

func (us StatusStorage) DeleteStatus(status entity.Status) error {
	err := us.db.Delete(&status, status.ID).Error
	return err
}


