package cloudassistant

type ActionListRequest struct {
	Locale    *string       `json:"-"`
	PageNo    *int32        `json:"pageNo,omitempty"`
	PageSize  *int32        `json:"pageSize,omitempty"`
	Sort      *string       `json:"sort,omitempty"`
	Ascending *bool         `json:"ascending,omitempty"`
	Action    *ActionFilter `json:"action,omitempty"`
}
