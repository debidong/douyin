package relation

import (
	"douyin/utils"
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
		// todo:
		return
	}
	actionType, err := utils.ParseParamToInt(c, "action_type")
	if err != nil {
		// todo:
		return
	}

	fmt.Println(toUserId)
	fmt.Println(user)
	fmt.Println(actionType)

	if actionType == 1 {
		// 自己的关注列表 follow
		//userFollow := UserFollow{
		//	UserID:         user.UserId,
		//	FollowID:       int64(toUserId),
		//	FollowUsername: "hh",
		//}
		//utils.DB.Save(userFollow)

		// 自己的关注数量+1

		// 对应的人的粉丝列表 follower

		// 对应的人的粉丝数量 + 1

	} else if actionType == 2 {

		// 自己的关注列表 follow 标志位 -1

		// 自己的关注数量-1

		// 对应的人的粉丝列表 follower 标志位 -1

		// 对应的人的粉丝数量 - 1
	} else {
		// 参数不合法
	}

}
