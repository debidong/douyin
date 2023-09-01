package relation

import (
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type actionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type userFollower struct {
	UserID     int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
	IsDelete   int64 `json:"is_delete"`
}

func Action(c *gin.Context) {
	// get token
	token := c.Query("token")
	// parse token to get username
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		rep := actionResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
		}
		c.JSON(http.StatusInternalServerError, rep)
		return
	}
	// get to_user_id from request
	toUserId, err := utils.ParseParamToInt(c, "to_user_id")
	if err != nil {
		resp := actionResponse{
			StatusCode: 1,
			StatusMsg:  "参数转换失败，传递信息错误",
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	actionType, err := utils.ParseParamToInt(c, "action_type")
	if err != nil {
		resp := actionResponse{
			StatusCode: 1,
			StatusMsg:  "参数转换失败，传递信息错误",
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	if actionType == 1 {
		// 1为关注选项
		uf := userFollower{
			UserID:     int64(toUserId),
			FollowerID: user.UserId,
		}
		result := utils.DB.First(&uf, "is_delete = ?", 1)
		if result.Error == gorm.ErrRecordNotFound {
			// 记录不存在，表示一定是没创建过
			uf.IsDelete = 0
			err = utils.DB.Create(&uf).Error
			if err != nil {
				resp := actionResponse{
					StatusCode: 1,
					StatusMsg:  "记录不存在，但插入失败",
				}
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		} else {
			uf.IsDelete = 0
			if err = utils.DB.Table("user_followers").Where("user_id = ?", toUserId).Where("follower_id = ?", user.UserId).Update("is_delete", 0).Error; err != nil {
				fmt.Println(err)
				resp := actionResponse{
					StatusCode: 1,
					StatusMsg:  "数据库已有数据，为删除数据，但更新失败",
				}
				c.JSON(http.StatusInternalServerError, resp)
				return
			}

		}

		// 使用Redis维护关注数量和粉丝数量
		// todo
		resp := actionResponse{
			StatusCode: 1,
			StatusMsg:  "关注成功",
		}
		c.JSON(http.StatusOK, resp)
	} else if actionType == 2 {
		// 2为取消关注
		if err = utils.DB.Table("user_followers").Where("user_id = ? and follower_id = ?", int64(toUserId), user.UserId).Update("is_delete", 1).Error; err != nil {
			resp := actionResponse{
				StatusCode: 1,
				StatusMsg:  "数据库更新失败",
			}
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		resp := actionResponse{
			StatusCode: 1,
			StatusMsg:  "取消成功",
		}
		c.JSON(http.StatusOK, resp)

	} else {
		// 参数不合法
		resp := actionResponse{
			StatusCode: 1,
			StatusMsg:  "参数不合法",
		}
		c.JSON(http.StatusInternalServerError, resp)
	}

}
