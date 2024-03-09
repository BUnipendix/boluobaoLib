package boluobaoLib

type Options interface {
	Apply(client *Client)
}
type OptionFunc func(client *Client)

func (f OptionFunc) Apply(client *Client) {
	f(client)
}

func WithCookie(cookie string) Options {
	return OptionFunc(func(client *Client) {
		client.Cookie = cookie
	})
}
func WithDeviceId(deviceId string) Options {
	return OptionFunc(func(client *Client) {
		client.DeviceId = deviceId
	})
}
func WithDebug() Options {
	return OptionFunc(func(client *Client) {
		client.Debug = true
	})
}

func WithOutputDebug() Options {
	return OptionFunc(func(client *Client) {
		//if client.API.HttpClient.OutputDebug {
		//	client.API.HttpClient.OutputDebug = false
		//} else {
		//	client.API.HttpClient.OutputDebug = true
		//}
	})
}
func WithProxyURLArray(proxyURLArray []string) Options {
	return OptionFunc(func(client *Client) {
		client.ProxyURLArray = proxyURLArray
	})
}
func WithProxyURL(proxyURL string) Options {
	return OptionFunc(func(client *Client) {
		client.ProxyURL = proxyURL
	})
}

func WithAPIBaseURL(apiBaseURL string) Options {
	return OptionFunc(func(client *Client) {
		client.baseURL = apiBaseURL
	})
}
func WithUserAgent(userAgent string) Options {
	return OptionFunc(func(client *Client) {
		client.UserAgent = userAgent
	})
}
func WithAuthorization(authorization string) Options {
	return OptionFunc(func(client *Client) {
		client.Authorization = authorization
	})
}
func WithAndroidApiKey(androidApiKey string) Options {
	return OptionFunc(func(client *Client) {
		client.AndroidApiKey = androidApiKey
	})
}
func WithSFCommunity(sfCommunity string) Options {
	return OptionFunc(func(client *Client) {
		client.Cookie = client.Cookie + ".SFCommunity=" + sfCommunity + ";"
	})
}

func WithSessionAPP(sessionApp string) Options {
	return OptionFunc(func(client *Client) {
		client.Cookie = client.Cookie + "session_APP=" + sessionApp + ";"
	})
}
