package middlewares

import (
	"errors"
	"strings"

	"github.com/caiomp87/sword-health-challenge/utils"
	"github.com/gin-gonic/gin"
)

func formatInterceptorValues(c *gin.Context) string {
	return c.Request.Method + c.FullPath()
}

func isPublic(interceptedRoutes map[string]string, interceptorValue string) bool {
	return interceptedRoutes[interceptorValue] == "public"
}

func isGranted(userType, interceptorValue string) (bool, error) {
	var isGranted bool
	switch {
	case strings.EqualFold(userType, utils.UserType.Manager):
		for _, permission := range managerPermissions {
			if interceptorValue == permission {
				isGranted = true
				break
			}
		}
	case strings.EqualFold(userType, utils.UserType.Technician):
		for _, permission := range technicianPermissions {
			if interceptorValue == permission {
				isGranted = true
				break
			}
		}
	default:
		return isGranted, errors.New("unsupported user type")
	}

	return isGranted, nil
}
