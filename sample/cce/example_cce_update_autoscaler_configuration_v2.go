package ccesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cce"
)

func UpdateAutoscalerConfigurationV2() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cce.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	CustomConfigs := make(map[string]string)
	updateAutoscalerConfigurationV2Request := &cce.UpdateAutoscalerConfigurationV2Request{
		ClusterID:                        util.PtrString(""),
		Expander:                         util.PtrString(""),
		InstanceGroups:                   []*map[string]interface{}{},
		KubeVersion:                      util.PtrString(""),
		MaxEmptyBulkDelete:               util.PtrInt32(int32(0)),
		ScaleDownDelayAfterAdd:           util.PtrInt32(int32(0)),
		ScaleDownEnabled:                 util.PtrBool(false),
		ScaleDownGPUUtilizationThreshold: util.PtrInt32(int32(0)),
		ScaleDownUnneededTime:            util.PtrInt32(int32(0)),
		ScaleDownUtilizationThreshold:    util.PtrInt32(int32(0)),
		SkipNodesWithLocalStorage:        util.PtrBool(false),
		SkipNodesWithSystemPods:          util.PtrBool(false),
		CustomConfigs:                    nil,
	}
	result, err := client.UpdateAutoscalerConfigurationV2(updateAutoscalerConfigurationV2Request)
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
