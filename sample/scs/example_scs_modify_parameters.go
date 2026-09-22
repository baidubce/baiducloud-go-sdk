package scssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func ModifyParameters() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Parameter := &scs.Parameter{
		ScsDefault:   util.PtrString(""),
		ForceRestart: util.PtrInt32(int32(0)),
		Name:         util.PtrString(""),
		Value:        util.PtrString(""),
	}
	modifyParametersRequest := &scs.ModifyParametersRequest{
		InstanceId: util.PtrString(""),
		Parameter:  Parameter,
	}
	err = client.ModifyParameters(modifyParametersRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
