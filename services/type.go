package services

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type ContainerStop struct {
	Ids string `json:"ids"`
}

type ImagesPull struct {
	Refstr string `json:"refstr"`
}
