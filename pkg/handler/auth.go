package handler

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body name true "User name"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input goTeam.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, token, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"token": token,
	})
}

// @Summary SignIn
// @Tags auth
// @Description get token for existing account
// @ID login
// @Accept  json
// @Produce  json
// @Param input body sig true "Signature"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]

func (h *Handler) signIn(c *gin.Context) {
	var input goTeam.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if !h.service.Authorization.CheckSig(input.Id, input.Sig) {
		newErrorResponse(c, http.StatusForbidden, "wrong signature")
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    input.Id,
		"token": token,
	})
}
