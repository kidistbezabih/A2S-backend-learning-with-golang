package router

import (
	"example/task-management/controller"
	"example/task-management/data"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, tm data.TaskManagement) {
	r.POST("/tasks", controller.CreateTask(tm))
	r.GET("/tasks", controller.GetAllTasks(tm))
	r.GET("/tasks/:id", controller.GetTaskById(tm))
	r.PUT("/tasks/:id", controller.UpdateTaskById(tm))
	r.DELETE("/tasks/:id", controller.DeleteTaskById(tm))
}
