package handler

import (
	_ "github.com/Princeofdispersia/goTeam/docs"
	"github.com/Princeofdispersia/goTeam/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		teams := api.Group("/teams")
		{
			teams.POST("/", h.createTeam)
			teams.GET("/", h.getAllTeams)
			teams.GET("/:teamId", h.getTeamById)
			teams.PUT("/:teamId", h.updateTeam)
			teams.DELETE("/:teamId", h.deleteTeam)

			getAll := teams.Group("/:teamId")
			{
				getAll.GET("/teamRoles", h.getAllTeamRoles)
				getAll.GET("/roles", h.getAllRoles)
				getAll.GET("/tasks", h.getAllTasks)
				getAll.GET("/jobs", h.getAllJobs)
				getAll.GET("/dones", h.getAllDones)
			}
			teamRoles := api.Group("/teamRoles")
			{
				teamRoles.POST("/", h.createTeamRole)
				teamRoles.GET("/:teamRoleId", h.getTeamRoleById)
				teamRoles.PUT("/:teamRoleId", h.updateTeamRole)
				teamRoles.DELETE("/:teamRoleId", h.deleteTeamRole)
			}

			roles := api.Group("/roles")
			{
				roles.POST("/", h.createRole)
				roles.GET("/:roleId", h.getRoleById)
				roles.PUT("/:roleId", h.updateRole)
				roles.DELETE("/:roleId", h.deleteRole)
			}

			tasks := api.Group("/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/:taskId", h.getTaskById)
				tasks.PUT("/:taskId", h.updateTask)
				tasks.DELETE("/:taskId", h.deleteTask)
			}

			jobs := api.Group("/jobs")
			{
				jobs.POST("/", h.createJob)
				jobs.GET("/:jobId", h.getJobById)
				//jobs.PUT("/:jobId", h.updateJob)
				jobs.DELETE("/:jobId", h.deleteJob)
			}

			dones := api.Group("/dones")
			{
				dones.POST("/", h.createDone)
				dones.GET("/:doneId", h.GetDoneById)
				dones.PUT("/:doneId", h.updateDone)
				dones.DELETE("/:doneId", h.deleteDone)
			}
		}

		invites := api.Group("/invites")
		{
			invites.POST("/", h.createInvite)
			invites.GET("/", h.getAllInvites)
			invites.GET("/:inviteId", h.GetInviteById)
			invites.PUT("/:inviteId", h.updateInvite)
			invites.DELETE("/:inviteId", h.deleteInvite)
		}

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
