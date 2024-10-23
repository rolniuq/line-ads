package lineads

import (
	"line-ads/configs"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LineAdsRequester_getEndpoint(t *testing.T) {
	configs := configs.ConfigMod.Resolve()

	endpoint := NewLineAdsRequest[any](configs.ClientId, configs.ClientSecret).
		WithUrl("https://api.line.me/v3/groups").
		WithParameters(LineAdsRequestParameters{
			GROUPS: "G08916310298",
		}).
		getEndpoint()

	require.NotNil(t, endpoint)

	url, err := url.Parse(endpoint)
	require.NoError(t, err)
	require.Equal(t, "https://api.line.me/v3/groups?groups=G08916310298", url.String())
}

func Test_LineAdsRequester_getToken(t *testing.T) {
	configs := configs.ConfigMod.Resolve()

	token := NewLineAdsRequest[any](configs.ClientId, configs.ClientSecret).getToken()

	require.NotNil(t, token)
}
