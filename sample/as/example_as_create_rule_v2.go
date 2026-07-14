package assample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/as"
)

func CreateRuleV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := as.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	AsAlarmRule := &as.AsAlarmRule{
		Id:    util.PtrInt32(int32(0)),
		Scope: util.PtrString(""),
		MonitorObject: &as.MonitorObject{
			AsType:    util.PtrString(""),
			Names:     []*string{},
			Resources: []*as.PolicyResource{},
			TypeName:  util.PtrString(""),
		},
		Rules:               [][]AlarmRule{},
		AlarmName:           util.PtrString(""),
		AliasName:           util.PtrString(""),
		InsufficientCycle:   util.PtrInt32(int32(0)),
		PolicyEnabled:       util.PtrBool(false),
		RuleContents:        []*string{},
		RuleContentsEn:      []*string{},
		Source:              util.PtrString(""),
		ComponentType:       util.PtrString(""),
		AlarmActions:        []*string{},
		OkActions:           []*string{},
		InsufficientActions: []*string{},
	}
	createRuleV2Request := &as.CreateRuleV2Request{
		RuleName:        util.PtrString(""),
		GroupId:         util.PtrString(""),
		State:           util.PtrString(""),
		AsType:          util.PtrString(""),
		ActionType:      util.PtrString(""),
		ActionNum:       util.PtrInt32(int32(0)),
		CronTime:        util.PtrString(""),
		CooldownInSec:   util.PtrInt32(int32(0)),
		PeriodType:      util.PtrString(""),
		PeriodValue:     util.PtrInt32(int32(0)),
		PeriodStartTime: util.PtrString(""),
		PeriodEndTime:   util.PtrString(""),
		AsAlarmRule:     AsAlarmRule,
	}
	result, err := client.CreateRuleV2(createRuleV2Request)
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
