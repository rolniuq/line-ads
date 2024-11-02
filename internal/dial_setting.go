package internal

import "net/http"

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PATCH  HttpMethod = "PATCH"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

type HttpBody map[string]any

type DialSettings struct {
	Header  http.Header
	Body    HttpBody
	Method  HttpMethod
	TimeOut int
}
