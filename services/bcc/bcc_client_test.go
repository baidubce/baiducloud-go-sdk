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

func TestClient_AcceptReservedInstanceTransfer(t *testing.T) {
	acceptReservedInstanceTransferRequest := &AcceptReservedInstanceTransferRequest{
		TransferRecordId: util.PtrString(""),
		EhcClusterId:     util.PtrString(""),
	}
	err := BCC_CLIENT.AcceptReservedInstanceTransfer(acceptReservedInstanceTransferRequest)
	ExpectEqual(t.Errorf, nil, err)
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
func TestClient_AttachAsp(t *testing.T) {
	attachAspRequest := &AttachAspRequest{
		AspId:              util.PtrString(""),
		VolumeIds:          []*string{},
		DeleteAutoSnapshot: util.PtrBool(false),
	}
	err := BCC_CLIENT.AttachAsp(attachAspRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AttachKeypair(t *testing.T) {
	attachKeypairRequest := &AttachKeypairRequest{
		KeypairId:   util.PtrString(""),
		InstanceIds: []*string{},
	}
	err := BCC_CLIENT.AttachKeypair(attachKeypairRequest)
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
func TestClient_AuthorizeSecurityGroupRule(t *testing.T) {
	Rule := &SecurityGroupRuleModel{
		Remark:              util.PtrString(""),
		Direction:           util.PtrString(""),
		Ethertype:           util.PtrString(""),
		PortRange:           util.PtrString(""),
		Protocol:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		SourceIp:            util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		DestIp:              util.PtrString(""),
		SecurityGroupId:     util.PtrString(""),
		SecurityGroupRuleId: util.PtrString(""),
		CreatedTime:         util.PtrString(""),
		UpdatedTime:         util.PtrString(""),
	}
	authorizeSecurityGroupRuleRequest := &AuthorizeSecurityGroupRuleRequest{
		SecurityGroupId: util.PtrString(""),
		SgVersion:       util.PtrInt32(int32(0)),
		Rule:            Rule,
	}
	err := BCC_CLIENT.AuthorizeSecurityGroupRule(authorizeSecurityGroupRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AuthorizeServerEvent(t *testing.T) {
	authorizeServerEventRequest := &AuthorizeServerEventRequest{
		Action:                        util.PtrString(""),
		ServerEventId:                 util.PtrString(""),
		InstanceId:                    util.PtrString(""),
		AuthorizeMaintenanceOperation: util.PtrString(""),
		ExecuteTime:                   util.PtrString(""),
	}
	result := &AuthorizeServerEventResponse{}
	result, err := BCC_CLIENT.AuthorizeServerEvent(authorizeServerEventRequest)
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
func TestClient_AutoRenewReservedInstance(t *testing.T) {
	autoRenewReservedInstanceRequest := &AutoRenewReservedInstanceRequest{
		ReservedInstanceIds: []*string{},
		AutoRenewTimeUnit:   util.PtrString(""),
		AutoRenewTime:       util.PtrInt32(int32(0)),
	}
	result := &AutoRenewReservedInstanceResponse{}
	result, err := BCC_CLIENT.AutoRenewReservedInstance(autoRenewReservedInstanceRequest)
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
func TestClient_AutoRenewVolumeCluster(t *testing.T) {
	autoRenewVolumeClusterRequest := &AutoRenewVolumeClusterRequest{
		ClusterId:     util.PtrString(""),
		RenewTimeUnit: util.PtrString(""),
		RenewTime:     util.PtrInt32(int32(0)),
	}
	err := BCC_CLIENT.AutoRenewVolumeCluster(autoRenewVolumeClusterRequest)
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
func TestClient_BindInstanceSecurityGroup(t *testing.T) {
	bindInstanceSecurityGroupRequest := &BindInstanceSecurityGroupRequest{
		InstanceIds:       []*string{},
		SecurityGroupIds:  []*string{},
		SecurityGroupType: util.PtrString(""),
	}
	result := &BindInstanceSecurityGroupResponse{}
	result, err := BCC_CLIENT.BindInstanceSecurityGroup(bindInstanceSecurityGroupRequest)
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
func TestClient_BindReservedInstanceToTags(t *testing.T) {
	bindReservedInstanceToTagsRequest := &BindReservedInstanceToTagsRequest{
		ReservedInstanceIds: []*string{},
		ChangeTags:          []*TagModel{},
	}
	err := BCC_CLIENT.BindReservedInstanceToTags(bindReservedInstanceToTagsRequest)
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
func TestClient_BindTagSnapchain(t *testing.T) {
	bindTagSnapchainRequest := &BindTagSnapchainRequest{
		ChainId:    util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.BindTagSnapchain(bindTagSnapchainRequest)
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
func TestClient_BindTagVolumeCluster(t *testing.T) {
	bindTagVolumeClusterRequest := &BindTagVolumeClusterRequest{
		ClusterId:  util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.BindTagVolumeCluster(bindTagVolumeClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelAutoRenewReservedInstance(t *testing.T) {
	cancelAutoRenewReservedInstanceRequest := &CancelAutoRenewReservedInstanceRequest{
		ReservedInstanceIds: []*string{},
	}
	result := &CancelAutoRenewReservedInstanceResponse{}
	result, err := BCC_CLIENT.CancelAutoRenewReservedInstance(cancelAutoRenewReservedInstanceRequest)
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
func TestClient_CancelAutoRenewVolumeCluster(t *testing.T) {
	cancelAutoRenewVolumeClusterRequest := &CancelAutoRenewVolumeClusterRequest{
		ClusterId: util.PtrString(""),
	}
	err := BCC_CLIENT.CancelAutoRenewVolumeCluster(cancelAutoRenewVolumeClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelBidOrder(t *testing.T) {
	cancelBidOrderRequest := &CancelBidOrderRequest{
		OrderId: util.PtrString(""),
	}
	result := &CancelBidOrderResponse{}
	result, err := BCC_CLIENT.CancelBidOrder(cancelBidOrderRequest)
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
func TestClient_CancelRemoteCopyImage(t *testing.T) {
	cancelRemoteCopyImageRequest := &CancelRemoteCopyImageRequest{
		ImageId: util.PtrString(""),
	}
	err := BCC_CLIENT.CancelRemoteCopyImage(cancelRemoteCopyImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelSnapshotShare(t *testing.T) {
	cancelSnapshotShareRequest := &CancelSnapshotShareRequest{
		SourceSnapshotId: util.PtrString(""),
		AccountIds:       []*string{},
		ShareSnapshotId:  util.PtrString(""),
	}
	result := &CancelSnapshotShareResponse{}
	result, err := BCC_CLIENT.CancelSnapshotShare(cancelSnapshotShareRequest)
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
func TestClient_CheckServerEvent(t *testing.T) {
	checkServerEventRequest := &CheckServerEventRequest{
		Action:                        util.PtrString(""),
		ServerEventId:                 util.PtrString(""),
		InstanceId:                    util.PtrString(""),
		CheckResult:                   util.PtrString(""),
		IssueEffect:                   util.PtrString(""),
		IssueDescription:              util.PtrString(""),
		AuthorizeMaintenanceOperation: util.PtrString(""),
	}
	result := &CheckServerEventResponse{}
	result, err := BCC_CLIENT.CheckServerEvent(checkServerEventRequest)
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
func TestClient_CreateAsp(t *testing.T) {
	createAspRequest := &CreateAspRequest{
		Name:           util.PtrString(""),
		TimePoints:     []*int32{},
		RepeatWeekdays: []*int32{},
		RetentionDays:  util.PtrString(""),
	}
	result := &CreateAspResponse{}
	result, err := BCC_CLIENT.CreateAsp(createAspRequest)
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
func TestClient_CreateAuthorizationRule(t *testing.T) {
	createAuthorizationRuleRequest := &CreateAuthorizationRuleRequest{
		Action:                         util.PtrString(""),
		EnableRule:                     util.PtrInt32(int32(0)),
		AuthorizeMaintenanceOperations: []*string{},
		RuleName:                       util.PtrString(""),
		ServerEventCategory:            util.PtrString(""),
	}
	result := &CreateAuthorizationRuleResponse{}
	result, err := BCC_CLIENT.CreateAuthorizationRule(createAuthorizationRuleRequest)
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
func TestClient_CreateBidInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createBidInstanceRequest := &CreateBidInstanceRequest{
		Spec:                      util.PtrString(""),
		ImageId:                   util.PtrString(""),
		Billing:                   Billing,
		BidModel:                  util.PtrString(""),
		BidPrice:                  util.PtrString(""),
		CpuCount:                  util.PtrInt32(int32(0)),
		MemoryCapacityInGB:        util.PtrInt32(int32(0)),
		RootDiskSizeInGb:          util.PtrInt32(int32(0)),
		RootDiskStorageType:       util.PtrString(""),
		CreateCdsList:             []*CreateCdsModel{},
		EphemeralDisks:            []*EphemeralDisk{},
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
		Tags:                      []*TagModel{},
		AspId:                     util.PtrString(""),
		FileSystems:               []*FileSystemModel{},
		IsEipAutoRelatedDelete:    util.PtrBool(false),
		ResGroupId:                util.PtrString(""),
	}
	result := &CreateBidInstanceResponse{}
	result, err := BCC_CLIENT.CreateBidInstance(createBidInstanceRequest)
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
func TestClient_CreateDeploySet(t *testing.T) {
	createDeploySetRequest := &CreateDeploySetRequest{
		Name:        util.PtrString(""),
		Desc:        util.PtrString(""),
		Strategy:    util.PtrString(""),
		Concurrency: util.PtrInt32(int32(0)),
	}
	result := &CreateDeploySetResponse{}
	result, err := BCC_CLIENT.CreateDeploySet(createDeploySetRequest)
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
func TestClient_CreateEhcCluster(t *testing.T) {
	createEhcClusterRequest := &CreateEhcClusterRequest{
		Name:        util.PtrString(""),
		ZoneName:    util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateEhcClusterResponse{}
	result, err := BCC_CLIENT.CreateEhcCluster(createEhcClusterRequest)
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
func TestClient_CreateKeypair(t *testing.T) {
	createKeypairRequest := &CreateKeypairRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateKeypairResponse{}
	result, err := BCC_CLIENT.CreateKeypair(createKeypairRequest)
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
func TestClient_CreateReservedInstanceTransfer(t *testing.T) {
	createReservedInstanceTransferRequest := &CreateReservedInstanceTransferRequest{
		ReservedInstanceIds: []*string{},
		RecipientAccountId:  util.PtrString(""),
	}
	err := BCC_CLIENT.CreateReservedInstanceTransfer(createReservedInstanceTransferRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateReservedInstances(t *testing.T) {
	createReservedInstancesRequest := &CreateReservedInstancesRequest{
		ReservedInstanceName:     util.PtrString(""),
		Scope:                    util.PtrString(""),
		ZoneName:                 util.PtrString(""),
		Spec:                     util.PtrString(""),
		OfferingType:             util.PtrString(""),
		InstanceCount:            util.PtrInt32(int32(0)),
		ReservedInstanceCount:    util.PtrInt32(int32(0)),
		ReservedInstanceTime:     util.PtrInt32(int32(0)),
		ReservedInstanceTimeUnit: util.PtrString(""),
		AutoRenew:                util.PtrBool(false),
		AutoRenewTimeUnit:        util.PtrString(""),
		AutoRenewTime:            util.PtrInt32(int32(0)),
		EffectiveTime:            util.PtrString(""),
		Tags:                     []*TagModel{},
		TicketId:                 util.PtrString(""),
		EhcClusterId:             util.PtrString(""),
	}
	result := &CreateReservedInstancesResponse{}
	result, err := BCC_CLIENT.CreateReservedInstances(createReservedInstancesRequest)
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
func TestClient_CreateSecurityGroup(t *testing.T) {
	createSecurityGroupRequest := &CreateSecurityGroupRequest{
		Name:  util.PtrString(""),
		Desc:  util.PtrString(""),
		VpcId: util.PtrString(""),
		Tags:  []*Tag{},
		Rules: []*SecurityGroupRuleModel{},
	}
	result := &CreateSecurityGroupResponse{}
	result, err := BCC_CLIENT.CreateSecurityGroup(createSecurityGroupRequest)
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
func TestClient_CreateSnapshot(t *testing.T) {
	createSnapshotRequest := &CreateSnapshotRequest{
		VolumeId:        util.PtrString(""),
		SnapshotName:    util.PtrString(""),
		Desc:            util.PtrString(""),
		Tags:            []*TagModel{},
		RetentionInDays: util.PtrInt32(int32(0)),
	}
	result := &CreateSnapshotResponse{}
	result, err := BCC_CLIENT.CreateSnapshot(createSnapshotRequest)
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
func TestClient_CreateSnapshotShare(t *testing.T) {
	createSnapshotShareRequest := &CreateSnapshotShareRequest{
		SnapshotId: util.PtrString(""),
		AccountIds: []*string{},
	}
	result := &CreateSnapshotShareResponse{}
	result, err := BCC_CLIENT.CreateSnapshotShare(createSnapshotShareRequest)
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
func TestClient_CreateVolumeCluster(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createVolumeClusterRequest := &CreateVolumeClusterRequest{
		ZoneName:        util.PtrString(""),
		ClusterName:     util.PtrString(""),
		ClusterSizeInGB: util.PtrInt32(int32(0)),
		StorageType:     util.PtrString(""),
		PurchaseCount:   util.PtrInt32(int32(0)),
		Billing:         Billing,
		Tags:            []*TagModel{},
	}
	result := &CreateVolumeClusterResponse{}
	result, err := BCC_CLIENT.CreateVolumeCluster(createVolumeClusterRequest)
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
func TestClient_DeleteAsp(t *testing.T) {
	deleteAspRequest := &DeleteAspRequest{
		AspId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteAsp(deleteAspRequest)
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
func TestClient_DeleteDeploySet(t *testing.T) {
	deleteDeploySetRequest := &DeleteDeploySetRequest{
		DeployId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteDeploySet(deleteDeploySetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteEhcCluster(t *testing.T) {
	deleteEhcClusterRequest := &DeleteEhcClusterRequest{
		EhcClusterIdList: []*string{},
	}
	err := BCC_CLIENT.DeleteEhcCluster(deleteEhcClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteImage(t *testing.T) {
	deleteImageRequest := &DeleteImageRequest{
		ImageId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteImage(deleteImageRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstUserOpAuthorizeRule(t *testing.T) {
	deleteInstUserOpAuthorizeRuleRequest := &DeleteInstUserOpAuthorizeRuleRequest{
		Action: util.PtrString(""),
		RuleId: util.PtrString(""),
	}
	result := &DeleteInstUserOpAuthorizeRuleResponse{}
	result, err := BCC_CLIENT.DeleteInstUserOpAuthorizeRule(deleteInstUserOpAuthorizeRuleRequest)
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
func TestClient_DeleteInstanceDeploySet(t *testing.T) {
	deleteInstanceDeploySetRequest := &DeleteInstanceDeploySetRequest{
		DeployId:       util.PtrString(""),
		InstanceIdList: []*string{},
	}
	err := BCC_CLIENT.DeleteInstanceDeploySet(deleteInstanceDeploySetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteKeypair(t *testing.T) {
	deleteKeypairRequest := &DeleteKeypairRequest{
		KeypairId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteKeypair(deleteKeypairRequest)
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
func TestClient_DeleteSecurityGroup(t *testing.T) {
	deleteSecurityGroupRequest := &DeleteSecurityGroupRequest{
		SecurityGroupId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteSecurityGroup(deleteSecurityGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSecurityGroupRule(t *testing.T) {
	deleteSecurityGroupRuleRequest := &DeleteSecurityGroupRuleRequest{
		SecurityGroupRuleId: util.PtrString(""),
		SgVersion:           util.PtrInt32(int32(0)),
	}
	err := BCC_CLIENT.DeleteSecurityGroupRule(deleteSecurityGroupRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSnapshot(t *testing.T) {
	deleteSnapshotRequest := &DeleteSnapshotRequest{
		SnapshotId: util.PtrString(""),
	}
	err := BCC_CLIENT.DeleteSnapshot(deleteSnapshotRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletesInstanceDeploySet(t *testing.T) {
	deletesInstanceDeploySetRequest := &DeletesInstanceDeploySetRequest{
		DeployId:       util.PtrString(""),
		InstanceIdList: []*string{},
	}
	err := BCC_CLIENT.DeletesInstanceDeploySet(deletesInstanceDeploySetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeAuthorizeRules(t *testing.T) {
	describeAuthorizeRulesRequest := &DescribeAuthorizeRulesRequest{
		Action:    util.PtrString(""),
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		RuleIds:   []*string{},
		RuleNames: []*string{},
	}
	result := &DescribeAuthorizeRulesResponse{}
	result, err := BCC_CLIENT.DescribeAuthorizeRules(describeAuthorizeRulesRequest)
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
func TestClient_DescribePlannedEventRecords(t *testing.T) {
	describePlannedEventRecordsRequest := &DescribePlannedEventRecordsRequest{
		Action:                   util.PtrString(""),
		ServerEventIds:           []*string{},
		InstanceIds:              []*string{},
		ProductCategory:          util.PtrString(""),
		ServerEventType:          util.PtrString(""),
		ServerEventLogTimeFilter: util.PtrString(""),
		PeriodStartTime:          util.PtrString(""),
		PeriodEndTime:            util.PtrString(""),
		MaxKeys:                  util.PtrInt32(int32(0)),
		Marker:                   util.PtrString(""),
	}
	result := &DescribePlannedEventRecordsResponse{}
	result, err := BCC_CLIENT.DescribePlannedEventRecords(describePlannedEventRecordsRequest)
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
func TestClient_DescribePlannedEvents(t *testing.T) {
	describePlannedEventsRequest := &DescribePlannedEventsRequest{
		Action:                   util.PtrString(""),
		ServerEventStatus:        util.PtrString(""),
		ServerEventIds:           []*string{},
		InstanceIds:              []*string{},
		ProductCategory:          util.PtrString(""),
		ServerEventType:          util.PtrString(""),
		ServerEventLogTimeFilter: util.PtrString(""),
		PeriodStartTime:          util.PtrString(""),
		PeriodEndTime:            util.PtrString(""),
		MaxKeys:                  util.PtrInt32(int32(0)),
		Marker:                   util.PtrString(""),
	}
	result := &DescribePlannedEventsResponse{}
	result, err := BCC_CLIENT.DescribePlannedEvents(describePlannedEventsRequest)
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
func TestClient_DescribeRegions(t *testing.T) {
	describeRegionsRequest := &DescribeRegionsRequest{
		Region: util.PtrString(""),
	}
	result := &DescribeRegionsResponse{}
	result, err := BCC_CLIENT.DescribeRegions(describeRegionsRequest)
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
func TestClient_DescribeUnplannedEventRecords(t *testing.T) {
	describeUnplannedEventRecordsRequest := &DescribeUnplannedEventRecordsRequest{
		Action:                   util.PtrString(""),
		ServerEventIds:           []*string{},
		InstanceIds:              []*string{},
		ProductCategory:          util.PtrString(""),
		ServerEventType:          util.PtrString(""),
		ServerEventLogTimeFilter: util.PtrString(""),
		PeriodStartTime:          util.PtrString(""),
		PeriodEndTime:            util.PtrString(""),
		MaxKeys:                  util.PtrInt32(int32(0)),
		Marker:                   util.PtrString(""),
	}
	result := &DescribeUnplannedEventRecordsResponse{}
	result, err := BCC_CLIENT.DescribeUnplannedEventRecords(describeUnplannedEventRecordsRequest)
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
func TestClient_DescribeUnplannedEvents(t *testing.T) {
	describeUnplannedEventsRequest := &DescribeUnplannedEventsRequest{
		Action:                   util.PtrString(""),
		ServerEventStatus:        util.PtrString(""),
		ServerEventIds:           []*string{},
		InstanceIds:              []*string{},
		ProductCategory:          util.PtrString(""),
		ServerEventType:          util.PtrString(""),
		ServerEventLogTimeFilter: util.PtrString(""),
		PeriodStartTime:          util.PtrString(""),
		PeriodEndTime:            util.PtrString(""),
		MaxKeys:                  util.PtrInt32(int32(0)),
		Marker:                   util.PtrString(""),
	}
	result := &DescribeUnplannedEventsResponse{}
	result, err := BCC_CLIENT.DescribeUnplannedEvents(describeUnplannedEventsRequest)
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
func TestClient_DetachAsp(t *testing.T) {
	detachAspRequest := &DetachAspRequest{
		AspId:              util.PtrString(""),
		VolumeIds:          []*string{},
		DeleteAutoSnapshot: util.PtrBool(false),
	}
	err := BCC_CLIENT.DetachAsp(detachAspRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DetachKeypair(t *testing.T) {
	detachKeypairRequest := &DetachKeypairRequest{
		KeypairId:   util.PtrString(""),
		InstanceIds: []*string{},
	}
	err := BCC_CLIENT.DetachKeypair(detachKeypairRequest)
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
func TestClient_EhcClusterList(t *testing.T) {
	ehcClusterListRequest := &EhcClusterListRequest{
		EhcClusterIdList: []*string{},
		NameList:         []*string{},
		ZoneName:         util.PtrString(""),
	}
	result := &EhcClusterListResponse{}
	result, err := BCC_CLIENT.EhcClusterList(ehcClusterListRequest)
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
func TestClient_GetAsp(t *testing.T) {
	getAspRequest := &GetAspRequest{
		AspId: util.PtrString(""),
	}
	result := &GetAspResponse{}
	result, err := BCC_CLIENT.GetAsp(getAspRequest)
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
func TestClient_GetBidInstancePrice(t *testing.T) {
	getBidInstancePriceRequest := &GetBidInstancePriceRequest{
		Spec:                  util.PtrString(""),
		RootDiskSizeInGb:      util.PtrInt32(int32(0)),
		RootDiskStorageType:   util.PtrString(""),
		CreateCdsList:         []*CreateCdsModel{},
		NetworkCapacityInMbps: util.PtrInt32(int32(0)),
		InternetChargeType:    util.PtrString(""),
		PurchaseCount:         util.PtrInt32(int32(0)),
		ZoneName:              util.PtrString(""),
	}
	result := &GetBidInstancePriceResponse{}
	result, err := BCC_CLIENT.GetBidInstancePrice(getBidInstancePriceRequest)
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
func TestClient_GetDeploySet(t *testing.T) {
	getDeploySetRequest := &GetDeploySetRequest{
		Id: util.PtrString(""),
	}
	result := &GetDeploySetResponse{}
	result, err := BCC_CLIENT.GetDeploySet(getDeploySetRequest)
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
func TestClient_GetPriceBySpec(t *testing.T) {
	getPriceBySpecRequest := &GetPriceBySpecRequest{
		SpecId:         util.PtrString(""),
		Spec:           util.PtrString(""),
		PaymentTiming:  util.PtrString(""),
		ZoneName:       util.PtrString(""),
		PurchaseCount:  util.PtrInt32(int32(0)),
		PurchaseLength: util.PtrInt32(int32(0)),
	}
	result := &GetPriceBySpecResponse{}
	result, err := BCC_CLIENT.GetPriceBySpec(getPriceBySpecRequest)
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
func TestClient_GetReservedInstance(t *testing.T) {
	getReservedInstanceRequest := &GetReservedInstanceRequest{
		Marker:                 util.PtrString(""),
		MaxKeys:                util.PtrInt32(int32(0)),
		ReservedInstanceIds:    []*string{},
		ReservedInstanceName:   util.PtrString(""),
		ZoneName:               util.PtrString(""),
		ReservedInstanceStatus: util.PtrString(""),
		Spec:                   util.PtrString(""),
		OfferingType:           util.PtrString(""),
		OsType:                 util.PtrString(""),
		InstanceId:             util.PtrString(""),
		InstanceName:           util.PtrString(""),
		IsDeduct:               util.PtrBool(false),
		EhcClusterId:           util.PtrString(""),
		SortKey:                util.PtrString(""),
		SortDir:                util.PtrString(""),
		ReservedInstanceSource: util.PtrString(""),
		Scope:                  util.PtrString(""),
	}
	result := &GetReservedInstanceResponse{}
	result, err := BCC_CLIENT.GetReservedInstance(getReservedInstanceRequest)
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
func TestClient_GetReservedInstancePrice(t *testing.T) {
	getReservedInstancePriceRequest := &GetReservedInstancePriceRequest{
		SpecId:                util.PtrString(""),
		Spec:                  util.PtrString(""),
		OfferingType:          util.PtrString(""),
		Scope:                 util.PtrString(""),
		ZoneName:              util.PtrString(""),
		ReservedInstanceCount: util.PtrInt32(int32(0)),
		PriceTimeUnit:         util.PtrString(""),
		ReservedInstanceTime:  util.PtrInt32(int32(0)),
		PurchaseNum:           util.PtrInt32(int32(0)),
	}
	result := &GetReservedInstancePriceResponse{}
	result, err := BCC_CLIENT.GetReservedInstancePrice(getReservedInstancePriceRequest)
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
func TestClient_GetSnapshot(t *testing.T) {
	getSnapshotRequest := &GetSnapshotRequest{
		SnapshotId: util.PtrString(""),
	}
	result := &GetSnapshotResponse{}
	result, err := BCC_CLIENT.GetSnapshot(getSnapshotRequest)
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
func TestClient_GetTask(t *testing.T) {
	getTaskRequest := &GetTaskRequest{
		TaskIds: []*string{},
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &GetTaskResponse{}
	result, err := BCC_CLIENT.GetTask(getTaskRequest)
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
func TestClient_GetVolumeCluster(t *testing.T) {
	getVolumeClusterRequest := &GetVolumeClusterRequest{
		ClusterId: util.PtrString(""),
	}
	result := &GetVolumeClusterResponse{}
	result, err := BCC_CLIENT.GetVolumeCluster(getVolumeClusterRequest)
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
func TestClient_GetZoneBySpec(t *testing.T) {
	getZoneBySpecRequest := &GetZoneBySpecRequest{
		InstanceType: util.PtrString(""),
		ProductType:  util.PtrString(""),
		Spec:         util.PtrString(""),
		SpecId:       util.PtrString(""),
	}
	result := &GetZoneBySpecResponse{}
	result, err := BCC_CLIENT.GetZoneBySpec(getZoneBySpecRequest)
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
func TestClient_ImportKeypair(t *testing.T) {
	importKeypairRequest := &ImportKeypairRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		PublicKey:   util.PtrString(""),
	}
	result := &ImportKeypairResponse{}
	result, err := BCC_CLIENT.ImportKeypair(importKeypairRequest)
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
func TestClient_KeypairDetail(t *testing.T) {
	keypairDetailRequest := &KeypairDetailRequest{
		KeypairId: util.PtrString(""),
	}
	result := &KeypairDetailResponse{}
	result, err := BCC_CLIENT.KeypairDetail(keypairDetailRequest)
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
func TestClient_ListAsps(t *testing.T) {
	listAspsRequest := &ListAspsRequest{
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		AspName:    util.PtrString(""),
		VolumeName: util.PtrString(""),
	}
	result := &ListAspsResponse{}
	result, err := BCC_CLIENT.ListAsps(listAspsRequest)
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
func TestClient_ListBidFlavor(t *testing.T) {
	result := &ListBidFlavorResponse{}
	result, err := BCC_CLIENT.ListBidFlavor()
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
func TestClient_ListDeploySet(t *testing.T) {
	result := &ListDeploySetResponse{}
	result, err := BCC_CLIENT.ListDeploySet()
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
func TestClient_ListFlavorSpec(t *testing.T) {
	listFlavorSpecRequest := &ListFlavorSpecRequest{
		ZoneName: util.PtrString(""),
		Specs:    util.PtrString(""),
		SpecIds:  util.PtrString(""),
	}
	result := &ListFlavorSpecResponse{}
	result, err := BCC_CLIENT.ListFlavorSpec(listFlavorSpecRequest)
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
func TestClient_ListKeypair(t *testing.T) {
	listKeypairRequest := &ListKeypairRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListKeypairResponse{}
	result, err := BCC_CLIENT.ListKeypair(listKeypairRequest)
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
func TestClient_ListReservedInstanceTransferIn(t *testing.T) {
	listReservedInstanceTransferInRequest := &ListReservedInstanceTransferInRequest{
		Marker:              util.PtrString(""),
		MaxKeys:             util.PtrInt32(int32(0)),
		ReservedInstanceIds: []*string{},
		TransferRecordIds:   []*string{},
		Spec:                util.PtrString(""),
		Status:              util.PtrString(""),
	}
	result := &ListReservedInstanceTransferInResponse{}
	result, err := BCC_CLIENT.ListReservedInstanceTransferIn(listReservedInstanceTransferInRequest)
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
func TestClient_ListReservedInstanceTransferOut(t *testing.T) {
	listReservedInstanceTransferOutRequest := &ListReservedInstanceTransferOutRequest{
		Marker:              util.PtrString(""),
		MaxKeys:             util.PtrInt32(int32(0)),
		ReservedInstanceIds: []*string{},
		TransferRecordIds:   []*string{},
		Spec:                util.PtrString(""),
		Status:              util.PtrString(""),
	}
	result := &ListReservedInstanceTransferOutResponse{}
	result, err := BCC_CLIENT.ListReservedInstanceTransferOut(listReservedInstanceTransferOutRequest)
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
func TestClient_ListSecurityGroups(t *testing.T) {
	listSecurityGroupsRequest := &ListSecurityGroupsRequest{
		Marker:           util.PtrString(""),
		MaxKeys:          util.PtrInt32(int32(0)),
		InstanceId:       util.PtrString(""),
		VpcId:            util.PtrString(""),
		SecurityGroupId:  util.PtrString(""),
		SecurityGroupIds: util.PtrString(""),
	}
	result := &ListSecurityGroupsResponse{}
	result, err := BCC_CLIENT.ListSecurityGroups(listSecurityGroupsRequest)
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
func TestClient_ListSnapchain(t *testing.T) {
	listSnapchainRequest := &ListSnapchainRequest{
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		PageSize: util.PtrInt32(int32(0)),
		PageNo:   util.PtrInt32(int32(0)),
		VolumeId: util.PtrString(""),
	}
	result := &ListSnapchainResponse{}
	result, err := BCC_CLIENT.ListSnapchain(listSnapchainRequest)
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
func TestClient_ListSnapshotShare(t *testing.T) {
	listSnapshotShareRequest := &ListSnapshotShareRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &ListSnapshotShareResponse{}
	result, err := BCC_CLIENT.ListSnapshotShare(listSnapshotShareRequest)
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
func TestClient_ListSnapshots(t *testing.T) {
	listSnapshotsRequest := &ListSnapshotsRequest{
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
		VolumeId: util.PtrString(""),
	}
	result := &ListSnapshotsResponse{}
	result, err := BCC_CLIENT.ListSnapshots(listSnapshotsRequest)
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
func TestClient_ListTask(t *testing.T) {
	listTaskRequest := &ListTaskRequest{
		TaskIds:     []*string{},
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		TaskAction:  util.PtrString(""),
		TaskStatus:  util.PtrString(""),
		StartTime:   util.PtrString(""),
		EndTime:     util.PtrString(""),
		ResourceIds: []*string{},
	}
	result := &ListTaskResponse{}
	result, err := BCC_CLIENT.ListTask(listTaskRequest)
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
func TestClient_ListVolumeClusters(t *testing.T) {
	listVolumeClustersRequest := &ListVolumeClustersRequest{
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
		ZoneName:    util.PtrString(""),
		ClusterName: util.PtrString(""),
	}
	result := &ListVolumeClustersResponse{}
	result, err := BCC_CLIENT.ListVolumeClusters(listVolumeClustersRequest)
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
func TestClient_ListZones(t *testing.T) {
	result := &ListZonesResponse{}
	result, err := BCC_CLIENT.ListZones()
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
func TestClient_ModifyEhcCluster(t *testing.T) {
	modifyEhcClusterRequest := &ModifyEhcClusterRequest{
		EhcClusterId: util.PtrString(""),
		Name:         util.PtrString(""),
		Description:  util.PtrString(""),
	}
	err := BCC_CLIENT.ModifyEhcCluster(modifyEhcClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstUserOpAuthorizeRuleAttribute(t *testing.T) {
	modifyInstUserOpAuthorizeRuleAttributeRequest := &ModifyInstUserOpAuthorizeRuleAttributeRequest{
		Action:                         util.PtrString(""),
		EnableRule:                     util.PtrInt32(int32(0)),
		AuthorizeMaintenanceOperations: []*string{},
		RuleName:                       util.PtrString(""),
		RuleId:                         util.PtrString(""),
	}
	result := &ModifyInstUserOpAuthorizeRuleAttributeResponse{}
	result, err := BCC_CLIENT.ModifyInstUserOpAuthorizeRuleAttribute(modifyInstUserOpAuthorizeRuleAttributeRequest)
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
func TestClient_ModifyReservedInstances(t *testing.T) {
	modifyReservedInstancesRequest := &ModifyReservedInstancesRequest{
		ReservedInstances: []*ReservedInstance{},
	}
	result := &ModifyReservedInstancesResponse{}
	result, err := BCC_CLIENT.ModifyReservedInstances(modifyReservedInstancesRequest)
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
func TestClient_PurchaseReservedVolumeCluster(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	purchaseReservedVolumeClusterRequest := &PurchaseReservedVolumeClusterRequest{
		ClusterId: util.PtrString(""),
		Billing:   Billing,
	}
	result := &PurchaseReservedVolumeClusterResponse{}
	result, err := BCC_CLIENT.PurchaseReservedVolumeCluster(purchaseReservedVolumeClusterRequest)
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
func TestClient_RefuseReservedInstanceTransfer(t *testing.T) {
	refuseReservedInstanceTransferRequest := &RefuseReservedInstanceTransferRequest{
		TransferRecordIds: []*string{},
	}
	err := BCC_CLIENT.RefuseReservedInstanceTransfer(refuseReservedInstanceTransferRequest)
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
func TestClient_RemoteCopySnapshot(t *testing.T) {
	remoteCopySnapshotRequest := &RemoteCopySnapshotRequest{
		SnapshotId:      util.PtrString(""),
		Uuid:            util.PtrString(""),
		DestRegionInfos: []*RemoteCopyRequest{},
	}
	result := &RemoteCopySnapshotResponse{}
	result, err := BCC_CLIENT.RemoteCopySnapshot(remoteCopySnapshotRequest)
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
func TestClient_RenameKeypair(t *testing.T) {
	renameKeypairRequest := &RenameKeypairRequest{
		KeypairId: util.PtrString(""),
		Name:      util.PtrString(""),
	}
	err := BCC_CLIENT.RenameKeypair(renameKeypairRequest)
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
func TestClient_RenewReservedInstance(t *testing.T) {
	renewReservedInstanceRequest := &RenewReservedInstanceRequest{
		ReservedInstanceIds:      []*string{},
		ReservedInstanceTime:     util.PtrString(""),
		ReservedInstanceTimeUnit: util.PtrString(""),
		AutoRenew:                util.PtrBool(false),
		AutoRenewTimeUnit:        util.PtrString(""),
		AutoRenewTime:            util.PtrInt32(int32(0)),
	}
	result := &RenewReservedInstanceResponse{}
	result, err := BCC_CLIENT.RenewReservedInstance(renewReservedInstanceRequest)
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
func TestClient_ReplaceInstanceSecurityGroup(t *testing.T) {
	replaceInstanceSecurityGroupRequest := &ReplaceInstanceSecurityGroupRequest{
		InstanceIds:       []*string{},
		SecurityGroupIds:  []*string{},
		SecurityGroupType: util.PtrString(""),
	}
	result := &ReplaceInstanceSecurityGroupResponse{}
	result, err := BCC_CLIENT.ReplaceInstanceSecurityGroup(replaceInstanceSecurityGroupRequest)
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
func TestClient_ResizeVolumeCluster(t *testing.T) {
	resizeVolumeClusterRequest := &ResizeVolumeClusterRequest{
		ClusterId:          util.PtrString(""),
		NewClusterSizeInGB: util.PtrInt32(int32(0)),
	}
	result := &ResizeVolumeClusterResponse{}
	result, err := BCC_CLIENT.ResizeVolumeCluster(resizeVolumeClusterRequest)
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
func TestClient_RevokeReservedInstanceTransfer(t *testing.T) {
	revokeReservedInstanceTransferRequest := &RevokeReservedInstanceTransferRequest{
		TransferRecordIds: []*string{},
	}
	err := BCC_CLIENT.RevokeReservedInstanceTransfer(revokeReservedInstanceTransferRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RevokeSecurityGroupRule(t *testing.T) {
	Rule := &SecurityGroupRuleModel{
		Remark:              util.PtrString(""),
		Direction:           util.PtrString(""),
		Ethertype:           util.PtrString(""),
		PortRange:           util.PtrString(""),
		Protocol:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		SourceIp:            util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		DestIp:              util.PtrString(""),
		SecurityGroupId:     util.PtrString(""),
		SecurityGroupRuleId: util.PtrString(""),
		CreatedTime:         util.PtrString(""),
		UpdatedTime:         util.PtrString(""),
	}
	revokeSecurityGroupRuleRequest := &RevokeSecurityGroupRuleRequest{
		SecurityGroupId: util.PtrString(""),
		SgVersion:       util.PtrInt32(int32(0)),
		Rule:            Rule,
	}
	err := BCC_CLIENT.RevokeSecurityGroupRule(revokeSecurityGroupRuleRequest)
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
func TestClient_UnbindInstanceSecurityGroup(t *testing.T) {
	unbindInstanceSecurityGroupRequest := &UnbindInstanceSecurityGroupRequest{
		InstanceIds:       []*string{},
		SecurityGroupIds:  []*string{},
		SecurityGroupType: util.PtrString(""),
	}
	result := &UnbindInstanceSecurityGroupResponse{}
	result, err := BCC_CLIENT.UnbindInstanceSecurityGroup(unbindInstanceSecurityGroupRequest)
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
func TestClient_UnbindReservedInstanceFromTags(t *testing.T) {
	unbindReservedInstanceFromTagsRequest := &UnbindReservedInstanceFromTagsRequest{
		ReservedInstanceIds: []*string{},
		ChangeTags:          []*TagModel{},
	}
	err := BCC_CLIENT.UnbindReservedInstanceFromTags(unbindReservedInstanceFromTagsRequest)
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
func TestClient_UnbindTagSnapchain(t *testing.T) {
	unbindTagSnapchainRequest := &UnbindTagSnapchainRequest{
		ChainId:    util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.UnbindTagSnapchain(unbindTagSnapchainRequest)
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
func TestClient_UnbindTagVolumeCluster(t *testing.T) {
	unbindTagVolumeClusterRequest := &UnbindTagVolumeClusterRequest{
		ClusterId:  util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.UnbindTagVolumeCluster(unbindTagVolumeClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAsp(t *testing.T) {
	updateAspRequest := &UpdateAspRequest{
		AspId:          util.PtrString(""),
		Name:           util.PtrString(""),
		TimePoints:     []*string{},
		RepeatWeekdays: []*string{},
		RetentionDays:  util.PtrString(""),
	}
	err := BCC_CLIENT.UpdateAsp(updateAspRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDeploySet(t *testing.T) {
	updateDeploySetRequest := &UpdateDeploySetRequest{
		DeployId: util.PtrString(""),
		Name:     util.PtrString(""),
		Desc:     util.PtrString(""),
	}
	err := BCC_CLIENT.UpdateDeploySet(updateDeploySetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateDeploySetRelation(t *testing.T) {
	updateDeploySetRelationRequest := &UpdateDeploySetRelationRequest{
		InstanceId:      util.PtrString(""),
		DeploysetIdList: []*string{},
	}
	result := &UpdateDeploySetRelationResponse{}
	result, err := BCC_CLIENT.UpdateDeploySetRelation(updateDeploySetRelationRequest)
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
func TestClient_UpdateKeypairDescription(t *testing.T) {
	updateKeypairDescriptionRequest := &UpdateKeypairDescriptionRequest{
		KeypairId:   util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := BCC_CLIENT.UpdateKeypairDescription(updateKeypairDescriptionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSecurityGroupRule(t *testing.T) {
	updateSecurityGroupRuleRequest := &UpdateSecurityGroupRuleRequest{
		SgVersion:           util.PtrInt32(int32(0)),
		SecurityGroupRuleId: util.PtrString(""),
		Remark:              util.PtrString(""),
		PortRange:           util.PtrString(""),
		SourceIp:            util.PtrString(""),
		SourceGroupId:       util.PtrString(""),
		DestIp:              util.PtrString(""),
		DestGroupId:         util.PtrString(""),
		Protocol:            util.PtrString(""),
	}
	err := BCC_CLIENT.UpdateSecurityGroupRule(updateSecurityGroupRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
