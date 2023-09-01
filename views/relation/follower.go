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

	// 验证参数是否和token对应相同
	var userID int
	if userID, err = utils.ParseParamToInt(c, "user_id"); err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "参数转换失败，传递信息错误",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 查询对应的followerIDs
	var followerIDs []int64
	if err = utils.DB.Table("user_followers").Select("follower_id").Where("user_id = ? and is_delete = ?", userID, 0).Find(&followerIDs).Error; err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "数据库中未找到对应数据",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 查询出user_id 对应的全部粉丝
	var followers []User
	if err = utils.DB.Table("users").Where("user_id in (?)", followerIDs).Find(&followers).Error; err != nil {
		resp := followerListResponse{
			StatusCode: 1,
			StatusMsg:  "数据库中未找到对应数据",
			UserList:   nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// 还需要查出粉丝是否是关注状态
	for index := range followers {
		var count int64
		if err = utils.DB.Table("user_followers").Where("user_id = (?) and follower_id = (?) and is_delete = ?", followers[index].UserID, userID, 0).Count(&count).Error; err != nil {
			resp := followerListResponse{
				StatusCode: 1,
				StatusMsg:  "查询失败，后段错误",
				UserList:   nil,
			}
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if count >= 1 {
			followers[index].IsFollow = true
		}
	}

	resp := followerListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		UserList:   followers,
	}
	c.JSON(http.StatusOK, resp)

}
