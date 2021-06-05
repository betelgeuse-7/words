package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// return exp time in string format from the payload of a token
func GetExpTimeFromTokenPayload(payload []byte) string {
	return strings.Split(strings.Split(string(payload), ",")[1], ":")[1]
}

// convert exp time in string format to unix time
func ConvertStringToUnix(value string) (time.Time, error) {
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	unixTime := time.Unix(intValue, 0)
	return unixTime, nil
}

// get user_id from a token's payload
func GetUserIdFromTokenPayload(payload []byte) (int, error) {
	stringPayload := string(payload)
	decodedStringPayload, err := jwt.DecodeSegment(stringPayload)
	if err != nil {
		return 0, err
	}
	decodedStringPayload = []byte(strings.Replace(string(decodedStringPayload), "{", "", 1))
	decodedStringPayload = []byte(strings.Replace(string(decodedStringPayload), "}", "", 1))

	userId, err := strconv.ParseInt(strings.Split(strings.Split(string(decodedStringPayload), ",")[2], ":")[1], 0, 64)
	if err != nil {
		return 0, err
	}

	return int(userId), nil
}

// get a token's payload
func GetTokenPayload(token string) string {
	return strings.Split(token, ".")[1]
}
