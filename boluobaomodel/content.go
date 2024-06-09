package boluobaomodel

type ContentData struct {
	ChapId     int    `json:"chapId"`
	NovelId    int    `json:"novelId"`
	VolumeId   int    `json:"volumeId"`
	CharCount  int    `json:"charCount"`
	RowNum     int    `json:"rowNum"`
	ChapOrder  int    `json:"chapOrder"`
	Title      string `json:"title"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
	Sno        int    `json:"sno"`
	IsVip      bool   `json:"isVip"`
	Expand     struct {
		NeedFireMoney       int    `json:"needFireMoney"`
		OriginNeedFireMoney int    `json:"originNeedFireMoney"`
		Content             string `json:"content"`
		Tsukkomi            []struct {
			Row   int `json:"row"`
			Count int `json:"count"`
		} `json:"tsukkomi"`
		ChatLines          []interface{} `json:"chatLines"`
		Volume             interface{}   `json:"volume"`
		AuthorTalk         string        `json:"authorTalk"`
		IsContentEncrypted bool          `json:"isContentEncrypted"`
		IsBranch           bool          `json:"isBranch"`
	} `json:"expand"`
	Ntitle      string `json:"ntitle"`
	IsRubbish   bool   `json:"isRubbish"`
	AuditStatus int    `json:"auditStatus"`
}
type Content struct {
	Status `json:"status"`
	Data   ContentData `json:"data"`
}
