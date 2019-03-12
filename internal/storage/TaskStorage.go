package storage

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
)

type TaskStorage struct {
	db gorm.DB
}

func NewTaskStorage(db gorm.DB) TaskStorage {
	return TaskStorage{db: db}
}

func (us TaskStorage) AddTask(task entity.Task) (entity.Task, error) {
	err := us.db.Save(&task).Error
	return task, err
}

func (us TaskStorage) GetTask(task entity.Task) (entity.Task, error) {
	err := us.db.First(&task, task.ID).Error
	return task, err
}

func (us TaskStorage) GetTasks() ([]entity.Task, error) {
	var tasks []entity.Task

	err := us.db.Find(&tasks).Error
	return tasks, err
}

func (us TaskStorage) UpdateTask(task entity.Task) (entity.Task, error) {
	err := us.db.Save(&task).Error
	return task, err
}

func (us TaskStorage) DeleteTask(task entity.Task) error {
	err := us.db.Delete(&task, task.ID).Error
	return err
}
