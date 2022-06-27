package v1

import (
	"github.com/gin-gonic/gin"
	"ini/pkg/api/controllers"
	"ini/pkg/api/middlewares"
)

func Init_v1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1.Use(middlewares.AuthMiddleware())
	controllers.UserRoutes(v1)
}
