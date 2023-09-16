package httpCode

type HttpCode struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	TracID  string `json:"trac_id"`
}
