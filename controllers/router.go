package controllers

import "github.com/gin-gonic/gin"
import "douyin/views/user"

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("user/login/", user.Login)
}
