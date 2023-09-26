package constant

const (
	Success            = 0
	Fail               = 1
	ErrorHaveNoAccess  = 2
	ErrorUnmarshalJSON = 3
	ErrorMarshalJSON   = 4
	ErrorDB            = 5
	ErrorNotLogin      = 6
	ErrorNetRequest    = 7
	ErrorIO            = 8

	// ErrorParams 参数错误
	ErrorParams = 400
	// ErrorAuthFail 相关业务逻辑校验失败
	ErrorAuthFail = 401
	// ErrorForbidden 权限不足
	ErrorForbidden = 403
	// ErrorNotFound 资源不存在
	ErrorNotFound = 404
	// ErrorInternetServer 代码错误
	ErrorInternetServer = 500
	// ErrorBadGateway 上游服务返回的内容无法识别
	ErrorBadGateway = 502
	// ErrorServiceUnavailable 上游服务暂时不可用
	ErrorServiceUnavailable = 503
	// ErrorGatewayTimeout 上游服务超时
	ErrorGatewayTimeout = 504

	ErrorUploadImage = 200201

	// ErrorDockerDataList docker
	ErrorDockerDataList = 200301
)

var ErrorMsg = map[int]string{
	Success:            "success",
	Fail:               "fail",
	ErrorHaveNoAccess:  "have no access",
	ErrorParams:        "request params err",
	ErrorUnmarshalJSON: "unmarshal json failed",
	ErrorMarshalJSON:   "marshal json failed",
	ErrorDB:            "db failed",
	ErrorNotLogin:      "need login",
	ErrorUploadImage:   "upload image failed",
}

type Error struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

var (
	// ErrorSuccess 接口正确返回，即可使用该错误码
	ErrorSuccess = &Error{ErrCode: Success, ErrMsg: ErrorMsg[Success]}
	ErrorParam   = &Error{ErrCode: ErrorParams, ErrMsg: ErrorMsg[ErrorParams]}

	ErrorUnmarshalJson = &Error{ErrCode: ErrorUnmarshalJSON, ErrMsg: ErrorMsg[ErrorUnmarshalJSON]}
	ErrorMarshalJson   = &Error{ErrCode: ErrorMarshalJSON, ErrMsg: ErrorMsg[ErrorMarshalJSON]}
	ErrorUpload        = &Error{ErrCode: ErrorUploadImage, ErrMsg: ErrorMsg[ErrorUploadImage]}
)
