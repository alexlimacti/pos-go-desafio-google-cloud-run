package utils

import "regexp"

func IsValidCEP(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	// Check if all characters are digits
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}
