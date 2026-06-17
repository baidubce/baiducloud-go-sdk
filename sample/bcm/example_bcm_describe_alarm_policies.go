package bcmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcm"
)

func DescribeAlarmPolicies() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	describeAlarmPoliciesRequest := &bcm.DescribeAlarmPoliciesRequest{
		PolicyName:      util.PtrString(""),
		PolicyId:        util.PtrString(""),
		Scope:           util.PtrString(""),
		ResourceType:    util.PtrString(""),
		Recursive:       util.PtrBool(false),
		SubResourceType: util.PtrString(""),
		NotifyEnabled:   util.PtrBool(false),
		BcmType:         util.PtrString(""),
		Order:           util.PtrString(""),
		OrderBy:         util.PtrString(""),
		PageNo:          util.PtrInt32(int32(0)),
		PageSize:        util.PtrInt32(int32(0)),
	}
	result, err := client.DescribeAlarmPolicies(describeAlarmPoliciesRequest)
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
