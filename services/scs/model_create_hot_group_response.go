package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateHotGroupResponse struct {
	bce.BaseResponse
	GroupId *string `json:"groupId,omitempty"`
}
