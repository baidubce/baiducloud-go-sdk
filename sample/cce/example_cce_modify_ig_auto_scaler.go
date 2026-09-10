package ccesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cce"
)

func ModifyIGAutoScaler() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cce.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	modifyIGAutoScalerRequest := &cce.ModifyIGAutoScalerRequest{
		ClusterID:            util.PtrString(""),
		InstanceGroupID:      util.PtrString(""),
		Enabled:              util.PtrBool(false),
		MinReplicas:          util.PtrInt32(int32(0)),
		MaxReplicas:          util.PtrInt32(int32(0)),
		ScalingGroupPriority: util.PtrInt32(int32(0)),
	}
	result, err := client.ModifyIGAutoScaler(modifyIGAutoScalerRequest)
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
