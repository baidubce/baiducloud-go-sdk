package rds

type GetParameterTemplateListRequest struct {
	PageNo    *int32  `json:"-"`
	PageSize  *int32  `json:"-"`
	Type      *string `json:"-"`
	DbType    *string `json:"-"`
	DbVersion *string `json:"-"`
}
