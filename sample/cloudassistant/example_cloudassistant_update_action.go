package cloudassistantsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cloudassistant"
)

func UpdateAction() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cloudassistant.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Action := &cloudassistant.Action{
		Id:                 util.PtrString(""),
		Ref:                util.PtrString(""),
		CloudassistantType: util.PtrString(""),
		Name:               util.PtrString(""),
		Alias:              util.PtrString(""),
		Description:        util.PtrString(""),
		TimeoutSecond:      util.PtrInt32(int32(0)),
		Command: &cloudassistant.Command{
			CloudassistantType: util.PtrString(""),
			Content:            util.PtrString(""),
			Scope:              util.PtrString(""),
			EnableParameter:    util.PtrBool(false),
			Parameters:         []*cloudassistant.Parameter{},
			User:               util.PtrString(""),
			WorkDir:            util.PtrString(""),
			ExecParams:         nil,
			WaitOnAgentMilli:   util.PtrInt32(int32(0)),
		},
		FileUpload: &cloudassistant.FileUpload{
			Os:            util.PtrString(""),
			Content:       util.PtrString(""),
			Filename:      util.PtrString(""),
			Filepath:      util.PtrString(""),
			BosBucketName: util.PtrString(""),
			BosFilePath:   util.PtrString(""),
			BosEtag:       util.PtrString(""),
			User:          util.PtrString(""),
			Group:         util.PtrString(""),
			Mode:          util.PtrString(""),
			Overwrite:     util.PtrBool(false),
		},
		SupportedInstanceTypes: []*string{},
		CreatedTimestamp:       util.PtrInt64(int64(0)),
		UpdatedTimestamp:       util.PtrInt64(int64(0)),
	}
	Parameters := make(map[string]string)
	TargetSelector := &cloudassistant.TargetSelector{
		InstanceType: util.PtrString(""),
		Tags:         []*cloudassistant.Tag{},
		ImportInstances: &cloudassistant.TargetImport{
			KeywordType: util.PtrString(""),
			Instances:   []*string{},
		},
	}
	updateActionRequest := &cloudassistant.UpdateActionRequest{
		Execution:          util.PtrString(""),
		Action:             Action,
		Parameters:         nil,
		TargetSelectorType: util.PtrString(""),
		Targets:            []*cloudassistant.Target{},
		TargetSelector:     TargetSelector,
	}
	result, err := client.UpdateAction(updateActionRequest)
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
