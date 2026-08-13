package cfwsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cfw"
)

func DisableCfwProtect() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cfw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	disableCfwProtectRequest := &cfw.DisableCfwProtectRequest{
		CfwId:      util.PtrString(""),
		InstanceId: util.PtrString(""),
		Role:       util.PtrString(""),
		MemberId:   util.PtrString(""),
	}
	err = client.DisableCfwProtect(disableCfwProtectRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
