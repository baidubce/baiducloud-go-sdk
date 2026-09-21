package vdbsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vdb"
)

func CreateInstanceUsingPOST() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vdb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	InstanceParam := &vdb.InstanceParam{
		AvailabilityZone:     util.PtrString(""),
		AzInfos:              []*vdb.AzInfo{},
		CloneDataAppBackupId: util.PtrString(""),
		CloneDataAppId:       util.PtrString(""),
		Components:           []*vdb.MilvusComponent{},
		DataNodeNum:          util.PtrInt32(int32(0)),
		DiskFlavor:           util.PtrInt32(int32(0)),
		DiskType:             util.PtrString(""),
		EnableEmbedding:      util.PtrBool(false),
		EnableEncryption:     util.PtrBool(false),
		EngineVersion:        util.PtrString(""),
		From:                 util.PtrString(""),
		InstanceName:         util.PtrString(""),
		InstanceNum:          util.PtrInt32(int32(0)),
		InstanceType:         util.PtrString(""),
		MasterNodeSpec:       util.PtrString(""),
		MasterNum:            util.PtrInt32(int32(0)),
		NodeSpec:             util.PtrString(""),
		NodeType:             util.PtrString(""),
		OrderId:              util.PtrString(""),
		Password:             util.PtrString(""),
		Port:                 util.PtrInt32(int32(0)),
		ProxyNodeSpec:        util.PtrString(""),
		ProxyNum:             util.PtrInt32(int32(0)),
		ReqSource:            util.PtrString(""),
		SubnetId:             util.PtrString(""),
		SwitchEntrance:       util.PtrString(""),
		VpcId:                util.PtrString(""),
	}
	createInstanceUsingPOSTRequest := &vdb.CreateInstanceUsingPOSTRequest{
		EngineType:        util.PtrString(""),
		AutoRenew:         util.PtrBool(false),
		AutoRenewTime:     util.PtrInt32(int32(0)),
		AutoRenewTimeUnit: util.PtrString(""),
		Components:        []*vdb.MilvusComponent{},
		Duration:          util.PtrInt32(int32(0)),
		Env:               util.PtrString(""),
		InstanceParam:     InstanceParam,
		ProductType:       util.PtrString(""),
		TimeUnit:          util.PtrString(""),
	}
	result, err := client.CreateInstanceUsingPOST(createInstanceUsingPOSTRequest)
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
