package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateResolverResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
