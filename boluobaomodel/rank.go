package boluobaomodel

type Rank struct {
	Status `json:"status"`
	Data   []BookInfoData `json:"data"`
}
