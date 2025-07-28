package http

import (
	ds "line-ads/internal/dial_settings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHttpDo(t *testing.T) {
	client := NewHttpClient("https://httpbin.org/get", WithMethod(ds.GET))
	res, err := Do[any](client)

	require.NoError(t, err)
	require.NotNil(t, res)
}
