package utils

import (
	"time"
)

// CheckAge check age for ageLimit
func CheckAge(date string, ageLimit int) bool {

	dateOfBirth, _ := time.Parse(time.RFC3339, date)

	years := time.Now().Year() - dateOfBirth.Year()

	if years < 18 {
		return false
	}

	return true
}
