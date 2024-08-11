package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kidistbezabih/task-manager/Delivery/controllers"
	infrastructure "github.com/kidistbezabih/task-manager/Infrastructure"
)

type UserRouter struct {
	usesrcontroller controllers.UserController
}

func NewUserRouter(uc controllers.UserController) UserRouter {
	return UserRouter{
		usesrcontroller: uc,
	}
}

func (ur *UserRouter) UserRouterRegister(rg *gin.RouterGroup) {
	userrout := rg.Group("/users")

	userrout.POST("/register", ur.usesrcontroller.RegisterUser)
	userrout.POST("/login", ur.usesrcontroller.Login)
	userrout.PATCH("/promote/:username", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), ur.usesrcontroller.Promote)
}
