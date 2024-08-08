package controllers

import (
	"net/http"
	"task_management/data"
	"task_management/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userservices data.UserServices
}

func NewUserController(userservices data.UserServices) UserController {
	return UserController{
		userservices: userservices,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.userservices.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}

func (uc *UserController) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	token, err := uc.userservices.Login(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "welcome", "token": token})
}

func (uc *UserController) Promote(c *gin.Context) {
	username := c.Param("username")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	err := uc.userservices.Promote(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "page not found"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully prompted"})
}
