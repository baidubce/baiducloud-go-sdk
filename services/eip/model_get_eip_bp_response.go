package eip

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetEipBpResponse struct {
	bce.BaseResponse
	Name                    *string   `json:"name,omitempty"`
	Id                      *string   `json:"id,omitempty"`
	BindType                *string   `json:"bindType,omitempty"`
	BandwidthInMbps         *int32    `json:"bandwidthInMbps,omitempty"`
	InstanceId              *string   `json:"instanceId,omitempty"`
	Eips                    []*string `json:"eips,omitempty"`
	InstanceBandwidthInMbps *int32    `json:"instanceBandwidthInMbps,omitempty"`
	CreateTime              *string   `json:"createTime,omitempty"`
	AutoReleaseTime         *string   `json:"autoReleaseTime,omitempty"`
	EipType                 *string   `json:"type,omitempty"`
	Region                  *string   `json:"region,omitempty"`
}
