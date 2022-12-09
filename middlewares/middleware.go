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

			c.Set("userID", userID)
			c.Set("userType", userType)

			var isGranted bool
			switch userType {
			case "manager":
				for _, permission := range managerPermissions {
					if interceptorValue == permission {
						isGranted = true
						break
					}
				}
			case "technician":
				for _, permission := range technicianPermissions {
					if interceptorValue == permission {
						isGranted = true
						break
					}
				}
			default:
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
		}

		c.Next()
	}
}
