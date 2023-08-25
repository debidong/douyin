package main

import (
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Service starting.")
	utils.InitDB()
	r := gin.Default()
	initRouter(r)
	err := r.Run()
	if err != nil {
		return
	}
}
