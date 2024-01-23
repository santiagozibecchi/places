package services

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"github.com/places/types"

	"github.com/places/models"
	"github.com/places/utils"
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

func updateViewsPerRequest(id string, field string) (error) {
    sqlStatement := fmt.Sprintf("UPDATE places SET %s = %s + 1 WHERE place_id = $1;", field, field)
    _, err := Db.Exec(sqlStatement, id)
    if err != nil {
		return fmt.Errorf("Unable to execute the query: %v\n Err %v", sqlStatement, err)
    }
	return nil
}

func GetPlaceById(id string) (models.Place, error) {

	// TODO! Deberia verificar si el ID existe primero!
	
	errMessage := updateViewsPerRequest(id, "total_view")
	if errMessage != nil {
		return models.Place{}, errMessage
	}

	errToUpdateView := determinateNumberOfViewsLastMinute(id)
	if errToUpdateView != nil {
		return models.Place{}, errMessage
	}

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
		&place.LatestView,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, fmt.Errorf("No se encontró ningún lugar con el ID: %s", id)
		}
		return models.Place{}, err
	}

	return place, nil
}

func determinateNumberOfViewsLastMinute(placeId string) (error){

	currentTimeByLocation := time.Now().Local()
	currentTimemilliByLocation := currentTimeByLocation.UnixMilli()

	if utils.ResetTimeInMilli > currentTimemilliByLocation {
		updateViewsPerRequest(placeId, "latest_view")
	} else {
		// medio trucheli, podria ser mejor lanzar una gorutina en main que se ejecute cada x tiempo
		// mi idea principal es tener las visitas por dia pero a fines practicos la hice cada 1 minuto.
		updateLastViewsToZero(placeId)
		utils.ResetTimeInMilli = utils.SetTheScheduleResetTime().UnixMilli()
	}

	return nil
}

func updateLastViewsToZero(placeId string) (error) {
	sqlStatement := "UPDATE places SET latest_view = 0;"

	_, err := Db.Exec(sqlStatement)
	if err != nil {
		return fmt.Errorf("Unable to execute the query: %s\n Error: %v", sqlStatement, err)
	}
	return nil
}

func GetAllPlaces(queryParams types.PlaceQueryParams) ([]models.Place, error) {
	var places []models.Place

	var sqlStatement string
	var args []interface{}

	// Sigo creyendo que se puede mejorar pero por ahora es lo que pude hacer.. 
	// Construir una query que se adapte a ciertos criterios de busquedas no esta tan sencillo
	if queryParams.Kind != "" && queryParams.Country != "" {
		sqlStatement = fmt.Sprintf("SELECT * FROM places WHERE country=$1 AND kind=$2 ORDER BY name %s", queryParams.Sort)
		args = append(args, queryParams.Country)
		args = append(args, queryParams.Kind)
	} else if queryParams.Country != "" {
		sqlStatement = fmt.Sprintf("SELECT * FROM places WHERE country=$1 ORDER BY name %s", queryParams.Sort)
		args = append(args, queryParams.Country)
	} else if queryParams.Kind != ""{
		sqlStatement = fmt.Sprintf("SELECT * FROM places WHERE kind=$1 ORDER BY name %s", queryParams.Sort)
		args = append(args, queryParams.Kind)
	} else {
		sqlStatement = fmt.Sprintf("SELECT * FROM places ORDER BY name %s", queryParams.Sort)
	}
	/*
	Db.Query es más adecuado cuando solo necesitas ejecutar una consulta
	sin necesidad de reutilizarla con diferentes parámetros.
	*/

	rows, err := Db.Query(sqlStatement, args...)
	if err != nil {
		return []models.Place{}, fmt.Errorf("Unable to execute the query: %s\n Error: %v", sqlStatement, err)
	}

	// close the statement
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
			&place.LatestView,
		)
		if err != nil {
			return []models.Place{}, fmt.Errorf("Unable to scan the row: %s\n Error: %v", sqlStatement, err)
		}

		// append the place in the places slice
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
			&updatedRow.TotalView,
			&updatedRow.LatestView)

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
		return []models.Place{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
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
			&place.LatestView,
		)
		if err != nil {
			return []models.Place{}, fmt.Errorf("Unable to scan the row %v.\nError: %v", sqlStatement, err)
		}
		places = append(places, place)
	}

	return places, nil
}




