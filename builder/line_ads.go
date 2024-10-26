package lineads

import (
	"context"
	"fmt"
)

const (
	AD_ACCOUNT   = "adaccount"
	AD_ACCOUNTS  = "adaccounts"
	CAMPAIGN     = "campaign"
	CHILDREN     = "children"
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

func (s *LineAdsService) SendLinkRequest(ctx context.Context, req ReqCreateLinkRequestDto) (*ResCreateLinkRequestDto, error) {
	res, err := MakeRequest[ReqCreateLinkRequestDto, ResCreateLinkRequestDto](ctx, s.lineAdsBuilder, req, POST, nil, req.GetPath()).Build()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LineAdsService) CreateCampaign(ctx context.Context, accountId string) error {
	return nil
}

func (s *LineAdsService) GetCampaigns(ctx context.Context, accountId string) error {
	return nil
}

func (s *LineAdsService) GetAdAccounts(ctx context.Context, req ReqGetAdAccountsDto) (*ResGetAdAccountsDto, error) {
	path := fmt.Sprintf("%s/%s/%s", GROUPS, req.GroupID, AD_ACCOUNTS)
	res, err := MakeRequest[ReqGetAdAccountsDto, ResGetAdAccountsDto](ctx, s.lineAdsBuilder, req, GET, nil, path).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to get ad accounts: %v", err)
	}

	return res, nil
}

func (s *LineAdsService) GetLinkRequests(ctx context.Context, req ReqGetLinkRequestsDto) (*ResGetLinkRequestsDto, error) {
	res, err := MakeRequest[ReqGetLinkRequestsDto, ResGetLinkRequestsDto](ctx, s.lineAdsBuilder, req, GET, req.GetParameters(), req.GetPath()).Build()
	if err != nil {
		return nil, err
	}

	return res, nil
}
