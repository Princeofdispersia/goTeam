package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type RolePostgres struct {
	db *sqlx.DB
}

func NewRolePostgres(db *sqlx.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) Create(userId int, role goTeam.Role) (int, error) {
	var id int
	createQuery := fmt.Sprintf("INSERT INTO %s (userId, teamRoleId) VALUES ($1, $2) RETURNING id",
		roleTable)
	err := r.db.Get(&id, createQuery, role.UserId, role.TeamRoleId)
	return id, err
}

func (r *RolePostgres) GetAll(teamId int) ([]goTeam.GetRoleResponse, error) {
	var roles []goTeam.GetRoleResponse
	getAllQuery := fmt.Sprintf(`SELECT r.id, userid, teamroleid, teamid, title AS teamRoleTitle, moderator
										FROM %s r JOIN %s tr ON r.teamroleid = tr.id 
										WHERE teamId = $1;`,
		roleTable, teamRoleTable)
	err := r.db.Select(&roles, getAllQuery, teamId)
	return roles, err
}

func (r *RolePostgres) GetById(roleId int) (goTeam.GetRoleResponse, error) {
	var role goTeam.GetRoleResponse
	getQuery := fmt.Sprintf(`SELECT r.id, userid, teamroleid, teamid, title AS teamRoleTitle, moderator
										FROM %s r JOIN %s tr ON r.teamroleid = tr.id 
										WHERE r.id = $1;`,
		roleTable, teamRoleTable)
	err := r.db.Get(&role, getQuery, roleId)
	return role, err
}

func (r *RolePostgres) Delete(roleId int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", roleTable)
	_, err := r.db.Exec(deleteQuery, roleId)

	return err
}

func (r *RolePostgres) Update(role goTeam.Role) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET (userId, teamRoleId) = ($2, $3) WHERE id = $1", roleTable)
	_, err := r.db.Exec(updateQuery, role.Id, role.UserId, role.TeamRoleId)

	return err
}
