package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateEnterpriseSecurityGroupResponse struct {
	bce.BaseResponse
	EnterpriseSecurityGroupId *string `json:"enterpriseSecurityGroupId,omitempty"`
}
