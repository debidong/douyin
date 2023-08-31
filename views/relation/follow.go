package relation

import (
	"douyin/models"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type followListResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg"`
	UserList   []models.User `json:"user_list"`
}

func FollowList(c *gin.Context) {
	// get token
	token := c.Query("token")
	username, err := auth.GetUserFromToken("token")
	if err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	userId := c.Query("user_id")
	fmt.Println(userId)
	fmt.Println(username)
	fmt.Println(token)
	resp := followListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList: []models.User{{
			Model:           gorm.Model{},
			Username:        "",
			UserId:          0,
			Password:        "",
			Subscribers:     nil,
			Followers:       nil,
			FavouriteVideos: nil,
			PublishedVideos: nil,
		}, {
			Model:           gorm.Model{},
			Username:        "",
			UserId:          1,
			Password:        "",
			Subscribers:     nil,
			Followers:       nil,
			FavouriteVideos: nil,
			PublishedVideos: nil,
		}},
	}
	c.JSON(http.StatusOK, resp)

}
