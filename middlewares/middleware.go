package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		interceptorValue := formatInterceptorValues(c)

		if !isPublic(interceptedRoutes, interceptorValue) {
			token := c.GetHeader("authorization")
			if token == "" {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": "token not provided",
				})
				return
			}

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			})

			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			userID := claims["user_id"].(string)
			userType := claims["user_type"].(string)

			isGranted, err := isGranted(userType, interceptorValue)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "unsupported user type",
				})
				return
			}

			if !isGranted {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "user not allowed",
				})
				return
			}

			c.Set("userID", userID)
			c.Set("userType", userType)
		}

		c.Next()
	}
}
