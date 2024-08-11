package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kidistbezabih/task-manager/Delivery/controllers"
	"github.com/kidistbezabih/task-manager/Delivery/routers"
	repositories "github.com/kidistbezabih/task-manager/Repositories"
	usecases "github.com/kidistbezabih/task-manager/Usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx        context.Context
	server     *gin.Engine
	userrouter routers.UserRouter
	taskrouter routers.TaskRouter
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")

	taskcollection := mongoclient.Database("taskdb").Collection("tasks")
	taskrepo := repositories.NewTaskService(taskcollection, ctx)
	taskusecase := usecases.NewTaskUsecase(taskrepo)
	taskcontroller := controllers.NewTaskController(taskusecase)
	taskrouter = routers.NewTaskRouter(taskcontroller)

	usercollection := mongoclient.Database("userdb").Collection("users")
	userrepo := repositories.CreateNewUser(usercollection, ctx)
	userusecase := usecases.NewUseCase(userrepo)
	usercontroller := controllers.NewUserController(userusecase)
	userrouter = routers.NewUserRouter(usercontroller)
	server = gin.Default()

}

func main() {

	basepath := server.Group("/v1")

	userrouter.UserRouterRegister(basepath)
	taskrouter.RegisterTaskRouter(basepath)
	log.Fatal(server.Run(":9090"))
}
