package user

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// user struct only for the response of userInfo,
// not the user struct in models/accounts.gp
type userResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	//followCount     int64
	IsFollow bool `json:"is_follow"`
	//avatar          string
	//backgroundImage string
	//signature       string
	//totalFavourited int64
	//workCount       int64
	//favouriteCount  int64
}

type userInfoResponse struct {
	StatusCode int32        `json:"status_code"`
	User       userResponse `json:"user"`
}

func Info(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	// find the user whose userId == user_id
	queryUser := models.User{UserId: userId}
	if err := utils.DB.First(&queryUser).Error; err != nil {
		response := userInfoResponse{
			StatusCode: -1,
		}
		fmt.Println(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// get the current user
	var isSubscriber bool
	user, err := auth.GetUserFromToken(c.PostForm("token"))
	if err != nil {
		fmt.Println(err)
		response := userInfoResponse{
			StatusCode: -1,
			User:       userResponse{},
		}
		c.JSON(http.StatusUnauthorized, response)
		return
	} else {
		if err := utils.DB.Preload("Subscribers").First(&user).Error; err != nil {
			fmt.Println("Error:", err)
			return
		} else {
			for _, subscriber := range user.Subscribers {
				if subscriber.UserId == queryUser.UserId {
					isSubscriber = true
					break
				}
			}
		}
	}
	response := userInfoResponse{
		StatusCode: 0,
		User: userResponse{
			Id:       userId,
			Name:     queryUser.Username,
			IsFollow: isSubscriber,
		},
	}
	c.JSON(http.StatusOK, response)
}
