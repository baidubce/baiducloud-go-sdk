package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbCollectionSpaceTrendResponse struct {
	bce.BaseResponse
	Result *TrendResultBase `json:"result,omitempty"`
}
