package iam

type Page struct {
	OrderBy    *string   `json:"orderBy,omitempty"`
	Order      *string   `json:"order,omitempty"`
	PageNo     *int32    `json:"pageNo,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty"`
	TotalCount *int32    `json:"totalCount,omitempty"`
	Result     []*APIKey `json:"result,omitempty"`
}
