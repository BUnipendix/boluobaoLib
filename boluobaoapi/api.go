package boluobaoapi

import (
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaomodel"
	"log"
	"strconv"
)

type API struct {
	HttpClient HttpsClient
}

func (sfacg *API) GetBookShelfInfo() ([]boluobaomodel.ShelfData, error) {
	var m boluobaomodel.InfoData
	_, err := sfacg.HttpClient.Get("user/Pockets", map[string]string{"expand": "novels"}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get book shelf information failed: %v", m.Status.Msg)
	}
	if len(m.Data) == 0 {
		return nil, fmt.Errorf("get book shelf information failed: no result")
	}
	return m.Data, nil
}

func (sfacg *API) GetBookInfo(bookId any) (*boluobaomodel.BookInfoData, error) {
	var m boluobaomodel.BookInfo
	_, err := sfacg.HttpClient.Get(fmt.Sprintf("novels/%v", bookId), map[string]string{"expand": bookInfoExpand}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get book information failed: %v", m.Status.Msg)
	}
	return &m.Data, nil
}

func (sfacg *API) GetCatalogue(bookId any) ([]boluobaomodel.VolumeList, error) {
	var m boluobaomodel.Catalogue
	_, err := sfacg.HttpClient.Get(fmt.Sprintf("novels/%v/dirs", bookId), map[string]string{"expand": "originNeedFireMoney"}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get catalogue failed: %v", m.Status.Msg)
	}
	return m.Data.VolumeList, nil

}

func (sfacg *API) GetChapterContent(chapterId any) (*boluobaomodel.ContentData, error) {
	var m boluobaomodel.Content
	params := map[string]string{"expand": contentInfoExpand, "autoOrder": "false"}
	_, err := sfacg.HttpClient.Get(fmt.Sprintf("Chaps/%v", chapterId), params, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get chapter content failed: %v", m.Status.Msg)
	}
	if m.Data.Expand.Content == "" {
		return nil, fmt.Errorf("get chapter content failed: no result")
	}
	return &m.Data, nil
}

func (sfacg *API) GetNewVipContent(bookId any, chapterId int) (bool, error) {
	var m boluobaomodel.Status
	_, err := sfacg.HttpClient.Post(fmt.Sprintf("novels/%v/orderedchaps", bookId), &BuyVipContentBody{
		OrderType: "readOrder",
		OrderAll:  false,
		AutoOrder: true,
		ChapIds:   []int{chapterId},
	}, &m)
	if err != nil {
		return false, fmt.Errorf("request failed: %v", err)
	}
	if m.ErrorCode != 200 {
		return false, fmt.Errorf("buy vip content failed: %v", m.Msg)
	}
	return true, nil
}

func (sfacg *API) GeContentTsukkomis(row, chapterId, page int) ([]boluobaomodel.TsukkomisData, error) {
	var m boluobaomodel.Tsukkomis
	params := map[string]string{"expand": "vipLevel,avatar,roleName,widgets,growup", "sort": "data", "page": strconv.Itoa(page), "size": "20", "row": strconv.Itoa(row)}
	_, err := sfacg.HttpClient.Get(fmt.Sprintf("chaps/0/%v/tsukkomis", chapterId), params, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get chapter tsukkomis failed: %v", m.Status.Msg)
	}
	return m.Data, nil
}
func (sfacg *API) GetUserInfo() (*boluobaomodel.AccountData, error) {
	var m boluobaomodel.Account
	_, err := sfacg.HttpClient.Get("user", map[string]string{"expand": userInfoExpand}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get account information failed: %v", m.Status.Msg)
	}
	return &m.Data, nil
}

func (sfacg *API) GetUserBuyBooksInfo(page int) ([]boluobaomodel.ConsumeData, error) {
	var m boluobaomodel.Consume
	_, err := sfacg.HttpClient.Get("user/consumeitems", map[string]string{"type": "novel", "page": strconv.Itoa(page), "size": "12"}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get user buyBooks information failed: %v", m.Status.Msg)
	}
	return m.Data, nil
}
func (sfacg *API) GetUserMoney() (*boluobaomodel.MoneyData, error) {
	var m boluobaomodel.Money
	_, err := sfacg.HttpClient.Get("user/money", map[string]string{"expand": userInfoExpand}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get account information failed: %v", m.Status.Msg)
	}
	return &m.Data, nil
}

func (sfacg *API) GetCurreyIp() (*boluobaomodel.AccountIpData, error) {
	var m boluobaomodel.AccountIp
	_, err := sfacg.HttpClient.Get("position", nil, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get account ip failed: %v", m.Status.Msg)
	}
	return &m.Data, nil
}

func (sfacg *API) getRankApi(date, rtype, ntype string, page int) ([]boluobaomodel.BookInfoData, error) {
	var m boluobaomodel.Rank
	params := map[string]string{"page": strconv.Itoa(page), "size": "50", "rtype": rtype, "ntype": ntype, "expand": bookInfoExpand}
	if params["rtype"] == "sale" {
		params["size"] = "40"
	}
	_, err := sfacg.HttpClient.Get(fmt.Sprintf("ranks/%s/novels", date), params, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get rank array failed: %v", m.Status.Msg)
	}
	if len(m.Data) == 0 {
		return nil, fmt.Errorf("get rank array failed: no result")
	}
	return m.Data, nil
}

func (sfacg *API) GetRankMonthArray(rtype string, page int) ([]boluobaomodel.BookInfoData, error) {
	return sfacg.getRankApi("month", rtype, "origin", page)
}
func (sfacg *API) GetRankWeekArray(rtype string, page int) ([]boluobaomodel.BookInfoData, error) {
	return sfacg.getRankApi("week", rtype, "origin", page)
}
func (sfacg *API) GetRankAllArray(rtype string, page int) ([]boluobaomodel.BookInfoData, error) {
	return sfacg.getRankApi("all", rtype, "origin", page)
}

func (sfacg *API) GetOtherUserInfo(accountId string) (*boluobaomodel.AccountData, error) {
	var m boluobaomodel.Users
	_, err := sfacg.HttpClient.Get("users/"+accountId, nil, &m)
	if err != nil {
		log.Println("get user information failed:", err)
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get user information failed: %v", m.Status.Msg)
	}
	return &m.Data, nil
}

func (sfacg *API) GetUserWorks(accountId string) ([]boluobaomodel.BookInfoData, error) {
	var m boluobaomodel.AuthorInfo
	_, err := sfacg.HttpClient.Get("users/"+accountId+"/novels", map[string]string{"expand": "typeName,sysTags,isbranch"}, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HTTPCode != 200 {
		return nil, fmt.Errorf("get user works failed: %v", m.Status.Msg)
	}
	if len(m.Data) == 0 {
		return nil, fmt.Errorf("get user works failed: no result")
	}
	return m.Data, nil
}

func (sfacg *API) Login(username string, password string) (string, error) {
	var m boluobaomodel.Status
	response, err := sfacg.HttpClient.Post("sessions", &LoginBody{Username: username, Password: password}, &m)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	if m.HTTPCode != 200 {
		return "", fmt.Errorf("login failed: %v", m.Msg)
	}
	var loginCookie string
	for _, cookie := range response.Cookies() {
		loginCookie += cookie.Name + "=" + cookie.Value + ";"
	}
	if loginCookie == "" {
		return "", fmt.Errorf("login failed: cookie is empty")
	}
	return loginCookie, nil
}

func (sfacg *API) GetSearch(keyword string, page int) ([]boluobaomodel.BookInfoData, error) {
	var m boluobaomodel.Search
	params := map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "15", "expand": bookInfoExpand}
	_, err := sfacg.HttpClient.Get("/search/novels/result", params, &m)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	if m.Status.HttpCode != 200 {
		return nil, fmt.Errorf("get search failed: %v", m.Status.Msg)
	}
	if len(m.Data.Novels) == 0 {
		return nil, fmt.Errorf("get search failed: no result")
	}
	return m.Data.Novels, nil
}
