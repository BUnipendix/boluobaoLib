package boluobaoLib

import (
	"crypto/rand"
	"fmt"
	"github.com/AlexiaVeronica/boluobaoLib/boluobaoapi"
)

type Client struct {
	API *boluobaoapi.API
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
	deviceId := generateDeviceId()
	c := &Client{
		API: &boluobaoapi.API{
			HttpClient: boluobaoapi.HttpsClient{
				Debug:         false,
				OutputDebug:   false,
				DeviceId:      deviceId,
				AndroidApiKey: "FMLxgOdsfxmN!Dt4",
				APIBaseURL:    "https://api.sfacg.com",
				Authorization: "Basic YW5kcm9pZHVzZXI6MWEjJDUxLXl0Njk7KkFjdkBxeHE=",
				UserAgent:     "boluobao/4.8.42(android;25)/XIAOMI/" + deviceId + "/OPPO",
			},
		},
	}
	for _, option := range options {
		option.Apply(c)
	}
	return c
}

type Options interface {
	Apply(client *Client)
}
type OptionFunc func(client *Client)

func (f OptionFunc) Apply(client *Client) {
	f(client)
}
func WithCookie(cookie string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.Cookie = cookie
	})
}
func WithDeviceId(deviceId string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.DeviceId = deviceId
	})
}
func WithDebug() Options {
	return OptionFunc(func(client *Client) {
		if client.API.HttpClient.Debug {
			client.API.HttpClient.Debug = false
		} else {
			client.API.HttpClient.Debug = true
		}
	})
}

func WithOutputDebug() Options {
	return OptionFunc(func(client *Client) {
		if client.API.HttpClient.OutputDebug {
			client.API.HttpClient.OutputDebug = false
		} else {
			client.API.HttpClient.OutputDebug = true
		}
	})
}
func WithProxyURLArray(proxyURLArray []string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.ProxyURLArray = proxyURLArray
	})
}
func WithProxyURL(proxyURL string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.ProxyURL = proxyURL
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.APIBaseURL = apiBaseURL
	})
}
func WithUserAgent(userAgent string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.UserAgent = userAgent
	})
}
func WithAuthorization(authorization string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.Authorization = authorization
	})
}
func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.API.HttpClient.AndroidApiKey = androidApiKey
	})
}
