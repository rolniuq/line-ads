package lineads

import (
	"context"
	"fmt"
)

const (
	host       = "ads.line.me"
	basePath   = "api"
	scheme     = "https"
	apiVersion = "v3"
)

type LineAdsBuilder struct {
	accessKey  string
	apiVersion string
	basePath   string
	host       string
	scheme     string
	secretKey  string
}

func NewLineAdsBuilder(accessKey, secretKey string) *LineAdsBuilder {
	return &LineAdsBuilder{
		accessKey:  accessKey,
		apiVersion: apiVersion,
		basePath:   basePath,
		host:       host,
		scheme:     scheme,
		secretKey:  secretKey,
	}
}

func MakeRequest[T, U any](
	ctx context.Context,
	s *LineAdsBuilder,
	body T,
	method LineAdsRequestMethod,
	parameters LineAdsRequestParameters,
	path string,
) *LineAdsRequest[U] {
	if s == nil {
		return nil
	}

	req := NewLineAdsRequest[U](s.accessKey, s.secretKey).
		WithContext(ctx).
		WithBody(body).
		WithMethod(method).
		WithUrl(fmt.Sprintf("%s://%s/%s/%s/%s", s.scheme, s.host, s.basePath, s.apiVersion, path)).
		WithParameters(parameters)

	return req
}
