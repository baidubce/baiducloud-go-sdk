package cpromsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cprom"
)

func UpdateAlert() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cprom.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	labels := make(map[string]string)
	annotations := make(map[string]string)
	updateAlertRequest := &cprom.UpdateAlertRequest{
		AlertingRuleId: util.PtrString(""),
		InstanceId:     util.PtrString(""),
		AlertName:      util.PtrString(""),
		Expr:           util.PtrString(""),
		CpromFor:       util.PtrString(""),
		Description:    util.PtrString(""),
		NotifyRuleId:   util.PtrString(""),
		Severity:       util.PtrString(""),
		Enable:         util.PtrBool(false),
		Labels:         &labels,
		Annotations:    &annotations,
	}
	err = client.UpdateAlert(updateAlertRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
