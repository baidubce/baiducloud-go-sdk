package ccrsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ccr"
)

func UpdateImageMigrationRule() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := ccr.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	SrcRegistry := &ccr.ReplicationRegistryRequest{
		Id: util.PtrInt32(int32(0)),
	}
	Trigger := &ccr.ReplicationTriggerRequest{
		CcrType: util.PtrString(""),
	}
	updateImageMigrationRuleRequest := &ccr.UpdateImageMigrationRuleRequest{
		InstanceId:      util.PtrString(""),
		PolicyId:        util.PtrString(""),
		Description:     util.PtrString(""),
		DestProjectName: util.PtrString(""),
		Filters:         []*ccr.ReplicationFilterRequest{},
		Name:            util.PtrString(""),
		Override:        util.PtrBool(false),
		SrcRegistry:     SrcRegistry,
		Trigger:         Trigger,
	}
	err = client.UpdateImageMigrationRule(updateImageMigrationRuleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
