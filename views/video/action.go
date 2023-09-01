package video

import (
	"douyin/models"
	"douyin/utils"
	"douyin/utils/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type publishActionResponse struct {
	StatusCode int32 `json:"status_code"`
}

func PublishVideo(c *gin.Context) {

	token := c.PostForm("token")
	user, err := auth.GetUserFromToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	videoContent, err := c.FormFile("data")
	if err != nil {
		fmt.Println(err)
		return
	}

	videoTitle := c.PostForm("title")
	videoID := rand.Int63()
	playURL := "videos/" + strconv.FormatInt(videoID, 10)

	// Save video metadata into db
	newVideo := models.Video{
		VideoID:       videoID,
		AuthorID:      user.UserId,
		PlayURL:       playURL,
		CoverURL:      "test",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         videoTitle,
	}
	utils.DB.Create(&newVideo)

	// save video into /videos
	err = c.SaveUploadedFile(videoContent, playURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// add videoID to author:publishedVideo
	publishedVideo := models.PublishedVideo{
		UserID:  user.UserId,
		VideoID: videoID,
	}
	utils.DB.Create(&publishedVideo)

	response := publishActionResponse{
		StatusCode: 0,
	}
	c.JSON(http.StatusAccepted, response)
}
