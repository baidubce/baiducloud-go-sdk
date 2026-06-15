package bcm

type DescribeReceiversRequest struct {
	BcmType  *string `json:"type,omitempty"`
	Name     *string `json:"name,omitempty"`
	PageNo   *int32  `json:"pageNo,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty"`
}
