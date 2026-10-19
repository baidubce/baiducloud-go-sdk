package rds

import "github.com/baidubce/baiducloud-go-sdk/bce"

type RetrieveInstanceAutoScalingConfigResponse struct {
	bce.BaseResponse
	AutoResizeDisk     *int32 `json:"autoResizeDisk,omitempty"`
	FreeSpaceThreshold *int32 `json:"freeSpaceThreshold,omitempty"`
	DiskMaxLimit       *int32 `json:"diskMaxLimit,omitempty"`
	ExtendStepPercent  *int32 `json:"extendStepPercent,omitempty"`
}
