package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateAuthGroupResponse struct {
	bce.BaseResponse
	AuthGroupId *string `json:"authGroupId,omitempty"`
}
