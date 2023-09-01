package message

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type messageListResponse struct {
	StatusCode  int32            `json:"status_code"`
	StatusMsg   string           `json:"status_msg"`
	MessageList []models.Message `json:"message_list"`
}

func MessageList(c *gin.Context) {
	// 解析获取用户
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	_, err := auth.GetUserFromToken(token)
	if err != nil {
		resp := messageListResponse{
			StatusCode:  1,
			StatusMsg:   "token解析失败,携带正确的token",
			MessageList: nil,
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	userId, err := utils.ParseParamToInt(c, "user_id")

	//获取双方所有消息
	var messages []models.Message
	if err := utils.DB.Table("message").Where("user_id = ?", userId).
		Where("to_user_id = ?", toUserId).Order("created_at").Find(&messages).Error; err != nil {
		response := messageListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	//封装查询成功结果
	resp := messageListResponse{
		StatusCode:  0,
		StatusMsg:   "查询成功",
		MessageList: messages,
	}
	c.JSON(http.StatusOK, resp)
}
