package video

import (
	"douyin/models"
	"douyin/utils"
	"douyin/views/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type publishedVideoListResponse struct {
	StatusCode int32           `json:"status_code"`
	VideoList  []videoResponse `json:"video_list"`
}

func GetPublishedVideo(c *gin.Context) {

	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	var queryUser models.User

	if err := utils.DB.Where("user_id = ?", userId).First(&queryUser).Error; err != nil {
		response := publishedVideoListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	var videoList []videoResponse

	var publishedVideos []models.PublishedVideo
	utils.DB.Where("user_id = ?", userId).Find(&publishedVideos)
	for _, publishedVideo := range publishedVideos {
		var video models.Video
		utils.DB.Where("video_id = ?", publishedVideo.VideoID).First(&video)

		author := models.User{UserId: video.AuthorID}
		if err := utils.DB.First(&author).Error; err != nil {
			response := publishedVideoListResponse{StatusCode: -1}
			c.JSON(http.StatusNotFound, response)
			return
		}

		authorResponse := user.UserResponse{
			Id:       author.UserId,
			Name:     author.Username,
			IsFollow: false,
		}

		videoResponse := videoResponse{
			VideoID:        video.VideoID,
			Author:         authorResponse,
			PlayUrl:        video.PlayURL,
			CoverUrl:       video.CoverURL,
			FavouriteCount: video.FavoriteCount,
			CommentCount:   video.CommentCount,
			IsFavourite:    false,
			Title:          video.Title,
		}

		videoList = append(videoList, videoResponse)
	}

	publishedVideoListResponse := publishedVideoListResponse{
		VideoList:  videoList,
		StatusCode: 0,
	}

	c.JSON(http.StatusOK, publishedVideoListResponse)
	return
}
