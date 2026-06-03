package ccr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetImageMigrationExecutionRecordDetailResponse struct {
	bce.BaseResponse
	EndTime    *string `json:"endTime,omitempty"`
	Failed     *int32  `json:"failed,omitempty"`
	Id         *int32  `json:"id,omitempty"`
	InProgress *int32  `json:"inProgress,omitempty"`
	PolicyId   *int32  `json:"policyId,omitempty"`
	StartTime  *string `json:"startTime,omitempty"`
	Status     *string `json:"status,omitempty"`
	StatusText *string `json:"statusText,omitempty"`
	Stopped    *int32  `json:"stopped,omitempty"`
	Succeed    *int32  `json:"succeed,omitempty"`
	Total      *int32  `json:"total,omitempty"`
	Trigger    *string `json:"trigger,omitempty"`
}
