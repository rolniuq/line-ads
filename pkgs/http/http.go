package http

import (
	"bytes"
	"encoding/json"
	"io"
	ds "line-ads/internal/dial_settings"
	"net/http"
	"net/url"
	"time"
)

type httpClient struct {
	url      *url.URL
	settings *ds.DialSettings
}

func NewHttpClient(urlStr string, ops ...Option) *httpClient {
	settings := GetDialSettings(ops)
	if settings == nil {
		settings = GetDefaultSettings()
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	return &httpClient{
		settings: settings,
		url:      url,
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

	req, err := http.NewRequest(string(c.settings.Method), c.url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header = c.settings.Header

	return req, nil
}

func Do[Res any](httpClient *httpClient) (*Res, error) {
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

	body, err := io.ReadAll(res.Body)
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
