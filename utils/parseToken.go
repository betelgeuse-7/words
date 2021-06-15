package utils

import (
	"log"
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

// TODO make this stable. (GetUserIdFromTokenPayload)

// get user_id from a token's payload
func GetUserIdFromTokenPayload(payload []byte) (int, error) {
	payload, err := jwt.DecodeSegment(getRidOfCurlies(string(payload)))
	if err != nil {
		return 0, err
	}
	stringPayload := string(payload)
	//  :O
	userIdStr := strings.Replace(strings.Split(stringPayload, ":")[len(strings.Split(stringPayload, ":"))-1], "}", "", 1)

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("gusid: ", err)
		return 0, err
	}

	return int(userId), nil
}

// get a token's payload
func GetTokenPayload(token string) string {
	return strings.Split(token, ".")[1]
}

func getRidOfCurlies(token string) string {
	new := strings.Replace(token, "{", "", 1)
	new = strings.Replace(new, "}", "", 1)

	return new
}
