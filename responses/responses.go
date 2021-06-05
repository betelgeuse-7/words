package responses

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Title, Message string
}

// send a response back, in json format.
func (r *response) Send(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(r)
}

var (
	CHECK_HEADER_FAIL    response = response{"err", "header invalid"}
	REGISTER_FAIL        response = response{"err", "register failed"}
	REGISTER_SUCCESS     response = response{"msg", "register successful"}
	LOGIN_FAIL           response = response{"err", "login failed. check credentials"}
	MISSING_CREDENTIALS  response = response{"err", "missing credentials"}
	EMAIL_INVALID        response = response{"err", "email invalid"}
	SERVER_ERROR         response = response{"err", "server error"}
	TOKEN_ERROR          response = response{"err", "token error"}
	ACCESS_TOKEN_TOO_OLD response = response{"err", "access_token expired more than 30 seconds ago"}
)
