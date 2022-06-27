package routes

import (
	"github.com/gin-gonic/gin"
	"ini/pkg/routes/auth"
	v1 "ini/pkg/routes/v1"
)

func StartHandler() {
	Router := gin.Default()
	api := Router.Group("/api")
	auth.InitAuth(api)
	v1.Init_v1(api)

	Router.Run(":8080")
}
