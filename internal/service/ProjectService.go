package service

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

type ProjectService struct {
	projectStorage storage.ProjectStorage
	userStorage    storage.UserStorage
	teamStorage    storage.TeamStorage
}

func NewProjectService(projectStorage storage.ProjectStorage,
	userStorage storage.UserStorage,
	teamStorage storage.TeamStorage) ProjectService {
	return ProjectService{
		projectStorage: projectStorage,
		userStorage:    userStorage,
		teamStorage:    teamStorage,
	}
}

func (us ProjectService) AddProject(project entity.Project) entity.Project {
	//TODO добавить хандлинг ошибок
	project.Team, _ = us.teamStorage.GetTeam(project.Team)
	project.Owner, _ = us.userStorage.GetUser(project.Owner)
	project.Members, _ = us.userStorage.GetUsers(project.Members)

	project, err := us.projectStorage.AddProject(project)
	if err != nil {
		panic(err)
	}
	return project
}

func (us ProjectService) GetProjects() []entity.Project {
	projects, err := us.projectStorage.GetProjects()
	if err != nil {
		panic(err)
	}

	return projects
}

func (us ProjectService) GetProject(project entity.Project) (entity.Project, error) {
	project, err := us.projectStorage.GetProject(project)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return project, err
}

func (us ProjectService) UpdateProject(project entity.Project) (entity.Project, error) {
	project, err := us.projectStorage.UpdateProject(project)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return project, err
}

func (us ProjectService) DeleteProject(project entity.Project) {
	err := us.projectStorage.DeleteProject(project)
	if err != nil {
		panic(err)
	}
}
