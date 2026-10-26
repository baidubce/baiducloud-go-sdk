package rdssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func CreateReadOnlyInstance() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &rds.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &rds.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createReadOnlyInstanceRequest := &rds.CreateReadOnlyInstanceRequest{
		ClientToken:           util.PtrString(""),
		Billing:               Billing,
		SourceInstanceId:      util.PtrString(""),
		InstanceName:          util.PtrString(""),
		CpuCount:              util.PtrInt32(int32(0)),
		MemoryCapacity:        util.PtrInt32(int32(0)),
		VolumeCapacity:        util.PtrInt32(int32(0)),
		DiskIoType:            util.PtrString(""),
		ZoneNames:             []*string{},
		VpcId:                 util.PtrString(""),
		BcmGroupName:          util.PtrString(""),
		IsDirectPay:           util.PtrBool(false),
		Subnets:               []*rds.SubnetMap{},
		Ovip:                  util.PtrString(""),
		EntryPort:             util.PtrInt32(int32(0)),
		ResourceGroupId:       util.PtrString(""),
		ReplicaType:           util.PtrString(""),
		IsInheritMasterAuthip: util.PtrBool(false),
		PurchaseCount:         util.PtrInt32(int32(0)),
		ResourceType:          util.PtrString(""),
		AutoRenewTimeUnit:     util.PtrString(""),
		AutoRenewTime:         util.PtrInt32(int32(0)),
	}
	result, err := client.CreateReadOnlyInstance(createReadOnlyInstanceRequest)
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
