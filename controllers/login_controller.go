package controllers

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/caiomp87/sword-health-challenge/models"
	"github.com/caiomp87/sword-health-challenge/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user, err := repository.UserRepository.FindByEmail(ctx, credentials.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found " + err.Error(),
		})
		return
	}

	decodedPwd, err := base64.StdEncoding.DecodeString(user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if credentials.Password != string(decodedPwd) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// call CreateToken
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func generateToken(user *models.User) (string, error) {
	accessSecret := os.Getenv("ACCESS_SECRET")

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["user_type"] = strings.ToLower(user.Type)
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString([]byte(accessSecret))
}
