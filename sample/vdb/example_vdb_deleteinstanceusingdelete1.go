package vdbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vdb"
)

func Deleteinstanceusingdelete1() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vdb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteinstanceusingdelete1Request := &vdb.Deleteinstanceusingdelete1Request{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err = client.Deleteinstanceusingdelete1(deleteinstanceusingdelete1Request)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
