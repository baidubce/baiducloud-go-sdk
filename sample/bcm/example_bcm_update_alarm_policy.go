package bcmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcm"
)

func UpdateAlarmPolicy() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Target := &bcm.AlarmTarget{
		BcmType:             util.PtrString(""),
		Instances:           []*bcm.TargetInstance{},
		Region:              util.PtrString(""),
		Tags:                []*bcm.Dimension{},
		InstanceGroups:      []*string{},
		IncludingDimensions: []*string{},
		ExcludingDimensions: []*string{},
	}
	updateAlarmPolicyRequest := &bcm.UpdateAlarmPolicyRequest{
		Id:                         util.PtrString(""),
		State:                      util.PtrString(""),
		Name:                       util.PtrString(""),
		Scope:                      util.PtrString(""),
		ResourceType:               util.PtrString(""),
		Target:                     Target,
		Rules:                      []*bcm.AlarmRule{},
		PendingCount:               util.PtrInt32(int32(0)),
		OnMissingData:              util.PtrString(""),
		NoDataNotifyPendingMinutes: util.PtrInt32(int32(0)),
		BcmType:                    util.PtrString(""),
		Level:                      util.PtrString(""),
		Actions:                    []*bcm.PolicyAction{},
		NotifyEnabled:              util.PtrBool(false),
		Callbacks:                  []*bcm.Callback{},
		RenotifyCount:              util.PtrInt32(int32(0)),
		RenotifyIntervalMinutes:    util.PtrInt32(int32(0)),
		NotifyMergeWindowSeconds:   util.PtrInt32(int32(0)),
	}
	result, err := client.UpdateAlarmPolicy(updateAlarmPolicyRequest)
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
