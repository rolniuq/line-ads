package lineads

type ReqCreateGroupDto struct {
	Id              *string
	Name            *string
	ParentGroupId   *string
	ParentGroupName *string
	Depth           *int64
}

type ResCreateGroupDto struct {
}

type ReqCreateLinkRequestDto struct {
	SourceGroupID     string
	TargetAdAccountId string `json:"targetAdaccountId"`
}

type ResCreateLinkRequestDto struct {
}

type ReqGetLinkRequestsDto struct {
	GroupID string
}

type ResGetLinkRequestsDto struct {
	Datas []LinkRequest `json:"datas"`
}
