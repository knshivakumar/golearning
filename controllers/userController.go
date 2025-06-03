package controllers

import (
	"learning-golang/models"
	"learning-golang/services"
	"learning-golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	//Convert input data to json
	var user models.User

	//convert the playload from postman to model
	err := ctx.ShouldBindJSON(&user)
	user.PasswordHash, err = utils.HassPassword(user.PasswordHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	id, err := c.service.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		ctx.JSON(http.StatusOK, id)
	}

}

func (c *UserController) Login(ctx *gin.Context) {
	//Convert input data to Json
	var user models.User

	//convert the playload from postman to model
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	results, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		ctx.JSON(http.StatusOK, results)
	}
}
