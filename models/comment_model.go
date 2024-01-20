package models

type Comment struct {
	CommentID int `json:"comment_id"`
	PlaceID   int `json:"place_id"`
	UserID    int `json:"user_id"`
}
