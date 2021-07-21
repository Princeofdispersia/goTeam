package handler

import (
	"github.com/Princeofdispersia/goTeam"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createJob(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input goTeam.Job
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Id != 0 {
		newErrorResponse(c, http.StatusBadRequest, "job id should not be specified")
		return
	}
	jobId, err := h.service.Job.Create(userId.(int), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"jobId": jobId,
	})

}

func (h *Handler) getAllJobs(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid job id param")
		return
	}

	jobs, err := h.service.Job.GetAll(userId.(int), jobId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func (h *Handler) getJobById(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid job id param")
		return
	}
	job, err := h.service.Job.GetById(userId.(int), jobId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, job)
}

/*
	func (h *Handler) updateJob(c *gin.Context) {

	type updateJobStruct struct {
		Title string `json:"title"`
		Description string `json:"description"`
		Deadline int `json:"deadline"`
	}
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	taskId, err:= strconv.Atoi(c.Param("taskId"))
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
*/

func (h *Handler) deleteJob(c *gin.Context) {
	userId, ok := c.Get(idCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	jobId, err := strconv.Atoi(c.Param("jobId"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid job id param")
		return
	}

	err = h.service.Job.Delete(userId.(int), jobId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, map[string]string{"Status": "Ok"})
	}
}
