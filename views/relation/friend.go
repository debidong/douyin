package relation

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type friendListResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	FriendList []UserFriend `json:"user_list"`
}

type UserFriend struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func FriendList(c *gin.Context) {
	// 解析获取用户
	token := c.Query("token")
	_, err := auth.GetUserFromToken(token)
	if err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	userId, err := utils.ParseParamToInt(c, "user_id")

	//获取该用户的粉丝ID
	var followerId []int64
	if err := utils.DB.Table("user_followers").Select("follower_id").
		Where("user_id = ?", userId).Find(&followerId).Error; err != nil {
		response := friendListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	//用粉丝id查subscribers获取好友（互关为好友
	var friendId []int64
	if err := utils.DB.Table("user_subscribers").Select("subscriber_id").Where("subscriber_id in (?)", followerId).
		Where("user_id = ?", userId).Find(&friendId).Error; err != nil {
		response := friendListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	//获取好友的用户信息
	var friends []models.User
	if err := utils.DB.Table("users").Where("user_id in (?)", friendId).
		Find(&friends).Error; err != nil {
		response := friendListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	//组装查询结果
	var users []UserFriend
	for _, friend := range friends {
		users = append(users, UserFriend{
			ID:   friend.UserId,
			Name: friend.Username,
		})
	}
	fmt.Println(len(friends))
	fmt.Println(len(users))

	//封装查询成功结果
	resp := friendListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		FriendList: users,
	}
	c.JSON(http.StatusOK, resp)
}
