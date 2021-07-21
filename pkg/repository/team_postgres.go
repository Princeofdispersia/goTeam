package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type TeamPostgres struct {
	db *sqlx.DB
}

func NewTeamPostgres(db *sqlx.DB) *TeamPostgres {
	return &TeamPostgres{db: db}
}

func (r *TeamPostgres) Create(userId int, team goTeam.Team) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var teamId int
	createTeamQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", teamTable)
	row := tx.QueryRow(createTeamQuery, team.Title)
	err = row.Scan(&teamId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var teamRoleId int
	createRoleCreator := fmt.Sprintf(
		"INSERT INTO %s (teamId, title, moderator) VALUES ($1, $2, $3) RETURNING id", teamRoleTable)
	row = tx.QueryRow(createRoleCreator, teamId, "Создатель", true)
	err = row.Scan(&teamRoleId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserRole := fmt.Sprintf("INSERT INTO %s (userId, teamRoleId) VALUES ($1, $2)", roleTable)
	_, err = tx.Exec(createUserRole, userId, teamRoleId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return teamId, tx.Commit()
}

func (r *TeamPostgres) GetAll(userId int) ([]goTeam.Team, error) {
	var teams []goTeam.Team
	getAllQuery := fmt.Sprintf(""+
		"SELECT t.id, t.title FROM %s r JOIN %s tr ON r.teamRoleId = tr.id JOIN %s t ON t.id=tr.teamId WHERE userId = $1;",
		roleTable, teamRoleTable, teamTable)
	err := r.db.Select(&teams, getAllQuery, userId)

	return teams, err
}

func (r *TeamPostgres) GetById(userId int, teamId int) (goTeam.GetTeamResponse, error) {
	var team goTeam.GetTeamResponse
	getByIdQuery := fmt.Sprintf(""+
		"SELECT t.id, t.title FROM %s r JOIN %s tr ON r.teamRoleId = tr.id JOIN %s t ON t.id=tr.teamId WHERE userId = $1 AND tr.teamId = $2;",
		roleTable, teamRoleTable, teamTable)
	err := r.db.Get(&team, getByIdQuery, userId, teamId)
	if err != nil {
		return goTeam.GetTeamResponse{}, err
	}
	getRolesQuery := fmt.Sprintf("SELECT t.id, t.title, t.moderator FROM %s t JOIN %s r ON r.teamroleid = t.id WHERE userId = $1 AND teamId = $2",
		teamRoleTable, roleTable)
	err = r.db.Select(&(team.TeamRoles), getRolesQuery, userId, teamId)

	return team, err
}

func (r *TeamPostgres) Delete(userId int, teamId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", teamTable)
	_, err := r.db.Exec(deleteQuery, teamId)

	return err
}

func (r *TeamPostgres) Update(userId int, team goTeam.Team) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET title = $1 WHERE id = $2", teamTable)
	_, err := r.db.Exec(updateQuery, team.Title, team.Id)

	return err
}

func (r *TeamPostgres) IsModerator(userId int, teamId int) (bool, error) {
	var count int
	isQuery := fmt.Sprintf("SELECT count(*) FROM %s tr JOIN %s r ON tr.id=r.teamRoleId WHERE r.userId = $1 AND tr.teamId = $2 AND tr.moderator = TRUE",
		teamRoleTable, roleTable)
	err := r.db.Get(&count, isQuery, userId, teamId)
	return count > 0, err
}

func (r *TeamPostgres) IsMember(userId int, teamId int) (bool, error) {
	var count int
	isQuery := fmt.Sprintf("SELECT count(*) FROM %s tr JOIN %s r ON tr.id=r.teamRoleId WHERE r.userId = $1 AND tr.teamId = $2",
		teamRoleTable, roleTable)
	err := r.db.Get(&count, isQuery, userId, teamId)
	return count > 0, err
}
