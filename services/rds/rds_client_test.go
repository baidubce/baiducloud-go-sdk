package rds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	RDS_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	// ==== AK/SK 鉴权 ====
	RDS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_BatchConfigurationOfInstances(t *testing.T) {
	batchConfigurationOfInstancesRequest := &BatchConfigurationOfInstancesRequest{
		InstanceId:     util.PtrString(""),
		CpuCount:       util.PtrInt32(int32(0)),
		MemoryCapacity: util.PtrInt32(int32(0)),
		VolumeCapacity: util.PtrInt32(int32(0)),
		ReadReplicas:   []*ReplicaRequest{},
		IsBatchResize:  util.PtrBool(false),
		IsDirectPay:    util.PtrBool(false),
		IsEnhanced:     util.PtrBool(false),
		EffectiveTime:  util.PtrString(""),
		MasterAzone:    util.PtrString(""),
		BackupAzone:    util.PtrString(""),
		DiskIoType:     util.PtrString(""),
		SubnetId:       util.PtrString(""),
		EdgeSubnetId:   util.PtrString(""),
		Subnets:        []*SubnetMap{},
		CouponId:       []*int64{},
	}
	result := &BatchConfigurationOfInstancesResponse{}
	result, err := RDS_CLIENT.BatchConfigurationOfInstances(batchConfigurationOfInstancesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindingTags(t *testing.T) {
	bindingTagsRequest := &BindingTagsRequest{
		Resources: []*Resource{},
	}
	err := RDS_CLIENT.BindingTags(bindingTagsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ClusterStatusCheck(t *testing.T) {
	clusterStatusCheckRequest := &ClusterStatusCheckRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ClusterStatusCheckResponse{}
	result, err := RDS_CLIENT.ClusterStatusCheck(clusterStatusCheckRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAPrimaryInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	InitialDataReference := &InitialDataReference{
		InstanceId:    util.PtrString(""),
		ReferenceType: util.PtrString(""),
		Datetime:      util.PtrString(""),
		SnapshotId:    util.PtrString(""),
	}
	createAPrimaryInstanceRequest := &CreateAPrimaryInstanceRequest{
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
		Subnets:              []*SubnetMap{},
		Tags:                 []*Tag{},
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
		Data:                 []*RecoveryToSourceInstanceModel{},
		LeaderInstanceId:     util.PtrString(""),
		ResourceType:         util.PtrString(""),
	}
	result := &CreateAPrimaryInstanceResponse{}
	result, err := RDS_CLIENT.CreateAPrimaryInstance(createAPrimaryInstanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAProxyInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	createAProxyInstanceRequest := &CreateAProxyInstanceRequest{
		ClientToken:      util.PtrString(""),
		Billing:          Billing,
		SourceInstanceId: util.PtrString(""),
		InstanceName:     util.PtrString(""),
		NodeAmount:       util.PtrInt32(int32(0)),
		ZoneNames:        []*string{},
		VpcId:            util.PtrString(""),
		IsDirectPay:      util.PtrBool(false),
		Subnets:          []*SubnetMap{},
		Ovip:             util.PtrString(""),
		EntryPort:        util.PtrInt32(int32(0)),
		ResourceGroupId:  util.PtrString(""),
	}
	result := &CreateAProxyInstanceResponse{}
	result, err := RDS_CLIENT.CreateAProxyInstance(createAProxyInstanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateReadOnlyInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation:   nil,
	}
	createReadOnlyInstanceRequest := &CreateReadOnlyInstanceRequest{
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
		Subnets:               []*SubnetMap{},
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
	result := &CreateReadOnlyInstanceResponse{}
	result, err := RDS_CLIENT.CreateReadOnlyInstance(createReadOnlyInstanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DialTestInterface(t *testing.T) {
	dialTestInterfaceRequest := &DialTestInterfaceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &DialTestInterfaceResponse{}
	result, err := RDS_CLIENT.DialTestInterface(dialTestInterfaceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DoesTheInstanceSupportEnablingAutomaticExpansion(t *testing.T) {
	doesTheInstanceSupportEnablingAutomaticExpansionRequest := &DoesTheInstanceSupportEnablingAutomaticExpansionRequest{
		InstanceId: util.PtrString(""),
	}
	result := &DoesTheInstanceSupportEnablingAutomaticExpansionResponse{}
	result, err := RDS_CLIENT.DoesTheInstanceSupportEnablingAutomaticExpansion(doesTheInstanceSupportEnablingAutomaticExpansionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetNewPurchasePrice(t *testing.T) {
	Instance := &Instance{
		InstanceId:       util.PtrString(""),
		InstanceShortId:  util.PtrString(""),
		InstanceName:     util.PtrString(""),
		Engine:           util.PtrString(""),
		EngineVersion:    util.PtrString(""),
		RdsMinorVersion:  util.PtrString(""),
		CharacterSetName: util.PtrString(""),
		Endpoint: &Endpoint{
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
		BackupPolicy: &SnapshotPolicy{
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
			DataBackupRetainStrategys:     []*DataBackupRetainStrategy{},
			EncryptStrategy: &EncryptStrategy{
				EncryptEnable:            util.PtrBool(false),
				KeyManagementServiceName: util.PtrString(""),
				KeyManagementType:        util.PtrString(""),
				SecretKeyID:              util.PtrString(""),
			},
		},
		BlbService: &BlbService{
			Status:  util.PtrString(""),
			Service: util.PtrString(""),
		},
		PubliclyAccessible: util.PtrBool(false),
		InstanceCreateTime: util.PtrString(""),
		InstanceExpireTime: util.PtrString(""),
		ProductType:        util.PtrString(""),
		PaymentTiming:      util.PtrString(""),
		Topology: &Topology{
			Master:               []*string{},
			ReadReplica:          []*string{},
			Rdsproxy:             []*string{},
			ReadReplicaIdMapping: []*ReadReplicaIdMapping{},
			MasterIdMapping: &MasterIdMapping{
				AppId:                util.PtrString(""),
				AppIdShort:           util.PtrString(""),
				Status:               util.PtrString(""),
				AllocatedMemoryInMB:  util.PtrInt32(int32(0)),
				CpuCount:             util.PtrInt32(int32(0)),
				AllocatedStorageInGB: util.PtrInt32(int32(0)),
				UsedStorageInGB:      util.PtrFloat64(float64(0)),
			},
			RdsproxyIdMapping: &ReadReplicaIdMapping{
				AppId:                util.PtrString(""),
				AppIdShort:           util.PtrString(""),
				Status:               util.PtrString(""),
				AllocatedMemoryInMB:  util.PtrInt32(int32(0)),
				CpuCount:             util.PtrInt32(int32(0)),
				AllocatedStorageInGB: util.PtrInt32(int32(0)),
				UsedStorageInGB:      util.PtrFloat64(float64(0)),
			},
			RdsproxyIdMappings: []*ReadReplicaIdMapping{},
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
		DccHostIds: &DccHostInfo{
			Master: &DccHostId{
				HostId: util.PtrString(""),
				Azone:  util.PtrString(""),
			},
			Backup: &DccHostId{
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
		NodeReadReplica: &RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeMaster: &RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeSlave: &RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeProxy: &RoleInfo{
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
		RoGroupList:       []*RoGroup{},
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
		NodeMasterReadReplica: &RoleInfo{
			Id:       util.PtrString(""),
			Azone:    util.PtrString(""),
			SubnetId: util.PtrString(""),
			Name:     util.PtrString(""),
			VpcCidr:  util.PtrString(""),
			ShortId:  util.PtrString(""),
		},
		NodeBackupReadReplica: &RoleInfo{
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
		Tags:             []*Tag{},
		Task:             util.PtrString(""),
		OrderStatus:      util.PtrString(""),
		VpcCidr:          util.PtrString(""),
		ZoneNames:        []*string{},
		DccHosts:         []*MachinePO{},
		Category:         util.PtrString(""),
		Name:             util.PtrString(""),
		ResourceGroup: &ResourceGroupsDetailFull{
			Groups:    []*ResourceInfo{},
			Name:      util.PtrString(""),
			RdsType:   util.PtrString(""),
			Region:    util.PtrString(""),
			Id:        util.PtrString(""),
			Uuid:      util.PtrString(""),
			Summary:   util.PtrString(""),
			Url:       util.PtrString(""),
			AccountId: util.PtrString(""),
			UserId:    util.PtrString(""),
			Tags:      []*Tag{},
			ResUuid:   util.PtrString(""),
		},
		ResourceGroupId:   util.PtrString(""),
		ResourceGroupName: util.PtrString(""),
		IsChargeProxy:     util.PtrBool(false),
		Single:            util.PtrBool(false),
	}
	getNewPurchasePriceRequest := &GetNewPurchasePriceRequest{
		Instance:    Instance,
		Duration:    util.PtrInt32(int32(0)),
		Number:      util.PtrInt32(int32(0)),
		ProductType: util.PtrString(""),
	}
	result := &GetNewPurchasePriceResponse{}
	result, err := RDS_CLIENT.GetNewPurchasePrice(getNewPurchasePriceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ObtainThePriceDifferenceOfPrepaidTransformer(t *testing.T) {
	obtainThePriceDifferenceOfPrepaidTransformerRequest := &ObtainThePriceDifferenceOfPrepaidTransformerRequest{
		InstanceId:           util.PtrString(""),
		CpuCount:             util.PtrInt32(int32(0)),
		AllocatedMemoryInGB:  util.PtrInt32(int32(0)),
		AllocatedStorageInGB: util.PtrInt32(int32(0)),
		DiskIoType:           util.PtrString(""),
		Category:             util.PtrString(""),
	}
	result := &ObtainThePriceDifferenceOfPrepaidTransformerResponse{}
	result, err := RDS_CLIENT.ObtainThePriceDifferenceOfPrepaidTransformer(obtainThePriceDifferenceOfPrepaidTransformerRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryInstanceList(t *testing.T) {
	queryInstanceListRequest := &QueryInstanceListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		VnetIp:  util.PtrString(""),
	}
	result := &QueryInstanceListResponse{}
	result, err := RDS_CLIENT.QueryInstanceList(queryInstanceListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySubnetList(t *testing.T) {
	querySubnetListRequest := &QuerySubnetListRequest{
		ZoneName: util.PtrString(""),
		VpcId:    util.PtrString(""),
	}
	result := &QuerySubnetListResponse{}
	result, err := RDS_CLIENT.QuerySubnetList(querySubnetListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTaskList(t *testing.T) {
	queryTaskListRequest := &QueryTaskListRequest{
		PageSize:     util.PtrString(""),
		PageNo:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		InstanceName: util.PtrString(""),
		TaskId:       util.PtrInt32(int32(0)),
		TaskType:     util.PtrString(""),
		TaskStatus:   util.PtrString(""),
		StartTime:    util.PtrString(""),
		EndTime:      util.PtrString(""),
	}
	result := &QueryTaskListResponse{}
	result, err := RDS_CLIENT.QueryTaskList(queryTaskListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheDetailsOfTheSpecifiedInstance(t *testing.T) {
	queryTheDetailsOfTheSpecifiedInstanceRequest := &QueryTheDetailsOfTheSpecifiedInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &QueryTheDetailsOfTheSpecifiedInstanceResponse{}
	result, err := RDS_CLIENT.QueryTheDetailsOfTheSpecifiedInstance(queryTheDetailsOfTheSpecifiedInstanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheListOfAvailableZones(t *testing.T) {
	result := &QueryTheListOfAvailableZonesResponse{}
	result, err := RDS_CLIENT.QueryTheListOfAvailableZones()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseInstance(t *testing.T) {
	releaseInstanceRequest := &ReleaseInstanceRequest{
		InstanceIds: util.PtrString(""),
	}
	result := &ReleaseInstanceResponse{}
	result, err := RDS_CLIENT.ReleaseInstance(releaseInstanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RestartTheInstance(t *testing.T) {
	restartTheInstanceRequest := &RestartTheInstanceRequest{
		InstanceId:    util.PtrString(""),
		EffectiveTime: util.PtrString(""),
	}
	err := RDS_CLIENT.RestartTheInstance(restartTheInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RetrieveInstanceAutoScalingConfig(t *testing.T) {
	retrieveInstanceAutoScalingConfigRequest := &RetrieveInstanceAutoScalingConfigRequest{
		InstanceId: util.PtrString(""),
	}
	result := &RetrieveInstanceAutoScalingConfigResponse{}
	result, err := RDS_CLIENT.RetrieveInstanceAutoScalingConfig(retrieveInstanceAutoScalingConfigRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TransformationExample(t *testing.T) {
	transformationExampleRequest := &TransformationExampleRequest{
		InstanceId:          util.PtrString(""),
		CpuCount:            util.PtrInt32(int32(0)),
		MemoryCapacity:      util.PtrInt32(int32(0)),
		VolumeCapacity:      util.PtrInt32(int32(0)),
		AllocatedMemoryInMB: util.PtrInt32(int32(0)),
		NodeAmount:          util.PtrInt32(int32(0)),
		IsDirectPay:         util.PtrBool(false),
		IsEnhanced:          util.PtrBool(false),
		EffectiveTime:       util.PtrString(""),
		MasterAzone:         util.PtrString(""),
		BackupAzone:         util.PtrString(""),
		DiskIoType:          util.PtrString(""),
		SubnetId:            util.PtrString(""),
		EdgeSubnetId:        util.PtrString(""),
		Subnets:             []*SubnetMap{},
	}
	result := &TransformationExampleResponse{}
	result, err := RDS_CLIENT.TransformationExample(transformationExampleRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAutomaticStorageExpansionConfiguration(t *testing.T) {
	updateAutomaticStorageExpansionConfigurationRequest := &UpdateAutomaticStorageExpansionConfigurationRequest{
		InstanceId:         util.PtrString(""),
		Action:             util.PtrString(""),
		FreeSpaceThreshold: util.PtrInt32(int32(0)),
		DiskMaxLimit:       util.PtrInt32(int32(0)),
	}
	err := RDS_CLIENT.UpdateAutomaticStorageExpansionConfiguration(updateAutomaticStorageExpansionConfigurationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceConnectionInformation(t *testing.T) {
	updateInstanceConnectionInformationRequest := &UpdateInstanceConnectionInformationRequest{
		InstanceId: util.PtrString(""),
		Address:    util.PtrString(""),
	}
	err := RDS_CLIENT.UpdateInstanceConnectionInformation(updateInstanceConnectionInformationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceName(t *testing.T) {
	updateInstanceNameRequest := &UpdateInstanceNameRequest{
		InstanceId:   util.PtrString(""),
		InstanceName: util.PtrString(""),
	}
	err := RDS_CLIENT.UpdateInstanceName(updateInstanceNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstancePublicNetworkAccessStatus(t *testing.T) {
	updateInstancePublicNetworkAccessStatusRequest := &UpdateInstancePublicNetworkAccessStatusRequest{
		InstanceId:   util.PtrString(""),
		PublicAccess: util.PtrBool(false),
	}
	err := RDS_CLIENT.UpdateInstancePublicNetworkAccessStatus(updateInstancePublicNetworkAccessStatusRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceSynchronizationMode(t *testing.T) {
	updateInstanceSynchronizationModeRequest := &UpdateInstanceSynchronizationModeRequest{
		InstanceId: util.PtrString(""),
		SyncMode:   util.PtrString(""),
	}
	err := RDS_CLIENT.UpdateInstanceSynchronizationMode(updateInstanceSynchronizationModeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateTimeWindow(t *testing.T) {
	updateTimeWindowRequest := &UpdateTimeWindowRequest{
		InstanceId:        util.PtrString(""),
		MaintainStartTime: util.PtrString(""),
		MaintainDuration:  util.PtrInt32(int32(0)),
	}
	err := RDS_CLIENT.UpdateTimeWindow(updateTimeWindowRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ZoneMigration(t *testing.T) {
	zoneMigrationRequest := &ZoneMigrationRequest{
		InstanceId:    util.PtrString(""),
		MasterAzone:   util.PtrString(""),
		BackupAzone:   util.PtrString(""),
		ZoneNames:     []*string{},
		Subnets:       []*SubnetMap{},
		EffectiveTime: util.PtrString(""),
	}
	err := RDS_CLIENT.ZoneMigration(zoneMigrationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
