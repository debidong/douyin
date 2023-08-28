package main

import (
	"douyin/controllers"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Service starting.")

	utils.InitDB()
	utils.InitRedis()

	r := gin.Default()
	controllers.InitRouter(r)
	err := r.Run()
	if err != nil {
		return
	}
}
