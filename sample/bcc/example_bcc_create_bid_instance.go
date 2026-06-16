package bccsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func CreateBidInstance() {
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
	createBidInstanceRequest := &bcc.CreateBidInstanceRequest{
		Spec:                      util.PtrString(""),
		ImageId:                   util.PtrString(""),
		Billing:                   Billing,
		BidModel:                  util.PtrString(""),
		BidPrice:                  util.PtrString(""),
		CpuCount:                  util.PtrInt32(int32(0)),
		MemoryCapacityInGB:        util.PtrInt32(int32(0)),
		RootDiskSizeInGb:          util.PtrInt32(int32(0)),
		RootDiskStorageType:       util.PtrString(""),
		CreateCdsList:             []*bcc.CreateCdsModel{},
		EphemeralDisks:            []*bcc.EphemeralDisk{},
		NetworkCapacityInMbps:     util.PtrInt32(int32(0)),
		InternetChargeType:        util.PtrString(""),
		EipName:                   util.PtrString(""),
		PurchaseCount:             util.PtrInt32(int32(0)),
		Name:                      util.PtrString(""),
		Hostname:                  util.PtrString(""),
		AutoSeqSuffix:             util.PtrBool(false),
		IsOpenHostnameDomain:      util.PtrBool(false),
		AdminPass:                 util.PtrString(""),
		KeypairId:                 util.PtrString(""),
		UserData:                  util.PtrString(""),
		ZoneName:                  util.PtrString(""),
		SubnetId:                  util.PtrString(""),
		SecurityGroupId:           util.PtrString(""),
		EnterpriseSecurityGroupId: util.PtrString(""),
		IsomerismCard:             util.PtrString(""),
		DeletionProtection:        util.PtrInt32(int32(0)),
		RelationTag:               util.PtrBool(false),
		IsOpenIpv6:                util.PtrBool(false),
		Tags:                      []*bcc.TagModel{},
		AspId:                     util.PtrString(""),
		FileSystems:               []*bcc.FileSystemModel{},
		IsEipAutoRelatedDelete:    util.PtrBool(false),
		ResGroupId:                util.PtrString(""),
	}
	result, err := client.CreateBidInstance(createBidInstanceRequest)
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
