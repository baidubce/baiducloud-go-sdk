package cfwsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cfw"
)

func UpdateCfwRule() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cfw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateCfwRuleRequest := &cfw.UpdateCfwRuleRequest{
		CfwId:         util.PtrString(""),
		CfwRuleId:     util.PtrString(""),
		IpVersion:     util.PtrInt32(int32(0)),
		Priority:      util.PtrInt32(int32(0)),
		Protocol:      util.PtrString(""),
		Direction:     util.PtrString(""),
		SourceAddress: util.PtrString(""),
		DestAddress:   util.PtrString(""),
		SourcePort:    util.PtrString(""),
		DestPort:      util.PtrString(""),
		Action:        util.PtrString(""),
		Description:   util.PtrString(""),
	}
	err = client.UpdateCfwRule(updateCfwRuleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
