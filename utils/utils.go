package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var UserType = struct {
	Manager    string
	Technician string
}{
	"Manager",
	"Technician",
}

func GetContextValue(c *gin.Context, key string) (string, error) {
	valueRaw, ok := c.Get(key)
	if !ok || valueRaw == nil {
		return "", errors.New(key + " not provided")
	}

	value, ok := valueRaw.(string)
	if !ok || value == "" {
		return "", errors.New(key + " not provided")
	}

	return value, nil
}
