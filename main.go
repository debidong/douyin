package main

import (
	"douyin/utils"
	"fmt"
)

func main() {
	fmt.Println("Service starting.")
	utils.InitDB()
}
