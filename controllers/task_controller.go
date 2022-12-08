package controllers

import "github.com/gin-gonic/gin"

func CreateTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "post",
	})
}

func GetTasks(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "gets",
	})
}

func GetTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get",
	})
}

func UpdateTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "patch",
	})
}

func DeleteTask(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "delete",
	})
}
