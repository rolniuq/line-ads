package lineads

import (
	"context"
	"line-ads/configs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinkRequests(t *testing.T) {
	t.Skip()
	configs := configs.ConfigMod.Resolve()

	lineService := NewLineAdsService(configs.ClientId, configs.ClientSecret)

	ctx := context.Background()

	res, err := lineService.GetLinkRequests(ctx, ReqGetLinkRequestsDto{
		GroupID: "G08916310298",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
}
