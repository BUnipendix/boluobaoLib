package boluobaomodel

type Search struct {
	Status struct {
		HttpCode  int         `json:"httpCode"`
		ErrorCode int         `json:"errorCode"`
		MsgType   int         `json:"msgType"`
		Msg       interface{} `json:"msg"`
	} `json:"status"`
	Data struct {
		Novels []BookInfoData `json:"novels"`
		Comics []interface{}  `json:"comics"`
		Albums []interface{}  `json:"albums"`
	} `json:"data"`
}
