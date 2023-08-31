package relation

import (
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type actionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Action(c *gin.Context) {
	// get token
	token := c.GetHeader("token")
	// parse token to get username
	username, err := auth.GetUserFromToken(token)
	if err != nil {
		rep := actionResponse{
			StatusCode: 1,
			StatusMsg:  "token解析失败,携带正确的token",
		}
		c.JSON(http.StatusInternalServerError, rep)
		return
	}
	// get to_user_id from request
	toUserId := c.Query("toUserId")
	actionType := c.Query("actionType")

	fmt.Println(username)
	fmt.Println(toUserId)
	fmt.Println(actionType)

}
