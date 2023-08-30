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

//message douyin_favorite_list_response {
//required int32 status_code = 1; // 状态码，0-成功，其他值-失败
//optional string status_msg = 2; // 返回状态描述
//repeated Video video_list = 3; // 用户点赞视频列表
//}

//message Video {
//required int64 id = 1; // 视频唯一标识
//required User author = 2; // 视频作者信息
//required string play_url = 3; // 视频播放地址
//required string cover_url = 4; // 视频封面地址
//required int64 favorite_count = 5; // 视频的点赞总数
//required int64 comment_count = 6; // 视频的评论总数
//required bool is_favorite = 7; // true-已点赞，false-未点赞
//required string title = 8; // 视频标题
//}

type favouriteVideoListResponse struct {
	StatusCode int32           `json:"status_code"`
	VideoList  []videoResponse `json:"video_list"`
}

func GetFavouriteVideo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	queryUser := models.User{UserId: userId}

	if err := utils.DB.First(&queryUser).Error; err != nil {
		response := favouriteVideoListResponse{StatusCode: -1}
		c.JSON(http.StatusNotFound, response)
		return
	}

	var videoList []videoResponse

	for _, VideoPointer := range queryUser.FavouriteVideos {
		favouriteVideo := *VideoPointer
		video := models.Video{VideoID: favouriteVideo.VideoID}
		if err := utils.DB.First(&video).Error; err != nil {
			response := favouriteVideoListResponse{StatusCode: -1}
			c.JSON(http.StatusNotFound, response)
			return
		}

		author := models.User{UserId: video.AuthorID}
		if err := utils.DB.First(&author).Error; err != nil {
			response := favouriteVideoListResponse{StatusCode: -1}
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

	favouriteVideoListResponse := favouriteVideoListResponse{
		VideoList:  videoList,
		StatusCode: 0,
	}

	c.JSON(http.StatusOK, favouriteVideoListResponse)
	return
}
