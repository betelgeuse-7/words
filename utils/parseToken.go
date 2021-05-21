package utils

import (
	"fmt"
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
		log.Println("convertToUnix: ", err)
		return time.Time{}, err
	}
	unixTime := time.Unix(intValue, 0)
	return unixTime, nil
}

func GetUserIdFromTokenPayload(payload []byte) (int, error) {
	stringPayload := string(payload)
	decodedStringPayload, err := jwt.DecodeSegment(stringPayload)
	if err != nil {
		fmt.Println("GetUserIdFromTokenPayload ERR: ", err)
		return 0, err
	}
	decodedStringPayload = []byte(strings.Replace(string(decodedStringPayload), "{", "", 1))
	decodedStringPayload = []byte(strings.Replace(string(decodedStringPayload), "}", "", 1))

	fmt.Println(strings.Split(strings.Split(string(decodedStringPayload), ",")[2], ":"))

	userId, err := strconv.ParseInt(strings.Split(strings.Split(string(decodedStringPayload), ",")[2], ":")[1], 0, 64)
	if err != nil {
		fmt.Println("parseToken.go#GetUserIdFromTokenPayload: ", err)
		return 0, err
	}

	fmt.Println("GetUserIdFromTokenPayload <msg> Got user_id: ", userId)

	return int(userId), nil
}

func GetTokenPayload(token string) string {
	return strings.Split(token, ".")[1]
}
