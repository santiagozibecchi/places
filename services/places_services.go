package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/places/types"

	"github.com/places/models"
	"github.com/places/utils"
)

func CreatePlace(placeLocation models.RequestPlaceLocation) (string, string, error) {

	if  placeLocation.Description == "" ||
		placeLocation.Address == ""     ||
		placeLocation.EndTime == ""     ||
		placeLocation.Kind == ""        ||
		placeLocation.PlaceName == ""   ||
		placeLocation.Country == ""     ||
		placeLocation.Location == ""    ||
		placeLocation.StartTime == "" {
			return "", "", errors.New("Todos los campos son obligatorios.")
	}

	// TODO: validar kind habilitados: restara

	// Inicio de la transacción
	// En este caso, primero es necesario crear una locacion para luego poder crear el lugar
	// si falla la creacion de la locacion tmb falla la creación de lugar
	tx, err := Db.Begin()
	if err != nil {
        log.Fatal(err)
    }

	// para que se ejecute al finalizar la query 
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			log.Fatal(err)
		} else {
			// Guardamos los cambios!!! Bieen! 
			tx.Commit()
		}
	}()

	var locationId int
	// CREATION OF LOCATION
    err = tx.QueryRow("INSERT INTO location (country, latitude, location, longitude) VALUES ($1, $2, $3, $4) RETURNING location_id;",
	placeLocation.Country, nil, placeLocation.Location, nil).Scan(&locationId)
    if err != nil {
		panic(err)
    }

	// CREATION OF THE PLACE
	var placeName string
	var placeId string
	err = tx.QueryRow("INSERT INTO place (location_id, description, end_time, kind, start_time, place_name, address) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING place_name, place_id;",
	locationId, placeLocation.Description, placeLocation.EndTime, placeLocation.Kind, placeLocation.StartTime, placeLocation.PlaceName, placeLocation.Address).Scan(&placeName, &placeId)
    if err != nil {
        panic(err)
    }

	// CREATION OF THE WEATHER
    _, err = tx.Exec("INSERT INTO weather (location_id, description, temperature, temperature_max, temperature_min) VALUES ($1, $2, $3, $4, $5) RETURNING weather_id;",
	locationId, nil, nil, nil, nil)
    if err != nil {
        panic(err)
    }

	return placeName, placeId, nil
}

func updateViewsPerRequest(id string, field string) (error) {
    sqlStatement := fmt.Sprintf("UPDATE place SET %s = %s + 1 WHERE place_id = $1;", field, field)
    _, err := Db.Exec(sqlStatement, id)
    if err != nil {
		return fmt.Errorf("Unable to execute the query: %v\n Err %v", sqlStatement, err)
    }
	return nil
}

func GetPlaceById(placeId string) (models.Place, error) {

	// TODO! Deberia verificar si el ID existe primero!
	
	errMessage := updateViewsPerRequest(placeId, "total_view")
	if errMessage != nil {
		return models.Place{}, errMessage
	}

	errToUpdateView := determinateNumberOfViewsLastMinute(placeId)
	if errToUpdateView != nil {
		return models.Place{}, errMessage
	}

	errUpdateWeather := UpdateWeatherCity(placeId)
	if errUpdateWeather != nil {
		return models.Place{}, errUpdateWeather
	} 

	var place models.Place

	sqlStatement := `SELECT * FROM place WHERE place_id=$1;`

	err := Db.QueryRow(sqlStatement, placeId).Scan(
		&place.PlaceId,
		&place.LocationId,
		&place.Description,
		&place.EndTime,
		&place.Kind,
		&place.LatestViews,
		&place.PlaceName,
		&place.StartTime,
		&place.TotalView,
		&place.Address,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, fmt.Errorf("No se encontró ningún lugar con el ID: %s", placeId)
		}
		return models.Place{}, err
	}

	return place, nil
}

func determinateNumberOfViewsLastMinute(placeId string) (error){

	currentTimeByLocation := time.Now().Local()
	currentTimemilliByLocation := currentTimeByLocation.UnixMilli()

	if utils.ResetTimeInMilli > currentTimemilliByLocation {
		updateViewsPerRequest(placeId, "latest_views")
	} else {
		// medio trucheli, podria ser mejor lanzar una gorutina en main que se ejecute cada x tiempo
		// mi idea principal es tener las visitas por dia pero a fines practicos la hice cada 1 minuto.
		updateLastViewsToZero(placeId)
		utils.ResetTimeInMilli = utils.SetTheScheduleResetTime().UnixMilli()
	}

	return nil
}

func updateLastViewsToZero(placeId string) (error) {
	sqlStatement := "UPDATE place SET latest_views = 0;"

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

	// INNER JOIN => unión(U) interna de las tablas

	// Sigo creyendo que se puede mejorar pero por ahora es lo que pude hacer.. 
	// Construir una query que se adapte a ciertos criterios de busquedas no esta tan sencillo
	if queryParams.Kind != "" && queryParams.Country != "" {
		sql := `SELECT p.place_id, p.location_id, p.description, p.end_time, p.kind, p.latest_views, p.place_name, p.start_time, p.total_view, address
		FROM place p INNER JOIN location l ON p.location_id = l.location_id WHERE l.country = $1 AND p.kind = $2 ORDER BY p.place_name %s`
		sqlStatement = fmt.Sprintf(sql, queryParams.Sort)
		args = append(args, queryParams.Country)
		args = append(args, queryParams.Kind)
	} else if queryParams.Country != "" {
		sql := `SELECT p.place_id, p.location_id, p.description, p.end_time, p.kind, p.latest_views, p.place_name, p.start_time, p.total_view, address 
		FROM place p INNER JOIN location l ON p.location_id = l.location_id WHERE l.country = $1 ORDER BY p.place_name %s`
		sqlStatement = fmt.Sprintf(sql, queryParams.Sort)
		args = append(args, queryParams.Country)
	} else if queryParams.Kind != ""{
		sqlStatement = fmt.Sprintf("SELECT * FROM place WHERE kind=$1 ORDER BY place_name %s", queryParams.Sort)
		args = append(args, queryParams.Kind)
	} else {
		sqlStatement = fmt.Sprintf("SELECT * FROM place ORDER BY place_name %s", queryParams.Sort)
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

	for rows.Next() {
		var place models.Place
		err = rows.Scan(
			&place.PlaceId,
			&place.LocationId,
			&place.Description,
			&place.EndTime,
			&place.Kind,
			&place.LatestViews,
			&place.PlaceName,
			&place.StartTime,
			&place.TotalView,
			&place.Address,
		)
		if err != nil {
			return []models.Place{}, fmt.Errorf("Unable to scan the row: %s\n Error: %v", sqlStatement, err)
		}

		// append the place in the places slice
		places = append(places, place)
	}

	return places, nil
}
	//* Como un lugar puede tener comentarios del usuario, una location e incluir el clima de la tabla weather
	//* por lo tanto es necesarios volar estos primeros para eliminar el lugar 
	//* Dev note: sql/03_not_implemented.sql #43

	// TODO!: LA TABLA LOCATION DEBERIA SER INDEPENDIENTE DEL LUGAR Y DEL CLIMA!!! REVEER ESTO (IMPORTANTE)
func DeleteByID(id string) (string, error) {
	var deletedPlaceName string

	// Eliminar comentarios relacionados al lugar
	_, err := Db.Exec("DELETE FROM comment WHERE place_id = $1", id)
	if err != nil {
		return "", err
	}

	// Obtener el location_id asociado al place_id
	var locationID int
	err = Db.QueryRow("SELECT location_id FROM place WHERE place_id = $1", id).Scan(&locationID)
	if err != nil {
		return "", err
	}

	// Registros relacionados en la tabla weather
	_, err = Db.Exec("DELETE FROM weather WHERE location_id = $1", locationID)
	if err != nil {
		return "", err
	}

	// Registros relacionados en la tabla place
	_, err = Db.Exec("DELETE FROM place WHERE place_id = $1", id)
	if err != nil {
		return "", err
	}

	// Registros en la tabla location
	_, err = Db.Exec("DELETE FROM location WHERE location_id = $1", locationID)
	if err != nil {
		return "", err
	}

	return deletedPlaceName, nil
}
	


// TODO: Refactorizar esta funcion en utils.go de ser posible... doesnt look easy.. 🤔
// TODO: NO me gusta, esto va a ser eliminado porque ni yo entiendo bien, ademas no hace falta, con devolver el ID ya esta.
// todo! VOLAR ESTE CODIGO QUE CONFUNDE
// Actualiza la fila con el ID proporcionado en la tabla place
func UpdatePlaceByID(id string, updatedPlace models.Place) (models.Place, error) {
    var updatedRow models.Place

    // Obtengo el tipo y el valor del struct REF: 
	// https://blog.friendsofgo.tech/posts/como-usar-reflection-en-golang/
	// https://pkg.go.dev/reflect

    t := reflect.TypeOf(updatedPlace)
    v := reflect.ValueOf(updatedPlace)

    sqlStatement := "UPDATE place SET "

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
        Scan(&updatedRow.PlaceId,
			&updatedRow.LocationId,
			&updatedRow.Description,
			&updatedRow.EndTime,
			&updatedRow.Kind,
			&updatedRow.LatestViews,
			&updatedRow.PlaceName,
			&updatedRow.StartTime,
			&updatedRow.TotalView,
			&updatedRow.Address,
		)

    if err != nil {
        return models.Place{}, err
    }

    return updatedRow, nil
}

func GetAllPlacesByName(placeName string) ([]models.Place, error) {
	var places []models.Place

	// TODO! Funciona pero me parece que funcionaria mejor una regrex
	sqlStatement := `SELECT * FROM place WHERE place_name ILIKE $1`
	searchCondition := "%" + placeName + "%"

	rows, err := Db.Query(sqlStatement, searchCondition)
	if err != nil {
		return []models.Place{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}

	defer rows.Close()

	var place models.Place
	for rows.Next() {
		err = rows.Scan(
			&place.PlaceId,
			&place.LocationId,
			&place.Description,
			&place.EndTime,
			&place.Kind,
			&place.LatestViews,
			&place.PlaceName,
			&place.StartTime,
			&place.TotalView,
			&place.Address,
		)
		if err != nil {
			return []models.Place{}, fmt.Errorf("Unable to scan the row %v.\nError: %v", sqlStatement, err)
		}
		places = append(places, place)
	}

	return places, nil
}




