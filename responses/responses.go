package responses

var (
	CHECK_HEADER_FAIL    = map[string]string{"err": "header invalid"}
	REGISTER_FAIL        = map[string]string{"err": "register failed. check credentials"}
	REGISTER_SUCCESS     = map[string]string{"msg": "register successfull"}
	LOGIN_FAIL           = map[string]string{"err": "login failed. check credentials"}
	MISSING_CREDENTIALS  = map[string]string{"err": "missing credentials"}
	EMAIL_INVALID        = map[string]string{"err": "e-mail invalid"}
	SERVER_ERROR         = map[string]string{"err": "server error"}
	TOKEN_ERROR          = map[string]string{"err": "token error"}
	ACCESS_TOKEN_TOO_OLD = map[string]string{"err": "access_token expired more than 30 seconds ago"}
)
