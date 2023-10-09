package boluobaoapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type HttpsClient struct {
	APIBaseURL    string
	UserAgent     string
	AndroidApiKey string
	Authorization string
	Cookie        string
	DeviceId      string
	Debug         bool
	OutputDebug   bool
	ProxyURL      string
	ProxyURLArray []string
}

func (sfacg *HttpsClient) security() string {
	t := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	uuId, newMd5 := strings.ToUpper(uuid.New().String()), md5.New()
	newMd5.Write([]byte(uuId + t + strings.ToUpper(sfacg.DeviceId) + sfacg.AndroidApiKey))
	return url.Values{
		"nonce":       {uuId},
		"timestamp":   {t},
		"devicetoken": {strings.ToUpper(sfacg.DeviceId)},
		"sign":        {strings.ToUpper(hex.EncodeToString(newMd5.Sum(nil)))},
	}.Encode()
}
func (sfacg *HttpsClient) defaultHeader() map[string]string {
	return map[string]string{"User-Agent": sfacg.UserAgent, "Authorization": sfacg.Authorization, "SFSecurity": sfacg.security()}

}
func (sfacg *HttpsClient) NewDefault() *req.Client {
	c := req.C().
		SetTimeout(30 * time.Second).
		SetBaseURL(sfacg.APIBaseURL).
		SetCommonHeaders(sfacg.defaultHeader())
	if sfacg.Cookie != "" {
		c.SetCommonHeader("Cookie", sfacg.Cookie)
	}
	if sfacg.ProxyURL != "" {
		c.SetProxyURL(sfacg.ProxyURL)
	}
	if len(sfacg.ProxyURLArray) > 0 {
		c.SetProxyURL(sfacg.ProxyURLArray[time.Now().UnixNano()%int64(len(sfacg.ProxyURLArray))])
	}
	if sfacg.Debug {
		c.DevMode()
		if sfacg.OutputDebug {
			c.EnableDumpAllToFile("ReqAllLogToFile.log")
		}
	}
	return c
}
func (sfacg *HttpsClient) Get(path string, params map[string]string, model any) (*req.Response, error) {
	newDefault := sfacg.NewDefault().R().SetQueryParams(params)
	response, err := newDefault.Get(path)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccessState() {
		return nil, fmt.Errorf("response is not success state: %v", response.String())
	}
	if model != nil {
		err = response.UnmarshalJson(model)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}
func (sfacg *HttpsClient) Post(path string, params any, model any) (*req.Response, error) {
	newDefault := sfacg.NewDefault().R().SetHeader("Content-Type", "application/json")
	if rv := reflect.ValueOf(params); rv.Kind() == reflect.Map {
		newDefault.SetFormData(params.(map[string]string))
	} else if rv.Kind() == reflect.Struct {
		newDefault.SetBody(params)
	}
	response, err := newDefault.Post(path)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccessState() {
		return nil, fmt.Errorf("response is not success state: %v", response.String())
	}
	if model != nil {
		err = response.UnmarshalJson(model)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}
