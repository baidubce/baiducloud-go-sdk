package vdbsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vdb"
)

func ResizeInstanceUsingPOST() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vdb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	resizeInstanceUsingPOSTRequest := &vdb.ResizeInstanceUsingPOSTRequest{
		EngineType:     util.PtrString(""),
		Components:     []*vdb.MilvusComponent{},
		DataNodeNum:    util.PtrInt32(int32(0)),
		DiskFlavor:     util.PtrInt32(int32(0)),
		DiskType:       util.PtrString(""),
		Env:            util.PtrString(""),
		InstanceId:     util.PtrString(""),
		MasterNodeSpec: util.PtrString(""),
		MasterNum:      util.PtrInt32(int32(0)),
		NodeSpec:       util.PtrString(""),
		NodeType:       util.PtrString(""),
		OrderId:        util.PtrString(""),
		ProxyNodeSpec:  util.PtrString(""),
		ProxyNum:       util.PtrInt32(int32(0)),
	}
	result, err := client.ResizeInstanceUsingPOST(resizeInstanceUsingPOSTRequest)
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
