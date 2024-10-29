package http

import "line-ads/internal"

type httpClient struct {
	settings internal.DialSettings
}

func NewHttpClient(url string, ops ...Option) *httpClient {
	settings := GetDialSettings(ops)
	if settings == nil {
		settings = GetDefaultSettings()
	}

	return &httpClient{
		settings: *settings,
	}
}
