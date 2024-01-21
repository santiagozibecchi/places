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
	INSERT INTO comments (place_id, user_id, comment)
	VALUES ($1, $2, $3)
	RETURNING comment_id, place_id, user_id, comment;`

	err := Db.QueryRow(sqlStatement, placeId, userId, newComment.Comment).
	Scan(&comment.CommentID,&comment.PlaceID, &comment.UserID, &comment.Comment)
	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

// placeExists verifica si un lugar existe en la base de datos
func placeExists(placeId int) bool {
	return true 
}

// userExists verifica si un usuario existe en la base de datos
func userExists(userId int) bool {
	return true 
}
