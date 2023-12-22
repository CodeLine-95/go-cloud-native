package constant

const (
	Success            = 0
	Fail               = 1
	ErrorHaveNoAccess  = 2
	ErrorUnmarshalJSON = 3
	ErrorMarshalJSON   = 4
	ErrorDB            = 5
	ErrorNotLogin      = 6
	ErrorJWT           = 7
	ErrorEtcd          = 8

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
	// ErrorDBRecordExist 记录已存在
	ErrorDBRecordExist = 200101
	// ErrorUploadImage 文件上传错误
	ErrorUploadImage = 200201

	// ErrorCreateContainer 创建容器错误
	ErrorCreateContainer = 300101
	// ErrorContainerList 获取容器列表错误
	ErrorContainerList = 300102
	// ErrorContainerLogs 获取容器日志错误
	ErrorContainerLogs = 300103
	// ErrorContainerStop 停止容器错误
	ErrorContainerStop = 300104
)

var ErrorMsg = map[int]string{
	Success:              "success",
	Fail:                 "fail",
	ErrorHaveNoAccess:    "have no access",
	ErrorParams:          "request params err",
	ErrorUnmarshalJSON:   "unmarshal json failed",
	ErrorMarshalJSON:     "marshal json failed",
	ErrorDB:              "db failed",
	ErrorNotLogin:        "need login",
	ErrorUploadImage:     "upload image failed",
	ErrorJWT:             "generate token failed",
	ErrorDBRecordExist:   "db record is exist",
	ErrorCreateContainer: "create container failed",
	ErrorContainerList:   "get container list failed",
	ErrorContainerLogs:   "get container ID %s logs failed",
	ErrorContainerStop:   "stop container ID %s failed",
	ErrorEtcd:            "etcd service failed",
}

type Error struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
