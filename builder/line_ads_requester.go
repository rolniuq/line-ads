package lineads

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"line-ads/utils"
	"net/http"
	"net/url"
	"time"
)

const (
	APPLICATION_JSON = "application/json"
)

type LineAdsRequestMethod string

type LineAdsRequestParameters map[string]string

func (s *LineAdsRequestParameters) Add(key string, value string) {
	if s == nil {
		return
	}

	(*s)[key] = value
}

func (s *LineAdsRequestParameters) String() string {
	if s == nil {
		return ""
	}

	values := url.Values{}
	for k, v := range *s {
		values.Add(k, v)
	}

	return values.Encode()
}

const (
	GET  LineAdsRequestMethod = "GET"
	POST LineAdsRequestMethod = "POST"
)

type LineAdsRequest[T any] struct {
	ctx        context.Context
	accessKey  string
	body       any
	method     LineAdsRequestMethod
	parameters LineAdsRequestParameters
	secretKey  string
	url        string
}

func NewLineAdsRequest[T any](accessKey, secretKey string) *LineAdsRequest[T] {
	return &LineAdsRequest[T]{
		accessKey:  accessKey,
		secretKey:  secretKey,
		parameters: make(LineAdsRequestParameters),
	}
}

func (s *LineAdsRequest[T]) WithContext(ctx context.Context) *LineAdsRequest[T] {
	s.ctx = ctx

	return s
}

func (s *LineAdsRequest[T]) WithBody(body any) *LineAdsRequest[T] {
	s.body = body

	return s
}

func (s *LineAdsRequest[T]) WithMethod(method LineAdsRequestMethod) *LineAdsRequest[T] {
	s.method = method

	return s
}

func (s *LineAdsRequest[T]) WithUrl(url string) *LineAdsRequest[T] {
	s.url = url

	return s
}

func (s *LineAdsRequest[T]) WithParameters(parameters LineAdsRequestParameters) *LineAdsRequest[T] {
	s.parameters = parameters

	return s
}

func (s *LineAdsRequest[T]) getEndpoint() string {
	if s == nil || s.parameters == nil {
		return s.url
	}

	return fmt.Sprintf("%s?%s", s.url, s.parameters.String())
}

func (s *LineAdsRequest[T]) getPath() (*string, error) {
	if s == nil || s.url == "" {
		return nil, nil
	}

	url, err := url.Parse(s.url)
	if err != nil {
		return nil, err
	}

	return &url.Path, nil
}

func (s *LineAdsRequest[T]) getRequestBody() ([]byte, error) {
	res, err := json.Marshal(s.body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	return res, nil
}

func (s *LineAdsRequest[T]) getPayloadStr() (*string, error) {
	if s == nil || s.body == nil {
		return nil, nil
	}

	reqBody, err := json.Marshal(s.body)
	if err != nil {
		return nil, err
	}

	contentType, err := s.getContentType()
	if err != nil {
		return nil, err
	}

	path, err := s.getPath()
	if err != nil {
		return nil, err
	}
	if path == nil {
		return nil, fmt.Errorf("path is required")
	}

	hexDigest := utils.CalcSHA256Digest(string(reqBody))
	payloadDate := time.Now().UTC().Format("20060102")

	payload := fmt.Sprintf("%s\n%s\n%s\n%s", hexDigest, contentType, payloadDate, *path)

	return &payload, nil
}

func (s *LineAdsRequest[T]) getToken() string {
	payload, err := s.getPayloadStr()
	if err != nil {
		return ""
	}
	if payload == nil {
		return ""
	}

	jwsHeader := utils.EncodeWithBase64([]byte(fmt.Sprintf(`{"alg":"HS256","kid":"%s","typ":"text/plain"}`, s.accessKey)))
	jwsPayload := utils.EncodeWithBase64([]byte(*payload))

	signingInput := fmt.Sprintf("%s.%s", jwsHeader, jwsPayload)
	signature := hmac.New(sha256.New, []byte(s.secretKey))
	signature.Write([]byte(signingInput))
	encodedSignature := utils.EncodeWithBase64(signature.Sum(nil))

	return fmt.Sprintf("%s.%s.%s", jwsHeader, jwsPayload, encodedSignature)
}

func (s *LineAdsRequest[T]) getContentType() (string, error) {
	reqBody, err := s.getRequestBody()
	if err != nil {
		return "", fmt.Errorf("failed to get request body: %v", err)
	}

	if reqBody == nil {
		return "", nil
	}

	return APPLICATION_JSON, nil
}

func (s *LineAdsRequest[T]) getHeaders(t time.Time) (map[string][]string, error) {
	token := s.getToken()
	if token == "" {
		return nil, fmt.Errorf("token signature is invalid")
	}

	contentType, err := s.getContentType()
	if err != nil {
		return nil, fmt.Errorf("failed to get content type: %v", err)
	}

	return http.Header{
		"Date":          {t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")},
		"Authorization": {"Bearer " + s.getToken()},
		"Content-Type":  {contentType},
	}, nil
}

func (s *LineAdsRequest[T]) makeRequest(t time.Time) (*http.Request, error) {
	endpoint := s.getEndpoint()
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}

	body, err := s.getRequestBody()
	if err != nil {
		return nil, fmt.Errorf("failed to get request body: %v", err)
	}
	if body == nil {
		return nil, fmt.Errorf("request body is required")
	}
	req, err := http.NewRequest(string(s.method), endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	header, err := s.getHeaders(t)
	if err != nil {
		return nil, fmt.Errorf("failed to get headers: %v", err)
	}

	req.Header = header

	return req, nil
}

func (s *LineAdsRequest[T]) Build() (*T, error) {
	req, err := s.makeRequest(time.Now())
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, s.getError(body)
	}
	var data T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
