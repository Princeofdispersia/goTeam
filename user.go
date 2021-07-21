package goTeam

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" binding:"required"`
	Sig  string `json:"sig"`
}
