package constant

type Response struct {
	// 错误码
	ErrNo int `json:"errno" example:"0"`
	// 错误信息
	Msg string `json:"msg" example:"json解析失败"`
	// 错误提示
	Info string `json:"info" example:"参数错误"`
	// log id
	TraceID string `json:"trace_id" example:"0"`
	// 返回信息
	Data any `json:"data"`
}

type RespPage struct {
	PageSize  int `json:"page_size"`
	PageNum   int `json:"page_num"`
	PageTotal int `json:"page_total"`
	TotalSize int `json:"total_size"`
}

func GetRespPage(count, size, current int) RespPage {
	return RespPage{
		PageSize:  size,
		PageNum:   current,
		PageTotal: GetTotalSize(count, size),
		TotalSize: count,
	}
}

func GetTotalSize(count, size int) int {
	return (count + size - 1) / size
}
