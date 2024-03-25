package controller

import (
	"mygram/helper"
	"mygram/model"
	"mygram/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser model.User

	if err := bindJSON(c, &newUser); err != nil {
		return
	}

	createdUser, err := service.UserService.Register(&newUser)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID,
		"username": createdUser.Username,
		"age":      createdUser.Age,
		"email":    createdUser.Email,
	})
}

func bindJSON(c *gin.Context, payload interface{}) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		errHandler := helper.UnprocessibleEntity("Invalid JSON body")
		c.JSON(errHandler.Status(), errHandler)
		return err
	}
	return nil
}

func Login(context *gin.Context) {
	var user model.LoginCredential

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Login(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": result})
}

func UpdateUser(context *gin.Context) {
	var update model.UserUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := helper.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	updatedUser, err := service.UserService.UpdateUser(userID, &update)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         updatedUser.ID,
		"email":      updatedUser.Email,
		"username":   updatedUser.Username,
		"age":        updatedUser.Age,
		"updated_at": updatedUser.UpdatedAt,
	})
}

func DeleteUser(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, err := service.UserService.DeleteUser(userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your Account has been successfully deleted",
	})
}
