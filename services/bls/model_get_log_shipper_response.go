package bls

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetLogShipperResponse struct {
	bce.BaseResponse
	Success        *bool       `json:"success,omitempty"`
	Code           *string     `json:"code,omitempty"`
	Status         *string     `json:"status,omitempty"`
	LogShipperName *string     `json:"logShipperName,omitempty"`
	Project        *string     `json:"project,omitempty"`
	LogStoreName   *string     `json:"logStoreName,omitempty"`
	StartTime      *string     `json:"startTime,omitempty"`
	DestType       *string     `json:"destType,omitempty"`
	DestConfig     *DestConfig `json:"destConfig,omitempty"`
}
