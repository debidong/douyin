package controllers

import (
	"douyin/views/message"
	"douyin/views/relation"
	"douyin/views/video"
	"github.com/gin-gonic/gin"
)
import "douyin/views/user"

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("user/login/", user.Login)
	apiRouter.GET("/user/", user.Info)

	apiRouter.GET("/favorite/list/", video.GetFavouriteVideo)
	apiRouter.GET("/publish/list/", video.GetPublishedVideo)
	apiRouter.POST("/publish/action/", video.PublishVideo)

	apiRouter.POST("/relation/action/", relation.Action)
	apiRouter.GET("/relation/follower/list/", relation.FollowerList)
	apiRouter.GET("/relation/follow/list/", relation.FollowList)
	apiRouter.GET("/relation/friend/list/", relation.FriendList)
	apiRouter.GET("/message/chat/", message.MessageList)
	apiRouter.POST("/message/action/", message.MessageAction)
}
