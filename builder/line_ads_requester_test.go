package lineads

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LineAdsRequester_getEndpoint(t *testing.T) {
	clientId := ""
	clientSecret := ""

	endpoint := NewLineAdsRequest[any](clientId, clientSecret).
		WithUrl("https://api.line.me/v2/bot/message/push").
		WithParameters(LineAdsRequestParameters{
			GROUPS: "G08916310298",
		}).
		getEndpoint()

	require.NotNil(t, endpoint)

	url, err := url.Parse(endpoint)
	require.NoError(t, err)
	require.Equal(t, "https://api.line.me/v2/bot/message/push?groups=G08916310298", url.String())
}
