package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddIpv6OnlyOutboundAndNoInboundPolicyResponse struct {
	bce.BaseResponse
	EgressOnlyRuleId *string `json:"egressOnlyRuleId,omitempty"`
}
