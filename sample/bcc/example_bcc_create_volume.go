package bccsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func CreateVolume() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &bcc.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &bcc.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	AutoSnapshotPolicy := &bcc.AutoSnapshotPolicyModel{
		Id:              util.PtrString(""),
		Name:            util.PtrString(""),
		TimePoints:      []*int32{},
		RepeatWeekdays:  []*int32{},
		Status:          util.PtrString(""),
		RetentionDays:   util.PtrInt32(int32(0)),
		CreatedTime:     util.PtrString(""),
		UpdatedTime:     util.PtrString(""),
		DeletedTime:     util.PtrString(""),
		LastExecuteTime: util.PtrString(""),
		VolumeCount:     util.PtrInt32(int32(0)),
	}
	createVolumeRequest := &bcc.CreateVolumeRequest{
		ZoneName:               util.PtrString(""),
		StorageType:            util.PtrString(""),
		CdsSizeInGB:            util.PtrInt32(int32(0)),
		CdsExtraIo:             util.PtrInt32(int32(0)),
		SnapshotId:             util.PtrString(""),
		ShareSnapshotId:        util.PtrString(""),
		EnableDeleteProtection: util.PtrString(""),
		InstanceId:             util.PtrString(""),
		EncryptKey:             util.PtrString(""),
		Name:                   util.PtrString(""),
		Description:            util.PtrString(""),
		RenewTimeUnit:          util.PtrString(""),
		RenewTime:              util.PtrInt32(int32(0)),
		RelationTag:            util.PtrBool(false),
		Tags:                   []*bcc.TagModel{},
		ResGroupId:             util.PtrString(""),
		Billing:                Billing,
		ClusterId:              util.PtrString(""),
		ChargeType:             util.PtrString(""),
		AutoSnapshotPolicy:     AutoSnapshotPolicy,
		DeleteWithInstance:     util.PtrBool(false),
		DeleteAutoSnapshot:     util.PtrBool(false),
		PurchaseCount:          util.PtrInt32(int32(0)),
	}
	result, err := client.CreateVolume(createVolumeRequest)
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
