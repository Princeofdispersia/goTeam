package handler

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createTeamRole(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input goTeam.TeamRole
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Id != 0 {
		newErrorResponse(c, http.StatusBadRequest, "teamRoleId should not be specified")
		return
	}
	teamRoleId, err := h.service.TeamRole.Create(userId.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"teamRoleId": teamRoleId,
	})

}

func (h *Handler) getAllTeamRoles(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamId, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid team id param")
		return
	}

	var teamRoles []goTeam.TeamRole
	teamRoles, err = h.service.TeamRole.GetAll(userId.(int), teamId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, teamRoles)
}

func (h *Handler) getTeamRoleById(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	teamRoleId, err := strconv.Atoi(c.Param("teamRoleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid teamRoleId id param")
		return
	}
	teamRole, err := h.service.TeamRole.GetById(userId.(int), teamRoleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, teamRole)
}

func (h *Handler) updateTeamRole(c *gin.Context) {
	type updateTeamRoleStruct struct {
		Title     string `json:"title"`
		Moderator bool   `json:"moderator"`
	}

	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamRolesId, err := strconv.Atoi(c.Param("teamRoleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid teamRole id param")
		return
	}

	var input updateTeamRoleStruct
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	uTeamRole := goTeam.TeamRole{
		Id:        teamRolesId,
		TeamId:    0,
		Title:     input.Title,
		Moderator: &input.Moderator,
	}

	err = h.service.TeamRole.Update(userId.(int), uTeamRole)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}

func (h *Handler) deleteTeamRole(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamRoleId, err := strconv.Atoi(c.Param("teamRoleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid team id param")
		return
	}

	err = h.service.TeamRole.Delete(userId.(int), teamRoleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}
