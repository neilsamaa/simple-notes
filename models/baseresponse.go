package models

type BaseResponse struct {
	Message string      `json:"massage"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
