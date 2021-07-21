package service

import (
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type TaskService struct {
	teamRepo repository.Team
	repo     repository.Task
}

func NewTaskService(repo repository.Task, teamRepo repository.Team) *TaskService {
	return &TaskService{repo: repo, teamRepo: teamRepo}

}

func (s *TaskService) Create(userId int, task goTeam.Task) (int, error) {
	mod, err := s.teamRepo.IsModerator(userId, task.TeamId)
	if err != nil {
		return 0, err
	}
	if !mod {
		return 0, errors.New("access denied: you are not a moderator")
	}
	return s.repo.Create(userId, task)
}

func (s *TaskService) GetAll(userId int, teamId int) ([]goTeam.Task, error) {
	memb, err := s.teamRepo.IsMember(userId, teamId)
	if err != nil {
		return nil, err
	}
	if !memb {
		return nil, errors.New("access denied: you are not a member")
	}

	return s.repo.GetAll(teamId)
}

func (s *TaskService) GetById(userId int, taskId int) (goTeam.Task, error) {
	task, err := s.repo.GetById(taskId)
	if err != nil {
		return goTeam.Task{}, err
	}
	memb, err := s.teamRepo.IsMember(userId, task.TeamId)
	if err != nil {
		return goTeam.Task{}, err
	}
	if !memb {
		return goTeam.Task{}, errors.New("access denied: you are not a member")
	}

	return task, nil

}

func (s *TaskService) Delete(userId int, taskId int) error {
	task, err := s.repo.GetById(taskId)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, task.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}

	return s.repo.Delete(taskId)

}

func (s *TaskService) Update(userId int, task goTeam.Task) error {
	oldTask, err := s.repo.GetById(task.Id)
	if err != nil {
		return err
	}
	memb, err := s.teamRepo.IsModerator(userId, oldTask.TeamId)
	if err != nil {
		return err
	}
	if !memb {
		return errors.New("access denied: you are not a moderator")
	}

	return s.repo.Update(task)
}
