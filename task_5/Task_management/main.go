package main

import (
	"context"
	"fmt"
	"log"
	"task_management/controller"
	"task_management/data"
	"task_management/router"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server               *gin.Engine
	ctx                  context.Context
	taskcollection       *mongo.Collection
	mongoclient          *mongo.Client
	err                  error
	taskmanagement       data.TaskManagement
	taskmanagementstruct controller.TaskManagementStruct
	taskmanager          router.TaskManager
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")

	taskcollection = mongoclient.Database("taskdb").Collection("tasks")
	taskmanagement = data.NewTaskService(taskcollection, ctx)
	taskmanagementstruct = controller.New(taskmanagement)
	taskmanager = router.NewTaskSer(&taskmanagementstruct)
	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	taskmanager.RegisterTaskRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
