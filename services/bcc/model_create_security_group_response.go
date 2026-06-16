package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateSecurityGroupResponse struct {
	bce.BaseResponse
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
}
