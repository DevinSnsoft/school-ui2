package e



var MsgFlags = map[int]string{
	SUCCESS:                         "OK",
	ERROR:                           "Fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token失效或不存在",
}
