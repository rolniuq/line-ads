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
	err := GetChildGroups()
	require.NoError(t, err)
}

func TestGetAdsAccount(t *testing.T) {
	err := GetListAdsAccounts()
	require.NoError(t, err)
}

func TestCreateChildGroup(t *testing.T) {
	err := CreateChildGroup()
	require.NoError(t, err)
}

func TestGetLinkRequest(t *testing.T) {
	err := GetLinkRequest()
	require.NoError(t, err)
}
