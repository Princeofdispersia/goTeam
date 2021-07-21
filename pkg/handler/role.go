package handler

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createRole(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input goTeam.Role
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Id != 0 {
		newErrorResponse(c, http.StatusBadRequest, "roleId should not be specified")
		return
	}
	roleId, err := h.service.Role.Create(userId.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"roleId": roleId,
	})

}

func (h *Handler) getAllRoles(c *gin.Context) {
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

	var roles []goTeam.GetRoleResponse
	roles, err = h.service.Role.GetAll(userId.(int), teamId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *Handler) getRoleById(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id param")
		return
	}
	role, err := h.service.Role.GetById(userId.(int), roleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, role)
}

func (h *Handler) updateRole(c *gin.Context) {
	type updateRoleStruct struct {
		TeamRoleId int `json:"teamRoleId"`
	}

	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	rolesId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id param")
		return
	}

	var input updateRoleStruct
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	uRole := goTeam.Role{
		Id:         rolesId,
		UserId:     0,
		TeamRoleId: input.TeamRoleId,
	}
	err = h.service.Role.Update(userId.(int), uRole)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}

func (h *Handler) deleteRole(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	roleId, err := strconv.Atoi(c.Param("roleId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid role id param")
		return
	}

	err = h.service.Role.Delete(userId.(int), roleId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}
