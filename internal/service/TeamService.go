package service

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

type TeamService struct {
	teamStorage storage.TeamStorage
	userStorage storage.UserStorage
}

func NewTeamService(teamStorage storage.TeamStorage, userStorage storage.UserStorage) TeamService {
	return TeamService{
		teamStorage: teamStorage,
		userStorage: userStorage,
	}
}

func (us TeamService) AddTeam(team entity.Team) entity.Team {
	//TODO добавить хандлинг ошибок
	team.Owner, _ = us.userStorage.GetUser(team.Owner)
	team.Users, _ = us.userStorage.GetUsers(team.Users)

	team, err := us.teamStorage.AddTeam(team)
	if err != nil {
		panic(err)
	}
	return team
}

func (us TeamService) GetTeams() []entity.Team {
	teams, err := us.teamStorage.GetTeams()
	if err != nil {
		panic(err)
	}

	return teams
}

func (us TeamService) GetTeam(team entity.Team) (entity.Team, error) {
	team, err := us.teamStorage.GetTeam(team)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return team, err
}

func (us TeamService) UpdateTeam(team entity.Team) (entity.Team, error) {
	team, err := us.teamStorage.UpdateTeam(team)
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	return team, err
}

func (us TeamService) DeleteTeam(team entity.Team) {
	err := us.teamStorage.DeleteTeam(team)
	if err != nil {
		panic(err)
	}
}
