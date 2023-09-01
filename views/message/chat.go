package message

import (
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type messageListResponse struct {
	StatusCode  int32     `json:"status_code"`
	StatusMsg   string    `json:"status_msg"`
	MessageList []Message `json:"message_list"`
}

type Message struct {
	ID          int `json:"id"`
	CreatedAt   time.Time
	MessageID   int64  `json:"message_id"`
	ToUserID    int64  `json:"to_user_id"`
	FromUserID  int64  `json:"from_user_iD"`
	Content     string `json:"content"`
	CreatedTime string `json:"created_time"`
}

func MessageList(c *gin.Context) {
	// 解析获取用户
	token := c.Query("token")
	user, err := auth.GetUserFromToken(token)
	toUserId, err := utils.ParseParamToInt(c, "to_user_id")
	/*createTime := c.Query("pre_msg_time")*/
	if err != nil {
		resp := messageActionResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
		}
		c.JSON(http.StatusInternalServerError, resp)
	}
	userId := user.UserId
	toId := int64(toUserId)

	//获取双方所有消息
	var messages []Message
	if err := utils.DB.Table("messages").Where("from_user_id in (?)", []int64{userId, toId}).
		/*Where("create_at > ").*/
		Where("to_user_id in (?)", []int64{userId, toId}).
		Order("created_at").
		Find(&messages).Error; err != nil {
		response := messageListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}
	//获取双方所有消息
	/*var messages []models.Message
	if err := utils.DB.Table("messages").Where("from_user_id = ?", userId).
		Where("to_user_id = ?", toId).Order("created_at").Find(&messages).Error; err != nil {
		response := messageListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}*/
	fmt.Println(len(messages))

	//封装查询成功结果
	for i := range messages {
		messages[i].CreatedTime = messages[i].CreatedAt.Format("15:04 2006/01/02")
	}

	resp := messageListResponse{
		StatusCode:  0,
		StatusMsg:   "查询成功",
		MessageList: messages,
	}
	c.JSON(http.StatusOK, resp)
}
