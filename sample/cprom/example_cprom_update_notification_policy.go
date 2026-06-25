package cpromsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cprom"
)

func UpdateNotificationPolicy() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cprom.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	RepeatNotifyConfig := &cprom.RepeatNotifyConfig{
		Enabled:      util.PtrBool(false),
		IntervalHour: util.PtrInt32(int32(0)),
		IntervalMin:  util.PtrInt32(int32(0)),
		MaxCount:     util.PtrInt32(int32(0)),
		Strategy:     util.PtrString(""),
	}
	updateNotificationPolicyRequest := &cprom.UpdateNotificationPolicyRequest{
		NotifyRuleId:       util.PtrString(""),
		NotifyRuleName:     util.PtrString(""),
		StartTime:          util.PtrString(""),
		EndTime:            util.PtrString(""),
		Channel:            []*string{},
		ReceiverType:       util.PtrString(""),
		Users:              []*cprom.User{},
		UserGroups:         []*cprom.UserGroup{},
		WebhookConfigList:  []*cprom.CallbackConfig{},
		EscalateConfigList: []*cprom.EscalateParam{},
		RepeatNotifyConfig: RepeatNotifyConfig,
	}
	err = client.UpdateNotificationPolicy(updateNotificationPolicyRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
