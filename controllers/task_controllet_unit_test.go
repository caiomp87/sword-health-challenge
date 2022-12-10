package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiomp87/sword-health-challenge/mock"
	"github.com/caiomp87/sword-health-challenge/models"
	"github.com/caiomp87/sword-health-challenge/repository"
	"github.com/caiomp87/sword-health-challenge/routes"
	"github.com/caiomp87/sword-health-challenge/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.ReleaseMode)

	t.Run("failed - bind json", func(t *testing.T) {
		expectMsgError := `{"error":"json: cannot unmarshal number into Go struct field Task.name of type string"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		app.Use(func(c *gin.Context) {})
		routes.AddRoutes(app)

		var jsonBody = struct {
			Name int
		}{
			123,
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Equal(expectMsgError, w.Body.String())
	})

	t.Run("failed - get context value", func(t *testing.T) {
		expectMsgError := `{"error":"userID not provided"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		app.Use(func(c *gin.Context) {})
		routes.AddRoutes(app)

		jsonBody := models.Task{
			ID:      uuid.New().String(),
			Name:    "Task 1",
			Summary: "Summary 1",
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusInternalServerError, w.Code)
		assert.Equal(expectMsgError, w.Body.String())
	})

	t.Run("failed - find user by id", func(t *testing.T) {
		expectMsgError := `{"error":"user not found no sql rows in result"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		c.Set("userID", "123")

		app.Use(func(c *gin.Context) {
			c.Set("userID", "123")
		})
		routes.AddRoutes(app)

		jsonBody := models.Task{
			ID:      uuid.New().String(),
			Name:    "Task 1",
			Summary: "Summary 1",
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		var user *models.User

		iUserMock := new(mock.IUser)
		iUserMock.On("FindByID", tmock.Anything, tmock.Anything).Return(user, errors.New("no sql rows in result"))
		repository.UserRepository = iUserMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusNotFound, w.Code)
		assert.Equal(expectMsgError, w.Body.String())
	})

	t.Run("failed - user type equal manager", func(t *testing.T) {
		expectMsgError := `{"error":"tasks can only be assigned to technicians"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		c.Set("userID", "123")

		app.Use(func(c *gin.Context) {
			c.Set("userID", "123")
		})
		routes.AddRoutes(app)

		jsonBody := models.Task{
			ID:      uuid.New().String(),
			Name:    "Task 1",
			Summary: "Summary 1",
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		user := &models.User{
			Type: utils.UserType.Manager,
		}

		iUserMock := new(mock.IUser)
		iUserMock.On("FindByID", tmock.Anything, tmock.Anything).Return(user, nil)
		repository.UserRepository = iUserMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Equal(expectMsgError, w.Body.String())
	})

	t.Run("failed - create task", func(t *testing.T) {
		expectMsgError := `{"error":"error to create in database"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		c.Set("userID", "123")

		app.Use(func(c *gin.Context) {
			c.Set("userID", "123")
		})
		routes.AddRoutes(app)

		jsonBody := models.Task{
			ID:      uuid.New().String(),
			Name:    "Task 1",
			Summary: "Summary 1",
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		user := &models.User{
			Type: utils.UserType.Technician,
		}

		iUserMock := new(mock.IUser)
		iUserMock.On("FindByID", tmock.Anything, tmock.Anything).Return(user, nil)
		repository.UserRepository = iUserMock

		iTaskMock := new(mock.ITask)
		iTaskMock.On("Create", tmock.Anything, tmock.Anything).Return(errors.New("error to create in database"))
		repository.TaskRepository = iTaskMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusInternalServerError, w.Code)
		assert.Equal(expectMsgError, w.Body.String())
	})

	t.Run("success - task created", func(t *testing.T) {
		expectMsg := `{"message":"task created successfully"}`

		w := httptest.NewRecorder()
		c, app := gin.CreateTestContext(w)
		c.Set("userID", "123")

		app.Use(func(c *gin.Context) {
			c.Set("userID", "123")
		})
		routes.AddRoutes(app)

		jsonBody := models.Task{
			ID:      uuid.New().String(),
			Name:    "Task 1",
			Summary: "Summary 1",
		}

		data, _ := json.Marshal(jsonBody)
		reader := bytes.NewReader(data)

		user := &models.User{
			Type: utils.UserType.Technician,
		}

		iUserMock := new(mock.IUser)
		iUserMock.On("FindByID", tmock.Anything, tmock.Anything).Return(user, nil)
		repository.UserRepository = iUserMock

		iTaskMock := new(mock.ITask)
		iTaskMock.On("Create", tmock.Anything, tmock.Anything).Return(nil)
		repository.TaskRepository = iTaskMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/v1/task/", reader)

		app.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusCreated, w.Code)
		assert.Equal(expectMsg, w.Body.String())
	})
}
