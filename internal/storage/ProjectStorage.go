package storage

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
)

type ProjectStorage struct {
	db gorm.DB
}

func NewProjectStorage(db gorm.DB) ProjectStorage {
	return ProjectStorage{db: db}
}

func (us ProjectStorage) AddProject(project entity.Project) (entity.Project, error) {
	err := us.db.Save(&project).Error
	return project, err
}

func (us ProjectStorage) GetProject(project entity.Project) (entity.Project, error) {
	err := us.db.First(&project, project.ID).Error
	return project, err
}

func (us ProjectStorage) GetProjects() ([]entity.Project, error) {
	var projects []entity.Project

	err := us.db.Find(&projects).Error
	return projects, err
}

func (us ProjectStorage) UpdateProject(project entity.Project) (entity.Project, error) {
	err := us.db.Save(&project).Error
	return project, err
}

func (us ProjectStorage) DeleteProject(project entity.Project) error {
	err := us.db.Delete(&project, project.ID).Error
	return err
}
