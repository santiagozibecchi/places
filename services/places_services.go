package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

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
	INSERT INTO places (kind, name, country, location, address, startTime, endTime, description) 
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

	fmt.Println("Lugar eliminado:", deletedPlaceName)

	return deletedPlaceName, nil
}
