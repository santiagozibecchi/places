package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/places/models"
	"github.com/places/types"
)

// Para el clima de la cuidad
func getAllCountries() ([]string, error) {
	
	sqlStatement := `SELECT country FROM place GROUP BY country;`

	rows, err := Db.Query(sqlStatement)

	defer rows.Close()

	var country string
	var countries []string

	for rows.Next() {
		err = rows.Scan(&country)

		if err != nil {
			return []string{}, fmt.Errorf("Unable to scan the row %v.\nError: %v", sqlStatement, err)
		}
		countries = append(countries, country)
	}

	return countries, nil 
}


func getMapboxDetailsByLocation(city string) ([]types.Feature, error) {
	baseURL := "https://api.mapbox.com/geocoding/v5/mapbox.places/"
	paramsMapbox := map[string]string{
		"access_token": os.Getenv("MAPBOX_KEY"),
		"limit": "5",
		"language": "es",
	}

	// Peticion HTTP a Mapbox
	instance := http.Client{}
	req, err := http.NewRequest("GET", baseURL + city + ".json", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range paramsMapbox {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := instance.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Features []types.Feature `json:"features"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Features, nil
}

func getWeatherLocation(lng float64, lat float64) (types.WeatherResponse, error) {

	baseURL := "https://api.openweathermap.org/data/2.5/weather"
	paramsOpenWeatherMap := map[string]interface{}{
		"appid": os.Getenv("OPENWEATHER_KEY"),
		"units": "metric",
		"lang":  "es",
		"lat":   lat,
		"lon":   lng,
	}

	instance := http.Client{}
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return types.WeatherResponse{}, err
	}

	q := req.URL.Query()
	for key, value := range paramsOpenWeatherMap {
		q.Add(key, fmt.Sprint(value))
	}

	req.URL.RawQuery = q.Encode()

	resp, err := instance.Do(req)
	if err != nil {
		return types.WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var response types.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return types.WeatherResponse{}, err
	}

	return response, nil
}

func getCityAndCountryByLocation(placeId string) (string, string, error)  {

	type cityAndCountry struct {
		city string
		country string
	}

	var location cityAndCountry
	sqlStatement := "SELECT l.location, l.country FROM place p INNER JOIN location l ON p.location_id = l.location_id WHERE place_id=$1;"
	err := Db.QueryRow(sqlStatement, placeId).Scan(&location.city, &location.country)

	if err == sql.ErrNoRows {
		return "", "", fmt.Errorf("No se encontró ningún lugar con el ID %s", placeId)
	} else if err != nil {
		return "", "", err
	}

	return location.city, location.country, nil
}

func getLocationId(placeId string) (string, error) {
	var locationId string
	sqlStatement := "SELECT l.location_id FROM place p INNER JOIN location l ON p.location_id = l.location_id WHERE place_id=$1;"

	err := Db.QueryRow(sqlStatement, placeId).Scan(&locationId)
	if err != nil {
		return "", err
	}

	return locationId, nil
}

func UpdateWeatherCity(placeId string) (error) {

	city, country, err := getCityAndCountryByLocation(placeId)
	if err != nil {
		return err
	}

	// ESTO DEBE SER REFACTORIZADO!
	locationId, err := getLocationId(placeId)
	if err != nil {
		return err
	}

	lng, lat, isLngAndLatCityInDB, err := getLngAndLatByPlaceId(placeId)
	if err != nil {
		return err
	}

	if !isLngAndLatCityInDB {
		fmt.Println("SE BUSCO LNG Y LAT MEDIANTE API EXTERNA")
		geocodingPlace, err := getGeocodingDetails(city, country)
		if err != nil {
			return err
		}

		weather, err := getWeatherLocation(geocodingPlace.Lng, geocodingPlace.Lat)
		if err != nil {
			return err
		}

		err = saveLngAndLatLocationToDB(geocodingPlace.Lng, geocodingPlace.Lat, city)
		if err != nil {
			return err
		}

		errAsignWeather := setWeather(weather, locationId)
		if errAsignWeather != nil {
			return err
		}
		
		return nil
	}
	
	fmt.Println("SE BUSCO LNG Y LAT MEDIANTE BASE DE DATOS!")

	weather, err := getWeatherLocation(lng, lat)
	if err != nil {
		return err
	}

	errAsignWeather := setWeather(weather, locationId)
	if errAsignWeather != nil {
		return err
	}
	
	return nil
}

func saveLngAndLatLocationToDB(lng, lat float64, city string) (error) {
	
	sqlStatement := `
	UPDATE location
	SET latitude = $1, longitude = $2
	WHERE location = $3;`

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(lat, lng, city)
	if err != nil {
		return err
	}

	return nil

}

func getLngAndLatByPlaceId(placeId string) (float64, float64, bool, error) {

	areNullFields, err := areLatLngNullByPlaceId(placeId)
	if err != nil {
		return 0, 0, false, err
	}
	if areNullFields {
		return 0, 0, false, nil
	}

	lng, lat, err := getLngAndLatFromDB(placeId)
	if err != nil {
		return 0, 0, false, err
	}

	return lng, lat, true, nil
}

func getLngAndLatFromDB(placeId string) (float64, float64, error) {

	type LngAndLat struct {
		Lng float64
		Lat float64
	}

	var geolocation LngAndLat
	
	sqlStatement := "SELECT l.longitude, l.latitude FROM place p INNER JOIN location l ON p.location_id = l.location_id WHERE place_id=$1;"

	err := Db.QueryRow(sqlStatement, placeId).Scan(&geolocation.Lng, &geolocation.Lat)

	if err == sql.ErrNoRows {
		return 0, 0, fmt.Errorf("Error al buscar datos de locacion del lugar por placeId: %s", placeId)
	} else if err != nil {
		return 0, 0, err
	}

	return geolocation.Lng, geolocation.Lat, nil
}

func areLatLngNullByPlaceId(placeId string) (bool, error) {
    var isLatLngNotNull bool

    // Verificar si latitude y longitude no son nulos en base a place_id
    query := `
        SELECT
            CASE WHEN l.latitude IS NULL OR l.longitude IS NULL THEN true ELSE false END
        FROM
            place p
        JOIN
            location l ON p.location_id = l.location_id
        WHERE
            p.place_id = $1
    `

    // Realizar la consulta y escanear el resultado
    err := Db.QueryRow(query, placeId).Scan(&isLatLngNotNull)
    if err != nil {
        return false, err
    }

    return isLatLngNotNull, nil
}


func getGeocodingDetails(city, country string) (models.GeocodingPlaceByLocation, error) {
	response, err := getMapboxDetailsByLocation(city)
	if err != nil {
		return models.GeocodingPlaceByLocation{}, err
	}

	geocodingPlace := getlatitudeAndLongitudeByLocation(response, country)
	return models.GeocodingPlaceByLocation{Lng: geocodingPlace.Lng, Lat: geocodingPlace.Lat}, nil
}

func setWeather(weather types.WeatherResponse, placeId string) error {
	type WeatherLocation struct {
		Description     string
		TemperatureMin  float64
		TemperatureMax  float64
		Temperature     float64
	}

	weatherLocation := WeatherLocation{
		Description:    weather.Weather[0].Description,
		TemperatureMin: weather.Main.TempMin,
		TemperatureMax: weather.Main.TempMax,
		Temperature:    weather.Main.Temp,
	}

	// Actualizar tabla del clima por cada petición
	sqlStatement := `
	UPDATE weather 
	SET temperature_min = $1, temperature_max = $2, temperature = $3, description = $4
	WHERE location_id = $5`

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(weatherLocation.TemperatureMin, weatherLocation.TemperatureMax, weatherLocation.Temperature, weatherLocation.Description, placeId)
	if err != nil {
		return err
	}

	return nil
}

func getlatitudeAndLongitudeByLocation(response []types.Feature, country string) types.MapboxGeocodingPlace {
	var mapboxGeocodingPlace types.MapboxGeocodingPlace

	for i := 0; i < len(response); i++ {
		context := response[i].Context
		for j := 0; j < len(context); j++ {
			if context[j].Text == country {
				mapboxGeocodingPlace = types.MapboxGeocodingPlace{
					Id: response[i].ID,
					Lng: response[i].Center[0],
					Lat: response[i].Center[1],
					Country: response[i].Context[j].Text,
				}
			}
		}
	}

	return mapboxGeocodingPlace
}