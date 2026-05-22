package cfw

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetCfwResponse struct {
	bce.BaseResponse
	CfwId           *string    `json:"cfwId,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Description     *string    `json:"description,omitempty"`
	CreatedTime     *string    `json:"createdTime,omitempty"`
	BindInstanceNum *int32     `json:"bindInstanceNum,omitempty"`
	CfwType         *int32     `json:"type,omitempty"`
	Border          *int32     `json:"border,omitempty"`
	CfwRules        []*CfwRule `json:"cfwRules,omitempty"`
}
