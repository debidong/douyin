package obseleted

//
//import (
//	"douyin/models"
//	"douyin/utils"
//	"douyin/views/user"
//	"douyin/views/video"
//)
//
//// getVideoListFromUser Get a list of videoResponses from a specific user.
//func getVideoListFromUser(queryUser models.User, videoList *[]video.videoResponse, option string) error {
//	for _, VideoPointer := range queryUser.FavouriteVideos {
//		videoReferenced := *VideoPointer
//		video := models.Video{VideoID: videoReferenced.VideoID}
//		if err := utils.DB.First(&video).Error; err != nil {
//			return err
//		}
//
//		author := models.User{UserId: video.AuthorID}
//		if err := utils.DB.First(&author).Error; err != nil {
//			return err
//		}
//
//		authorResponse := user.UserResponse{
//			Id:       author.UserId,
//			Name:     author.Username,
//			IsFollow: false, // will implement later... too many codes here
//		}
//
//		videoResponse := video.videoResponse{
//			VideoID:        video.VideoID,
//			Author:         authorResponse,
//			PlayUrl:        video.PlayURL,
//			CoverUrl:       video.CoverURL,
//			FavouriteCount: video.FavoriteCount,
//			CommentCount:   video.CommentCount,
//			IsFavourite:    true,
//			Title:          video.Title,
//		}
//
//		videoList = append(videoList, videoResponse)
//	}
//}
