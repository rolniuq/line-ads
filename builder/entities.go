package lineads

type LinkRequestStatus string

const (
	WAITING_APPROVAL LinkRequestStatus = "WAITING_APPROVAL"
)

type LinkRequestTargetType string

const (
	ADACCOUNT LinkRequestTargetType = "ADACCOUNT"
)

type LinkRequest struct {
	ID                  int                   `json:"id"`
	SourceGroupID       string                `json:"sourceGroupId"`
	SourceGroupName     string                `json:"sourceGroupName"`
	TargetAdaccountID   string                `json:"targetAdaccountId"`
	TargetAdaccountName string                `json:"targetAdaccountName"`
	Status              LinkRequestStatus     `json:"status"`
	TargetType          LinkRequestTargetType `json:"targetType"`
	CreatedDate         string                `json:"createdDate"`
	ModifiedDate        string                `json:"modifiedDate"`
	RemovedDate         string                `json:"removedDate"`
}
