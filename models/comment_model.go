package models

type Comment struct {
	CommentID int `json:"comment_id"`
	PlaceID   int `json:"place_id"`
	UserID    int `json:"user_id"`
	Comment    string `json:"comment"`
}

type ExpandComment struct {
    CommentID    int    `json:"comment_id"`
    PlaceID      int    `json:"place_id"`
    UserID       int    `json:"user_id"`
    Comment      string `json:"comment"`
    UserName     string `json:"name"`
    UserLastName string `json:"lastname"`
    Username     string `json:"username"`
	Name        string `json:"place_name"`
	Location    string `json:"location"`
}

