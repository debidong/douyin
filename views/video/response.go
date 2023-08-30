package video

import "douyin/views/user"

type videoResponse struct {
	VideoID        int64             `json:"id"`
	Author         user.UserResponse `json:"author"`
	PlayUrl        string            `json:"play_url"`
	CoverUrl       string            `json:"cover_url"`
	FavouriteCount int64             `json:"favourite_count"`
	CommentCount   int64             `json:"comment_count"`
	IsFavourite    bool              `json:"is_favourite"`
	Title          string            `json:"title"`
}
