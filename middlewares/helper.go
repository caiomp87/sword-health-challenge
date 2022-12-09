package middlewares

import (
	"github.com/gin-gonic/gin"
)

func formatInterceptorValues(c *gin.Context) string {
	return c.Request.Method + c.FullPath()
}

func isPublic(interceptedRoutes map[string]string, interceptorValue string) bool {
	return interceptedRoutes[interceptorValue] == "public"
}
