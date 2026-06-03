package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DeleteInstanceResponse struct {
	bce.BaseResponse
	Token *string `json:"token,omitempty"`
}
