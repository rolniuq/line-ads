package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendLinkRequest(t *testing.T) {
	// t.Skip()
	res := SendLinkRequest()
	require.NoError(t, res)
}

func TestGetGroups(t *testing.T) {
	t.Skip()
	err := GetChildGroups()
	require.NoError(t, err)
}

func TestGetAdsAccount(t *testing.T) {
	// t.Skip()
	err := GetListAdsAccounts()
	require.NoError(t, err)
}

func TestCreateChildGroup(t *testing.T) {
	t.Skip()
	err := CreateChildGroup()
	require.NoError(t, err)
}

func TestGetLinkRequest(t *testing.T) {
	// t.Skip()
	err := GetLinkRequest()
	require.NoError(t, err)
}
