package errors

var (
	ErrInvalidParam = newResponse(400, "参数无效，请检查入参")
)
