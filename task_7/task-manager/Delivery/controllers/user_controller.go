package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidistbezabih/task-manager/Domain"
	usecases "github.com/kidistbezabih/task-manager/Usecases"
)

type UserController struct {
	userusecase usecases.UserUseCase
}

func NewUserController(userusecase usecases.UserUseCase) UserController {
	return UserController{
		userusecase: userusecase,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user Domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.userusecase.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var user Domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := uc.userusecase.Login(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "welcome", "token": token})
}

func (uc *UserController) Promote(c *gin.Context) {
	username := c.Param("username")
	var user Domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err := uc.userusecase.Promote(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully prompted"})
}
