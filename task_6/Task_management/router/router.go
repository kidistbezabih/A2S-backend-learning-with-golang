package router

import (
	"task_management/controllers"
	"task_management/middleware"

	"github.com/gin-gonic/gin"
)

type TaskManager struct {
	taskmanager *controllers.TaskManagementStruct
}

func NewTaskSer(tm *controllers.TaskManagementStruct) TaskManager {
	return TaskManager{
		taskmanager: tm,
	}
}

type UserManager struct {
	taskmanager *controllers.UserController
}

func UserTaskSer(uc *controllers.UserController) UserManager {
	return UserManager{
		taskmanager: uc,
	}
}

// RegisterTaskRoutes registers routes related to task management.
func (tm *TaskManager) RegisterTaskRoutes(rg *gin.RouterGroup) {
	taskRoute := rg.Group("/tasks")

	taskRoute.POST("/", middleware.AuthMiddleware(), middleware.AdminMidleware(), tm.taskmanager.CreateTask)
	taskRoute.GET("/getall", middleware.AuthMiddleware(), tm.taskmanager.GetAllTasks)
	taskRoute.PATCH("/update/:id", middleware.AuthMiddleware(), middleware.AdminMidleware(), tm.taskmanager.UpdateTaskById)
	taskRoute.DELETE("/delete/:id", middleware.AuthMiddleware(), middleware.AdminMidleware(), tm.taskmanager.DeleteTaskById)
}

func (um *UserManager) RegisterUserRoutes(rg *gin.RouterGroup) {
	userrout := rg.Group("/users")

	userrout.POST("/register", um.taskmanager.RegisterUser)
	userrout.POST("/login", um.taskmanager.Login)
	userrout.PATCH("/promote/:username", um.taskmanager.Promote)
}
