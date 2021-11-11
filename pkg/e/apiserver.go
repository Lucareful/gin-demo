package e

const (
	// SUCCESS - 200: OK.
	SUCCESS int = iota + 10001

	// ERROR - 500: Fail.
	ERROR

	// INVALID_PARAMS - 400: Request parameter error.
	INVALID_PARAMS

	// ERROR_EXIST_TAG - 401: The Tag already exists.
	ERROR_EXIST_TAG

	// ERROR_NOT_EXIST_TAG - 401: The Tag does not exist.
	ERROR_NOT_EXIST_TAG

	// ERROR_NOT_EXIST_ARTICLE - 401: The article does not exist.
	ERROR_NOT_EXIST_ARTICLE
)

const (

	// ERROR_AUTH_CHECK_TOKEN_FAIL - 403: Token鉴权失败.
	ERROR_AUTH_CHECK_TOKEN_FAIL int = iota + 20001

	// BIND_PARAMS_FAIL - 403: Bind 参数失败.
	BIND_PARAMS_FAIL

	// ERROR_AUTH_CHECK_TOKEN_TIMEOUT - 403: Token已超时.
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT

	// ERROR_AUTH_TOKEN - 403: Token生成失败.
	ERROR_AUTH_TOKEN

	// ERROR_AUTH - 403: Token错误.
	ERROR_AUTH
)

const (
	// VALIDARION_ERRORS - 403: Illegal field.
	VALIDARION_ERRORS int = iota + 30001
)
