package cpromsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cprom"
)

func CreateServiceMonitor() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cprom.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Metadata := &cprom.ObjectMeta{
		Name:      util.PtrString(""),
		Namespace: util.PtrString(""),
	}
	Spec := &cprom.ServiceMonitorSpec{
		Endpoints: []*cprom.ServiceMonitorEndpoint{},
		NamespaceSelector: &cprom.NamespaceSelector{
			MatchNames: []*string{},
		},
		Selector: &cprom.LabelSelector{
			MatchLabels: nil,
		},
	}
	createServiceMonitorRequest := &cprom.CreateServiceMonitorRequest{
		InstanceId: util.PtrString(""),
		AgentId:    util.PtrString(""),
		ApiVersion: util.PtrString(""),
		Kind:       util.PtrString(""),
		Metadata:   Metadata,
		Spec:       Spec,
	}
	result, err := client.CreateServiceMonitor(createServiceMonitorRequest)
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
