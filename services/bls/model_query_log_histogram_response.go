package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryLogHistogramResponse struct {
	bce.BaseResponse
	SearchInfo      *SearchInfo      `json:"searchInfo,omitempty"`
	SearchStatistic *SearchStatistic `json:"searchStatistic,omitempty"`
}
