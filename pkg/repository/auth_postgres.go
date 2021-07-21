package repository

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user goTeam.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
