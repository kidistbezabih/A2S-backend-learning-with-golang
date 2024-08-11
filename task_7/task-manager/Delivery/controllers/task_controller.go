package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidistbezabih/task-manager/Domain"
	usecases "github.com/kidistbezabih/task-manager/Usecases"
)

type TaskController struct {
	taskusecase usecases.TaskUsecase
}

func NewTaskController(u usecases.TaskUsecase) TaskController {
	return TaskController{
		taskusecase: u,
	}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	username, ok := c.Get("Username")
	role, roleExists := c.Get("Role")
	var tasks []Domain.Task
	var err error

	if !roleExists {
		c.JSON(http.StatusNotFound, gin.H{"message": "role not found"})
		return
	}

	if role == "admin" {
		tasks, err = tc.taskusecase.GetAllTasks()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	} else {
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		tasks, err = tc.taskusecase.GetAUserTasks(username.(string))
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}
	}
	c.JSON(http.StatusOK, tasks)
}

func (tm *TaskController) UpdateTaskById(controllers *gin.Context) {
	id := controllers.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		controllers.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	var task Domain.Task
	if err := controllers.ShouldBindJSON(&task); err != nil {
		controllers.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = tm.taskusecase.UpdateTaskById(task)
	if err != nil {
		controllers.JSON(http.StatusNotFound, err.Error())
		return
	}
	controllers.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (tm *TaskController) DeleteTaskById(controllers *gin.Context) {
	id := controllers.Param("id")
	number, err := strconv.Atoi(id)
	if err != nil {
		controllers.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var task Domain.Task
	if err := controllers.ShouldBindJSON(&task); err != nil {
		controllers.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = tm.taskusecase.DeleteTaskById(number)
	if err != nil {
		controllers.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	controllers.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (tm *TaskController) CreateTask(ctx *gin.Context) {
	var task Domain.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "status bad request"})
		return
	}
	err := tm.taskusecase.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
