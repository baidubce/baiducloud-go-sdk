package oos

type GetOperatorListV2Request struct {
	Locale    *string      `json:"-"`
	Operator  *interface{} `json:"operator,omitempty"`
	Sort      *string      `json:"sort,omitempty"`
	Ascending *bool        `json:"ascending,omitempty"`
	PageNo    *int32       `json:"pageNo,omitempty"`
	PageSize  *int32       `json:"pageSize,omitempty"`
}
