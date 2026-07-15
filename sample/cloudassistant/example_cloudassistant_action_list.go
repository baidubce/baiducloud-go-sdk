package cloudassistantsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cloudassistant"
)

func ActionList() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cloudassistant.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Action := &cloudassistant.ActionFilter{
		CloudassistantType: util.PtrString(""),
		Command: &cloudassistant.CommandFilter{
			Scope:              util.PtrString(""),
			Name:               util.PtrString(""),
			CloudassistantType: util.PtrString(""),
		},
		InstanceType: util.PtrString(""),
		Id:           util.PtrString(""),
		Name:         util.PtrString(""),
		FileUpload:   nil,
	}
	actionListRequest := &cloudassistant.ActionListRequest{
		Locale:    util.PtrString(""),
		PageNo:    util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
		Sort:      util.PtrString(""),
		Ascending: util.PtrBool(false),
		Action:    Action,
	}
	result, err := client.ActionList(actionListRequest)
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
