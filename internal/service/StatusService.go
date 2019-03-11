package service

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

type StatusService struct {
	statusStorage storage.StatusStorage
}

func NewStatusService(statusStorage storage.StatusStorage) StatusService {
	return StatusService{statusStorage: statusStorage}
}

func (us StatusService) AddStatus(status entity.Status) entity.Status {
	status, err := us.statusStorage.AddStatus(status)
	if err != nil {
		panic(err)
	}
	return status
}

func (us StatusService) GetStatuses() []entity.Status {
	statuses, err := us.statusStorage.GetStatuses()
	if err != nil {
		panic(err)
	}

	return statuses
}

func (us StatusService) GetStatus(status entity.Status) (entity.Status, error) {
	status, err := us.statusStorage.GetStatus(status)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return status, err
}

func (us StatusService) UpdateStatus(status entity.Status) (entity.Status, error) {
	status, err := us.statusStorage.UpdateStatus(status)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return status, err
}

func (us StatusService) DeleteStatus(status entity.Status) {
	err := us.statusStorage.DeleteStatus(status)
	if err != nil {
		panic(err)
	}
}

