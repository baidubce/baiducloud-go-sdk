package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func RevokeRegularSecurityGroupRulesV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := vpc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Rule := &vpc.SecurityGroupRuleModel{
		Remark:              util.PtrString(""),
		Direction:           util.PtrString(""),
		Ethertype:           util.PtrString(""),
		PortRange:           util.PtrString(""),
		Protocol:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		SourceIp:            util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		DestIp:              util.PtrString(""),
		SecurityGroupId:     util.PtrString(""),
		SecurityGroupRuleId: util.PtrString(""),
		CreatedTime:         util.PtrString(""),
		UpdatedTime:         util.PtrString(""),
	}
	revokeRegularSecurityGroupRulesV2Request := &vpc.RevokeRegularSecurityGroupRulesV2Request{
		SecurityGroupId: util.PtrString(""),
		ClientToken:     util.PtrString(""),
		SgVersion:       util.PtrInt64(int64(0)),
		Rule:            Rule,
	}
	err = client.RevokeRegularSecurityGroupRulesV2(revokeRegularSecurityGroupRulesV2Request)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
