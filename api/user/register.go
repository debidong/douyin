package user

import (
	"douyin/models"
	"douyin/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type response struct {
	StatusCode int32  `json:"status_code"`
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Generate UID (random string with length = 64)
	uid := rand.Int63()

	user := models.User{
		Username: username,
		Password: password,
		UID:      uid,
	}
	utils.DB.Create(&user)

	response := response{
		StatusCode: 0,
		UserId:     uid,
		Token:      "test",
	}
	c.JSON(http.StatusOK, response)
}
