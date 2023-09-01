package relation

import (
	"douyin/utils"
	"douyin/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type followListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserList   []User `json:"user_list"`
}

type User struct {
	UserID          int64  `json:"id"`
	Username        string `json:"name"`
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

func FollowList(c *gin.Context) {
	// get token
	token := c.Query("token")
	_, err := auth.GetUserFromToken(token)
	if err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	var userID int
	if userID, err = utils.ParseParamToInt(c, "user_id"); err != nil {
		resp := followListResponse{
			StatusCode: 1,
			StatusMsg:  "参数转换失败，传递信息错误",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 根据userid查询出所有的关注者id
	var followIDs []int64
	if err = utils.DB.Table("user_followers").Select("user_id").Where("follower_id = ?", userID).
		Find(&followIDs).Error; err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "数据库中未找到对应数据",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 根据所有的关注者id，查询user全部信息
	var followUsers []User
	if err := utils.DB.Table("users").Where("user_id in (?)", followIDs).
		Find(&followUsers).Error; err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "数据库中未找到对应数据",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 封装返回类型
	for i := range followUsers {
		followUsers[i].IsFollow = true
	}

	resp := followListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   followUsers,
	}
	c.JSON(http.StatusOK, resp)
}
