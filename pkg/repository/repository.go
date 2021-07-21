package repository

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user goTeam.User) (int, error)
}

type Team interface {
	Create(userId int, team goTeam.Team) (int, error)
	GetAll(userId int) ([]goTeam.Team, error)
	GetById(userId int, teamId int) (goTeam.GetTeamResponse, error)
	Delete(userId int, teamId int) error
	Update(userId int, team goTeam.Team) error
	IsModerator(userId int, teamId int) (bool, error)
	IsMember(userId int, teamId int) (bool, error)
}

type TeamRole interface {
	Create(userId int, teamRole goTeam.TeamRole) (int, error)
	GetAll(teamId int) ([]goTeam.TeamRole, error)
	GetById(teamRoleId int) (goTeam.TeamRole, error)
	Delete(teamRoleId int) error
	Update(teamRole goTeam.TeamRole) error
}

type Role interface {
	Create(userId int, role goTeam.Role) (int, error)
	GetAll(teamId int) ([]goTeam.GetRoleResponse, error)
	GetById(teamRoleId int) (goTeam.GetRoleResponse, error)
	Delete(teamRoleId int) error
	Update(teamRole goTeam.Role) error
}

type Task interface {
	Create(userId int, task goTeam.Task) (int, error)
	GetAll(teamId int) ([]goTeam.Task, error)
	GetById(taskId int) (goTeam.Task, error)
	Delete(taskId int) error
	Update(task goTeam.Task) error
}

type Job interface {
	Create(userId int, job goTeam.Job) (int, error)
	GetAll(teamId int) ([]goTeam.Job, error)
	GetById(jobId int) (goTeam.Job, error)
	Delete(jobId int) error
	//Update(job goTeam.Job) error

}

type Done interface {
}

type Repository struct {
	Authorization
	Team
	TeamRole
	Role
	Task
	Job
	Done
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Team:          NewTeamPostgres(db),
		TeamRole:      NewTeamRolePostgres(db),
		Role:          NewRolePostgres(db),
		Task:          NewTaskPostgres(db),
		Job:           NewJobPostgres(db),
	}
}
