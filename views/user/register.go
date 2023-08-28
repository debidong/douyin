package user

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
)

type registerResponse struct {
	StatusCode int32  `json:"status_code"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response := registerResponse{StatusCode: -3}
		c.JSON(http.StatusInternalServerError, response)
	}

	// Generate UID (random string with length = 64)
	uid := rand.Int63()

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		UID:      uid,
	}
	utils.DB.Create(&user)

	token, err := auth.SetToken(username)

	if err != nil {
		fmt.Println(err)
		return
	}
	response := registerResponse{
		StatusCode: 0,
		UserId:     uid,
		Token:      token,
	}
	c.JSON(http.StatusOK, response)
	return
}
