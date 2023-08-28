package user

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginResponse struct {
	StatusCode int32  `json:"status_code"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{Username: username}

	// error retrieving user, user doesn't exist
	if err := utils.DB.First(&user).Error; err != nil {
		response := loginResponse{StatusCode: -1}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// error checking password, username and password mismatch
	if err != nil {
		response := loginResponse{StatusCode: -2}
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := auth.SetToken(username)

	if err != nil {
		fmt.Println(err)
		return
	}
	response := loginResponse{
		StatusCode: 0,
		UserId:     user.UserId,
		Token:      token,
	}

	c.JSON(http.StatusOK, response)
	return
}
