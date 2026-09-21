package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlSpaceSummaryResponse struct {
	bce.BaseResponse
	TotalDiskSize   *map[string]interface{} `json:"totalDiskSize,omitempty"`
	UsedSpace       *float32                `json:"usedSpace,omitempty"`
	AvailableSpace  *float32                `json:"availableSpace,omitempty"`
	AvailableDays   *float32                `json:"availableDays,omitempty"`
	DayGrowSpaceAvg *float32                `json:"dayGrowSpaceAvg,omitempty"`
}
