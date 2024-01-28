package services

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/places/models"
)

// Devuelve el nombre de la pelicula
func CreateUser(place models.User) (string, string, error) {

	if  place.Name == ""      ||
		place.LastName == ""      ||
		place.Email == ""   ||
		place.Username == ""  ||
		place.Gender == "" {
			return "", "", errors.New("Todos los campos son obligatorios.")
	}

	stmt, err := Db.Prepare(`
	INSERT INTO user_account (user_name, user_lastname, email, username, gender) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING user_name, user_id`)

	if err != nil {
		return "", "", err
	}

	var newUserName string
	var userId string
	err = stmt.QueryRow(place.Name, place.LastName, place.Email, place.Username ,place.Gender).Scan(&newUserName, &userId)
	if err != nil {
		return "", "", err
	}

	return newUserName, userId, nil 
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	sqlStatement := `SELECT * FROM user_account`

	rows, err := Db.Query(sqlStatement)
	if err != nil {
		return []models.User{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}

	// close the statement
	defer rows.Close()

	var user models.User
	for rows.Next() {

		// unmarshal the row object to movie
		err = rows.Scan(
			&user.UserID,
			&user.Name,
			&user.LastName,
			&user.Email,
			&user.Username,
			&user.Gender,
		)

		if err != nil {
			return []models.User{}, fmt.Errorf("Unable to scan the row => %v.\nError: %v", sqlStatement, err)
		}

		// append the movie in the movies slice
		users = append(users, user)

	}

	return users, nil
}

func DeleteUserByID(id string) (string, error) {
	var deletedUserName string
	sqlStatement := `DELETE FROM user_account WHERE user_id=$1 RETURNING user_name;`

	err := Db.QueryRow(sqlStatement, id).Scan(&deletedUserName)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("No se encontró ningún usuario con el ID %s", id)
	} else if err != nil {
		return "", err
	}

	return deletedUserName, nil
}

// TODO: Refactorizar esta funcion en utils.go de ser posible... doesnt look easy.. 🤔
/*
	* Devnotes:
	* Debe aceptar n modelos => notar que el modelo que recibe por parametro es el mismo que retorna
	* Debe recibir por parámetro la tabla en cuestion
	* Debe recibir nombre del campo que corresponde al id: place_id, user_id (FK)
*/
func UpdateUserByID(id string, updatedPlace models.User) (models.User, error) {
    var updatedRow models.User

    // Obtengo el tipo y el valor del struct REF: 
	// https://blog.friendsofgo.tech/posts/como-usar-reflection-en-golang/
	// https://pkg.go.dev/reflect

    t := reflect.TypeOf(updatedPlace)
    v := reflect.ValueOf(updatedPlace)

    sqlStatement := "UPDATE user_account SET "

    // Almacenar los valores de los campos que se actualizarán
    var sqlValues []interface{}

    // Campos del struct 
    for i := 0; i < t.NumField(); i++ {
        fieldName := t.Field(i).Name
        fieldValue := v.Field(i).Interface()

        // Agregar el campo a la consulta solo si el valor no es cero
        if fieldValue != reflect.Zero(v.Field(i).Type()).Interface() {
            sqlStatement += fieldName + "=$" + strconv.Itoa(len(sqlValues)+1) + ", "
            sqlValues = append(sqlValues, fieldValue)
        }
    }

    // Eliminar la coma adicional al final de la declaración SQL
    sqlStatement = strings.TrimSuffix(sqlStatement, ", ")
    sqlStatement += " WHERE user_id=$" + strconv.Itoa(len(sqlValues)+1) + " RETURNING *;"

	fmt.Println(sqlStatement)

    // Agregar el ID al final de consulta SQL
    sqlValues = append(sqlValues, id)

    err := Db.QueryRow(sqlStatement, sqlValues...).
	// ! Y esto hay que ver si se puede refactorizar para que sea generico... 🤔🤔🤔
        Scan(&updatedRow.UserID,
			&updatedRow.Name,
			&updatedRow.LastName,
			&updatedRow.Email,
			&updatedRow.Username,
			&updatedRow.Gender)

    if err != nil {
        return models.User{}, err
    }

    return updatedRow, nil
}




