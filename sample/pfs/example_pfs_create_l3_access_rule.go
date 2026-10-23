package pfssample

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func CreateL3AccessRule() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createL3AccessRuleRequest := &pfs.CreateL3AccessRuleRequest{
		AccessGroupName: util.PtrString(""),
		ZoneName:        util.PtrString(""),
		AccessRules: []*pfs.AccessRuleCreateInfo{
			{
				Ip:          util.PtrString(""),
				Mask:        util.PtrInt32(int32(0)),
				WriteAccess: util.PtrString(""),
				UserAccess:  util.PtrString(""),
				Priority:    util.PtrInt32(int32(0)),
			},
		},
	}
	result, err := client.CreateL3AccessRule(createL3AccessRuleRequest)
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
