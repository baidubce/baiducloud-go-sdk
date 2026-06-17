package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AsyncSearchResponse struct {
	bce.BaseResponse
	StartTimeInMillis      *int32    `json:"start_time_in_millis,omitempty"`
	ExpirationTimeInMillis *int32    `json:"expiration_time_in_millis,omitempty"`
	Response               *Response `json:"response,omitempty"`
	BlsError               *Error    `json:"error,omitempty"`
	Status                 *int32    `json:"status,omitempty"`
}
