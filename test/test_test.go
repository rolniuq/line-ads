package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendLinkRequest(t *testing.T) {
	res := SendLinkRequest()
	require.NoError(t, res)
}

func TestGetGroups(t *testing.T) {
	GetChildGroups()
}

func TestGetAdsAccount(t *testing.T) {
	err := GetLinkRequest()
	require.NoError(t, err)
}
