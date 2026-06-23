package cpromsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cprom"
)

func CreateInstance() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cprom.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createInstanceRequest := &cprom.CreateInstanceRequest{
		InstanceName:         util.PtrString(""),
		InstanceType:         util.PtrString(""),
		InstanceSpec:         util.PtrString(""),
		RetentionPeriod:      util.PtrString(""),
		NeedGrafana:          util.PtrBool(false),
		GrafanaName:          util.PtrString(""),
		GrafanaAdminPassword: util.PtrString(""),
	}
	err = client.CreateInstance(createInstanceRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
