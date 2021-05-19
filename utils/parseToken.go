package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
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

func GetUserIdFromTokenPayload(payload []byte) string {
	return strings.Split(strings.Split(string(payload), ",")[2], ":")[2]
}
