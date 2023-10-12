package constant

const (
	EnvKey    = "app.env"
	EnvOnline = "release"
	EnvDev    = "dev"
	EnvTest   = "test"

	UserName = "user"
)

const (
	CloudNative = "cloud_native"
)

// UploadFileMaxSize 上传文件的最大限制
const (
	UploadFileMaxSize int64 = 20 << 20
)

// http协议
const (
	HttpProtocol  = "http://"
	HttpsProtocol = "https://"
)

// 客户端类型
const (
	ClientTypePc     = "pc"
	ClientTypeMobile = "mobile"
)

// 列表分页数据默认值
const (
	ListPageSize    = 10
	ListCurrentPage = 1
)

// 未知错误
const ErrorUnknow = -1

// 请求太过频繁
const ErrorRequestTooOften = -2

// 设置数据返回类型
const RT_AJAX = 1
const RT_PROTO = 2
const RT_DOWNLOAD = 3
const RT_ORIGINAL = 4

// 常见错误
const ErrorInvalidParams = 10001
const ErrorMarshal = 10002
const ErrorProtobufNodata = 10003
const ErrorNeedLogin = 10004
const ErrorBadReferrer = 10005
const ErrorDownloadNodata = 10006
const ErrorSomethingWrong = 10007
const ActivityTimeError = 10008
const ErrorRedis = 10013
