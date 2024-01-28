package services

import (
	"fmt"

	"github.com/places/models"
)

func CreateComment(placeId int, userId int, newComment models.Comment) (models.Comment, error) {
	// TODO: Verificar que el lugar
	if !placeExists(placeId) {
		return models.Comment{}, fmt.Errorf("El lugar con placeId: %d no existe", placeId)
	} 
	// TODO: Verificar que el usuario
	if !userExists(userId) {
		return models.Comment{}, fmt.Errorf("El usuario con userId: %d no existe", userId)
	} 

	var comment models.Comment

	sqlStatement := `
	INSERT INTO comment (place_id, user_id, comment)
	VALUES ($1, $2, $3)
	RETURNING comment_id, place_id, user_id, comment;`

	err := Db.QueryRow(sqlStatement, placeId, userId, newComment.Comment).
	Scan(&comment.CommentID,&comment.PlaceID, &comment.UserID, &comment.Comment)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func GetCommentsByUserId(userId int) ([]models.Comment, error) {
	
	sqlStatement := "SELECT * FROM comment WHERE user_id=$1;"
	
	rows, err := Db.Query(sqlStatement, userId)
	if err != nil {
		return []models.Comment{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}
	
	defer rows.Close()
	
	var commets []models.Comment
	var comment models.Comment
	
	for rows.Next() {
		err = rows.Scan(&comment.CommentID, &comment.PlaceID, &comment.UserID, &comment.Comment)

		if err != nil {
			return []models.Comment{}, fmt.Errorf("Unable to scan the row => %v.\nError: %v", sqlStatement, err)
		}

		commets = append(commets, comment)
	}

	return commets, nil 
}

func GetExpandCommentsByUserIdAndPlaceId(placeId, userId int) ([]models.ExpandComment, error) {
	
	sqlStatement := `
	SELECT 
		c.comment_id, c.place_id, c.user_id, c.comment,
		s.user_name, s.user_lastname, s.username,
		p.place_name
	FROM comment c
	JOIN user_account s ON c.user_id = s.user_id
	JOIN place p ON c.place_id = p.place_id
	WHERE c.user_id=$1 AND c.place_id=$2;`
	
	rows, err := Db.Query(sqlStatement, userId, placeId)
	if err != nil {
		return []models.ExpandComment{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}
	
	defer rows.Close()
	
	var commets []models.ExpandComment
	var comment models.ExpandComment
	
	for rows.Next() {
		err = rows.Scan(
			&comment.CommentID, &comment.PlaceID, &comment.UserID, &comment.Comment,
			&comment.UserName, &comment.UserLastName, &comment.Username,
			&comment.Name,
		)

		if err != nil {
			return []models.ExpandComment{}, fmt.Errorf("Unable to scan the row => %v.\nError: %v", sqlStatement, err)
		}

		commets = append(commets, comment)
	}

	return commets, nil 
}

// placeExists verifica si un lugar existe en la base de datos
func placeExists(placeId int) bool {
	return true 
}

// userExists verifica si un usuario existe en la base de datos
func userExists(userId int) bool {
	return true 
}
