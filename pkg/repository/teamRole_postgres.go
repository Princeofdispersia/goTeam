package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type TeamRolePostgres struct {
	db *sqlx.DB
}

func NewTeamRolePostgres(db *sqlx.DB) *TeamRolePostgres {
	return &TeamRolePostgres{db: db}
}

func (r *TeamRolePostgres) Create(userId int, teamRole goTeam.TeamRole) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (teamId, title, moderator) VALUES ($1, $2, $3) RETURNING id",
		teamRoleTable)
	err := r.db.Get(&id, createQuery, teamRole.TeamId, teamRole.Title, teamRole.Moderator)
	return id, err
}

func (r *TeamRolePostgres) GetAll(teamId int) ([]goTeam.TeamRole, error) {
	var teamRoles []goTeam.TeamRole
	getAllQuery := fmt.Sprintf("SELECT * FROM %s tr WHERE teamId = $1", teamRoleTable)
	err := r.db.Select(&teamRoles, getAllQuery, teamId)
	return teamRoles, err
}

func (r *TeamRolePostgres) GetById(teamRoleId int) (goTeam.TeamRole, error) {
	var teamRole goTeam.TeamRole
	getQuery := fmt.Sprintf("SELECT * FROM %s tr WHERE id = $1", teamRoleTable)
	err := r.db.Get(&teamRole, getQuery, teamRoleId)
	return teamRole, err
}

func (r *TeamRolePostgres) Delete(teamRoleId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", teamRoleTable)
	_, err := r.db.Exec(deleteQuery, teamRoleId)

	return err
}

func (r *TeamRolePostgres) Update(teamRole goTeam.TeamRole) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET (teamId, title, moderator) = ($2, $3, $4) WHERE id = $1", teamRoleTable)
	_, err := r.db.Exec(updateQuery, teamRole.Id, teamRole.TeamId, teamRole.Title, *teamRole.Moderator)

	return err
}
