package boluobaomodel

import "net/http"

type Status struct {
	HTTPCode  int    `json:"httpCode"`
	ErrorCode int    `json:"errorCode"`
	MsgType   int    `json:"msgType"`
	Msg       string `json:"msg"`
}
type LoginStatus struct {
	Status
	Cookie string `json:"cookie"`
}

func (status *Status) GetCode() int {
	return status.HTTPCode
}

func (status *Status) GetTip() string {
	return status.Msg
}

func (status *Status) IsSuccess() bool {
	return status.HTTPCode == http.StatusOK
}
