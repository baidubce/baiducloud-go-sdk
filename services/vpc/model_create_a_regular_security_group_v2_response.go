package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateARegularSecurityGroupV2Response struct {
	bce.BaseResponse
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}
