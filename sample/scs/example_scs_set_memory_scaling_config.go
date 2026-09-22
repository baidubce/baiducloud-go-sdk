package scssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func SetMemoryScalingConfig() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	MemSpec := &scs.MemSpec{
		MemUsageUpperThreshold:        util.PtrInt32(int32(0)),
		MemUsageDownThreshold:         util.PtrInt32(int32(0)),
		MaxNodeType:                   util.PtrString(""),
		MinNodeType:                   util.PtrString(""),
		ObservationWindowSizeForUpper: util.PtrString(""),
		ObservationWindowSizeForDown:  util.PtrString(""),
	}
	setMemoryScalingConfigRequest := &scs.SetMemoryScalingConfigRequest{
		InstanceId: util.PtrString(""),
		MemSpec:    MemSpec,
	}
	err = client.SetMemoryScalingConfig(setMemoryScalingConfigRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
