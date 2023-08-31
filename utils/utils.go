package utils

import (
	"douyin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
)

var DB *gorm.DB
var Client *redis.Client

func InitDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect to DB.")
	}

	DB = database
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func ParseParamToInt(c *gin.Context, paramName string) (int, error) {
	paramStr := c.Query(paramName)
	return ParseStringToInt(paramStr)
}

func ParseStringToInt(str string) (int, error) {
	paramInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("无效的 %s 参数", paramInt)
	}
	return paramInt, nil
}
