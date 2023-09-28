package constant

const (
	EnvKey    = "app.env"
	EnvOnline = "release"
	EnvDev    = "dev"
	EnvTest   = "test"

	UserName = "user"
)

const (
	CloudNative = "cloudNative"
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
