package boluobaoLib

import (
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaomodel"
	"github.com/AlexiaVeronica/req/v3"
	"strconv"
)

type API struct {
	HttpRequest *req.Request
}

type Request[T any] struct {
	HttpRequest *req.Request
}

func (request *Request[T]) handlePostResponse(url string, body any) (*T, error) {
	res, err := request.HttpRequest.SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}

	data := new(T)
	if err = res.UnmarshalJson(data); err != nil {
		return nil, err
	}

	if response, ok := any(data).(interface {
		GetCode() int
		GetTip() string
		IsSuccess() bool
	}); ok && !response.IsSuccess() {
		return nil, fmt.Errorf("error: %s", response.GetTip())
	} else if !ok {
		return nil, fmt.Errorf("response does not implement required methods")
	}
	switch v := any(data).(type) {
	case boluobaomodel.LoginStatus:
		for _, cookie := range res.Cookies() {
			v.Cookie += cookie.Name + "=" + cookie.Value + ";"
		}
		if v.Cookie == "" {
			return nil, fmt.Errorf("login failed: cookie is empty")
		}
	}
	return data, nil
}

func (request *Request[T]) handleGetResponse(url string, params map[string]string) (*T, error) {
	res, err := request.HttpRequest.SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("response is nil")
	}

	data := new(T)
	if err = res.UnmarshalJson(data); err != nil {
		return nil, err
	}

	if response, ok := any(data).(interface {
		GetCode() int
		GetTip() string
		IsSuccess() bool
	}); ok && !response.IsSuccess() {
		return nil, fmt.Errorf("error: %s", response.GetTip())
	} else if !ok {
		return nil, fmt.Errorf("response does not implement required methods")
	}
	switch v := any(data).(type) {
	case boluobaomodel.Content:
		if v.Data.Expand.Content == "" {
			return nil, fmt.Errorf("get chapter content failed: no result")
		} else {
			v.Data.Expand.Content = decodeContent(v.Data.Expand.Content)
		}
	}

	return data, nil
}

func newRequest[T any](HttpRequest *req.Request) *Request[T] {
	return &Request[T]{HttpRequest: HttpRequest}
}

func (sfacg *API) GetBookShelfInfo() (*boluobaomodel.InfoData, error) {
	return newRequest[boluobaomodel.InfoData](sfacg.HttpRequest).handleGetResponse("user/Pockets", map[string]string{"expand": "novels"})
}

func (sfacg *API) GetBookInfo(bookId any) (*boluobaomodel.BookInfo, error) {
	return newRequest[boluobaomodel.BookInfo](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("novels/%v", bookId), map[string]string{"expand": bookInfoExpand})
}

func (sfacg *API) GetCatalogue(bookId any) (*boluobaomodel.Catalogue, error) {
	return newRequest[boluobaomodel.Catalogue](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("novels/%v/dirs", bookId), map[string]string{"expand": "originNeedFireMoney"})

}

func (sfacg *API) GetChapterContent(chapterId any) (*boluobaomodel.Content, error) {
	params := map[string]string{"expand": contentInfoExpand, "autoOrder": "false"}
	return newRequest[boluobaomodel.Content](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("Chaps/%v", chapterId), params)
}

func (sfacg *API) GetNewVipContent(bookId any, chapterIds ...int) (*boluobaomodel.Status, error) {
	return newRequest[boluobaomodel.Status](sfacg.HttpRequest).handlePostResponse(fmt.Sprintf("novels/%v/orderedchaps", bookId), struct {
		OrderType string `json:"orderType"`
		OrderAll  bool   `json:"orderAll"`
		AutoOrder bool   `json:"autoOrder"`
		ChapIds   []int  `json:"chapIds"`
	}{OrderType: "readOrder", OrderAll: false, AutoOrder: true, ChapIds: chapterIds})
}

func (sfacg *API) GeContentTsukkomis(row, chapterId, page int) (*boluobaomodel.Tsukkomis, error) {
	params := map[string]string{"expand": "vipLevel,avatar,roleName,widgets,growup", "sort": "data", "page": strconv.Itoa(page), "size": "20", "row": strconv.Itoa(row)}
	return newRequest[boluobaomodel.Tsukkomis](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("chaps/0/%v/tsukkomis", chapterId), params)
}
func (sfacg *API) GetUserInfo() (*boluobaomodel.Account, error) {
	return newRequest[boluobaomodel.Account](sfacg.HttpRequest).handleGetResponse("user", map[string]string{"expand": userInfoExpand})
}

func (sfacg *API) GetUserBuyBooksInfo(page int) (*boluobaomodel.ConsumeData, error) {
	params := map[string]string{"type": "novel", "page": strconv.Itoa(page), "size": "12"}
	return newRequest[boluobaomodel.ConsumeData](sfacg.HttpRequest).handleGetResponse("user/consumeitems", params)
}
func (sfacg *API) GetUserMoney() (*boluobaomodel.Money, error) {
	return newRequest[boluobaomodel.Money](sfacg.HttpRequest).handleGetResponse("user/money", map[string]string{"expand": userInfoExpand})
}

func (sfacg *API) GetCurreyIp() (*boluobaomodel.AccountIp, error) {
	return newRequest[boluobaomodel.AccountIp](sfacg.HttpRequest).handleGetResponse("position", nil)
}

func (sfacg *API) getRankApi(date, rtype, ntype string, page int) (*boluobaomodel.Rank, error) {
	params := map[string]string{"page": strconv.Itoa(page), "size": "50", "rtype": rtype, "ntype": ntype, "expand": bookInfoExpand}
	if params["rtype"] == "sale" {
		params["size"] = "40"
	}
	return newRequest[boluobaomodel.Rank](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("ranks/%s/novels", date), params)
}

func (sfacg *API) GetRankMonthArray(rtype string, page int) (*boluobaomodel.Rank, error) {
	return sfacg.getRankApi("month", rtype, "origin", page)
}
func (sfacg *API) GetRankWeekArray(rtype string, page int) (*boluobaomodel.Rank, error) {
	return sfacg.getRankApi("week", rtype, "origin", page)
}
func (sfacg *API) GetRankAllArray(rtype string, page int) (*boluobaomodel.Rank, error) {
	return sfacg.getRankApi("all", rtype, "origin", page)
}

func (sfacg *API) GetOtherUserInfo(accountId string) (*boluobaomodel.Users, error) {
	return newRequest[boluobaomodel.Users](sfacg.HttpRequest).handleGetResponse("users/"+accountId, nil)
}

func (sfacg *API) GetUserWorks(accountId string) (*boluobaomodel.AuthorInfo, error) {
	params := map[string]string{"expand": "typeName,sysTags,isbranch"}
	return newRequest[boluobaomodel.AuthorInfo](sfacg.HttpRequest).handleGetResponse("users/"+accountId+"/novels", params)
}

func (sfacg *API) Login(username string, password string) (*boluobaomodel.LoginStatus, error) {
	return newRequest[boluobaomodel.LoginStatus](sfacg.HttpRequest).handlePostResponse("sessions", struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{Username: username, Password: password})
}

func (sfacg *API) GetSearch(keyword string, page int) (*boluobaomodel.Search, error) {
	params := map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "15", "expand": bookInfoExpand}
	return newRequest[boluobaomodel.Search](sfacg.HttpRequest).handleGetResponse("search/novels/result", params)
}
