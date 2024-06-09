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
	return newRequest[boluobaomodel.Users](sfacg.HttpRequest).handleGetResponse(fmt.Sprintf("users/%v", accountId), nil)
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
