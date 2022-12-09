package routes

import (
	"github.com/caiomp87/sword-health-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoutes(app *gin.Engine) {
	v1 := app.Group("v1")
	{
		login := v1.Group("login")
		{
			login.POST("/", controllers.Login)
		}

		user := v1.Group("user")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/", controllers.GetUsers)
			user.GET("/:id", controllers.GetUser)
		}

		task := v1.Group("task")
		{
			task.POST("/", controllers.CreateTask)
			task.GET("/", controllers.GetTasks)
			task.GET("/:id", controllers.GetTask)
			task.PATCH("/:id", controllers.UpdateTask)
			task.DELETE("/:id", controllers.DeleteTask)
			task.PATCH("/done/:id", controllers.DoneTask)
		}
	}
}
