package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateElasticNetworkCardResponse struct {
	bce.BaseResponse
	EniId *string `json:"eniId,omitempty"`
}
