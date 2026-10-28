package bcc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RenewElasticityAssuranceResponse struct {
	bce.BaseResponse
	OrderId                *string   `json:"orderId,omitempty"`
	ElasticityAssuranceIds []*string `json:"elasticityAssuranceIds,omitempty"`
}
