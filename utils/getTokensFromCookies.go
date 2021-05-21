package utils

import (
	"errors"
	"net/http"

	"github.com/betelgeuse-7/words/constants"
)

// get access and refresh token from cookies.
// return them respectively.
// return non-nil error if one of the cookies are missing.
func GetTokensFromCookies(r *http.Request) ([]string, error) {
	cookies := r.Cookies()

	var tokens []string
	var accessToken, refreshToken string

	for _, v := range cookies {
		if v.Name == constants.ACCESS_TOKEN_COOKIE_NAME {
			accessToken = v.Value
			tokens = append(tokens, accessToken)
		}
		if v.Name == constants.REFRESH_TOKEN_COOKIE_NAME {
			refreshToken = v.Value
			tokens = append(tokens, refreshToken)
		}
	}

	if err := LenGreaterThanZero(accessToken, refreshToken); err != nil {
		return []string{}, errors.New("cookie missing")
	}

	return tokens, nil
}
