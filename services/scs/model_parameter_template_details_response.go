package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ParameterTemplateDetailsResponse struct {
	bce.BaseResponse
	TemplateId     *int32        `json:"templateId,omitempty"`
	TemplateShowId *string       `json:"templateShowId,omitempty"`
	TemplateName   *string       `json:"templateName,omitempty"`
	ParameterNum   *int32        `json:"parameterNum,omitempty"`
	ClusterType    *string       `json:"clusterType,omitempty"`
	Engine         *string       `json:"engine,omitempty"`
	EngineVersion  *string       `json:"engineVersion,omitempty"`
	TemplateType   *int32        `json:"templateType,omitempty"`
	NeedReboot     *int32        `json:"needReboot,omitempty"`
	Comment        *string       `json:"comment,omitempty"`
	CreateTime     *string       `json:"createTime,omitempty"`
	UpdateTime     *string       `json:"updateTime,omitempty"`
	Parameters     []*Parameters `json:"parameters,omitempty"`
}
