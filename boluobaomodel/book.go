package boluobaomodel

type SysTag struct {
	TagID   int    `json:"tagId"`
	TagName string `json:"tagName"`
}
type BookInfoData struct {
	AuthorId       int     `json:"authorId"`
	LastUpdateTime string  `json:"lastUpdateTime"`
	MarkCount      int     `json:"markCount"`
	NovelCover     string  `json:"novelCover"`
	BgBanner       string  `json:"bgBanner"`
	NovelId        int     `json:"novelId"`
	NovelName      string  `json:"novelName"`
	Point          float64 `json:"point"`
	IsFinish       bool    `json:"isFinish"`
	AuthorName     string  `json:"authorName"`
	CharCount      int     `json:"charCount"`
	ViewTimes      int     `json:"viewTimes"`
	TypeId         int     `json:"typeId"`
	AllowDown      bool    `json:"allowDown"`
	AddTime        string  `json:"addTime"`
	IsSensitive    bool    `json:"isSensitive"`
	SignStatus     string  `json:"signStatus"`
	CategoryId     int     `json:"categoryId"`
	Expand         struct {
		ChapterCount  int      `json:"chapterCount"`
		BigBgBanner   string   `json:"bigBgBanner"`
		BigNovelCover string   `json:"bigNovelCover"`
		TypeName      string   `json:"typeName"`
		Intro         string   `json:"intro"`
		Fav           int      `json:"fav"`
		Ticket        int      `json:"ticket"`
		PointCount    int      `json:"pointCount"`
		Tags          []string `json:"tags"`
		SysTags       []struct {
			SysTagId int    `json:"sysTagId"`
			TagName  string `json:"tagName"`
		} `json:"sysTags"`
		SignLevel          string  `json:"signLevel"`
		Discount           float64 `json:"discount"`
		DiscountExpireDate string  `json:"discountExpireDate"`
		TotalNeedFireMoney int     `json:"totalNeedFireMoney"`
		Rankinglist        struct {
			RankinglistStr string `json:"rankinglistStr,omitempty"`
			Rank           int    `json:"rank,omitempty"`
			Desc           string `json:"desc"`
			Type           int    `json:"type"`
			DateRange      int    `json:"dateRange"`
		} `json:"rankinglist"`
		OriginTotalNeedFireMoney int `json:"originTotalNeedFireMoney"`
		FirstChapter             struct {
			Title   string `json:"title"`
			ChapId  int    `json:"chapId"`
			AddTime string `json:"addTime"`
		} `json:"firstChapter"`
		LatestChapter struct {
			Title   string `json:"title"`
			ChapId  int    `json:"chapId"`
			AddTime string `json:"addTime"`
		} `json:"latestChapter"`
		LatestCommentDate string `json:"latestCommentDate"`
		EssayTag          *struct {
			EssayTagId int    `json:"essayTagId"`
			TagName    string `json:"tagName"`
		} `json:"essayTag"`
		AuditCover         string        `json:"auditCover"`
		PreOrderInfo       interface{}   `json:"preOrderInfo"`
		CustomTag          []interface{} `json:"customTag"`
		Topic              interface{}   `json:"topic"`
		UnauditedCustomtag []interface{} `json:"unauditedCustomtag"`
		HomeFlag           []struct {
			Desc      string `json:"desc"`
			Type      int    `json:"type"`
			DateRange int    `json:"dateRange"`
			Num       int    `json:"num"`
		} `json:"homeFlag"`
		IsBranch    bool `json:"isBranch"`
		EssayAwards []struct {
			SysTagId int         `json:"sysTagId"`
			TagName  string      `json:"tagName"`
			Link     string      `json:"link"`
			IntroUrl interface{} `json:"introUrl"`
		} `json:"essayAwards"`
	} `json:"expand"`
}

type BookInfo struct {
	Status `json:"status"`
	Data   BookInfoData `json:"data"`
}
