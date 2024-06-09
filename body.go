package boluobaoLib

type BuyVipContentBody struct {
	OrderType string `json:"orderType"`
	OrderAll  bool   `json:"orderAll"`
	AutoOrder bool   `json:"autoOrder"`
	ChapIds   []int  `json:"chapIds"`
}
type LoginBody struct {
	Username interface{} `json:"username"`
	Password interface{} `json:"password"`
}
