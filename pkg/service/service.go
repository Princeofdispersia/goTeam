package service

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/Princeofdispersia/goTeam/pkg/repository"
)

type Authorization interface {
	CreateUser(user goTeam.User) (int, string, error)
	GenerateToken(id int) (string, error)
	CheckSig(id int, sig string) bool
	ParseToken(token string) (int, error)
}

type Invite interface {
}

type Team interface {
	Create(userId int, team goTeam.Team) (int, error)
	GetAll(userId int) ([]goTeam.Team, error)
	GetById(userId int, teamId int) (goTeam.GetTeamResponse, error)
	Delete(userId int, teamId int) error
	Update(userId int, team goTeam.Team) error
}

type TeamRole interface {
	Create(userId int, teamRole goTeam.TeamRole) (int, error)
	GetAll(userId int, teamId int) ([]goTeam.TeamRole, error)
	GetById(userId int, teamRoleId int) (goTeam.TeamRole, error)
	Delete(userId int, teamRoleId int) error
	Update(userId int, teamRole goTeam.TeamRole) error
}

type Role interface {
	Create(userId int, role goTeam.Role) (int, error)
	GetAll(userId int, teamId int) ([]goTeam.GetRoleResponse, error)
	GetById(userId int, roleId int) (goTeam.GetRoleResponse, error)
	Delete(userId int, roleId int) error
	Update(userId int, role goTeam.Role) error
}

type Task interface {
	Create(userId int, task goTeam.Task) (int, error)
	GetAll(userId int, teamId int) ([]goTeam.Task, error)
	GetById(userId int, taskId int) (goTeam.Task, error)
	Delete(userId int, taskId int) error
	Update(userId int, task goTeam.Task) error
}

type Job interface {
	Create(userId int, job goTeam.Job) (int, error)
	GetAll(userId int, jobId int) ([]goTeam.Job, error)
	GetById(userId int, jobId int) (goTeam.Job, error)
	Delete(userId int, jobId int) error
	//Update(userId int, task goTeam.Task) error

}

type Done interface {
}

type Service struct {
	Authorization
	Invite
	Team
	TeamRole
	Role
	Task
	Job
	Done
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Team:          NewTeamService(repos.Team),
		TeamRole:      NewTeamRoleService(repos.TeamRole, repos.Team),
		Role:          NewRoleService(repos.Role, repos.Team, repos.TeamRole),
		Task:          NewTaskService(repos.Task, repos.Team),
		Job:           NewJobService(repos.Job, repos.Team, repos.Task),
	}
}
