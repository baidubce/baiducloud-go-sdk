package oos

type GetTemplateListV2Request struct {
	Locale                *string `json:"-"`
	Namespace             *string `json:"namespace,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Id                    *string `json:"id,omitempty"`
	OosType               *string `json:"type,omitempty"`
	Sort                  *string `json:"sort,omitempty"`
	Ascending             *bool   `json:"ascending,omitempty"`
	PageNo                *int32  `json:"pageNo,omitempty"`
	PageSize              *int32  `json:"pageSize,omitempty"`
	SupportedInstanceType *string `json:"supportedInstanceType,omitempty"`
}
