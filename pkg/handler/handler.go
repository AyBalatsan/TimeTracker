package handler

import (
	"github.com/AyBalatsan/TimeTracker/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	users := router.Group("/users")
	{
		users.POST("/", h.createUser)
		users.GET("/", h.getAllUsers)
		// users.GET("/:id", h.getUserByPassport)
		// users.GET("/:id", h.getUserTasksList)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}
	// tasks := router.Group("/tasks")
	// {
	// 	tasks.POST("/", h.createTask)
	// 	tasks.GET("/", h.getAllTasks)
	// 	tasks.PUT("/:id", h.updateTask)
	// 	tasks.DELETE("/:id", h.deleteTask)

	// 	management := tasks.Group(":id/management")
	// 	{
	// 		management.POST("/", h.userInstall)
	// 		management.PUT("/id/start", h.userStartTask)
	// 		management.PUT("/:id/stop", h.userStopTask)
	// 		management.PUT("/:id/done", h.userDoneTask)

	// 	}
	// }
	return router
}
