package router

import (
	"task_management/controller"

	"github.com/gin-gonic/gin"
)

type TaskManager struct {
	taskmanager *controller.TaskManagementStruct
}

func NewTaskSer(tm *controller.TaskManagementStruct) TaskManager {
	return TaskManager{
		taskmanager: tm,
	}
}

// RegisterTaskRoutes registers routes related to task management.
func (tm *TaskManager) RegisterTaskRoutes(rg *gin.RouterGroup) {
	taskRoute := rg.Group("/tasks")

	taskRoute.POST("/", tm.taskmanager.CreateTask)
	taskRoute.GET("/get/:id", tm.taskmanager.GetTaskById)
	taskRoute.GET("/getall", tm.taskmanager.GetAllTasks)
	taskRoute.PATCH("/update", tm.taskmanager.UpdateTaskById)
	taskRoute.DELETE("/delete/:id", tm.taskmanager.DeleteTaskById)
}
