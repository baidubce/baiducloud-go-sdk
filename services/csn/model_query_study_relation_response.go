package csn

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryStudyRelationResponse struct {
	bce.BaseResponse
	Propagations []*CsnRtPropagation `json:"propagations,omitempty"`
}
