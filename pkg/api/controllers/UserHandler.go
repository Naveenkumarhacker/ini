package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ini/pkg/api/models"
	"ini/pkg/api/services"
	"ini/pkg/utils"
	"net/http"
)

func UserRoutes(v1 *gin.RouterGroup) {
	v1.GET("/users", GetUsers)
	v1.GET("/user/:id", GetUserById)
	v1.POST("/user", CreateUser)
	v1.PUT("/user", UpdateUser)
	v1.DELETE("/user/:id", DeleteUser)
}

var validate = validator.New()

func GetUsers(c *gin.Context) {

	var users, err = services.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user, err = services.UserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})

}

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&newUser); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	var user, err = services.UserService.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})

}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}

	var updatedUser, err = services.UserService.UpdateUser(user)

	if errors.Is(err, utils.UserNotFoundErr) {
		c.JSON(http.StatusNotFound, models.UserResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var err = services.UserService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "user data deleted successfully"}})

}
