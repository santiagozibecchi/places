package utils

import "fmt"

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
