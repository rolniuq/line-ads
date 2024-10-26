package http

import "line-ads/internal"

type httpClient struct {
	settings internal.DialSettings
}

func NewHttpClient(url string, ops ...Option) *httpClient {
	return &httpClient{}
}
