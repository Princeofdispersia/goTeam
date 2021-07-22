package handler

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create team
// @Security ApiKeyAuth
// @Tags teams
// @Description create new team
// @ID create-team
// @Accept  json
// @Produce  json
// @Param input body goTeam.Team true "team info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teams [post]
func (h *Handler) createTeam(c *gin.Context) {
	id, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input goTeam.Team
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	teamId, err := h.service.Team.Create(id.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"teamId": teamId,
	})
}

// @Summary Get teams
// @Security ApiKeyAuth
// @Tags teams
// @Description get all teams
// @ID get-teams
// @Accept  json
// @Produce  json
// @Success 200 {object} []goTeam.Team
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teams [get]
func (h *Handler) getAllTeams(c *gin.Context) {
	id, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	resp, err := h.service.Team.GetAll(id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Get team
// @Security ApiKeyAuth
// @Tags teams
// @Description get team by id
// @ID get-team-by-id
// @Accept  json
// @Produce  json
// @Param id path integer true "Team id"
// @Success 200 {object} goTeam.GetTeamResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/team/{id} [get]
func (h *Handler) getTeamById(c *gin.Context) {
	id, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamId, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	team, err := h.service.Team.GetById(id.(int), teamId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, team)
}

// @Summary Update team
// @Security ApiKeyAuth
// @Tags teams
// @Description update team info
// @ID update-team
// @Accept  json
// @Produce  json
// @Param id path integer true "Team id"
// @Param input body goTeam.Team true "update team info"
// @Success 200 {object} goTeam.StatusOk
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/teams/{id} [put]
func (h *Handler) updateTeam(c *gin.Context) {
	id, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamId, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input goTeam.Team
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Id != 0 && input.Id != teamId {
		newErrorResponse(c, http.StatusBadRequest, "invalid request")
		return
	}
	input.Id = teamId
	err = h.service.Team.Update(id.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}

}

// @Summary Delete team
// @Security ApiKeyAuth
// @Tags teams
// @Description delete team by id
// @ID delete-team-by-id
// @Accept  json
// @Produce  json
// @Param id path integer true "Team id"
// @Success 200 {object} goTeam.StatusOk
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/team/{id} [delete]
func (h *Handler) deleteTeam(c *gin.Context) {
	id, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	teamId, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.service.Team.Delete(id.(int), teamId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}

}
