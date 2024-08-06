package controller

import (
	"net/http"
	"strconv"
	"task_management/data"
	"task_management/model"

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
	tasks, err := tm.TaskManagement.GetAllTasks()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (tm *TaskManagementStruct) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	number, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}
	task, err := tm.TaskManagement.GetTaskById(number)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "no task with this id"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

///// done

func (tm *TaskManagementStruct) UpdateTaskById(ctx *gin.Context) {
	// id := ctx.Param("id")
	// _, err := strconv.Atoi(id)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
	// 	return
	// }

	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := tm.TaskManagement.UpdateTaskById(&task)
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

	var task model.Task
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
	var task model.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := tm.TaskManagement.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
