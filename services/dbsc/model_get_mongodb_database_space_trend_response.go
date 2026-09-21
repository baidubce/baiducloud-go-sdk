package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMongodbDatabaseSpaceTrendResponse struct {
	bce.BaseResponse
	Result *TrendResultBase `json:"result,omitempty"`
}
