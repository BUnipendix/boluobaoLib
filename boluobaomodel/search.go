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

func (search *Search) EachBook(f func(int, BookInfoData)) {
	if search.Data.Novels != nil {
		for i, book := range search.Data.Novels {
			f(i, book)
		}
	}

}
func (search *Search) GetBook(index int) *BookInfoData {
	if search.Data.Novels != nil && index < len(search.Data.Novels) {
		return &search.Data.Novels[index]
	}
	return nil
}
