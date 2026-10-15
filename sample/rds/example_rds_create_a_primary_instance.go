package rdssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func CreateAPrimaryInstance() {
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
		Reservation:   nil,
	}
	InitialDataReference := &rds.InitialDataReference{
		InstanceId:    util.PtrString(""),
		ReferenceType: util.PtrString(""),
		Datetime:      util.PtrString(""),
		SnapshotId:    util.PtrString(""),
	}
	createAPrimaryInstanceRequest := &rds.CreateAPrimaryInstanceRequest{
		ClientToken:          util.PtrString(""),
		Billing:              Billing,
		PurchaseCount:        util.PtrInt32(int32(0)),
		InstanceName:         util.PtrString(""),
		Engine:               util.PtrString(""),
		EngineVersion:        util.PtrString(""),
		Category:             util.PtrString(""),
		CpuCount:             util.PtrInt32(int32(0)),
		MemoryCapacity:       util.PtrInt32(int32(0)),
		VolumeCapacity:       util.PtrInt32(int32(0)),
		DiskIoType:           util.PtrString(""),
		ZoneNames:            []*string{},
		BcmGroupName:         util.PtrString(""),
		VpcId:                util.PtrString(""),
		IsDirectPay:          util.PtrBool(false),
		Subnets:              []*rds.SubnetMap{},
		Tags:                 []*rds.Tag{},
		BgwGroupId:           util.PtrString(""),
		BgwGroupExclusive:    util.PtrBool(false),
		CharacterSetName:     util.PtrString(""),
		AutoRenewTimeUnit:    util.PtrString(""),
		AutoRenewTime:        util.PtrInt32(int32(0)),
		LowerCaseTableNames:  util.PtrInt32(int32(0)),
		ParameterTemplateId:  util.PtrString(""),
		Ovip:                 util.PtrString(""),
		EntryPort:            util.PtrInt32(int32(0)),
		ReplicationType:      util.PtrString(""),
		ResourceGroupId:      util.PtrString(""),
		InitialDataReference: InitialDataReference,
		Data:                 []*rds.RecoveryToSourceInstanceModel{},
		LeaderInstanceId:     util.PtrString(""),
		ResourceType:         util.PtrString(""),
	}
	result, err := client.CreateAPrimaryInstance(createAPrimaryInstanceRequest)
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
