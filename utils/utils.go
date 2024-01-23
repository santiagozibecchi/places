package utils

import (
	"fmt"
	"time"
)

func DeterminateValidPlaceKind(kind string) (bool, error) {
	// Esto debería estar guardado en la DB
	validKind := []string{"pubs", "restaurants", "parties"}
	
	if contains(validKind, kind) {
		return true, nil
	}

	return false, fmt.Errorf("No es un tipo válido de lugar")
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