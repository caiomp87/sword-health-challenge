package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/caiomp87/sword-health-challenge/models"
	"github.com/caiomp87/sword-health-challenge/repository"
	"github.com/caiomp87/sword-health-challenge/utils"
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

	userID, err := utils.GetContextValue(c, "userID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := repository.UserRepository.FindByID(ctx, userID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found " + err.Error(),
		})
		return
	}

	if strings.EqualFold(user.Type, utils.UserType.Manager) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tasks can only be assigned to technicians",
		})
		return
	}

	task := &models.Task{
		ID:      uuid.New().String(),
		Name:    taskRaw.Name,
		Summary: taskRaw.Summary,
		UserID:  userID,
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
	userType, err := utils.GetContextValue(c, "userType")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var (
		tasks   []*models.Task
		userID  string
		findErr error
	)

	switch userType {
	case strings.ToLower(utils.UserType.Manager):
		tasks, findErr = repository.TaskRepository.FindAll(ctx)
	case strings.ToLower(utils.UserType.Technician):
		userID, findErr = utils.GetContextValue(c, "userID")
		if findErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": findErr.Error(),
			})
			return
		}

		tasks, findErr = repository.TaskRepository.FindAllByUserID(ctx, userID)
	}

	if findErr != nil {
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

	userType, err := utils.GetContextValue(c, "userType")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var (
		task    *models.Task
		userID  string
		findErr error
	)
	switch userType {
	case strings.ToLower(utils.UserType.Manager):
		task, findErr = repository.TaskRepository.FindByID(ctx, id)
	case strings.ToLower(utils.UserType.Technician):
		userID, findErr = utils.GetContextValue(c, "userID")
		if findErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": findErr.Error(),
			})
			return
		}

		task, findErr = repository.TaskRepository.FindByIDAndUserID(ctx, id, userID)
	}

	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": findErr.Error(),
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

	userID, err := utils.GetContextValue(c, "userID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repository.TaskRepository.UpdateByID(ctx, id, userID, task); err != nil {
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

	userID, err := utils.GetContextValue(c, "userID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := repository.TaskRepository.Done(ctx, id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task done successfully",
	})
}
