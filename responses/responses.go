package responses

var CHECK_HEADER_FAIL = map[string]string{"err": "header invalid"}
var REGISTER_FAIl = map[string]string{"err": "register failed. check credentials"}
var REGISTER_SUCCESS = map[string]string{"err": "register successfull"}
var MISSING_CREDENTIALS = map[string]string{"err": "missing credentials"}
var EMAIL_INVALID = map[string]string{"err": "e-mail invalid"}
var SERVER_ERROR = map[string]string{"err": "server error"}
