package http

import (
	ds "line-ads/internal/dial_settings"
	"net/http"
)

type Option interface {
	Apply(*ds.DialSettings)
}

type withMethod ds.HttpMethod

func (w withMethod) Apply(s *ds.DialSettings) {
	s.Method = ds.HttpMethod(w)
}

func WithMethod(method ds.HttpMethod) Option {
	return withMethod(method)
}

type withHeader http.Header

func (w withHeader) Apply(s *ds.DialSettings) {
	s.Header = http.Header(w)
}

func WithHeader(header http.Header) Option {
	return withHeader(header)
}

type withBody ds.HttpBody

func (w withBody) Apply(s *ds.DialSettings) {
	s.Body = ds.HttpBody(w)
}

func WithBody(body ds.HttpBody) Option {
	return withBody(body)
}

type withTimeOut int

func (w withTimeOut) Apply(s *ds.DialSettings) {
	s.TimeOut = int(w)
}

func WithTimeOut(t int) Option {
	return withTimeOut(t)
}

func GetDialSettings(ops []Option) *ds.DialSettings {
	res := &ds.DialSettings{}

	for _, op := range ops {
		if op == nil {
			continue
		}

		op.Apply(res)
	}

	return res
}

func GetDefaultSettings() *ds.DialSettings {
	return &ds.DialSettings{
		Method: ds.GET,
	}
}
