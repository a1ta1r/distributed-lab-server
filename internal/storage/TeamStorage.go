package storage

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
)

type TeamStorage struct {
	db gorm.DB
}

func NewTeamStorage(db gorm.DB) TeamStorage {
	return TeamStorage{db: db}
}

func (us TeamStorage) AddTeam(team entity.Team) (entity.Team, error) {
	err := us.db.Save(&team).Error
	return team, err
}

func (us TeamStorage) GetTeam(team entity.Team) (entity.Team, error) {
	err := us.db.First(&team, team.ID).Error
	return team, err
}

func (us TeamStorage) GetTeams() ([]entity.Team, error) {
	var teams []entity.Team

	err := us.db.Find(&teams).Error
	return teams, err
}

func (us TeamStorage) UpdateTeam(team entity.Team) (entity.Team, error) {
	err := us.db.Save(&team).Error
	return team, err
}

func (us TeamStorage) DeleteTeam(team entity.Team) error {
	err := us.db.Delete(&team, team.ID).Error
	return err
}

