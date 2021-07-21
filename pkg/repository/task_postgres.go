package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) Create(userId int, task goTeam.Task) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (teamId, title, description, deadline) VALUES ($1, $2, $3, $4) RETURNING id",
		taskTable)
	err := r.db.Get(&id, createQuery, task.TeamId, task.Title, task.Description, task.Deadline)
	return id, err
}

func (r *TaskPostgres) GetAll(teamId int) ([]goTeam.Task, error) {
	var tasks []goTeam.Task
	getAllQuery := fmt.Sprintf("SELECT * FROM %s WHERE teamId = $1", taskTable)
	err := r.db.Select(&tasks, getAllQuery, teamId)
	return tasks, err
}

func (r *TaskPostgres) GetById(taskId int) (goTeam.Task, error) {
	var task goTeam.Task
	getQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", taskTable)
	err := r.db.Get(&task, getQuery, taskId)
	return task, err
}

func (r *TaskPostgres) Delete(taskId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", taskTable)
	_, err := r.db.Exec(deleteQuery, taskId)

	return err
}

func (r *TaskPostgres) Update(task goTeam.Task) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET (title, description, deadline) = ($2, $3, $4) WHERE id = $1", taskTable)
	_, err := r.db.Exec(updateQuery, task.Id, task.Title, task.Description, task.Deadline)

	return err
}
