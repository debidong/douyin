package relation

import (
	"douyin/models"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FriendListResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	FriendUser []models.User `json:"user_list"`
}

func FriendList(c *gin.Context) {
	// get token
	token := c.Query("token")
	username, err := auth.GetUserFromToken(token)
	if err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	id := c.Query("user_id")
	fmt.Println(id)
	fmt.Println(username)
	fmt.Println(token)
	resp := followListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   nil,
	}
	c.JSON(http.StatusOK, resp)

}
