package cce

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetPackageListV2Response struct {
	bce.BaseResponse
	MachineSpecList []*MachineSpecStatus `json:"machineSpecList,omitempty"`
}
