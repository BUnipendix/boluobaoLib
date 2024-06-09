package boluobaomodel

type TsukkomisData struct {
	TsukkomiId  int    `json:"tsukkomiId"`
	Row         int    `json:"row"`
	FavNum      int    `json:"favNum"`
	ReplyNum    int    `json:"replyNum"`
	Content     string `json:"content"`
	PostDate    string `json:"postDate"`
	IsCanDelete bool   `json:"isCanDelete"`
	PostUser    struct {
		AccountId int    `json:"accountId"`
		UserName  string `json:"userName"`
		NickName  string `json:"nickName"`
		Expand    struct {
			VipLevel int    `json:"vipLevel"`
			Avatar   string `json:"avatar"`
			Widgets  struct {
				AvatarFrame interface{}   `json:"avatarFrame"`
				Badge       []interface{} `json:"badge"`
			} `json:"widgets"`
			Growup struct {
				Lv   int    `json:"lv"`
				Exp  int    `json:"exp"`
				Name string `json:"name"`
				Icon string `json:"icon"`
			} `json:"growup"`
		} `json:"expand"`
	} `json:"postUser"`
	RoleName   string      `json:"roleName"`
	ReplyInfo  interface{} `json:"replyInfo"`
	ReplyInfos interface{} `json:"replyInfos"`
	Expand     struct {
		ChapterTitle   interface{} `json:"chapterTitle"`
		ChapterContent interface{} `json:"chapterContent"`
		Replys         interface{} `json:"replys"`
		Comic          interface{} `json:"comic"`
		Novel          interface{} `json:"novel"`
	} `json:"expand"`
	EntityID     int `json:"entityID"`
	EntityType   int `json:"entityType"`
	ChapterID    int `json:"chapterID"`
	WarningModel struct {
		WarningStatus       int  `json:"warningStatus"`
		WarningExpireSecond int  `json:"warningExpireSecond"`
		IsForbidden         bool `json:"isForbidden"`
	} `json:"warningModel"`
}
type Tsukkomis struct {
	Status `json:"status"`
	Data   []TsukkomisData `json:"data"`
}
