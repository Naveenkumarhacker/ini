package auth

import (
	"github.com/gin-gonic/gin"
	"ini/pkg/api/controllers"
)

func InitAuth(route *gin.RouterGroup) {
	controllers.AuthRoutes(route)
}
