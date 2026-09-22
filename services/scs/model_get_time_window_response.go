package scs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetTimeWindowResponse struct {
	bce.BaseResponse
	MaintainTime *MaintainTime `json:"maintainTime,omitempty"`
}
