package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/places/types"
)

// Para el clima de la cuidad
func getAllCountries() ([]string, error) {
	
	sqlStatement := `SELECT country FROM places GROUP BY country;`

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

func getCityAndCountryByLocation(id string) (string, string, error)  {

	type cityAndCountry struct {
		city string
		country string
	}

	var location cityAndCountry
	sqlStatement := "SELECT location, country FROM places WHERE place_id=$1;"
	err := Db.QueryRow(sqlStatement, id).Scan(&location.city, &location.country)

	if err == sql.ErrNoRows {
		return "", "", fmt.Errorf("No se encontró ningún lugar con el ID %s", id)
	} else if err != nil {
		return "", "", err
	}

	return location.city, location.country, nil
}

func UpdateWeatherCity(placeId string) (error) {

	city, country, err := getCityAndCountryByLocation(placeId)
	if err != nil {
		return err
	}
	
	response, err  :=  getMapboxDetailsByLocation(city)
	if err != nil {
		return err
	}

	// TODO: Una vez obtengo la geolocalización puedo almacenar estas en otra
	// TODO: TABLA para no estar realizando a cada rato consultas a este endpoint
	// TODO: que me puede cobrar!
	geocodingPlace := getlatitudeAndLongitudeByLocation(response, country)

	weather, err := getWeatherLocation(geocodingPlace.Lng, geocodingPlace.Lat)
	if err != nil {
		return err
	}

	errAsignWeather := setWeatherToLocation(weather, placeId)
	if errAsignWeather != nil {
		return err
	}
	
	return nil
}

func setWeatherToLocation(weather types.WeatherResponse, placeId string) error {
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

	sqlStatement := `
	INSERT INTO weathers (place_id, temperature_min, temperature_max, temperature, description)
	VALUES ($1, $2, $3, $4, $5)`

	stmt, err := Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(placeId, weatherLocation.TemperatureMin, weatherLocation.TemperatureMax, weatherLocation.Temperature, weatherLocation.Description)
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