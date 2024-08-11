package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kidistbezabih/task-manager/Delivery/controllers"
	infrastructure "github.com/kidistbezabih/task-manager/Infrastructure"
)

type TaskRouter struct {
	taskcontroller controllers.TaskController
}

func NewTaskRouter(taskcontroller controllers.TaskController) TaskRouter {
	return TaskRouter{
		taskcontroller: taskcontroller,
	}
}

func (tr *TaskRouter) RegisterTaskRouter(rg *gin.RouterGroup) {
	taskrout := rg.Group("/tasks")

	taskrout.POST("/", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), tr.taskcontroller.CreateTask)
	taskrout.GET("/getall", infrastructure.AuthMiddleware(), tr.taskcontroller.GetAllTasks)
	taskrout.PATCH("/update/:id", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), tr.taskcontroller.UpdateTaskById)
	taskrout.DELETE("/delete/:id", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), tr.taskcontroller.DeleteTaskById)
}
