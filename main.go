package main

import (
	"fmt"
	"line-ads/test"
)

func main() {
	// sample.ReadGroups()
	// test.CreateChildGroup()

	// l := NewLineClient()
	// l.Auth()

	// test.GetListAdsAccounts()
	err := test.SendLinkRequest()
	if err != nil {
		fmt.Errorf("failed to send link request: %v", err)
	}
}
