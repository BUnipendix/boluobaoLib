package boluobaoLib

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaoapi"
	"github.com/AlexiaVeronica/req/v3"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	HttpClient    *req.Client
	baseURL       string
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

func generateDeviceId() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func NewClient(options ...Options) *Client {
	c := &Client{
		HttpClient:    req.NewClient().SetTimeout(timeout * time.Second),
		DeviceId:      generateDeviceId(),
		baseURL:       baseURL,
		AndroidApiKey: androidApiKey,
	}

	c.UserAgent = fmt.Sprintf(userAgent, c.DeviceId)
	for _, option := range options {
		option.Apply(c)
	}
	c.HttpClient.SetCommonBasicAuth(username, password)
	c.HttpClient.SetCommonHeader("User-Agent", c.UserAgent)

	c.HttpClient.SetBaseURL(c.baseURL)
	c.HttpClient.SetCommonCookies()
	if c.ProxyURL != "" {
		c.HttpClient.SetProxyURL(c.ProxyURL)
	} else if len(c.ProxyURLArray) > 0 {
		c.HttpClient.SetProxyURL(c.ProxyURLArray[time.Now().UnixNano()%int64(len(c.ProxyURLArray))])
	}
	if c.Debug {
		c.HttpClient.DevMode()
	}
	if c.Cookie != "" {
		c.HttpClient.SetCommonHeader("Cookie", c.Cookie)

	}
	return c
}
func (client *Client) API() *boluobaoapi.API {
	return &boluobaoapi.API{

		HttpRequest: client.HttpClient.R().SetHeader("SFSecurity", client.security()),
	}
}

func (client *Client) security() string {
	t := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	uuId, newMd5 := strings.ToUpper(generateDeviceId()), md5.New()
	newMd5.Write([]byte(uuId + t + strings.ToUpper(client.DeviceId) + client.AndroidApiKey))
	return url.Values{
		"nonce":       {uuId},
		"timestamp":   {t},
		"devicetoken": {strings.ToUpper(client.DeviceId)},
		"sign":        {strings.ToUpper(hex.EncodeToString(newMd5.Sum(nil)))},
	}.Encode()
}
