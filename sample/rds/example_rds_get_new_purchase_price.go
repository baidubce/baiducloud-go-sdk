package rdssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func GetNewPurchasePrice() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Instance := &rds.Instance{
		InstanceId:       util.PtrString(""),
		InstanceShortId:  util.PtrString(""),
		InstanceName:     util.PtrString(""),
		Engine:           util.PtrString(""),
		EngineVersion:    util.PtrString(""),
		RdsMinorVersion:  util.PtrString(""),
		CharacterSetName: util.PtrString(""),
		Endpoint: &rds.Endpoint{
			Port:     util.PtrInt32(int32(0)),
			VnetIp:   util.PtrString(""),
			InetIp:   util.PtrString(""),
			Address:  util.PtrString(""),
			InetIpV6: util.PtrString(""),
			VnetIpV6: util.PtrString(""),
		},
		InstanceClass:        util.PtrString(""),
		AllocatedMemoryInMB:  util.PtrInt32(int32(0)),
		AllocatedMemoryInGB:  util.PtrInt32(int32(0)),
		MemoryCapacity:       util.PtrFloat64(float64(0)),
		AllocatedStorageInGB: util.PtrInt32(int32(0)),
		VolumeCapacity:       util.PtrInt32(int32(0)),
		TotalStorageInGB:     util.PtrInt32(int32(0)),
		UsedStorageInMB:      util.PtrInt32(int32(0)),
		UsedStorageInGB:      util.PtrFloat64(float64(0)),
		UsedStorage:          util.PtrFloat64(float64(0)),
		CpuCount:             util.PtrInt32(int32(0)),
		InstanceType:         util.PtrString(""),
		SourceInstanceId:     util.PtrString(""),
		ReadReplicaNum:       util.PtrInt32(int32(0)),
		ReadReplica:          []*string{},
		InstanceStatus:       util.PtrString(""),
		LockMode:             util.PtrString(""),
		EipStatus:            util.PtrString(""),
		PublicAccessStatus:   util.PtrString(""),
		SuperUserFlag:        util.PtrString(""),
		ReplicationType:      util.PtrString(""),
		SyncMode:             util.PtrString(""),
		SqlserverBrType:      util.PtrString(""),
		BackupPolicy: &rds.SnapshotPolicy{
			BackupTime:                    util.PtrString(""),
			BackupDays:                    util.PtrString(""),
			Persistent:                    util.PtrBool(false),
			ExpireInDays:                  util.PtrInt32(int32(0)),
			LogBackupRetainDays:           util.PtrInt32(int32(0)),
			FreeSpaceInGB:                 util.PtrInt32(int32(0)),
			DataBackupType:                util.PtrString(""),
			FreeSpace:                     util.PtrInt32(int32(0)),
			IncrementalDataBackupEnable:   util.PtrBool(false),
			TimeUnit:                      util.PtrString(""),
			DataBackupCopyEnable:          util.PtrBool(false),
			IncrementalDataBackupInterval: util.PtrInt32(int32(0)),
			LatestDataBackupRetainDays:    util.PtrInt32(int32(0)),
			RetentionPeriod:               util.PtrInt32(int32(0)),
			DataBackupRetainStrategys:     []*rds.DataBackupRetainStrategy{},
			EncryptStrategy: &rds.EncryptStrategy{
				EncryptEnable:            util.PtrBool(false),
				KeyManagementServiceName: util.PtrString(""),
				KeyManagementType:        util.PtrString(""),
				SecretKeyID:              util.PtrString(""),
			},
		},
		BlbService: &rds.BlbService{
			Status:  util.PtrString(""),
			Service: util.PtrString(""),
		},
		PubliclyAccessible: util.PtrBool(false),
		InstanceCreateTime: util.PtrString(""),
		InstanceExpireTime: util.PtrString(""),
		ProductType:        util.PtrString(""),
		PaymentTiming:      util.PtrString(""),
		Topology: &rds.Topology{
			Master:               []*string{},
			ReadReplica:          []*string{},
			Rdsproxy:             []*string{},
			ReadReplicaIdMapping: []*rds.ReadReplicaIdMapping{},
			MasterIdMapping: &rds.MasterIdMapping{
				AppId:                util.PtrString(""),
				AppIdShort:           util.PtrString(""),
				Status:               util.PtrString(""),
				AllocatedMemoryInMB:  util.PtrInt32(int32(0)),
				CpuCount:             util.PtrInt32(int32(0)),
				AllocatedStorageInGB: util.PtrInt32(int32(0)),
				UsedStorageInGB:      util.PtrFloat64(float64(0)),
			},
			RdsproxyIdMapping: &rds.ReadReplicaIdMapping{
				AppId:                util.PtrString(""),
				AppIdShort:           util.PtrString(""),
				Status:               util.PtrString(""),
				AllocatedMemoryInMB:  util.PtrInt32(int32(0)),
				CpuCount:             util.PtrInt32(int32(0)),
				AllocatedStorageInGB: util.PtrInt32(int32(0)),
				UsedStorageInGB:      util.PtrFloat64(float64(0)),
			},
			RdsproxyIdMappings: []*rds.ReadReplicaIdMapping{},
		},
		Azone:           util.PtrString(""),
		VpcId:           util.PtrString(""),
		SubnetId:        nil,
		PhysicalZone:    nil,
		Region:          util.PtrString(""),
		SourceRegion:    util.PtrString(""),
		BillingStatus:   util.PtrString(""),
		ApplicationType: util.PtrString(""),
		NodeAmount:      util.PtrInt32(int32(0)),
		OldInstance:     util.PtrInt32(int32(0)),
		OnlineStatus:    util.PtrInt32(int32(0)),
		IsSingle:        util.PtrBool(false),
		DccHostIds: &rds.DccHostInfo{
			Master: &rds.DccHostId{
				HostId: util.PtrString(""),
				Azone:  util.PtrString(""),
			},
			Backup: &rds.DccHostId{
				HostId: util.PtrString(""),
				Azone:  util.PtrString(""),
			},
		},
		MachineType:     util.PtrInt32(int32(0)),
		Weight:          util.PtrInt32(int32(0)),
		NodeType:        util.PtrString(""),
		RepairStartTime: util.PtrString(""),
		RepairEndTime:   util.PtrString(""),
		QuotStatus:      util.PtrString(""),
		NodeReadReplica: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeMaster: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeSlave: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeProxy: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		DiskIoType:        util.PtrString(""),
		GroupId:           util.PtrString(""),
		GroupName:         util.PtrString(""),
		TdeStatus:         util.PtrString(""),
		UserId:            util.PtrString(""),
		BgwGroupId:        util.PtrString(""),
		BgwGroupExclusive: util.PtrBool(false),
		RoGroupList:       []*rds.RoGroup{},
		RoGroupAbnormal:   util.PtrBool(false),
		DiskType:          util.PtrString(""),
		CdsType:           util.PtrString(""),
		MaintainStartTime: util.PtrString(""),
		MaintainDuration:  util.PtrInt32(int32(0)),
		MaintainPeriod:    util.PtrString(""),
		InstanceSubStatus: util.PtrString(""),
		BcmSendCycle:      util.PtrString(""),
		AutoResizeDisk:    util.PtrInt32(int32(0)),
		AutoResizeCpu:     util.PtrInt32(int32(0)),
		EdgeRegion:        util.PtrString(""),
		EdgeRegionName:    util.PtrString(""),
		EdgeVpcId:         util.PtrString(""),
		EdgeSubnetId:      util.PtrString(""),
		GtidType:          util.PtrInt32(int32(0)),
		NodeMasterReadReplica: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeBackupReadReplica: &rds.RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		HaStrategy:       util.PtrInt32(int32(0)),
		IsDBRSEnable:     util.PtrBool(false),
		ReplicaType:      util.PtrString(""),
		ResourceType:     util.PtrString(""),
		CdsEncrypt:       util.PtrInt32(int32(0)),
		CdsEncryptKey:    util.PtrString(""),
		ResourcePlatform: util.PtrString(""),
		Location:         util.PtrString(""),
		HasSlave:         util.PtrBool(false),
		ExpireDate:       util.PtrInt32(int32(0)),
		HasProxy:         util.PtrBool(false),
		VpcName:          util.PtrString(""),
		ResourceUuid:     util.PtrString(""),
		Tags:             []*rds.Tag{},
		Task:             util.PtrString(""),
		OrderStatus:      util.PtrString(""),
		VpcCidr:          util.PtrString(""),
		ZoneNames:        []*string{},
		DccHosts:         []*rds.MachinePO{},
		Category:         util.PtrString(""),
		Name:             util.PtrString(""),
		ResourceGroup: &rds.ResourceGroupsDetailFull{
			Groups:    []*rds.ResourceInfo{},
			Name:      util.PtrString(""),
			RdsType:   util.PtrString(""),
			Region:    util.PtrString(""),
			Id:        util.PtrString(""),
			Uuid:      util.PtrString(""),
			Summary:   util.PtrString(""),
			Url:       util.PtrString(""),
			AccountId: util.PtrString(""),
			UserId:    util.PtrString(""),
			Tags:      []*rds.Tag{},
			ResUuid:   util.PtrString(""),
		},
		ResourceGroupId:   util.PtrString(""),
		ResourceGroupName: util.PtrString(""),
		IsChargeProxy:     util.PtrBool(false),
		Single:            util.PtrBool(false),
	}
	getNewPurchasePriceRequest := &rds.GetNewPurchasePriceRequest{
		Instance:    Instance,
		Duration:    util.PtrInt32(int32(0)),
		Number:      util.PtrInt32(int32(0)),
		ProductType: util.PtrString(""),
	}
	result, err := client.GetNewPurchasePrice(getNewPurchasePriceRequest)
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
