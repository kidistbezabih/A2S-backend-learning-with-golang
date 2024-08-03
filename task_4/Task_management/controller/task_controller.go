package controller

import (
	"example/task-management/data"
	"example/task-management/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tm data.TaskManagement) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, tm.GetTasks())
	}
}

func GetTaskById(tm data.TaskManagement) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")

		number, err := strconv.Atoi(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return
		}

		task, exist := tm.GetTaskById(number)

		if exist != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}

		ctx.JSON(http.StatusOK, task)
	}
}

func UpdateTaskById(tm data.TaskManagement) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		number, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return
		}

		var updatedTask model.Task

		if err := c.BindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = tm.UpdateTaskById(number, updatedTask)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		}
		c.JSON(http.StatusOK, gin.H{"message": "task updated"})
	}
}

func DeleteTaskById(ts data.TaskManagement) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		number, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
			return

		}

		err = ts.DeleteTaskById(number)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		}

		c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
	}
}

func CreateTask(st data.TaskManagement) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task model.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		task = st.CreateTask(task)
		c.JSON(http.StatusCreated, gin.H{"message": "task created"})
	}
}
