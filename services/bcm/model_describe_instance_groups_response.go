package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeInstanceGroupsResponse struct {
	bce.BaseResponse
	Success        *bool                   `json:"success,omitempty"`
	Code           *string                 `json:"code,omitempty"`
	Message        *string                 `json:"message,omitempty"`
	InstanceGroups []*InstanceGroupSummary `json:"instanceGroups,omitempty"`
	PageNo         *int32                  `json:"pageNo,omitempty"`
	PageSize       *int32                  `json:"pageSize,omitempty"`
	TotalCount     *int32                  `json:"totalCount,omitempty"`
}
