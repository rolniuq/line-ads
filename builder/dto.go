package lineads

import (
	"fmt"
)

type ReqLineAdsPaging struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}

type ResLineAdsPaging struct {
	Page          *int     `json:"page"`
	Size          *int     `json:"size"`
	TotalElements *int     `json:"totalElements"`
	Sorts         []string `json:"sorts"`
}

type ReqCreateLinkRequestDto struct {
	SourceGroupID     string `json:"sourceGroupId"`
	TargetAdaccountID string `json:"targetAdaccountId"`
}

func (r *ReqCreateLinkRequestDto) GetPath() string {
	if r == nil || r.SourceGroupID == "" {
		return ""
	}

	return fmt.Sprintf("%s/%s/%s/%s", GROUPS, r.SourceGroupID, LINK_REQUEST, AD_ACCOUNT)
}

type ResCreateLinkRequestDto struct {
	ID                  int    `json:"id"`
	SourceGroupID       string `json:"sourceGroupId"`
	SourceGroupName     string `json:"sourceGroupName"`
	TargetAdaccountID   string `json:"targetAdaccountId"`
	TargetAdaccountName string `json:"targetAdaccountName"`
	Status              string `json:"status"`
	TargetType          string `json:"targetType"`
	CreatedDate         string `json:"createdDate"`
}

type ReqGetAdAccountsDto struct {
	GroupID string `json:"groupId"`
}

type LineAccount struct {
	Name   string `json:"name"`
	LineID string `json:"line_id"`
}

type DeliveryStatusReason struct {
	Code string `json:"code"`
}

type Adaccount struct {
	ID                         string                 `json:"id"`
	Name                       string                 `json:"name"`
	ConfiguredStatus           string                 `json:"configured_status"`
	ProductType                string                 `json:"product_type"`
	AvailableCampaignObjective []string               `json:"available_campaign_objective"`
	Currency                   string                 `json:"currency"`
	Timezone                   string                 `json:"timezone"`
	Country                    string                 `json:"country"`
	LineAccount                LineAccount            `json:"line_account"`
	DeliveryStatus             string                 `json:"delivery_status"`
	DeliveryStatusReasons      []DeliveryStatusReason `json:"delivery_status_reasons"`
	CreatedDate                string                 `json:"created_date"`
	ModifiedDate               string                 `json:"modified_date"`
	RemovedDate                string                 `json:"removed_date"`
}

type ResGetAdAccountsDto struct {
	Datas []Adaccount `json:"datas"`
}

type ReqGetLinkRequestsDto struct {
	ReqLineAdsPaging
	GroupID string
}

func (r *ReqGetLinkRequestsDto) GetPath() string {
	if r == nil || r.GroupID == "" {
		return ""
	}

	return fmt.Sprintf("%s/%s/%s", GROUPS, r.GroupID, LINK_REQUEST)
}

func (r *ReqGetLinkRequestsDto) GetParameters() LineAdsRequestParameters {
	if r == nil || r.Page == nil || r.Size == nil {
		return nil
	}

	params := make(LineAdsRequestParameters)
	params.Add("page", fmt.Sprintf("%d", *r.Page))
	params.Add("size", fmt.Sprintf("%d", *r.Size))

	return params
}

type ResGetLinkRequestsDto struct {
	Paging ResLineAdsPaging `json:"paging"`
	Datas  []LinkRequest    `json:"datas"`
}
