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
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	FriendList []models.User `json:"user_list"`
}

func FriendList(c *gin.Context) {
	// get token
	token := c.Query("token")
	userId := c.Query("user_id")
	username, err := auth.GetUserFromToken("token")
	if err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	fmt.Println(userId)
	fmt.Println(username)
	fmt.Println(token)
	var userList []models.User
	if err := utils.DB.Find(&userList).Error; err != nil {
		response := friendListResponse{StatusCode: -2}
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	resp := friendListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		FriendList: userList,
	}
	c.JSON(http.StatusOK, resp)
}
