package boluobaomodel

type Search struct {
	Status `json:"status"`
	Data   struct {
		Novels []BookInfoData `json:"novels"`
		Comics []interface{}  `json:"comics"`
		Albums []interface{}  `json:"albums"`
	} `json:"data"`
}

func (search *Search) EachBook(f func(int, BookInfoData)) {
	for i, book := range search.Data.Novels {
		f(i, book)
	}
}

func (search *Search) GetBook(index int) *BookInfoData {
	if index >= 0 && index < len(search.Data.Novels) {
		return &search.Data.Novels[index]
	}
	return nil
}
