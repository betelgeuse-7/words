package utils

import "net/http"

// range over a map of wanted header parameters and compare it with the actual header
// return true if desires are met. otherwise false
// also returns a message string
func CheckRequestHeader(desiredHeaderParams map[string]string, header http.Header) (bool, string) {
	for k, v := range desiredHeaderParams {
		if header.Get(k) == "" {
			return false, "header-key-not-present"
		}
		if header.Get(k) != v {
			return false, "unwanted-header-value"
		}
	}
	return true, "ok"
}
