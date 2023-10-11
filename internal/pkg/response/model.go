package response

type Response struct {
	RequestId string `json:"requestId,omitempty"`
	Code      int32  `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Status    string `json:"status,omitempty"`
}

type response struct {
	Response
	Data any `json:"data"`
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int32) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	}
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type page struct {
	Page
	List any `json:"list"`
}

func (e *response) SetData(data any) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}
