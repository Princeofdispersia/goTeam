package service

import (
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) Create(userId int, team goTeam.Team) (int, error) {
	return s.repo.Create(userId, team)
}

func (s *TeamService) GetAll(userId int) ([]goTeam.Team, error) {
	return s.repo.GetAll(userId)

}

func (s *TeamService) GetById(userId int, teamId int) (goTeam.GetTeamResponse, error) {
	return s.repo.GetById(userId, teamId)
}

func (s *TeamService) Delete(userId int, teamId int) error {
	_, err := s.GetById(userId, teamId) // Проверяем, что команда существует и принадлежит пользователю
	if err != nil {
		return err
	}
	return s.repo.Delete(userId, teamId)
}

func (s *TeamService) Update(userId int, team goTeam.Team) error {
	_, err := s.GetById(userId, team.Id) // Проверяем, что команда существует и принадлежит пользователю
	if err != nil {
		return err
	}
	is, err := s.repo.IsModerator(userId, team.Id) // Проверяем, что пользователь – админ
	if err != nil || !is {
		return errors.New("access forbidden")
	}

	return s.repo.Update(userId, team)
}
