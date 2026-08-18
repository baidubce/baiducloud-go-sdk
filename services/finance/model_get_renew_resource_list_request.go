package finance

type GetRenewResourceListRequest struct {
	QueryAccountId     *string   `json:"queryAccountId,omitempty"`
	ServiceType        *string   `json:"serviceType,omitempty"`
	Region             *string   `json:"region,omitempty"`
	ExpiredDays        *int32    `json:"expiredDays,omitempty"`
	ShortOrInstanceIds []*string `json:"shortOrInstanceIds,omitempty"`
	PageNo             *int32    `json:"pageNo,omitempty"`
	PageSize           *int32    `json:"pageSize,omitempty"`
}
