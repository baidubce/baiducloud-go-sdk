package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func UpdateSecurityGroupRules() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateSecurityGroupRulesRequest := &vpc.UpdateSecurityGroupRulesRequest{
		ClientToken:         util.PtrString(""),
		SgVersion:           util.PtrInt64(int64(0)),
		SecurityGroupRuleId: util.PtrString(""),
		Remark:              util.PtrString(""),
		PortRange:           util.PtrString(""),
		SourceIp:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		DestIp:              util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		Protocol:            util.PtrString(""),
	}
	err = client.UpdateSecurityGroupRules(updateSecurityGroupRulesRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
