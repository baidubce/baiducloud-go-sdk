package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeLogStoreViewResponse struct {
	bce.BaseResponse
	Project          *string     `json:"project,omitempty"`
	Name             *string     `json:"name,omitempty"`
	Logstores        []*LogStore `json:"logstores,omitempty"`
	CreatedTimestamp *string     `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp *string     `json:"updatedTimestamp,omitempty"`
}
