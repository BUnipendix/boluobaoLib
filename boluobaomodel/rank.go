package boluobaomodel

type Rank struct {
	Status `json:"status"`
	Data   []BookInfoData `json:"data"`
}

var RankQuery = struct {
	SaleRank    string
	ViewRank    string
	NewBookRank string
	MarkRank    string
	bonusRank   string
}{
	SaleRank:    "sale",
	ViewRank:    "view",
	NewBookRank: "newhit",
	MarkRank:    "mark",
	bonusRank:   "bonus",
}
