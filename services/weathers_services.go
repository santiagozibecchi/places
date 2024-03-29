package services

import (
	"database/sql"
	"fmt"

	"github.com/places/models"
)



func GetWeatherByPlaceName(placeName string) (models.Weather, error) {

	sqlStatement := `
	SELECT 
    	w.weather_id,
    	w.location_id,
    	w.description,
    	w.temperature,
    	w.temperature_max,
    	w.temperature_min 
	FROM place p 
	JOIN weather w ON p.location_id = w.location_id 
	WHERE p.place_name = $1
	LIMIT 1;`
	
	var weather models.Weather
	err := Db.QueryRow(sqlStatement, placeName).Scan(
		&weather.WeatherId,
		&weather.LocationId,
		&weather.Description,
		&weather.Temperature,
		&weather.TemperatureMax,
		&weather.TemperatureMin,
	)

	if err == sql.ErrNoRows {
		return models.Weather{}, fmt.Errorf("No se encontró ningún clima con el nombre del lugar: %s", placeName)
	} else if err != nil {
		return models.Weather{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}

	return weather, nil
}



func GetWeatherByLocationId(locationId int) (models.WeatherWithLocation, error) {
	sqlStatement := fmt.Sprintf(`
		SELECT 
			w.weather_id,
			w.location_id,
			w.description,
			w.temperature,
			w.temperature_max,
			w.temperature_min,
			l.location
		FROM 
			location l 
		JOIN 
			weather w ON l.location_id = w.location_id
		WHERE 
			l.location_id = %v 
		LIMIT 1;
	`, locationId)

	var weather models.WeatherWithLocation
	err := Db.QueryRow(sqlStatement).Scan(
		&weather.WeatherId,
		&weather.LocationId,
		&weather.Description,
		&weather.Temperature,
		&weather.TemperatureMax,
		&weather.TemperatureMin,
		&weather.Location,
	)

	if err == sql.ErrNoRows {
		return models.WeatherWithLocation{}, fmt.Errorf("No se encontró ningún clima con el location_id: %v", locationId)
	} else if err != nil {
		return models.WeatherWithLocation{}, fmt.Errorf("Unable to execute the query: %v.\nError: %v", sqlStatement, err)
	}

	return weather, nil
}