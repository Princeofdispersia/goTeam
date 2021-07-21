package service

import (
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type TeamRoleService struct {
	teamRepo repository.Team
	repo     repository.TeamRole
}

func NewTeamRoleService(repo repository.TeamRole, teamRepo repository.Team) *TeamRoleService {
	return &TeamRoleService{repo: repo, teamRepo: teamRepo}

}

func (s *TeamRoleService) Create(userId int, teamRole goTeam.TeamRole) (int, error) {
	mod, err := s.teamRepo.IsModerator(userId, teamRole.TeamId)
	if err != nil {
		return 0, err
	}
	if !mod {
		return 0, errors.New("access denied: you are not a moderator")
	}

	return s.repo.Create(userId, teamRole)
}

func (s *TeamRoleService) GetAll(userId int, teamId int) ([]goTeam.TeamRole, error) {
	memb, err := s.teamRepo.IsMember(userId, teamId)
	if err != nil {
		return nil, err
	}
	if !memb {
		return nil, errors.New("access denied: you are not a member")
	}

	return s.repo.GetAll(teamId)
}

func (s *TeamRoleService) GetById(userId int, teamRoleId int) (goTeam.TeamRole, error) {
	teamRole, err := s.repo.GetById(teamRoleId)
	if err != nil {
		return goTeam.TeamRole{}, err
	}
	memb, err := s.teamRepo.IsMember(userId, teamRole.TeamId)
	if err != nil {
		return goTeam.TeamRole{}, err
	}
	if !memb {
		return goTeam.TeamRole{}, errors.New("access denied: you are not a member")
	}

	return teamRole, nil

}

func (s *TeamRoleService) Delete(userId int, teamRoleId int) error {
	teamRole, err := s.repo.GetById(teamRoleId)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, teamRole.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}

	return s.repo.Delete(teamRoleId)

}

func (s *TeamRoleService) Update(userId int, teamRole goTeam.TeamRole) error {
	oldTeamRole, err := s.repo.GetById(teamRole.Id)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, oldTeamRole.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}
	if teamRole.TeamId == 0 {
		teamRole.TeamId = oldTeamRole.TeamId
	}
	if teamRole.Title == "" {
		teamRole.Title = oldTeamRole.Title
	}
	if teamRole.Moderator == nil {
		teamRole.Moderator = oldTeamRole.Moderator
	}

	return s.repo.Update(teamRole)
}
