package apmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/apm"
)

func ApmUpdateAlarmPolicy() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := apm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Target := &apm.AlarmTarget{
		ApmType:  util.PtrString(""),
		Tags:     []*apm.Tag{},
		Services: []*string{},
	}
	Rule := &apm.AlarmRule{
		Operator:        util.PtrString(""),
		Rules:           []*apm.AlarmRule{},
		Metric:          util.PtrString(""),
		WindowInSeconds: util.PtrInt32(int32(0)),
		Aggregate:       util.PtrString(""),
		DisplayValue:    util.PtrFloat32(float32(0)),
		DisplayUnit:     util.PtrString(""),
	}
	apmUpdateAlarmPolicyRequest := &apm.ApmUpdateAlarmPolicyRequest{
		Id:                                   util.PtrString(""),
		Name:                                 util.PtrString(""),
		Target:                               Target,
		MetricKind:                           util.PtrString(""),
		Rule:                                 Rule,
		Filters:                              []*apm.AlarmFilter{},
		PendingCount:                         util.PtrInt32(int32(0)),
		RenotifyIntervalInMinutes:            util.PtrInt32(int32(0)),
		RenotifyCount:                        util.PtrInt32(int32(0)),
		NotifyRecovery:                       util.PtrBool(false),
		OnMissingData:                        util.PtrString(""),
		NoDataNotifyPendingIntervalInMinutes: util.PtrInt32(int32(0)),
		Level:                                util.PtrString(""),
		Actions:                              []*apm.AlarmAction{},
	}
	result, err := client.ApmUpdateAlarmPolicy(apmUpdateAlarmPolicyRequest)
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
