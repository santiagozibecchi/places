package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/places/models"
)

// Devuelve el nombre de la pelicula
func CreatePlace(place models.Place) (string, error) {

	if  place.Kind == ""      ||
		place.Name == ""      ||
		place.Country == ""   ||
		place.Location == ""  ||
		place.Address == ""   ||
		place.StartTime == "" ||
		place.EndTime == ""   ||
		place.Description == "" {
			return "", errors.New("Todos los campos son obligatorios.")
	}

	/*
	El uso de Db.Prepare es beneficioso cuando se planea ejecutar la misma consulta
	varias veces con diferentes valores de parámetros, ya que la consulta se compila
	una vez y luego se puede ejecutar eficientemente con diferentes valores
	*/ 
	stmt, err := Db.Prepare(`
	INSERT INTO places (kind, name, country, location, address, start_time, end_time, description) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING name`)

	if err != nil {
		return "", err
	}

	var newPlace string
	err = stmt.QueryRow(place.Kind, place.Name, place.Country, place.Location ,place.Address, place.StartTime, place.EndTime, place.Description).Scan(&newPlace)
	if err != nil {
		return "", err
	}

	return newPlace, nil 
}

func updateViewsPerRequest(id string) {
    sqlStatement := "UPDATE places SET total_view = total_view + 1 WHERE place_id = $1;"
    _, err := Db.Exec(sqlStatement, id)
    if err != nil {
        log.Fatalf("Unable to execute the query. %v\n Err %v", sqlStatement, err)
    }
}

func GetPlaceById(id string) (models.Place, error) {
	
	updateViewsPerRequest(id)

	var place models.Place

	sqlStatement := `SELECT * FROM places WHERE place_id=$1;`

	err := Db.QueryRow(sqlStatement, id).Scan(
		&place.PlaceID,
		&place.Name,
		&place.Kind,
		&place.Country,
		&place.Location,
		&place.Address,
		&place.StartTime,
		&place.EndTime,
		&place.Description,
		&place.TotalView,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, fmt.Errorf("No se encontró ningún lugar con el ID: %s", id)
		}
		return models.Place{}, err
	}

	return place, nil
}

func GetAllPlaces() ([]models.Place, error) {
	var places []models.Place

	sqlStatement := `SELECT * FROM places`
	/*
	Db.Query es más adecuado cuando solo necesitas ejecutar una consulta
	sin necesidad de reutilizarla con diferentes parámetros.
	*/
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	var place models.Place
	for rows.Next() {

		// unmarshal the row object to movie
		err = rows.Scan(
			&place.PlaceID,
			&place.Name,
			&place.Kind,
			&place.Country,
			&place.Location,
			&place.Address,
			&place.StartTime,
			&place.EndTime,
			&place.Description,
			&place.TotalView,
		)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the movie in the movies slice
		places = append(places, place)

	}

	return places, nil
}

func DeleteByID(id string) (string, error) {
	var deletedPlaceName string
	sqlStatement := `DELETE FROM places WHERE place_id=$1 RETURNING name;`

	err := Db.QueryRow(sqlStatement, id).Scan(&deletedPlaceName)

	if err == sql.ErrNoRows {
		// Cuando no se encuentra ninguna fila con el ID proporcionado
		return "", fmt.Errorf("No se encontró ningún lugar con el ID %s", id)
	} else if err != nil {
		return "", err
	}

	return deletedPlaceName, nil
}

// TODO: Refactorizar esta funcion en utils.go de ser posible... doesnt look easy.. 🤔
// Actualiza la fila con el ID proporcionado en la tabla places
func UpdatePlaceByID(id string, updatedPlace models.Place) (models.Place, error) {
    var updatedRow models.Place

    // Obtengo el tipo y el valor del struct REF: 
	// https://blog.friendsofgo.tech/posts/como-usar-reflection-en-golang/
	// https://pkg.go.dev/reflect

    t := reflect.TypeOf(updatedPlace)
    v := reflect.ValueOf(updatedPlace)

    sqlStatement := "UPDATE places SET "

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
    sqlStatement += " WHERE place_id=$" + strconv.Itoa(len(sqlValues)+1) + " RETURNING *;"

	fmt.Println(sqlStatement)

    // Agregar el ID al final de consulta SQL
    sqlValues = append(sqlValues, id)

    err := Db.QueryRow(sqlStatement, sqlValues...).
        Scan(&updatedRow.PlaceID,
			&updatedRow.Name,
			&updatedRow.Kind,
			&updatedRow.Country,
			&updatedRow.Location,
			&updatedRow.Address,
			&updatedRow.StartTime,
			&updatedRow.EndTime,
			&updatedRow.Description,
			&updatedRow.TotalView)

    if err != nil {
        return models.Place{}, err
    }

    return updatedRow, nil
}

func GetAllPlacesByName(placeName string) ([]models.Place, error) {
	var places []models.Place

	// TODO! Funciona pero me parece que funcionaria mejor una regrex
	sqlStatement := `SELECT * FROM places WHERE name ILIKE $1`
	searchCondition := "%" + placeName + "%"

	rows, err := Db.Query(sqlStatement, searchCondition)
	if err != nil {
		log.Fatalf("Unable to execute the query => %v. %v", sqlStatement, err)
		return []models.Place{}, err
	}

	defer rows.Close()

	var place models.Place
	for rows.Next() {
		err = rows.Scan(
			&place.PlaceID,
			&place.Name,
			&place.Kind,
			&place.Country,
			&place.Location,
			&place.Address,
			&place.StartTime,
			&place.EndTime,
			&place.Description,
			&place.TotalView,
		)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		places = append(places, place)
	}

	return places, nil
}




