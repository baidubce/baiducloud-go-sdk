package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheStatusOfTheElasticNetworkCardResponse struct {
	bce.BaseResponse
	Status *string `json:"status,omitempty"`
}
