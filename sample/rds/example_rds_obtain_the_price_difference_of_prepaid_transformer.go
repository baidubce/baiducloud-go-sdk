package rdssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func ObtainThePriceDifferenceOfPrepaidTransformer() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	obtainThePriceDifferenceOfPrepaidTransformerRequest := &rds.ObtainThePriceDifferenceOfPrepaidTransformerRequest{
		InstanceId:           util.PtrString(""),
		CpuCount:             util.PtrInt32(int32(0)),
		AllocatedMemoryInGB:  util.PtrInt32(int32(0)),
		AllocatedStorageInGB: util.PtrInt32(int32(0)),
		DiskIoType:           util.PtrString(""),
		Category:             util.PtrString(""),
	}
	result, err := client.ObtainThePriceDifferenceOfPrepaidTransformer(obtainThePriceDifferenceOfPrepaidTransformerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
}
