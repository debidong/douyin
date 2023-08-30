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

	queryUser := models.User{UserId: userId}

	if err := utils.DB.First(&queryUser).Error; err != nil {
		response := publishedVideoListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	var videoList []videoResponse

	for _, VideoPointer := range queryUser.PublishedVideos {
		publishedVideo := *VideoPointer
		video := models.Video{VideoID: publishedVideo.VideoID}
		if err := utils.DB.First(&video).Error; err != nil {
			response := publishedVideoListResponse{StatusCode: -1}
			c.JSON(http.StatusNotFound, response)
			return
		}

		author := models.User{UserId: video.AuthorID}
		if err := utils.DB.First(&author).Error; err != nil {
			response := publishedVideoListResponse{StatusCode: -1}
			c.JSON(http.StatusNotFound, response)
			return
		}

		authorResponse := user.UserResponse{
			Id:       author.UserId,
			Name:     author.Username,
			IsFollow: false, // will implement later... too many codes here
		}

		videoResponse := videoResponse{
			VideoID:        video.VideoID,
			Author:         authorResponse,
			PlayUrl:        video.PlayURL,
			CoverUrl:       video.CoverURL,
			FavouriteCount: video.FavoriteCount,
			CommentCount:   video.CommentCount,
			IsFavourite:    true,
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
