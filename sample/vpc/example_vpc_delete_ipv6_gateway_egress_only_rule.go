package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func DeleteIpv6GatewayEgressOnlyRule() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := vpc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteIpv6GatewayEgressOnlyRuleRequest := &vpc.DeleteIpv6GatewayEgressOnlyRuleRequest{
		GatewayId:        util.PtrString(""),
		EgressOnlyRuleId: util.PtrString(""),
		ClientToken:      util.PtrString(""),
	}
	err = client.DeleteIpv6GatewayEgressOnlyRule(deleteIpv6GatewayEgressOnlyRuleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
