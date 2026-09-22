package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateParameterTemplateResponse struct {
	bce.BaseResponse
	TemplateId     *int32  `json:"templateId,omitempty"`
	TemplateShowId *string `json:"templateShowId,omitempty"`
}
