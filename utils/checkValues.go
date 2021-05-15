package utils

import (
	"errors"
	"strings"
)

func LenGreaterThanZero(params ...string) error {
	for _, v := range params {
		if len(v) < 1 {
			return errors.New("len smaller than 1")
		}
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}
	if len(email) < 1 {
		return errors.New("email non-existent (length smaller than 1)")
	}
	return nil
}
