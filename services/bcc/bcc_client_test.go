package bcc

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
	BCC_CLIENT *Client
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

	BCC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddIpv6(t *testing.T) {
	addIpv6Request := &AddIpv6Request{
		InstanceId:  util.PtrString(""),
		Ipv6Address: util.PtrString(""),
		Reboot:      util.PtrBool(false),
	}
	result := &AddIpv6Response{}
	result, err := BCC_CLIENT.AddIpv6(addIpv6Request)
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
func TestClient_AttachVolume(t *testing.T) {
	attachVolumeRequest := &AttachVolumeRequest{
		VolumeId:           util.PtrString(""),
		InstanceId:         util.PtrString(""),
		DeleteWithInstance: util.PtrBool(false),
		DeleteAutoSnapshot: util.PtrBool(false),
	}
	result := &AttachVolumeResponse{}
	result, err := BCC_CLIENT.AttachVolume(attachVolumeRequest)
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
func TestClient_AutoReleaseInstance(t *testing.T) {
	autoReleaseInstanceRequest := &AutoReleaseInstanceRequest{
		InstanceId:             util.PtrString(""),
		IsEipAutoRelatedDelete: util.PtrBool(false),
		ReleaseTime:            util.PtrString(""),
	}
	err := BCC_CLIENT.AutoReleaseInstance(autoReleaseInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchAddIp(t *testing.T) {
	batchAddIpRequest := &BatchAddIpRequest{
		InstanceId:                     util.PtrString(""),
		SecondaryPrivateIpAddressCount: util.PtrInt32(int32(0)),
		PrivateIps:                     []*string{},
		AllocateMultiIpv6Addr:          util.PtrBool(false),
	}
	err := BCC_CLIENT.BatchAddIp(batchAddIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchChangeToPostpaid(t *testing.T) {
	batchChangeToPostpaidRequest := &BatchChangeToPostpaidRequest{
		Config: []*PostpayConfig{},
	}
	result := &BatchChangeToPostpaidResponse{}
	result, err := BCC_CLIENT.BatchChangeToPostpaid(batchChangeToPostpaidRequest)
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
func TestClient_BatchChangeToPrepaid(t *testing.T) {
	batchChangeToPrepaidRequest := &BatchChangeToPrepaidRequest{
		Config: []*PrepayConfig{},
	}
	result := &BatchChangeToPrepaidResponse{}
	result, err := BCC_CLIENT.BatchChangeToPrepaid(batchChangeToPrepaidRequest)
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
func TestClient_BatchDeleteIp(t *testing.T) {
	batchDeleteIpRequest := &BatchDeleteIpRequest{
		InstanceId: util.PtrString(""),
		PrivateIps: []*string{},
	}
	err := BCC_CLIENT.BatchDeleteIp(batchDeleteIpRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchRefundResource(t *testing.T) {
	batchRefundResourceRequest := &BatchRefundResourceRequest{
		InstanceIds:           []*string{},
		RelatedReleaseFlag:    util.PtrBool(false),
		DeleteCdsSnapshotFlag: util.PtrBool(false),
		DeleteRelatedEnisFlag: util.PtrBool(false),
	}
	result := &BatchRefundResourceResponse{}
	result, err := BCC_CLIENT.BatchRefundResource(batchRefundResourceRequest)
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
func TestClient_BatchStartInstance(t *testing.T) {
	batchStartInstanceRequest := &BatchStartInstanceRequest{
		InstanceIds: []*string{},
	}
	err := BCC_CLIENT.BatchStartInstance(batchStartInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchStopInstance(t *testing.T) {
	batchStopInstanceRequest := &BatchStopInstanceRequest{
		InstanceIds:      []*string{},
		ForceStop:        util.PtrBool(false),
		StopWithNoCharge: util.PtrBool(false),
	}
	err := BCC_CLIENT.BatchStopInstance(batchStopInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindInstanceToSecurityGroup(t *testing.T) {
	bindInstanceToSecurityGroupRequest := &BindInstanceToSecurityGroupRequest{
		InstanceId:      util.PtrString(""),
		SecurityGroupId: util.PtrString(""),
	}
	err := BCC_CLIENT.BindInstanceToSecurityGroup(bindInstanceToSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindInstanceToTags(t *testing.T) {
	bindInstanceToTagsRequest := &BindInstanceToTagsRequest{
		InstanceId:               util.PtrString(""),
		ChangeTags:               []*TagModel{},
		AttachRelatedResourceTag: util.PtrBool(false),
	}
	err := BCC_CLIENT.BindInstanceToTags(bindInstanceToTagsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindRole(t *testing.T) {
	bindRoleRequest := &BindRoleRequest{
		RoleName:  util.PtrString(""),
		Instances: []*InstancePassRoleModel{},
	}
	result := &BindRoleResponse{}
	result, err := BCC_CLIENT.BindRole(bindRoleRequest)
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
func TestClient_BindTagImage(t *testing.T) {
	bindTagImageRequest := &BindTagImageRequest{
		ImageId:    util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.BindTagImage(bindTagImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindTagVolume(t *testing.T) {
	bindTagVolumeRequest := &BindTagVolumeRequest{
		VolumeId:   util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.BindTagVolume(bindTagVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelRemoteCopyImage(t *testing.T) {
	cancelRemoteCopyImageRequest := &CancelRemoteCopyImageRequest{
		ImageId: util.PtrString(""),
	}
	err := BCC_CLIENT.CancelRemoteCopyImage(cancelRemoteCopyImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ChangeToPrepaid(t *testing.T) {
	changeToPrepaidRequest := &ChangeToPrepaidRequest{
		InstanceId:      util.PtrString(""),
		Duration:        util.PtrInt32(int32(0)),
		AutoRenew:       util.PtrBool(false),
		AutoRenewPeriod: util.PtrInt32(int32(0)),
		RelationCds:     util.PtrBool(false),
	}
	result := &ChangeToPrepaidResponse{}
	result, err := BCC_CLIENT.ChangeToPrepaid(changeToPrepaidRequest)
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
func TestClient_ChangeVpc(t *testing.T) {
	changeVpcRequest := &ChangeVpcRequest{
		InstanceId:                 util.PtrString(""),
		SubnetId:                   util.PtrString(""),
		InternalIp:                 util.PtrString(""),
		SecurityGroupIds:           []*string{},
		EnterpriseSecurityGroupIds: []*string{},
		Reboot:                     util.PtrBool(false),
	}
	err := BCC_CLIENT.ChangeVpc(changeVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAutoRenewRule(t *testing.T) {
	createAutoRenewRuleRequest := &CreateAutoRenewRuleRequest{
		InstanceId:    util.PtrString(""),
		RenewTimeUnit: util.PtrString(""),
		RenewTime:     util.PtrInt32(int32(0)),
		RenewEip:      util.PtrBool(false),
	}
	err := BCC_CLIENT.CreateAutoRenewRule(createAutoRenewRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateImage(t *testing.T) {
	createImageRequest := &CreateImageRequest{
		ImageName:  util.PtrString(""),
		InstanceId: util.PtrString(""),
		SnapshotId: util.PtrString(""),
		EncryptKey: util.PtrString(""),
		RelateCds:  util.PtrBool(false),
		Detection:  util.PtrBool(false),
	}
	result := &CreateImageResponse{}
	result, err := BCC_CLIENT.CreateImage(createImageRequest)
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
func TestClient_CreateInstanceBySpec(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createInstanceBySpecRequest := &CreateInstanceBySpecRequest{
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
		CreateCdsList:              []*CreateCdsModel{},
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
		Tags:                       []*TagModel{},
		FileSystems:                []*FileSystemModel{},
		EphemeralDisks:             []*EphemeralDisk{},
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
	result := &CreateInstanceBySpecResponse{}
	result, err := BCC_CLIENT.CreateInstanceBySpec(createInstanceBySpecRequest)
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
func TestClient_CreateVolume(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	AutoSnapshotPolicy := &AutoSnapshotPolicyModel{
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
	createVolumeRequest := &CreateVolumeRequest{
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
		Tags:                   []*TagModel{},
		ResGroupId:             util.PtrString(""),
		Billing:                Billing,
		ClusterId:              util.PtrString(""),
		ChargeType:             util.PtrString(""),
		AutoSnapshotPolicy:     AutoSnapshotPolicy,
		DeleteWithInstance:     util.PtrBool(false),
		DeleteAutoSnapshot:     util.PtrBool(false),
		PurchaseCount:          util.PtrInt32(int32(0)),
	}
	result := &CreateVolumeResponse{}
	result, err := BCC_CLIENT.CreateVolume(createVolumeRequest)
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
func TestClient_DelIpv6(t *testing.T) {
	delIpv6Request := &DelIpv6Request{
		InstanceId:  util.PtrString(""),
		Ipv6Address: util.PtrString(""),
		Reboot:      util.PtrBool(false),
	}
	err := BCC_CLIENT.DelIpv6(delIpv6Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAutoRenewRule(t *testing.T) {
	deleteAutoRenewRuleRequest := &DeleteAutoRenewRuleRequest{
		InstanceId: util.PtrString(""),
		RenewEip:   util.PtrBool(false),
	}
	err := BCC_CLIENT.DeleteAutoRenewRule(deleteAutoRenewRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteImage(t *testing.T) {
	deleteImageRequest := &DeleteImageRequest{
		ImageId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteImage(deleteImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstanceDeploySet(t *testing.T) {
	deleteInstanceDeploySetRequest := &DeleteInstanceDeploySetRequest{
		DeployId:       util.PtrString(""),
		InstanceIdList: []*string{},
	}
	err := BCC_CLIENT.DeleteInstanceDeploySet(deleteInstanceDeploySetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletePrepayInstance(t *testing.T) {
	deletePrepayInstanceRequest := &DeletePrepayInstanceRequest{
		InstanceId:            util.PtrString(""),
		RelatedReleaseFlag:    util.PtrBool(false),
		DeleteCdsSnapshotFlag: util.PtrBool(false),
		DeleteRelatedEnisFlag: util.PtrBool(false),
	}
	result := &DeletePrepayInstanceResponse{}
	result, err := BCC_CLIENT.DeletePrepayInstance(deletePrepayInstanceRequest)
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
func TestClient_DeleteRecycledInstance(t *testing.T) {
	deleteRecycledInstanceRequest := &DeleteRecycledInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteRecycledInstance(deleteRecycledInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetachVolume(t *testing.T) {
	detachVolumeRequest := &DetachVolumeRequest{
		VolumeId:   util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	err := BCC_CLIENT.DetachVolume(detachVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnterRescueMode(t *testing.T) {
	enterRescueModeRequest := &EnterRescueModeRequest{
		InstanceId: util.PtrString(""),
		ForceStop:  util.PtrBool(false),
		Password:   util.PtrString(""),
	}
	result := &EnterRescueModeResponse{}
	result, err := BCC_CLIENT.EnterRescueMode(enterRescueModeRequest)
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
func TestClient_ExitRescueMode(t *testing.T) {
	exitRescueModeRequest := &ExitRescueModeRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ExitRescueModeResponse{}
	result, err := BCC_CLIENT.ExitRescueMode(exitRescueModeRequest)
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
func TestClient_GetAvailableImagesBySpec(t *testing.T) {
	getAvailableImagesBySpecRequest := &GetAvailableImagesBySpecRequest{
		Spec:    util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		OsName:  util.PtrString(""),
	}
	result := &GetAvailableImagesBySpecResponse{}
	result, err := BCC_CLIENT.GetAvailableImagesBySpec(getAvailableImagesBySpecRequest)
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
func TestClient_GetCdsPrice(t *testing.T) {
	getCdsPriceRequest := &GetCdsPriceRequest{
		StorageType:    util.PtrString(""),
		CdsSizeInGB:    util.PtrInt32(int32(0)),
		PaymentTiming:  util.PtrString(""),
		ZoneName:       util.PtrString(""),
		PurchaseCount:  util.PtrInt32(int32(0)),
		PurchaseLength: util.PtrInt32(int32(0)),
		CdsExtraIo:     util.PtrInt32(int32(0)),
	}
	result := &GetCdsPriceResponse{}
	result, err := BCC_CLIENT.GetCdsPrice(getCdsPriceRequest)
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
func TestClient_GetDiskQuota(t *testing.T) {
	getDiskQuotaRequest := &GetDiskQuotaRequest{
		ZoneName: util.PtrString(""),
	}
	result := &GetDiskQuotaResponse{}
	result, err := BCC_CLIENT.GetDiskQuota(getDiskQuotaRequest)
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
func TestClient_GetImage(t *testing.T) {
	getImageRequest := &GetImageRequest{
		ImageId: util.PtrString(""),
	}
	result := &GetImageResponse{}
	result, err := BCC_CLIENT.GetImage(getImageRequest)
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
func TestClient_GetInstance(t *testing.T) {
	getInstanceRequest := &GetInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetInstanceResponse{}
	result, err := BCC_CLIENT.GetInstance(getInstanceRequest)
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
func TestClient_GetInstanceNoChargeList(t *testing.T) {
	getInstanceNoChargeListRequest := &GetInstanceNoChargeListRequest{
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		InternalIp:  util.PtrString(""),
		KeypairId:   util.PtrString(""),
		ZoneName:    util.PtrString(""),
		InstanceIds: util.PtrString(""),
	}
	result := &GetInstanceNoChargeListResponse{}
	result, err := BCC_CLIENT.GetInstanceNoChargeList(getInstanceNoChargeListRequest)
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
func TestClient_GetInstanceUserDataInfo(t *testing.T) {
	getInstanceUserDataInfoRequest := &GetInstanceUserDataInfoRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetInstanceUserDataInfoResponse{}
	result, err := BCC_CLIENT.GetInstanceUserDataInfo(getInstanceUserDataInfoRequest)
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
func TestClient_GetInstanceVnc(t *testing.T) {
	getInstanceVncRequest := &GetInstanceVncRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetInstanceVncResponse{}
	result, err := BCC_CLIENT.GetInstanceVnc(getInstanceVncRequest)
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
func TestClient_GetRoleList(t *testing.T) {
	result := &GetRoleListResponse{}
	result, err := BCC_CLIENT.GetRoleList()
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
func TestClient_GetVolume(t *testing.T) {
	getVolumeRequest := &GetVolumeRequest{
		VolumeId: util.PtrString(""),
	}
	result := &GetVolumeResponse{}
	result, err := BCC_CLIENT.GetVolume(getVolumeRequest)
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
func TestClient_GetVolumeResizeProgress(t *testing.T) {
	getVolumeResizeProgressRequest := &GetVolumeResizeProgressRequest{
		VolumeId: util.PtrString(""),
	}
	result := &GetVolumeResizeProgressResponse{}
	result, err := BCC_CLIENT.GetVolumeResizeProgress(getVolumeResizeProgressRequest)
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
func TestClient_ImportImage(t *testing.T) {
	importImageRequest := &ImportImageRequest{
		OsName:         util.PtrString(""),
		OsArch:         util.PtrString(""),
		OsType:         util.PtrString(""),
		OsVersion:      util.PtrString(""),
		Name:           util.PtrString(""),
		BosUrl:         util.PtrString(""),
		Detection:      util.PtrBool(false),
		GenerationType: util.PtrString(""),
	}
	result := &ImportImageResponse{}
	result, err := BCC_CLIENT.ImportImage(importImageRequest)
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
func TestClient_InstanceBatchResizeBySpec(t *testing.T) {
	instanceBatchResizeBySpecRequest := &InstanceBatchResizeBySpecRequest{
		Spec:             util.PtrString(""),
		InstanceIdList:   []*string{},
		EnableJumboFrame: util.PtrBool(false),
		SubnetId:         util.PtrString(""),
		LogicalZone:      util.PtrString(""),
	}
	result := &InstanceBatchResizeBySpecResponse{}
	result, err := BCC_CLIENT.InstanceBatchResizeBySpec(instanceBatchResizeBySpecRequest)
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
func TestClient_InstanceDeletionProtection(t *testing.T) {
	instanceDeletionProtectionRequest := &InstanceDeletionProtectionRequest{
		InstanceId:         util.PtrString(""),
		DeletionProtection: util.PtrInt32(int32(0)),
	}
	err := BCC_CLIENT.InstanceDeletionProtection(instanceDeletionProtectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_InstanceRecovery(t *testing.T) {
	instanceRecoveryRequest := &InstanceRecoveryRequest{
		InstanceIds: []*InstanceIdItem{},
	}
	err := BCC_CLIENT.InstanceRecovery(instanceRecoveryRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAvailableResizeSpec(t *testing.T) {
	listAvailableResizeSpecRequest := &ListAvailableResizeSpecRequest{
		Spec:           util.PtrString(""),
		SpecId:         util.PtrString(""),
		Zone:           util.PtrString(""),
		InstanceIdList: []*string{},
	}
	result := &ListAvailableResizeSpecResponse{}
	result, err := BCC_CLIENT.ListAvailableResizeSpec(listAvailableResizeSpecRequest)
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
func TestClient_ListImages(t *testing.T) {
	listImagesRequest := &ListImagesRequest{
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		ImageType: util.PtrString(""),
		ImageName: util.PtrString(""),
	}
	result := &ListImagesResponse{}
	result, err := BCC_CLIENT.ListImages(listImagesRequest)
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
func TestClient_ListInstanceByIds(t *testing.T) {
	listInstanceByIdsRequest := &ListInstanceByIdsRequest{
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		InstanceIds: []*string{},
	}
	result := &ListInstanceByIdsResponse{}
	result, err := BCC_CLIENT.ListInstanceByIds(listInstanceByIdsRequest)
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
func TestClient_ListInstanceEnis(t *testing.T) {
	listInstanceEnisRequest := &ListInstanceEnisRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ListInstanceEnisResponse{}
	result, err := BCC_CLIENT.ListInstanceEnis(listInstanceEnisRequest)
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
func TestClient_ListInstances(t *testing.T) {
	listInstancesRequest := &ListInstancesRequest{
		Marker:            util.PtrString(""),
		MaxKeys:           util.PtrInt32(int32(0)),
		InternalIp:        util.PtrString(""),
		DedicatedHostId:   util.PtrString(""),
		ZoneName:          util.PtrString(""),
		ShowRdmaTopo:      util.PtrString(""),
		InstanceIds:       util.PtrString(""),
		InstanceNames:     util.PtrString(""),
		FuzzyInstanceName: util.PtrString(""),
		VolumeIds:         util.PtrString(""),
		DeploySetIds:      util.PtrString(""),
		SecurityGroupIds:  util.PtrString(""),
		PaymentTiming:     util.PtrString(""),
		Status:            util.PtrString(""),
		Tags:              util.PtrString(""),
		VpcId:             util.PtrString(""),
		PrivateIps:        util.PtrString(""),
		EhcClusterId:      util.PtrString(""),
	}
	result := &ListInstancesResponse{}
	result, err := BCC_CLIENT.ListInstances(listInstancesRequest)
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
func TestClient_ListOs(t *testing.T) {
	listOsRequest := &ListOsRequest{
		InstanceIds: []*string{},
	}
	result := &ListOsResponse{}
	result, err := BCC_CLIENT.ListOs(listOsRequest)
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
func TestClient_ListRecycleInstance(t *testing.T) {
	listRecycleInstanceRequest := &ListRecycleInstanceRequest{
		Marker:        util.PtrString(""),
		MaxKeys:       util.PtrInt32(int32(0)),
		InstanceId:    util.PtrString(""),
		Name:          util.PtrString(""),
		PaymentTiming: util.PtrString(""),
		RecycleBegin:  util.PtrString(""),
		RecycleEnd:    util.PtrString(""),
	}
	result := &ListRecycleInstanceResponse{}
	result, err := BCC_CLIENT.ListRecycleInstance(listRecycleInstanceRequest)
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
func TestClient_ListSharedUser(t *testing.T) {
	listSharedUserRequest := &ListSharedUserRequest{
		ImageId: util.PtrString(""),
	}
	result := &ListSharedUserResponse{}
	result, err := BCC_CLIENT.ListSharedUser(listSharedUserRequest)
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
func TestClient_ListVolumes(t *testing.T) {
	listVolumesRequest := &ListVolumesRequest{
		Marker:          util.PtrString(""),
		MaxKeys:         util.PtrInt32(int32(0)),
		InstanceId:      util.PtrString(""),
		ZoneName:        util.PtrString(""),
		ClusterId:       util.PtrString(""),
		ChargeFilter:    util.PtrString(""),
		UsageFilter:     util.PtrString(""),
		Name:            util.PtrString(""),
		ProductCategory: util.PtrString(""),
	}
	result := &ListVolumesResponse{}
	result, err := BCC_CLIENT.ListVolumes(listVolumesRequest)
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
func TestClient_ModifyCdsAttribute(t *testing.T) {
	modifyCdsAttributeRequest := &ModifyCdsAttributeRequest{
		VolumeId:           util.PtrString(""),
		CdsName:            util.PtrString(""),
		Desc:               util.PtrString(""),
		DeleteWithInstance: util.PtrBool(false),
		DeleteAutoSnapshot: util.PtrBool(false),
	}
	err := BCC_CLIENT.ModifyCdsAttribute(modifyCdsAttributeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceAttributes(t *testing.T) {
	modifyInstanceAttributesRequest := &ModifyInstanceAttributesRequest{
		InstanceId:       util.PtrString(""),
		Name:             util.PtrString(""),
		EnableJumboFrame: util.PtrBool(false),
		NetEthQueueCount: util.PtrString(""),
	}
	err := BCC_CLIENT.ModifyInstanceAttributes(modifyInstanceAttributesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceDesc(t *testing.T) {
	modifyInstanceDescRequest := &ModifyInstanceDescRequest{
		InstanceId: util.PtrString(""),
		Desc:       util.PtrString(""),
	}
	err := BCC_CLIENT.ModifyInstanceDesc(modifyInstanceDescRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceHostname(t *testing.T) {
	modifyInstanceHostnameRequest := &ModifyInstanceHostnameRequest{
		InstanceId:            util.PtrString(""),
		Reboot:                util.PtrBool(false),
		IsOpenHostnameDomain:  util.PtrBool(false),
		Hostname:              util.PtrString(""),
		IsAllowHostnameRepeat: util.PtrBool(false),
	}
	err := BCC_CLIENT.ModifyInstanceHostname(modifyInstanceHostnameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstancePassword(t *testing.T) {
	modifyInstancePasswordRequest := &ModifyInstancePasswordRequest{
		InstanceId: util.PtrString(""),
		AdminPass:  util.PtrString(""),
	}
	err := BCC_CLIENT.ModifyInstancePassword(modifyInstancePasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyRelatedDeletePolicy(t *testing.T) {
	modifyRelatedDeletePolicyRequest := &ModifyRelatedDeletePolicyRequest{
		InstanceId:             util.PtrString(""),
		IsEipAutoRelatedDelete: util.PtrBool(false),
	}
	err := BCC_CLIENT.ModifyRelatedDeletePolicy(modifyRelatedDeletePolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyVolumeChargeType(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	modifyVolumeChargeTypeRequest := &ModifyVolumeChargeTypeRequest{
		VolumeId:      util.PtrString(""),
		Billing:       Billing,
		EffectiveType: util.PtrString(""),
	}
	err := BCC_CLIENT.ModifyVolumeChargeType(modifyVolumeChargeTypeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PurchaseReservedInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	purchaseReservedInstanceRequest := &PurchaseReservedInstanceRequest{
		InstanceId:       util.PtrString(""),
		RelatedRenewFlag: util.PtrString(""),
		Billing:          Billing,
		CdsCustomPeriod:  []*CdsCustomPeriod{},
	}
	result := &PurchaseReservedInstanceResponse{}
	result, err := BCC_CLIENT.PurchaseReservedInstance(purchaseReservedInstanceRequest)
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
func TestClient_PurchaseReservedVolume(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	purchaseReservedVolumeRequest := &PurchaseReservedVolumeRequest{
		VolumeId:   util.PtrString(""),
		Billing:    Billing,
		InstanceId: util.PtrString(""),
	}
	result := &PurchaseReservedVolumeResponse{}
	result, err := BCC_CLIENT.PurchaseReservedVolume(purchaseReservedVolumeRequest)
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
func TestClient_RebootInstance(t *testing.T) {
	rebootInstanceRequest := &RebootInstanceRequest{
		InstanceId: util.PtrString(""),
		ForceStop:  util.PtrBool(false),
	}
	err := BCC_CLIENT.RebootInstance(rebootInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RebuildBatchInstance(t *testing.T) {
	rebuildBatchInstanceRequest := &RebuildBatchInstanceRequest{
		ImageId:           util.PtrString(""),
		KeepImageLogin:    util.PtrBool(false),
		IsPreserveData:    util.PtrBool(false),
		AdminPass:         util.PtrString(""),
		IsOpenHostEye:     util.PtrBool(false),
		SysRootSize:       util.PtrInt32(int32(0)),
		KeypairId:         util.PtrString(""),
		DataPartitionType: util.PtrString(""),
		RootPartitionType: util.PtrString(""),
		RaidId:            util.PtrString(""),
		UserData:          util.PtrString(""),
		UseLastUserData:   util.PtrBool(false),
		CleanLastUserData: util.PtrBool(false),
		InstanceIds:       []*string{},
	}
	err := BCC_CLIENT.RebuildBatchInstance(rebuildBatchInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RebuildInstance(t *testing.T) {
	rebuildInstanceRequest := &RebuildInstanceRequest{
		InstanceId:        util.PtrString(""),
		ImageId:           util.PtrString(""),
		KeepImageLogin:    util.PtrBool(false),
		IsPreserveData:    util.PtrBool(false),
		AdminPass:         util.PtrString(""),
		IsOpenHostEye:     util.PtrBool(false),
		SysRootSize:       util.PtrInt32(int32(0)),
		KeypairId:         util.PtrString(""),
		DataPartitionType: util.PtrString(""),
		RootPartitionType: util.PtrString(""),
		RaidId:            util.PtrString(""),
		UserData:          util.PtrString(""),
		UseLastUserData:   util.PtrBool(false),
		CleanLastUserData: util.PtrBool(false),
	}
	err := BCC_CLIENT.RebuildInstance(rebuildInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseInstanceByPost(t *testing.T) {
	releaseInstanceByPostRequest := &ReleaseInstanceByPostRequest{
		InstanceId:            util.PtrString(""),
		RelatedReleaseFlag:    util.PtrBool(false),
		DeleteCdsSnapshotFlag: util.PtrBool(false),
		DeleteRelatedEnisFlag: util.PtrBool(false),
		BccRecycleFlag:        util.PtrBool(false),
		CdsAttributeActive:    util.PtrBool(false),
	}
	err := BCC_CLIENT.ReleaseInstanceByPost(releaseInstanceByPostRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseMultipleInstanceByPost(t *testing.T) {
	releaseMultipleInstanceByPostRequest := &ReleaseMultipleInstanceByPostRequest{
		InstanceIds:           []*string{},
		RelatedReleaseFlag:    util.PtrBool(false),
		DeleteRelatedCdsFlag:  util.PtrBool(false),
		DeleteCdsSnapshotFlag: util.PtrBool(false),
		DeleteRelatedEnisFlag: util.PtrBool(false),
		BccRecycleFlag:        util.PtrBool(false),
	}
	err := BCC_CLIENT.ReleaseMultipleInstanceByPost(releaseMultipleInstanceByPostRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseVolume(t *testing.T) {
	releaseVolumeRequest := &ReleaseVolumeRequest{
		VolumeId:           util.PtrString(""),
		AutoSnapshot:       util.PtrString(""),
		ManualSnapshot:     util.PtrString(""),
		CdsAttributeActive: util.PtrBool(false),
		Recycle:            util.PtrString(""),
	}
	err := BCC_CLIENT.ReleaseVolume(releaseVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoteCopyImage(t *testing.T) {
	remoteCopyImageRequest := &RemoteCopyImageRequest{
		ImageId:    util.PtrString(""),
		Name:       util.PtrString(""),
		DestRegion: []*string{},
	}
	result := &RemoteCopyImageResponse{}
	result, err := BCC_CLIENT.RemoteCopyImage(remoteCopyImageRequest)
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
func TestClient_RenameImage(t *testing.T) {
	renameImageRequest := &RenameImageRequest{
		ImageIds: []*string{},
		Name:     util.PtrString(""),
	}
	err := BCC_CLIENT.RenameImage(renameImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenameVolume(t *testing.T) {
	renameVolumeRequest := &RenameVolumeRequest{
		VolumeId: util.PtrString(""),
		Name:     util.PtrString(""),
	}
	err := BCC_CLIENT.RenameVolume(renameVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeInstanceBySpec(t *testing.T) {
	resizeInstanceBySpecRequest := &ResizeInstanceBySpecRequest{
		InstanceId:       util.PtrString(""),
		Action:           util.PtrString(""),
		Spec:             util.PtrString(""),
		EnableJumboFrame: util.PtrBool(false),
	}
	err := BCC_CLIENT.ResizeInstanceBySpec(resizeInstanceBySpecRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeVolume(t *testing.T) {
	resizeVolumeRequest := &ResizeVolumeRequest{
		VolumeId:       util.PtrString(""),
		NewCdsSizeInGB: util.PtrInt32(int32(0)),
		NewExtraIO:     util.PtrInt32(int32(0)),
		NewVolumeType:  util.PtrString(""),
	}
	result := &ResizeVolumeResponse{}
	result, err := BCC_CLIENT.ResizeVolume(resizeVolumeRequest)
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
func TestClient_RollbackVolume(t *testing.T) {
	rollbackVolumeRequest := &RollbackVolumeRequest{
		VolumeId:   util.PtrString(""),
		SnapshotId: util.PtrString(""),
	}
	err := BCC_CLIENT.RollbackVolume(rollbackVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ShareImage(t *testing.T) {
	shareImageRequest := &ShareImageRequest{
		ImageId:   util.PtrString(""),
		Account:   util.PtrString(""),
		AccountId: util.PtrString(""),
		UcAccount: util.PtrString(""),
	}
	err := BCC_CLIENT.ShareImage(shareImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StartInstance(t *testing.T) {
	startInstanceRequest := &StartInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	err := BCC_CLIENT.StartInstance(startInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StopInstance(t *testing.T) {
	stopInstanceRequest := &StopInstanceRequest{
		InstanceId:       util.PtrString(""),
		ForceStop:        util.PtrBool(false),
		StopWithNoCharge: util.PtrBool(false),
	}
	err := BCC_CLIENT.StopInstance(stopInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnShareImage(t *testing.T) {
	unShareImageRequest := &UnShareImageRequest{
		ImageId:   util.PtrString(""),
		Account:   util.PtrString(""),
		AccountId: util.PtrString(""),
		UcAccount: util.PtrString(""),
	}
	err := BCC_CLIENT.UnShareImage(unShareImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindInstanceFromSecurityGroup(t *testing.T) {
	unbindInstanceFromSecurityGroupRequest := &UnbindInstanceFromSecurityGroupRequest{
		InstanceId:      util.PtrString(""),
		SecurityGroupId: util.PtrString(""),
	}
	err := BCC_CLIENT.UnbindInstanceFromSecurityGroup(unbindInstanceFromSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindInstanceFromTags(t *testing.T) {
	unbindInstanceFromTagsRequest := &UnbindInstanceFromTagsRequest{
		InstanceId:               util.PtrString(""),
		ChangeTags:               []*TagModel{},
		AttachRelatedResourceTag: util.PtrBool(false),
	}
	err := BCC_CLIENT.UnbindInstanceFromTags(unbindInstanceFromTagsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindRole(t *testing.T) {
	unbindRoleRequest := &UnbindRoleRequest{
		RoleName:  util.PtrString(""),
		Instances: []*InstancePassRoleModel{},
	}
	result := &UnbindRoleResponse{}
	result, err := BCC_CLIENT.UnbindRole(unbindRoleRequest)
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
func TestClient_UnbindTagImage(t *testing.T) {
	unbindTagImageRequest := &UnbindTagImageRequest{
		ImageId:    util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.UnbindTagImage(unbindTagImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindTagVolume(t *testing.T) {
	unbindTagVolumeRequest := &UnbindTagVolumeRequest{
		VolumeId:   util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.UnbindTagVolume(unbindTagVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceSubnet(t *testing.T) {
	updateInstanceSubnetRequest := &UpdateInstanceSubnetRequest{
		InstanceId:                 util.PtrString(""),
		SubnetId:                   util.PtrString(""),
		InternalIp:                 util.PtrString(""),
		SecurityGroupIds:           []*string{},
		EnterpriseSecurityGroupIds: []*string{},
	}
	err := BCC_CLIENT.UpdateInstanceSubnet(updateInstanceSubnetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
