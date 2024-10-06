package lineads

import (
	"context"
	"fmt"
)

const (
	AD_ACCOUNT   = "adaccount"
	CAMPAIGN     = "campaign"
	GROUPS       = "groups"
	LINK_REQUEST = "link-request"
)

type LineAdsService struct {
	lineAdsBuilder *LineAdsBuilder
}

func NewLineAdsService(accessKey, secretKey string) *LineAdsService {
	lineAdsBuilder := NewLineAdsBuilder(accessKey, secretKey)

	return &LineAdsService{
		lineAdsBuilder: lineAdsBuilder,
	}
}

func (s *LineAdsService) CreateGroup(_ context.Context, req ReqCreateGroupDto) (*ResCreateGroupDto, error) {
	res, err := MakeRequest[ReqCreateGroupDto, ResCreateGroupDto](s.lineAdsBuilder, req, POST, nil, GROUPS).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %v", err)
	}

	return res, nil
}

func (s *LineAdsService) SendLinkRequest(_ context.Context, req ReqCreateLinkRequestDto) (*ResCreateLinkRequestDto, error) {
	path := fmt.Sprintf("/%s/%s/%s/%s", GROUPS, req.SourceGroupID, LINK_REQUEST, AD_ACCOUNT)
	res, err := MakeRequest[ReqCreateLinkRequestDto, ResCreateLinkRequestDto](s.lineAdsBuilder, req, POST, nil, path).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to send link request: %v", err)
	}

	return res, nil
}

func (s *LineAdsService) CreateCampaign(accountId string) error {
	return nil
}

func (s *LineAdsService) GetCampaigns(accountId string) error {
	return nil
}
