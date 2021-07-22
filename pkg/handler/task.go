package handler

import (
	"fmt"
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"net/http"
	"strconv"
)

// @Summary Create task
// @Security ApiKeyAuth
// @Tags tasks
// @Description create new task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param input body goTeam.Task true "task info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks [post]
func (h *Handler) createTask(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input goTeam.Task
	err := c.BindJSON(&input)
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Id != 0 {
		newErrorResponse(c, http.StatusBadRequest, "task id should not be specified")
		return
	}
	taskId, err := h.service.Task.Create(userId.(int), input)
	if err != nil {
		fmt.Println(errors.Wrap(err, 2).ErrorStack())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"taskId": taskId,
	})

}

// @Summary Get all tasks
// @Security ApiKeyAuth
// @Tags tasks
// @Description get all tasks
// @ID get-tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} goTeam.Task
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks [get]
func (h *Handler) getAllTasks(c *gin.Context) {
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

	tasks, err := h.service.Task.GetAll(userId.(int), teamId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// @Summary Get task
// @Security ApiKeyAuth
// @Tags tasks
// @Description get task by id
// @ID get-task-by-id
// @Accept  json
// @Produce  json
// @Param id path integer true "Task id"
// @Success 200 {object} goTeam.Task
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [get]
func (h *Handler) getTaskById(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	task, err := h.service.Task.GetById(userId.(int), taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

type updateTaskStruct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int    `json:"deadline"`
}

// @Summary Update tasks
// @Security ApiKeyAuth
// @Tags tasks
// @Description update task information
// @ID update-tasks
// @Accept  json
// @Produce  json
// @Param id path integer true "Task id"
// @Param input body updateTaskStruct true "update task info"
// @Success 200 {object} goTeam.StatusOk
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [put]
func (h *Handler) updateTask(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	var input updateTaskStruct
	err = c.Bind(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	uTask := goTeam.Task{
		Id:          taskId,
		TeamId:      0,
		Title:       input.Title,
		Description: input.Description,
		Deadline:    input.Deadline,
	}
	err = h.service.Task.Update(userId.(int), uTask)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}

// @Summary Delete task
// @Security ApiKeyAuth
// @Tags tasks
// @Description delete task by id
// @ID delete-task-by-id
// @Accept  json
// @Produce  json
// @Param id path integer true "Task id"
// @Success 200 {object} goTeam.StatusOk
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/tasks/{id} [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	err = h.service.Task.Delete(userId.(int), taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}
