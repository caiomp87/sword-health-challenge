package routes

import (
	"github.com/caiomp87/sword-health-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(app *gin.Engine) {
	app.POST("/task", controllers.CreateTask)
	app.GET("/task/:id", controllers.GetTask)
	app.GET("/task", controllers.GetTasks)
	app.PATCH("/task/:id", controllers.UpdateTask)
	app.DELETE("/task/:id", controllers.DeleteTask)
}
