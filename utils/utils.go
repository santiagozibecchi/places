package utils

import (
	"fmt"
	"time"
)

func DeterminateValidPlaceKind(kind string) (bool, error) {
	// TODO: crear tabla
	// Esto debería estar guardado en la DB
	validKind := []string{"pub", "restaurant", "party"}
	
	if contains(validKind, kind) {
		return true, nil
	}

	return false, fmt.Errorf("No es un tipo válido de lugar")
}

func DeterminateValidCountry(kind string) (bool, error) {
	// TODO: crear tabla
	// Esto debería estar guardado en la DB
	validKind := []string{"Argentina", "Peru", "Mexico", "Brasil"}
	
	if contains(validKind, kind) {
		return true, nil
	}

	return false, fmt.Errorf("No corresponde a un País válido")
}

func DetermineValidSortOrder(sort string) (bool, error) {
	sortType := map[string]bool{
		"asc": true,
		"desc": true,
	}

	if sortType[sort] {
		return true, nil 
	}

	return false, fmt.Errorf("Tipo de orden no válido")
}

// Determinate if slice contains the element
func contains(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

var ResetTimeInMilli int64 = SetTheScheduleResetTime().UnixMilli()

// Obtengo el horario actual y le sumo x|(h,m,s)
func SetTheScheduleResetTime() *time.Time {
	currentTimeByLocation := time.Now().Local()
	refreshMinutes := 1

	h := currentTimeByLocation.Hour()
	m := currentTimeByLocation.Minute() + refreshMinutes
	s := currentTimeByLocation.Second()

	year, month, day := currentTimeByLocation.Date()

	// Current time plus refreshMinutes
	startTimeOfDay := time.Date(year, month, day, h, m, s, 0, time.Local)

	return &startTimeOfDay
}