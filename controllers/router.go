package controllers

import (
	"douyin/views/video"
	"github.com/gin-gonic/gin"
)
import "douyin/views/user"

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("user/login/", user.Login)
	apiRouter.GET("/user/", user.Info)

	apiRouter.GET("/favorite/list", video.GetFavouriteVideo)
	apiRouter.GET("/publish/list", video.GetPublishedVideo)

}
