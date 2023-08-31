package relation

import (
	"douyin/utils"
	"douyin/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type followerListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type UserFollower struct {
	ID               int64  `json:"id"`
	UserID           int64  `json:"user_id"`
	FollowerID       int64  `json:"follower_id"`
	FollowerUsername string `json:"follower_username"`
}

func FollowerList(c *gin.Context) {
	// get token
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
	// todo

	var followers []UserFollower
	utils.DB.Table("user").Where("user_id = ?", userId).Find(&followers)

	var userFollowerResponses []User
	var userList []User
	for _, follower := range followers {
		userList = append(userFollowerResponses, User{
			ID:              follower.FollowerID,
			Name:            follower.FollowerUsername,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        false,
			Avatar:          "",
			BackgroundImage: "",
			Signature:       "",
			TotalFavorited:  0,
			WorkCount:       0,
			FavoriteCount:   0,
		})
	}
	resp := followerListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   userList,
	}
	c.JSON(http.StatusOK, resp)

}
