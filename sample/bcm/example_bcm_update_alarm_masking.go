package bcmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcm"
)

func UpdateAlarmMasking() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAlarmMaskingRequest := &bcm.UpdateAlarmMaskingRequest{
		Id:                  util.PtrString(""),
		State:               util.PtrString(""),
		Name:                util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		PolicyId:            util.PtrString(""),
		Instances:           []*bcm.TargetInstance{},
		Region:              util.PtrString(""),
		MetricNames:         []*string{},
		PeriodType:          util.PtrString(""),
		BeginTime:           util.PtrString(""),
		EndTime:             util.PtrString(""),
		Tz:                  util.PtrString(""),
		DailyBeginTimestamp: util.PtrInt32(int32(0)),
		DailyEndTimestamp:   util.PtrInt32(int32(0)),
	}
	result, err := client.UpdateAlarmMasking(updateAlarmMaskingRequest)
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
