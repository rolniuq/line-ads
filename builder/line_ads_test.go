package lineads

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendLinkInvite(t *testing.T) {
	clientId := ""
	clientSecret := ""

	lineService := NewLineAdsService(clientId, clientSecret)

	ctx := context.Background()

	res, err := lineService.SendLinkRequest(ctx, ReqCreateLinkRequestDto{
		TargetAdAccountId: "A08655312340",
	})
	require.NoError(t, err)
	require.NotNil(t, res)
}
