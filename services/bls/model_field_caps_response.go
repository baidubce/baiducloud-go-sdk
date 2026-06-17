package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type FieldCapsResponse struct {
	bce.BaseResponse
	Indices  []*string                    `json:"indices,omitempty"`
	Fields   *map[string]map[string]Field `json:"fields,omitempty"`
	BlsError *Error                       `json:"error,omitempty"`
	Status   *int32                       `json:"status,omitempty"`
}
