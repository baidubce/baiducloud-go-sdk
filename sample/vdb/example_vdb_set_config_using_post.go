package vdbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vdb"
)

func SetConfigUsingPOST() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vdb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	setConfigUsingPOSTRequest := &vdb.SetConfigUsingPOSTRequest{
		InstanceId:        util.PtrString(""),
		EngineType:        util.PtrString(""),
		AutoBackupConfig:  util.PtrString(""),
		AutoBackupEnabled: util.PtrBool(false),
		IsEncrypt:         util.PtrString(""),
	}
	err = client.SetConfigUsingPOST(setConfigUsingPOSTRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
