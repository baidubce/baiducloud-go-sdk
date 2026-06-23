package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type BindClusterResponse struct {
	bce.BaseResponse
	BindingStatus *string `json:"bindingStatus,omitempty"`
}
