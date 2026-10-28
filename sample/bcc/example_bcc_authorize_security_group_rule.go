package bccsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func AuthorizeSecurityGroupRule() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := bcc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Rule := &bcc.SecurityGroupRuleModel{
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
	authorizeSecurityGroupRuleRequest := &bcc.AuthorizeSecurityGroupRuleRequest{
		SecurityGroupId: util.PtrString(""),
		SgVersion:       util.PtrInt32(int32(0)),
		Rule:            Rule,
	}
	err = client.AuthorizeSecurityGroupRule(authorizeSecurityGroupRuleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
