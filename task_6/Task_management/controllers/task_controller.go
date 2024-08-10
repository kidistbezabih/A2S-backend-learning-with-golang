package controllers

import (
	"net/http"
	"strconv"
	"task_management/data"
	"task_management/models"

	"github.com/gin-gonic/gin"
)

type TaskManagementStruct struct {
	TaskManagement data.TaskManagement
}

func New(TaskManagement data.TaskManagement) TaskManagementStruct {
	return TaskManagementStruct{
		TaskManagement: TaskManagement,
	}
}

func (tm *TaskManagementStruct) GetAllTasks(ctx *gin.Context) {
	// var username string
	username, ok := ctx.Get("Username")
	role, roleExists := ctx.Get("Role")
	var tasks []models.Task
	var err error

	if !roleExists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "status not found"})
		return
	}

	if role == "admin" {
		tasks, err = tm.TaskManagement.GetAllTasks()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		}
	} else {
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		}
		tasks, err = tm.TaskManagement.GetAUserTasks(username.(string))
		if err != nil {
			ctx.JSON(http.StatusNotFound, err)
		}
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (tm *TaskManagementStruct) UpdateTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = tm.TaskManagement.UpdateTaskById(&task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (tm *TaskManagementStruct) DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	number, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = tm.TaskManagement.DeleteTaskById(number)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (tm *TaskManagementStruct) CreateTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "status bad request"})
		return
	}
	err := tm.TaskManagement.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
