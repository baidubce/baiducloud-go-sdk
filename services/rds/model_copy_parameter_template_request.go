package rds

type CopyParameterTemplateRequest struct {
	Name       *string `json:"name,omitempty"`
	Desc       *string `json:"desc,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
}
