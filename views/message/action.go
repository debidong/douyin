package message

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type messageActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func MessageAction(c *gin.Context) {
	// 解析获取用户
	token := c.Query("token")
	user, err := auth.GetUserFromToken(token)
	toUserId, err := utils.ParseParamToInt(c, "to_user_id")
	content := c.Query("content")
	if err != nil {
		resp := messageActionResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
		}
		c.JSON(http.StatusInternalServerError, resp)
	}

	rand.Seed(time.Now().UnixNano())
	mid := rand.Int63()

	message := models.Message{
		ToUserID:   int64(toUserId),
		FromUserID: user.UserId,
		Content:    content,
		MessageID:  mid,
	}

	if err = utils.DB.Create(&message).Error; err != nil {
		resp := messageActionResponse{
			StatusCode: 2,
			StatusMsg:  "插入失败",
		}
		c.JSON(http.StatusInternalServerError, resp)
	}

	//封装查询成功结果
	resp := messageActionResponse{
		StatusCode: 0,
		StatusMsg:  "发出成功",
	}
	c.JSON(http.StatusOK, resp)
}
