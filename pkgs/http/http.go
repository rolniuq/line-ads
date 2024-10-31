package http

import (
	"bytes"
	"encoding/json"
	"io"
	"line-ads/internal"
	"net/http"
	"time"
)

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

func (c *httpClient) makeRequest() (*http.Request, error) {
	var body io.Reader
	if c.settings.Body != nil {
		b, err := json.Marshal(c.settings.Body)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(c.settings.Url, string(c.settings.Method), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func Do[Req, Res any](httpClient *httpClient) (*Res, any) {
	req, err := httpClient.makeRequest()
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: http.DefaultClient.Timeout * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var t Res
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
