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

func (r *JobPostgres) Create(userId int, task goTeam.Task) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (teamId, title, description, deadline) VALUES ($1, $2, $3, $4) RETURNING id",
		taskTable)
	err := r.db.Get(&id, createQuery, task.TeamId, task.Title, task.Description, task.Deadline)
	return id, err
}

func (r *JobPostgres) GetAll(teamId int) ([]goTeam.Task, error) {
	var tasks []goTeam.Task
	getAllQuery := fmt.Sprintf("SELECT * FROM %s WHERE teamId = $1", taskTable)
	err := r.db.Select(&tasks, getAllQuery, teamId)
	return tasks, err
}

func (r *JobPostgres) GetById(taskId int) (goTeam.Task, error) {
	var task goTeam.Task
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", taskTable)
	err := r.db.Get(&task, getQuery, taskId)
	return task, err
}

func (r *JobPostgres) Delete(taskId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", taskTable)
	_, err := r.db.Exec(deleteQuery, taskId)

	return err
}

/*
func (r *JobPostgres) Update(task goTeam.Task) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET (title, description, deadline) = ($2, $3, $4) WHERE id = $1", taskTable)
	_, err := r.db.Exec(updateQuery, task.Id, task.Title, task.Description, task.Deadline)

	return err
}
*/
