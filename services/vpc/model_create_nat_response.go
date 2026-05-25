package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateNatResponse struct {
	bce.BaseResponse
	NatId *string `json:"natId,omitempty"`
}
