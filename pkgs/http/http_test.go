package http

import (
	"line-ads/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHttpDo(t *testing.T) {
	client := NewHttpClient("https://httpbin.org/get", WithMethod(internal.GET))
	res, err := Do[any](client)

	require.NoError(t, err)
	require.NotNil(t, res)
}
