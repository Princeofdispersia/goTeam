package service

import (
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type RoleService struct {
	teamRoleRepo repository.TeamRole
	teamRepo     repository.Team
	repo         repository.Role
}

func NewRoleService(repo repository.Role, teamRepo repository.Team, teamRoleRepo repository.TeamRole) *RoleService {
	return &RoleService{repo: repo, teamRepo: teamRepo, teamRoleRepo: teamRoleRepo}

}

func (s *RoleService) Create(userId int, role goTeam.Role) (int, error) {
	teamRole, err := s.teamRoleRepo.GetById(role.TeamRoleId)
	if err != nil {
		return 0, err
	}
	mod, err := s.teamRepo.IsModerator(userId, teamRole.TeamId)
	if err != nil {
		return 0, err
	}
	if !mod {
		return 0, errors.New("access denied: you are not a moderator")
	}

	return s.repo.Create(userId, role)
}

func (s *RoleService) GetAll(userId int, teamId int) ([]goTeam.GetRoleResponse, error) {
	memb, err := s.teamRepo.IsMember(userId, teamId)
	if err != nil {
		return nil, err
	}
	if !memb {
		return nil, errors.New("access denied: you are not a member")
	}

	return s.repo.GetAll(teamId)
}

func (s *RoleService) GetById(userId int, roleId int) (goTeam.GetRoleResponse, error) {
	role, err := s.repo.GetById(roleId)
	if err != nil {
		return goTeam.GetRoleResponse{}, err
	}
	memb, err := s.teamRepo.IsMember(userId, role.TeamId)
	if err != nil {
		return goTeam.GetRoleResponse{}, err
	}
	if !memb {
		return goTeam.GetRoleResponse{}, errors.New("access denied: you are not a member")
	}

	return role, nil

}

func (s *RoleService) Delete(userId int, roleId int) error {
	role, err := s.repo.GetById(roleId)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, role.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}

	return s.repo.Delete(roleId)

}

func (s *RoleService) Update(userId int, role goTeam.Role) error {
	oldRole, err := s.repo.GetById(role.Id)
	if err != nil {
		return err
	}
	oldMemb, err := s.teamRepo.IsModerator(userId, oldRole.TeamId)
	if err != nil {
		return err
	}
	if !oldMemb {
		return errors.New("access denied: you are not a moderator")
	}
	team, err := s.teamRoleRepo.GetById(role.TeamRoleId)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, team.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}
	if role.UserId == 0 {
		role.UserId = oldRole.UserId
	}
	if role.TeamRoleId == 0 {
		role.TeamRoleId = oldRole.TeamRoleId
	}

	return s.repo.Update(role)
}
