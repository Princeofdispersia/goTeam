package goTeam

type Team struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" binding:"required" db:"title"`
}

type Role struct {
	Id         int `json:"id"`
	UserId     int `json:"userId"`
	TeamRoleId int `json:"teamRoleId"`
}

type Task struct {
	Id          int    `json:"id"`
	TeamId      int    `json:"teamId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
	IsActual    *bool  `json:"isActual"`
}

type TeamRole struct {
	Id        int    `json:"id"`
	TeamId    int    `json:"teamId"`
	Title     string `json:"title"`
	Moderator *bool  `json:"moderator"`
}

type Job struct {
	Id     int `json:"id"`
	TaskId int `json:"taskId"`
	RoleId int `json:"roleId"`
}

type Done struct {
	Id     int  `json:"id"`
	JobId  int  `json:"jobId"`
	UserId int  `json:"userId"`
	IsDone bool `json:"isDone"`
}
