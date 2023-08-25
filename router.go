package main

import "github.com/gin-gonic/gin"
import "douyin/api/user"

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("user/login/", user.Login)
}
