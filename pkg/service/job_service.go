package service

import (
	"errors"
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type JobService struct {
	repo     repository.Job
	teamRepo repository.Team
	taskRepo repository.Task
}

func NewJobService(repo repository.Job, teamRepo repository.Team, taskRepo repository.Task) *JobService {
	return &JobService{repo: repo, teamRepo: teamRepo, taskRepo: taskRepo}

}

func (s *JobService) Create(userId int, job goTeam.Job) (int, error) {
	task, err := s.taskRepo.GetById(job.TaskId)
	if err != nil {
		return 0, err
	}
	mod, err := s.teamRepo.IsModerator(userId, task.TeamId)
	if err != nil {
		return 0, err
	}
	if !mod {
		return 0, errors.New("access denied: you are not a moderator")
	}
	return s.repo.Create(userId, job)
}

func (s *JobService) GetAll(userId int, teamId int) ([]goTeam.Job, error) {
	memb, err := s.teamRepo.IsMember(userId, teamId)
	if err != nil {
		return nil, err
	}
	if !memb {
		return nil, errors.New("access denied: you are not a member")
	}

	return s.repo.GetAll(teamId)
}

func (s *JobService) GetById(userId int, jobId int) (goTeam.Job, error) {
	job, err := s.repo.GetById(jobId)
	if err != nil {
		return goTeam.Job{}, err
	}

	task, err := s.taskRepo.GetById(job.TaskId)
	if err != nil {
		return goTeam.Job{}, err
	}

	memb, err := s.teamRepo.IsMember(userId, task.TeamId)
	if err != nil {
		return goTeam.Job{}, err
	}

	if !memb {
		return goTeam.Job{}, errors.New("access denied: you are not a member")
	}

	return job, nil

}

func (s *JobService) Delete(userId int, jobId int) error {
	job, err := s.repo.GetById(jobId)
	if err != nil {
		return err
	}

	task, err := s.taskRepo.GetById(job.TaskId)
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

	return s.repo.Delete(jobId)

}

/*
func (s *JobService) Update(userId int, task goTeam.Task) error {
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
*/
