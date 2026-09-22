package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ViewSecurityGroupResponse struct {
	bce.BaseResponse
	Groups      []*Group `json:"groups,omitempty"`
	ActiveRules []*Rule  `json:"activeRules,omitempty"`
}
