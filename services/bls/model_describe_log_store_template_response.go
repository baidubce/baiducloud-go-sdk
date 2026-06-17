package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLogStoreTemplateResponse struct {
	bce.BaseResponse
	Success          *bool     `json:"success,omitempty"`
	Code             *string   `json:"code,omitempty"`
	Message          *string   `json:"message,omitempty"`
	Name             *string   `json:"name,omitempty"`
	ProjectPatterns  []*string `json:"projectPatterns,omitempty"`
	LogstorePatterns []*string `json:"logstorePatterns,omitempty"`
	Priority         *int32    `json:"priority,omitempty"`
	CreatedTimestamp *string   `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp *string   `json:"updatedTimestamp,omitempty"`
	Template         *Template `json:"template,omitempty"`
}
