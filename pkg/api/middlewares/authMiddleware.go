package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"ini/pkg/api/models"
	"ini/pkg/api/services"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//authHeader := c.Request.Header.Get("Authorization")
		tokenString, err := c.Cookie("token")

		if err != nil {

			if errors.Is(err, http.ErrNoCookie) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status_code": 401,
					"message":     "Cookie not found",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 400,
				"message":     "Bad Request",
			})
			c.Abort()
			return
		}

		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return services.JwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status_code": 401,
					"message":     "Invalid Jwt token",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": 400,
				"message":     "Bad Request",
			})
			c.Abort()
			return
		}

		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"message":     "invalid JWT token",
			})
			c.Abort()
			return
		}

		//if authHeader == "" {
		//	c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
		//		"status_code": 203,
		//		"message":     "Request header `Authorization` Empty",
		//	})
		//	c.Abort()
		//	return
		//}

		//c.Set("uuid", uuid)

		c.Next()
	}
}
