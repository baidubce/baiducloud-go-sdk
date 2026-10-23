package pfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func UpdateL3AccessRule() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateL3AccessRuleRequest := &pfs.UpdateL3AccessRuleRequest{
		Action:          util.PtrString(""),
		AccessGroupName: util.PtrString(""),
		ZoneName:        util.PtrString(""),
		AccessRuleId:    util.PtrInt32(int32(0)),
		Ip:              util.PtrString(""),
		Mask:            util.PtrInt32(int32(0)),
		WriteAccess:     util.PtrString(""),
		UserAccess:      util.PtrString(""),
		Priority:        util.PtrInt32(int32(0)),
	}
	result, err := client.UpdateL3AccessRule(updateL3AccessRuleRequest)
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
