package ccesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cce"
)

func CreateNodeGroupV2() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cce.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	InstanceTemplate := &cce.InstanceTemplate{
		MachineType:               util.PtrString(""),
		InstanceType:              util.PtrString(""),
		InstanceName:              util.PtrString(""),
		VpcConfig:                 nil,
		InstanceResource:          nil,
		CheckGPUDriver:            util.PtrBool(false),
		ImageID:                   util.PtrString(""),
		UserData:                  nil,
		InstanceOS:                nil,
		ScaleDownDisabled:         util.PtrBool(false),
		IsOpenHostnameDomain:      util.PtrBool(false),
		NeedEIP:                   util.PtrBool(false),
		EipOption:                 nil,
		IamRole:                   nil,
		DeployCustomConfig:        nil,
		RuntimeType:               util.PtrString(""),
		RuntimeVersion:            util.PtrString(""),
		DeploySetIDs:              []*string{},
		Labels:                    nil,
		Annotations:               nil,
		Tags:                      []*string{},
		Taints:                    []*string{},
		RelationTag:               util.PtrBool(false),
		InstancePreChargingOption: nil,
	}
	ClusterAutoscalerSpec := &cce.ClusterAutoscalerSpec{
		Enabled:              util.PtrBool(false),
		MinReplicas:          util.PtrInt32(int32(0)),
		MaxReplicas:          util.PtrInt32(int32(0)),
		ScalingGroupPriority: util.PtrInt32(int32(0)),
	}
	createNodeGroupV2Request := &cce.CreateNodeGroupV2Request{
		ClusterID:             util.PtrString(""),
		InstanceGroupName:     util.PtrString(""),
		ClusterRole:           util.PtrString(""),
		ShrinkPolicy:          util.PtrString(""),
		UpdatePolicy:          util.PtrString(""),
		CleanPolicy:           util.PtrString(""),
		InstanceTemplate:      InstanceTemplate,
		Replicas:              util.PtrInt32(int32(0)),
		ClusterAutoscalerSpec: ClusterAutoscalerSpec,
	}
	result, err := client.CreateNodeGroupV2(createNodeGroupV2Request)
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
