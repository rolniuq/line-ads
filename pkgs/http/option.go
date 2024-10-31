package http

import (
	"line-ads/internal"
	"net/http"
)

type Option interface {
	Apply(*internal.DialSettings)
}

type withMethod internal.HttpMethod

func (w withMethod) Apply(s *internal.DialSettings) {
	s.Method = internal.HttpMethod(w)
}

func WithMethod(method internal.HttpMethod) Option {
	return withMethod(method)
}

type withHeader http.Header

func (w withHeader) Apply(s *internal.DialSettings) {
	s.Header = http.Header(w)
}

func WithHeader(header http.Header) Option {
	return withHeader(header)
}

type withBody internal.HttpBody

func (w withBody) Apply(s *internal.DialSettings) {
	s.Body = internal.HttpBody(w)
}

func WithBody(body internal.HttpBody) Option {
	return withBody(body)
}

type withTimeOut int

func (w withTimeOut) Apply(s *internal.DialSettings) {
	s.TimeOut = int(w)
}

func WithTimeOut(t int) Option {
	return withTimeOut(t)
}

func GetDialSettings(ops []Option) *internal.DialSettings {
	res := &internal.DialSettings{}

	for _, op := range ops {
		if op == nil {
			continue
		}

		op.Apply(res)
	}

	return res
}

func GetDefaultSettings() *internal.DialSettings {
	return &internal.DialSettings{
		Method: internal.GET,
	}
}
