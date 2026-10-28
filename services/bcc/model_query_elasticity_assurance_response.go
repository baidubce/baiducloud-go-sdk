package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryElasticityAssuranceResponse struct {
	bce.BaseResponse
	Marker               *string                     `json:"marker,omitempty"`
	MaxKeys              *int32                      `json:"maxKeys,omitempty"`
	NextMarker           *string                     `json:"nextMarker,omitempty"`
	IsTruncated          *bool                       `json:"isTruncated,omitempty"`
	ElasticityAssurances []*ElasticityAssuranceModel `json:"elasticityAssurances,omitempty"`
}
