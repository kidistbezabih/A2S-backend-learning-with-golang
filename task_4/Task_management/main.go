package main

import (
	"example/task-management/data"
	"example/task-management/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	s := data.NewTaskService()

	router.SetupRouter(r, s)

	r.Run("localhost:8080")
}
