package models

type User struct {
	UserID   	int    `json:"user_id"`
	Name     	string `json:"user_name"`
	LastName 	string `json:"user_lastname"`
	Email    	string `json:"email"`
	Username 	string `json:"username"`
	Gender   	string `json:"gender"`
}
