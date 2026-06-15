package bccsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func CreateInstanceBySpec() {
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
	createInstanceBySpecRequest := &bcc.CreateInstanceBySpecRequest{
		KeepImageLogin:             util.PtrBool(false),
		BccCreateWithScript:        util.PtrString(""),
		Name:                       util.PtrString(""),
		CpuThreadConfig:            util.PtrString(""),
		NumaConfig:                 util.PtrString(""),
		EnableDeleteProtection:     util.PtrBool(false),
		Hostname:                   util.PtrString(""),
		AutoSeqSuffix:              util.PtrBool(false),
		IsOpenHostnameDomain:       util.PtrBool(false),
		AdminPass:                  util.PtrString(""),
		KeypairId:                  util.PtrString(""),
		AspId:                      util.PtrString(""),
		SpecId:                     util.PtrString(""),
		EnableJumboFrame:           util.PtrBool(false),
		UserData:                   util.PtrString(""),
		DeletionProtection:         util.PtrString(""),
		AutoRenewTimeUnit:          util.PtrString(""),
		AutoRenewTime:              util.PtrInt32(int32(0)),
		HosteyeType:                util.PtrString(""),
		EnableNuma:                 util.PtrBool(false),
		DataPartitionType:          util.PtrString(""),
		RootPartitionType:          util.PtrString(""),
		CdsAutoRenew:               util.PtrBool(false),
		CreateCdsList:              []*bcc.CreateCdsModel{},
		ImageId:                    util.PtrString(""),
		Spec:                       util.PtrString(""),
		RoleName:                   util.PtrString(""),
		BidModel:                   util.PtrString(""),
		BidPrice:                   util.PtrString(""),
		RootDiskSizeInGb:           util.PtrInt32(int32(0)),
		RootDiskExtraIo:            util.PtrString(""),
		RootDiskStorageType:        util.PtrString(""),
		NetworkCapacityInMbps:      util.PtrInt32(int32(0)),
		EhcClusterId:               util.PtrString(""),
		PurchaseCount:              util.PtrInt32(int32(0)),
		PurchaseMinCount:           util.PtrInt32(int32(0)),
		DedicatedHostId:            util.PtrString(""),
		RelationTag:                util.PtrBool(false),
		Tags:                       []*bcc.TagModel{},
		FileSystems:                []*bcc.FileSystemModel{},
		EphemeralDisks:             []*bcc.EphemeralDisk{},
		SecurityGroupId:            util.PtrString(""),
		EnterpriseSecurityGroupId:  util.PtrString(""),
		SecurityGroupIds:           []*string{},
		EnterpriseSecurityGroupIds: []*string{},
		SubnetId:                   util.PtrString(""),
		DeployId:                   util.PtrString(""),
		DeployIdList:               []*string{},
		EniIds:                     []*string{},
		DisableRootDiskSerial:      util.PtrString(""),
		ZoneName:                   util.PtrString(""),
		InternalIps:                []*string{},
		ResGroupId:                 util.PtrString(""),
		IsEipAutoRelatedDelete:     util.PtrBool(false),
		NetworkPurchaseType:        util.PtrString(""),
		InstanceType:               util.PtrString(""),
		InternetChargeType:         util.PtrString(""),
		EipName:                    util.PtrString(""),
		IsOpenHostEye:              util.PtrBool(false),
		EnableHt:                   util.PtrBool(false),
		Billing:                    Billing,
		IsOpenIpv6:                 util.PtrBool(false),
	}
	result, err := client.CreateInstanceBySpec(createInstanceBySpecRequest)
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
