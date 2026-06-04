package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateResolverRuleResponse struct {
	bce.BaseResponse
	Id *string `json:"id,omitempty"`
}
