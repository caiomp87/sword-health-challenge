package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/caiomp87/sword-health-challenge/models"
	"github.com/caiomp87/sword-health-challenge/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTask(c *gin.Context) {
	var taskRaw *models.Task
	if err := c.BindJSON(&taskRaw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := repository.UserRepository.FindByID(ctx, taskRaw.UserID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found " + err.Error(),
		})
		return
	}

	if strings.EqualFold(user.Type, "manager") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tasks can only be assigned to technicians",
		})
		return
	}

	task := &models.Task{
		ID:      uuid.New().String(),
		Name:    taskRaw.Name,
		Summary: taskRaw.Summary,
		UserID:  taskRaw.UserID,
	}

	if err := repository.TaskRepository.Create(ctx, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "task created successfully",
	})
}

func GetTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	tasks, err := repository.TaskRepository.FindAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	task, err := repository.TaskRepository.FindByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	var task *models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repository.TaskRepository.UpdateByID(ctx, id, task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repository.TaskRepository.DeleteByID(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}

func DoneTask(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repository.TaskRepository.Done(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task done successfully",
	})
}
