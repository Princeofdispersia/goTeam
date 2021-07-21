package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type JobPostgres struct {
	db *sqlx.DB
}

func NewJobPostgres(db *sqlx.DB) *JobPostgres {
	return &JobPostgres{db: db}
}

func (r *JobPostgres) Create(userId int, job goTeam.Job) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (taskId, teamRoleId) VALUES ($1, $2) RETURNING id",
		jobTable)
	err := r.db.Get(&id, createQuery, job.TaskId, job.RoleId)
	return id, err
}

func (r *JobPostgres) GetAll(teamId int) ([]goTeam.Job, error) {
	var jobs []goTeam.Job
	getAllQuery := fmt.Sprintf("SELECT * FROM %s WHERE teamId = $1", jobTable)
	err := r.db.Select(&jobs, getAllQuery, teamId)
	return jobs, err
}

func (r *JobPostgres) GetById(taskId int) (goTeam.Job, error) {
	var job goTeam.Job
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", jobTable)
	err := r.db.Get(&job, getQuery, taskId)
	return job, err
}

func (r *JobPostgres) Delete(jobId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", jobTable)
	_, err := r.db.Exec(deleteQuery, jobId)

	return err
}

/*
func (r *JobPostgres) Update(task goTeam.Task) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET (title, description, deadline) = ($2, $3, $4) WHERE id = $1", taskTable)
	_, err := r.db.Exec(updateQuery, task.Id, task.Title, task.Description, task.Deadline)

	return err
}
*/
