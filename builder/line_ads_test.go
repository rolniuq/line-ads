package lineads

import (
	"context"
	"line-ads/configs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendLinkInvite(t *testing.T) {
	configs := configs.ConfigMod.Resolve()

	lineService := NewLineAdsService(configs.ClientId, configs.ClientSecret)

	ctx := context.Background()

	res, err := lineService.SendLinkRequest(ctx, ReqCreateLinkRequestDto{
		SourceGroupID:     "G08916310298",
		TargetAdAccountId: "A08655312340",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
}
