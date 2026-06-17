package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAlarmPoliciesResponse struct {
	bce.BaseResponse
	Success    *bool          `json:"success,omitempty"`
	Code       *string        `json:"code,omitempty"`
	Message    *string        `json:"message,omitempty"`
	Policies   []*AlarmPolicy `json:"policies,omitempty"`
	PageNo     *int32         `json:"pageNo,omitempty"`
	PageSize   *int32         `json:"pageSize,omitempty"`
	TotalCount *int32         `json:"totalCount,omitempty"`
}
