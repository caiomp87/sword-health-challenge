package routes

import (
	"github.com/caiomp87/sword-health-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(app *gin.Engine) {
	v1 := app.Group("v1")
	{
		user := v1.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}

		login := v1.Group("login")
		{
			login.POST("/", controllers.Login)
		}

		task := v1.Group("task")
		{
			task.POST("/", controllers.CreateTask)
			task.GET("/:id", controllers.GetTask)
			task.GET("/", controllers.GetTasks)
			task.PATCH("/:id", controllers.UpdateTask)
			task.DELETE("/:id", controllers.DeleteTask)
			task.PATCH("/done/:id", controllers.DoneTask)
		}
	}
}
