package scssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func CreateAnInstance() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &scs.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &scs.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createAnInstanceRequest := &scs.CreateAnInstanceRequest{
		ClientToken:       util.PtrString(""),
		Billing:           Billing,
		InstanceName:      util.PtrString(""),
		NodeType:          util.PtrString(""),
		Port:              util.PtrInt32(int32(0)),
		Engine:            util.PtrInt32(int32(0)),
		EngineVersion:     util.PtrString(""),
		StoreType:         util.PtrInt32(int32(0)),
		EnableReadOnly:    util.PtrInt32(int32(0)),
		PurchaseCount:     util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		ProxyNum:          util.PtrInt32(int32(0)),
		ClusterType:       util.PtrString(""),
		DiskFlavor:        util.PtrInt32(int32(0)),
		DiskType:          util.PtrString(""),
		VpcId:             util.PtrString(""),
		ReplicationInfo:   []*scs.ReplicationMap{},
		AutoRenewTimeUnit: util.PtrString(""),
		AutoRenewTime:     util.PtrInt32(int32(0)),
		BgwGroupId:        util.PtrString(""),
		ClientAuth:        util.PtrString(""),
		Tags:              []*scs.Tag{},
		ConfTpl:           util.PtrString(""),
		ResourceGroupId:   util.PtrString(""),
		AutoBackupConfig:  util.PtrString(""),
		DeployIdList:      []*string{},
	}
	result, err := client.CreateAnInstance(createAnInstanceRequest)
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
