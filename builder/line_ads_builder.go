package lineads

import "fmt"

const (
	host       = "ads.line.me"
	basePath   = "/api"
	scheme     = "https"
	apiVersion = "v3"
)

type LineAdsBuilder struct {
	accessKey  string
	basePath   string
	host       string
	scheme     string
	secretKey  string
	apiVersion string
}

func NewLineAdsBuilder(accessKey, secretKey string) *LineAdsBuilder {
	return &LineAdsBuilder{
		accessKey:  accessKey,
		basePath:   basePath,
		host:       host,
		scheme:     scheme,
		secretKey:  secretKey,
		apiVersion: apiVersion,
	}
}

func MakeRequest[T, U any](
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
		WithBody(body).
		WithMethod(method).
		WithUrl(fmt.Sprintf("%s://%s%s/%s/%s", s.scheme, s.host, s.basePath, s.apiVersion, path)).
		WithParameters(parameters)

	return req
}
