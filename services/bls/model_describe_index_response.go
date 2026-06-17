package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeIndexResponse struct {
	bce.BaseResponse
	Success        *bool                   `json:"success,omitempty"`
	Code           *string                 `json:"code,omitempty"`
	Fulltext       *bool                   `json:"fulltext,omitempty"`
	Fields         *map[string]*IndexField `json:"fields,omitempty"`
	CaseSensitive  *bool                   `json:"caseSensitive,omitempty"`
	IncludeChinese *bool                   `json:"includeChinese,omitempty"`
	Separators     *string                 `json:"separators,omitempty"`
}
