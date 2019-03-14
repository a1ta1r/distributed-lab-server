package service

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

type TaskService struct {
	taskStorage    storage.TaskStorage
	userStorage    storage.UserStorage
	projectStorage storage.ProjectStorage
	statusStorage  storage.StatusStorage
}

func NewTaskService(taskStorage storage.TaskStorage,
	userStorage storage.UserStorage,
	projectStorage storage.ProjectStorage,
	statusStorage storage.StatusStorage) TaskService {
	return TaskService{
		taskStorage:    taskStorage,
		userStorage:    userStorage,
		projectStorage: projectStorage,
		statusStorage:  statusStorage,
	}
}

func (us TaskService) AddTask(task entity.Task) entity.Task {
	task, err := us.taskStorage.AddTask(task)
	if err != nil {
		panic(err)
	}
	return task
}

func (us TaskService) GetTasks() []entity.Task {
	tasks, err := us.taskStorage.GetTasks()
	if err != nil {
		panic(err)
	}

	return tasks
}

func (us TaskService) GetTask(task entity.Task) (entity.Task, error) {
	task, err := us.taskStorage.GetTask(task)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return task, err
}

func (us TaskService) UpdateTask(task entity.Task) (entity.Task, error) {
	task, err := us.taskStorage.UpdateTask(task)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return task, err
}

func (us TaskService) DeleteTask(task entity.Task) {
	err := us.taskStorage.DeleteTask(task)
	if err != nil {
		panic(err)
	}
}
