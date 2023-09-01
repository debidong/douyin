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

// UserResponse user struct only for the response of userInfo,
// not the user struct in models/accounts.gp
type UserResponse struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	//avatar          string
	//backgroundImage string
	//signature       string
	//totalFavourited int64
	//workCount       int64
	//favouriteCount  int64
}

type userInfoResponse struct {
	StatusCode int32        `json:"status_code"`
	User       UserResponse `json:"user"`
}

func Info(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userId)
	// find the user whose userId == user_id
	var queryUser models.User
	if err := utils.DB.Where("user_id = ?", userId).First(&queryUser).Error; err != nil {
		response := userInfoResponse{
			StatusCode: -1,
		}
		fmt.Println(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	// get the current user
	var isSubscriber bool
	token := c.Query("token")
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		fmt.Println(err)
		response := userInfoResponse{
			StatusCode: -1,
			User:       UserResponse{},
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
		User: UserResponse{
			Id:       userId,
			Name:     queryUser.Username,
			IsFollow: isSubscriber,
		},
	}
	c.JSON(http.StatusOK, response)
}
