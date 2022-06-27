package controllers

import (
	"github.com/gin-gonic/gin"
	"ini/pkg/api/models"
	"ini/pkg/api/services"
	"net/http"
)

func AuthRoutes(route *gin.RouterGroup) {
	route.POST("/login", Login)
	route.POST("/signup", SignUp)
}

func Login(c *gin.Context) {
	var credential models.Credential
	if err := c.BindJSON(&credential); err != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	jwtToken, err := services.AuthService.GenerateJWT(credential)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.SetCookie("token", jwtToken.Token, jwtToken.ExpireAt.Second(), "/", "localhost", false, true)
}

func SignUp(c *gin.Context) {

}
