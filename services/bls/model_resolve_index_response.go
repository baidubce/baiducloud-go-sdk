package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type ResolveIndexResponse struct {
	bce.BaseResponse
	DataStreams *DataStreams `json:"data_streams,omitempty"`
	BlsError    *Error       `json:"error,omitempty"`
	Status      *int32       `json:"status,omitempty"`
}
