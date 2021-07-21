package goTeam

type GetTeamResponseRole struct {
	Id        int    `json:"roleId"`
	Title     string `json:"roleTitle"`
	Moderator bool   `json:"moderator"`
}

type GetTeamResponse struct {
	Id        int                   `json:"teamId"`
	Title     string                `json:"title"`
	TeamRoles []GetTeamResponseRole `json:"teamRoles"`
}

type GetRoleResponse struct {
	Id            int    `json:"roleId"`
	UserId        int    `json:"userId"`
	TeamRoleId    int    `json:"teamRoleId"`
	TeamId        int    `json:"teamId"`
	TeamRoleTitle string `json:"teamRoleTitle"`
	Moderator     bool   `json:"moderator"`
}
