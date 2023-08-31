package relation

import (
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type followListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type UserFollow struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	FollowID       int64  `json:"follow_id"`
	FollowUsername string `json:"follow_username"`
}

func FollowList(c *gin.Context) {
	// get token
	token := c.Query("token")
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	userId, err := utils.ParseParamToInt(c, "user_id")
	//todo

	fmt.Println(userId)
	fmt.Println(user)
	fmt.Println(token)

	// 根据userid查询出所有的
	var userFollows []UserFollow
	utils.DB.Table("users").Where("user_id = ?", userId).Find(&userFollows)

	var userFollowResponses []User
	var resultList []User
	for _, userFollow := range userFollows {
		resultList = append(userFollowResponses, User{
			ID:              userFollow.FollowID,
			Name:            userFollow.FollowUsername,
			FollowCount:     0,
			FollowerCount:   0,
			IsFollow:        true,
			Avatar:          "",
			BackgroundImage: "",
			Signature:       "",
			TotalFavorited:  0,
			WorkCount:       0,
			FavoriteCount:   0,
		})
	}

	resp := followListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   resultList,
	}
	c.JSON(http.StatusOK, resp)

}
