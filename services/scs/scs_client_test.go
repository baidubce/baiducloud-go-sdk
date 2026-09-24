package scs

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
	SCS_CLIENT *Client
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
	SCS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_AccountList(t *testing.T) {
	accountListRequest := &AccountListRequest{
		InstanceId: util.PtrString(""),
	}
	result := &AccountListResponse{}
	result, err := SCS_CLIENT.AccountList(accountListRequest)
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
func TestClient_AddIpWhitelist(t *testing.T) {
	addIpWhitelistRequest := &AddIpWhitelistRequest{
		InstanceId:  util.PtrString(""),
		SecurityIps: []*string{},
	}
	err := SCS_CLIENT.AddIpWhitelist(addIpWhitelistRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddParametersToParameterTemplate(t *testing.T) {
	addParametersToParameterTemplateRequest := &AddParametersToParameterTemplateRequest{
		TemplateShowId: util.PtrString(""),
		Parameters:     []*Parameters{},
	}
	err := SCS_CLIENT.AddParametersToParameterTemplate(addParametersToParameterTemplateRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ApplicationParameterTemplate(t *testing.T) {
	applicationParameterTemplateRequest := &ApplicationParameterTemplateRequest{
		TemplateShowId:   util.PtrString(""),
		Extra:            util.PtrString(""),
		CacheClusterList: []*CacheClusterShowIdItem{},
		RebootType:       util.PtrInt32(int32(0)),
		Parameters:       []*Parameters{},
	}
	err := SCS_CLIENT.ApplicationParameterTemplate(applicationParameterTemplateRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AuditLogSwitch(t *testing.T) {
	auditLogSwitchRequest := &AuditLogSwitchRequest{
		InstanceId: util.PtrString(""),
		Action:     util.PtrString(""),
	}
	err := SCS_CLIENT.AuditLogSwitch(auditLogSwitchRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchRestoreInstances(t *testing.T) {
	batchRestoreInstancesRequest := &BatchRestoreInstancesRequest{
		InstanceIds: []*string{},
	}
	err := SCS_CLIENT.BatchRestoreInstances(batchRestoreInstancesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindSecurityGroup(t *testing.T) {
	bindSecurityGroupRequest := &BindSecurityGroupRequest{
		InstanceId:       util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	result := &BindSecurityGroupResponse{}
	result, err := SCS_CLIENT.BindSecurityGroup(bindSecurityGroupRequest)
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
func TestClient_BindTags(t *testing.T) {
	bindTagsRequest := &BindTagsRequest{
		InstanceId: util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := SCS_CLIENT.BindTags(bindTagsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CancelPrepaidToPostpaid(t *testing.T) {
	cancelPrepaidToPostpaidRequest := &CancelPrepaidToPostpaidRequest{
		InstanceIds: []*string{},
	}
	err := SCS_CLIENT.CancelPrepaidToPostpaid(cancelPrepaidToPostpaidRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ChangeAccessPassword(t *testing.T) {
	changeAccessPasswordRequest := &ChangeAccessPasswordRequest{
		InstanceId: util.PtrString(""),
		Password:   util.PtrString(""),
	}
	err := SCS_CLIENT.ChangeAccessPassword(changeAccessPasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ChangeAccountPassword(t *testing.T) {
	changeAccountPasswordRequest := &ChangeAccountPasswordRequest{
		InstanceId: util.PtrString(""),
		UserName:   util.PtrString(""),
		ClientAuth: util.PtrString(""),
	}
	err := SCS_CLIENT.ChangeAccountPassword(changeAccountPasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ChangeConfiguration(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	changeConfigurationRequest := &ChangeConfigurationRequest{
		InstanceId:    util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Billing:       Billing,
		EngineVersion: util.PtrString(""),
		NodeType:      util.PtrString(""),
		ShardNum:      util.PtrInt32(int32(0)),
		DiskFlavor:    util.PtrInt32(int32(0)),
	}
	result := &ChangeConfigurationResponse{}
	result, err := SCS_CLIENT.ChangeConfiguration(changeConfigurationRequest)
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
func TestClient_ClearInstance(t *testing.T) {
	clearInstanceRequest := &ClearInstanceRequest{
		InstanceId:     util.PtrString(""),
		Password:       util.PtrString(""),
		DbIndex:        util.PtrInt32(int32(0)),
		IsFlushExpired: util.PtrBool(false),
		IsDefer:        util.PtrBool(false),
	}
	err := SCS_CLIENT.ClearInstance(clearInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ClusterStatusCheck(t *testing.T) {
	clusterStatusCheckRequest := &ClusterStatusCheckRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ClusterStatusCheckResponse{}
	result, err := SCS_CLIENT.ClusterStatusCheck(clusterStatusCheckRequest)
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
func TestClient_ClusterTypeUpgrade(t *testing.T) {
	clusterTypeUpgradeRequest := &ClusterTypeUpgradeRequest{
		InstanceId:      util.PtrString(""),
		IsDefer:         util.PtrBool(false),
		NodeType:        util.PtrString(""),
		ShardNum:        util.PtrInt32(int32(0)),
		ReplicationInfo: []*ReplicationItem{},
	}
	result := &ClusterTypeUpgradeResponse{}
	result, err := SCS_CLIENT.ClusterTypeUpgrade(clusterTypeUpgradeRequest)
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
func TestClient_CreateAccount(t *testing.T) {
	createAccountRequest := &CreateAccountRequest{
		InstanceId: util.PtrString(""),
		UserName:   util.PtrString(""),
		ClientAuth: util.PtrString(""),
		Extra:      util.PtrString(""),
		UserType:   util.PtrInt32(int32(0)),
	}
	err := SCS_CLIENT.CreateAccount(createAccountRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAnInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createAnInstanceRequest := &CreateAnInstanceRequest{
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
		ReplicationInfo:   []*ReplicationMap{},
		AutoRenewTimeUnit: util.PtrString(""),
		AutoRenewTime:     util.PtrInt32(int32(0)),
		BgwGroupId:        util.PtrString(""),
		ClientAuth:        util.PtrString(""),
		Tags:              []*Tag{},
		ConfTpl:           util.PtrString(""),
		ResourceGroupId:   util.PtrString(""),
		AutoBackupConfig:  util.PtrString(""),
		DeployIdList:      []*string{},
	}
	result := &CreateAnInstanceResponse{}
	result, err := SCS_CLIENT.CreateAnInstance(createAnInstanceRequest)
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
func TestClient_CreateDeploymentSet(t *testing.T) {
	createDeploymentSetRequest := &CreateDeploymentSetRequest{
		Name:        util.PtrString(""),
		Desc:        util.PtrString(""),
		Strategy:    util.PtrString(""),
		Concurrency: util.PtrInt32(int32(0)),
	}
	result := &CreateDeploymentSetResponse{}
	result, err := SCS_CLIENT.CreateDeploymentSet(createDeploymentSetRequest)
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
func TestClient_CreateEntrance(t *testing.T) {
	createEntranceRequest := &CreateEntranceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &CreateEntranceResponse{}
	result, err := SCS_CLIENT.CreateEntrance(createEntranceRequest)
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
func TestClient_CreateHotGroup(t *testing.T) {
	Leader := &Leader{
		GroupName:         util.PtrString(""),
		LeaderId:          util.PtrString(""),
		LeaderRegion:      util.PtrString(""),
		ClusterName:       util.PtrString(""),
		ClusterShowId:     util.PtrString(""),
		Region:            util.PtrString(""),
		Status:            util.PtrString(""),
		TotalCapacityInGB: util.PtrFloat32(float32(0)),
		UsedCapacityInGB:  util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		Flavor:            util.PtrInt32(int32(0)),
		QpsWrite:          util.PtrInt64(int64(0)),
		QpsRead:           util.PtrInt64(int64(0)),
		StaleReadable:     util.PtrBool(false),
		ForbidWrite:       util.PtrInt32(int32(0)),
		AvailabilityZone:  util.PtrString(""),
		ExpiredTime:       util.PtrString(""),
	}
	createHotGroupRequest := &CreateHotGroupRequest{
		Leader: Leader,
	}
	result := &CreateHotGroupResponse{}
	result, err := SCS_CLIENT.CreateHotGroup(createHotGroupRequest)
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
func TestClient_CreateInstanceWhiteGroup(t *testing.T) {
	createInstanceWhiteGroupRequest := &CreateInstanceWhiteGroupRequest{
		InstanceId:    util.PtrString(""),
		GroupName:     util.PtrString(""),
		ClusterIpList: []*string{},
	}
	err := SCS_CLIENT.CreateInstanceWhiteGroup(createInstanceWhiteGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateParameterTemplate(t *testing.T) {
	createParameterTemplateRequest := &CreateParameterTemplateRequest{
		Name:          util.PtrString(""),
		Engine:        util.PtrString(""),
		EngineVersion: util.PtrString(""),
		ClusterType:   util.PtrString(""),
		TemplateType:  util.PtrInt32(int32(0)),
		Comment:       util.PtrString(""),
		Parameters:    []*Parameters{},
	}
	result := &CreateParameterTemplateResponse{}
	result, err := SCS_CLIENT.CreateParameterTemplate(createParameterTemplateRequest)
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
func TestClient_CreateSyncGroup(t *testing.T) {
	createSyncGroupRequest := &CreateSyncGroupRequest{
		SyncGroupName: util.PtrString(""),
		Members:       []*Member{},
	}
	result := &CreateSyncGroupResponse{}
	result, err := SCS_CLIENT.CreateSyncGroup(createSyncGroupRequest)
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
func TestClient_DeleteAccount(t *testing.T) {
	deleteAccountRequest := &DeleteAccountRequest{
		InstanceId: util.PtrString(""),
		UserName:   util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteAccount(deleteAccountRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDeploymentSet(t *testing.T) {
	deleteDeploymentSetRequest := &DeleteDeploymentSetRequest{
		DeploySetId: util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteDeploymentSet(deleteDeploymentSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstanceWhiteGroup(t *testing.T) {
	deleteInstanceWhiteGroupRequest := &DeleteInstanceWhiteGroupRequest{
		InstanceId: util.PtrString(""),
		GroupName:  util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteInstanceWhiteGroup(deleteInstanceWhiteGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstances(t *testing.T) {
	deleteInstancesRequest := &DeleteInstancesRequest{
		InstanceIds: []*string{},
	}
	err := SCS_CLIENT.DeleteInstances(deleteInstancesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpWhitelist(t *testing.T) {
	deleteIpWhitelistRequest := &DeleteIpWhitelistRequest{
		InstanceId:  util.PtrString(""),
		SecurityIps: []*string{},
	}
	err := SCS_CLIENT.DeleteIpWhitelist(deleteIpWhitelistRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteManualBackup(t *testing.T) {
	deleteManualBackupRequest := &DeleteManualBackupRequest{
		InstanceId: util.PtrString(""),
		BatchId:    util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteManualBackup(deleteManualBackupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteMemoryScalingConfig(t *testing.T) {
	deleteMemoryScalingConfigRequest := &DeleteMemoryScalingConfigRequest{
		InstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteMemoryScalingConfig(deleteMemoryScalingConfigRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteParameterTemplate(t *testing.T) {
	deleteParameterTemplateRequest := &DeleteParameterTemplateRequest{
		TemplateShowId: util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteParameterTemplate(deleteParameterTemplateRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSyncGroup(t *testing.T) {
	deleteSyncGroupRequest := &DeleteSyncGroupRequest{
		SyncGroupShowId: util.PtrString(""),
	}
	err := SCS_CLIENT.DeleteSyncGroup(deleteSyncGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisconnectEntrance(t *testing.T) {
	disconnectEntranceRequest := &DisconnectEntranceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &DisconnectEntranceResponse{}
	result, err := SCS_CLIENT.DisconnectEntrance(disconnectEntranceRequest)
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
func TestClient_DomainNameExchange(t *testing.T) {
	domainNameExchangeRequest := &DomainNameExchangeRequest{
		SourceInstanceId: util.PtrString(""),
		TargetInstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.DomainNameExchange(domainNameExchangeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GePriceForResizeInstance(t *testing.T) {
	gePriceForResizeInstanceRequest := &GePriceForResizeInstanceRequest{
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		NodeType:       util.PtrString(""),
		ShardNum:       util.PtrInt32(int32(0)),
		ReplicationNum: util.PtrInt32(int32(0)),
		DiskFlavor:     util.PtrInt32(int32(0)),
		ChargeType:     util.PtrString(""),
		Period:         util.PtrInt32(int32(0)),
		ChangeType:     util.PtrString(""),
	}
	result := &GePriceForResizeInstanceResponse{}
	result, err := SCS_CLIENT.GePriceForResizeInstance(gePriceForResizeInstanceRequest)
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
func TestClient_GetApplicationParameterTemplateRecords(t *testing.T) {
	getApplicationParameterTemplateRecordsRequest := &GetApplicationParameterTemplateRecordsRequest{
		TemplateShowId: util.PtrString(""),
		Marker:         util.PtrString(""),
		MaxKeys:        util.PtrInt32(int32(0)),
	}
	result := &GetApplicationParameterTemplateRecordsResponse{}
	result, err := SCS_CLIENT.GetApplicationParameterTemplateRecords(getApplicationParameterTemplateRecordsRequest)
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
func TestClient_GetAvailableZones(t *testing.T) {
	result := &GetAvailableZonesResponse{}
	result, err := SCS_CLIENT.GetAvailableZones()
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
func TestClient_GetBackUpUrl(t *testing.T) {
	getBackUpUrlRequest := &GetBackUpUrlRequest{
		InstanceId: util.PtrString(""),
		BackupId:   util.PtrString(""),
	}
	result := &GetBackUpUrlResponse{}
	result, err := SCS_CLIENT.GetBackUpUrl(getBackUpUrlRequest)
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
func TestClient_GetBackUpUsage(t *testing.T) {
	getBackUpUsageRequest := &GetBackUpUsageRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetBackUpUsageResponse{}
	result, err := SCS_CLIENT.GetBackUpUsage(getBackUpUsageRequest)
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
func TestClient_GetBackupList(t *testing.T) {
	getBackupListRequest := &GetBackupListRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetBackupListResponse{}
	result, err := SCS_CLIENT.GetBackupList(getBackupListRequest)
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
func TestClient_GetBackupStrategy(t *testing.T) {
	getBackupStrategyRequest := &GetBackupStrategyRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetBackupStrategyResponse{}
	result, err := SCS_CLIENT.GetBackupStrategy(getBackupStrategyRequest)
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
func TestClient_GetClusterBlbStatus(t *testing.T) {
	getClusterBlbStatusRequest := &GetClusterBlbStatusRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetClusterBlbStatusResponse{}
	result, err := SCS_CLIENT.GetClusterBlbStatus(getClusterBlbStatusRequest)
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
func TestClient_GetDeploymentSetList(t *testing.T) {
	getDeploymentSetListRequest := &GetDeploymentSetListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &GetDeploymentSetListResponse{}
	result, err := SCS_CLIENT.GetDeploymentSetList(getDeploymentSetListRequest)
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
func TestClient_GetHotGroupDetail(t *testing.T) {
	getHotGroupDetailRequest := &GetHotGroupDetailRequest{
		GroupId: util.PtrString(""),
	}
	result := &GetHotGroupDetailResponse{}
	result, err := SCS_CLIENT.GetHotGroupDetail(getHotGroupDetailRequest)
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
func TestClient_GetHotGroupList(t *testing.T) {
	getHotGroupListRequest := &GetHotGroupListRequest{
		PageSize: util.PtrInt32(int32(0)),
		PageNo:   util.PtrInt32(int32(0)),
	}
	result := &GetHotGroupListResponse{}
	result, err := SCS_CLIENT.GetHotGroupList(getHotGroupListRequest)
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
func TestClient_GetInstanceDetail(t *testing.T) {
	getInstanceDetailRequest := &GetInstanceDetailRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetInstanceDetailResponse{}
	result, err := SCS_CLIENT.GetInstanceDetail(getInstanceDetailRequest)
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
func TestClient_GetInstanceList(t *testing.T) {
	getInstanceListRequest := &GetInstanceListRequest{
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrString(""),
		InstanceIds: util.PtrString(""),
		VnetIp:      util.PtrString(""),
	}
	result := &GetInstanceListResponse{}
	result, err := SCS_CLIENT.GetInstanceList(getInstanceListRequest)
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
func TestClient_GetInstanceSpecList(t *testing.T) {
	result := &GetInstanceSpecListResponse{}
	result, err := SCS_CLIENT.GetInstanceSpecList()
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
func TestClient_GetInstanceWhiteGroup(t *testing.T) {
	getInstanceWhiteGroupRequest := &GetInstanceWhiteGroupRequest{
		InstanceId: util.PtrString(""),
		GroupName:  util.PtrString(""),
	}
	result := &GetInstanceWhiteGroupResponse{}
	result, err := SCS_CLIENT.GetInstanceWhiteGroup(getInstanceWhiteGroupRequest)
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
func TestClient_GetParameterList(t *testing.T) {
	getParameterListRequest := &GetParameterListRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetParameterListResponse{}
	result, err := SCS_CLIENT.GetParameterList(getParameterListRequest)
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
func TestClient_GetParameterTemplateList(t *testing.T) {
	getParameterTemplateListRequest := &GetParameterTemplateListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &GetParameterTemplateListResponse{}
	result, err := SCS_CLIENT.GetParameterTemplateList(getParameterTemplateListRequest)
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
func TestClient_GetPriceForCreateInstance(t *testing.T) {
	getPriceForCreateInstanceRequest := &GetPriceForCreateInstanceRequest{
		Engine:            util.PtrInt32(int32(0)),
		ClusterType:       util.PtrString(""),
		NodeType:          util.PtrString(""),
		CacheInstanceType: util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		ReplicationNum:    util.PtrInt32(int32(0)),
		InstanceNum:       util.PtrInt32(int32(0)),
		DiskType:          util.PtrString(""),
		DiskFlavor:        util.PtrInt32(int32(0)),
		ChargeType:        util.PtrString(""),
		Period:            util.PtrInt32(int32(0)),
		TimeUnit:          util.PtrString(""),
	}
	result := &GetPriceForCreateInstanceResponse{}
	result, err := SCS_CLIENT.GetPriceForCreateInstance(getPriceForCreateInstanceRequest)
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
func TestClient_GetRecycleList(t *testing.T) {
	getRecycleListRequest := &GetRecycleListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &GetRecycleListResponse{}
	result, err := SCS_CLIENT.GetRecycleList(getRecycleListRequest)
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
func TestClient_GetSubnetList(t *testing.T) {
	getSubnetListRequest := &GetSubnetListRequest{
		VpcId:    util.PtrString(""),
		ZoneName: util.PtrString(""),
	}
	result := &GetSubnetListResponse{}
	result, err := SCS_CLIENT.GetSubnetList(getSubnetListRequest)
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
func TestClient_GetSyncGroupStatus(t *testing.T) {
	getSyncGroupStatusRequest := &GetSyncGroupStatusRequest{
		SyncGroupShowId: util.PtrString(""),
	}
	result := &GetSyncGroupStatusResponse{}
	result, err := SCS_CLIENT.GetSyncGroupStatus(getSyncGroupStatusRequest)
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
func TestClient_GetSystemParameterList(t *testing.T) {
	getSystemParameterListRequest := &GetSystemParameterListRequest{
		Engine:        util.PtrString(""),
		EngineVersion: util.PtrString(""),
		ClusterType:   util.PtrString(""),
	}
	result := &GetSystemParameterListResponse{}
	result, err := SCS_CLIENT.GetSystemParameterList(getSystemParameterListRequest)
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
func TestClient_GetTimeWindow(t *testing.T) {
	getTimeWindowRequest := &GetTimeWindowRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetTimeWindowResponse{}
	result, err := SCS_CLIENT.GetTimeWindow(getTimeWindowRequest)
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
func TestClient_GetTlsCert(t *testing.T) {
	getTlsCertRequest := &GetTlsCertRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetTlsCertResponse{}
	result, err := SCS_CLIENT.GetTlsCert(getTlsCertRequest)
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
func TestClient_HotGroupAddCluster(t *testing.T) {
	hotGroupAddClusterRequest := &HotGroupAddClusterRequest{
		GroupId:        util.PtrString(""),
		FollowerId:     util.PtrString(""),
		FollowerRegion: util.PtrString(""),
		SyncMaster:     util.PtrString(""),
	}
	err := SCS_CLIENT.HotGroupAddCluster(hotGroupAddClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupChangeMasterRole(t *testing.T) {
	hotGroupChangeMasterRoleRequest := &HotGroupChangeMasterRoleRequest{
		GroupId:    util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.HotGroupChangeMasterRole(hotGroupChangeMasterRoleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupForbidWrite(t *testing.T) {
	hotGroupForbidWriteRequest := &HotGroupForbidWriteRequest{
		GroupId:         util.PtrString(""),
		ForbidWriteFlag: util.PtrBool(false),
	}
	err := SCS_CLIENT.HotGroupForbidWrite(hotGroupForbidWriteRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupModifyName(t *testing.T) {
	hotGroupModifyNameRequest := &HotGroupModifyNameRequest{
		GroupId:   util.PtrString(""),
		GroupName: util.PtrString(""),
	}
	err := SCS_CLIENT.HotGroupModifyName(hotGroupModifyNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupPreCheck(t *testing.T) {
	Leader := &Leader{
		GroupName:         util.PtrString(""),
		LeaderId:          util.PtrString(""),
		LeaderRegion:      util.PtrString(""),
		ClusterName:       util.PtrString(""),
		ClusterShowId:     util.PtrString(""),
		Region:            util.PtrString(""),
		Status:            util.PtrString(""),
		TotalCapacityInGB: util.PtrFloat32(float32(0)),
		UsedCapacityInGB:  util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		Flavor:            util.PtrInt32(int32(0)),
		QpsWrite:          util.PtrInt64(int64(0)),
		QpsRead:           util.PtrInt64(int64(0)),
		StaleReadable:     util.PtrBool(false),
		ForbidWrite:       util.PtrInt32(int32(0)),
		AvailabilityZone:  util.PtrString(""),
		ExpiredTime:       util.PtrString(""),
	}
	hotGroupPreCheckRequest := &HotGroupPreCheckRequest{
		Leader:    Leader,
		Followers: []*FollowersItem{},
	}
	result := &HotGroupPreCheckResponse{}
	result, err := SCS_CLIENT.HotGroupPreCheck(hotGroupPreCheckRequest)
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
func TestClient_HotGroupRemoveCluster(t *testing.T) {
	hotGroupRemoveClusterRequest := &HotGroupRemoveClusterRequest{
		GroupId:    util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.HotGroupRemoveCluster(hotGroupRemoveClusterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupSetFlowControlRules(t *testing.T) {
	hotGroupSetFlowControlRulesRequest := &HotGroupSetFlowControlRulesRequest{
		GroupId:       util.PtrString(""),
		ClusterShowId: util.PtrString(""),
		QpsWrite:      util.PtrInt32(int32(0)),
		QpsRead:       util.PtrInt32(int32(0)),
	}
	err := SCS_CLIENT.HotGroupSetFlowControlRules(hotGroupSetFlowControlRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupStaleReadable(t *testing.T) {
	hotGroupStaleReadableRequest := &HotGroupStaleReadableRequest{
		GroupId:       util.PtrString(""),
		FollowerId:    util.PtrString(""),
		StaleReadable: util.PtrBool(false),
	}
	err := SCS_CLIENT.HotGroupStaleReadable(hotGroupStaleReadableRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HotGroupSyncStatus(t *testing.T) {
	hotGroupSyncStatusRequest := &HotGroupSyncStatusRequest{
		GroupId: util.PtrString(""),
	}
	result := &HotGroupSyncStatusResponse{}
	result, err := SCS_CLIENT.HotGroupSyncStatus(hotGroupSyncStatusRequest)
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
func TestClient_InstanceVersionUpgrade(t *testing.T) {
	instanceVersionUpgradeRequest := &InstanceVersionUpgradeRequest{
		InstanceId:    util.PtrString(""),
		KernelVersion: util.PtrString(""),
		IsDefer:       util.PtrBool(false),
	}
	err := SCS_CLIENT.InstanceVersionUpgrade(instanceVersionUpgradeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_LogDetails(t *testing.T) {
	logDetailsRequest := &LogDetailsRequest{
		InstanceId:   util.PtrString(""),
		LogId:        util.PtrString(""),
		ValidSeconds: util.PtrString(""),
	}
	result := &LogDetailsResponse{}
	result, err := SCS_CLIENT.LogDetails(logDetailsRequest)
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
func TestClient_LogList(t *testing.T) {
	logListRequest := &LogListRequest{
		InstanceId: util.PtrString(""),
		FileType:   util.PtrString(""),
		StartTime:  util.PtrString(""),
		EndTime:    util.PtrString(""),
	}
	result := &LogListResponse{}
	result, err := SCS_CLIENT.LogList(logListRequest)
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
func TestClient_ManualBackup(t *testing.T) {
	manualBackupRequest := &ManualBackupRequest{
		InstanceId: util.PtrString(""),
		Comment:    util.PtrString(""),
	}
	err := SCS_CLIENT.ManualBackup(manualBackupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ManuallyModifyBandwidth(t *testing.T) {
	manuallyModifyBandwidthRequest := &ManuallyModifyBandwidthRequest{
		InstanceId:         util.PtrString(""),
		ShardBandwidthInfo: []*ShardBandwidth{},
	}
	result := &ManuallyModifyBandwidthResponse{}
	result, err := SCS_CLIENT.ManuallyModifyBandwidth(manuallyModifyBandwidthRequest)
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
func TestClient_MasterSlaveSwitch(t *testing.T) {
	masterSlaveSwitchRequest := &MasterSlaveSwitchRequest{
		InstanceId: util.PtrString(""),
		Shards:     []*SwitchMasterSlaveShard{},
	}
	err := SCS_CLIENT.MasterSlaveSwitch(masterSlaveSwitchRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyBackupComment(t *testing.T) {
	modifyBackupCommentRequest := &ModifyBackupCommentRequest{
		InstanceId: util.PtrString(""),
		BatchId:    util.PtrString(""),
		Comment:    util.PtrString(""),
	}
	err := SCS_CLIENT.ModifyBackupComment(modifyBackupCommentRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyDeploymentSet(t *testing.T) {
	modifyDeploymentSetRequest := &ModifyDeploymentSetRequest{
		DeploySetId: util.PtrString(""),
		Desc:        util.PtrString(""),
		Name:        util.PtrString(""),
		Concurrency: util.PtrInt32(int32(0)),
	}
	err := SCS_CLIENT.ModifyDeploymentSet(modifyDeploymentSetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyEntrance(t *testing.T) {
	modifyEntranceRequest := &ModifyEntranceRequest{
		InstanceId: util.PtrString(""),
		IsDefer:    util.PtrBool(false),
	}
	err := SCS_CLIENT.ModifyEntrance(modifyEntranceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceDomainName(t *testing.T) {
	modifyInstanceDomainNameRequest := &ModifyInstanceDomainNameRequest{
		InstanceId: util.PtrString(""),
		Domain:     util.PtrString(""),
	}
	err := SCS_CLIENT.ModifyInstanceDomainName(modifyInstanceDomainNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceName(t *testing.T) {
	modifyInstanceNameRequest := &ModifyInstanceNameRequest{
		InstanceId:   util.PtrString(""),
		InstanceName: util.PtrString(""),
	}
	err := SCS_CLIENT.ModifyInstanceName(modifyInstanceNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyParameterTemplateName(t *testing.T) {
	modifyParameterTemplateNameRequest := &ModifyParameterTemplateNameRequest{
		TemplateShowId: util.PtrString(""),
		Name:           util.PtrString(""),
	}
	err := SCS_CLIENT.ModifyParameterTemplateName(modifyParameterTemplateNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyParameters(t *testing.T) {
	Parameter := &Parameter{
		ScsDefault:   util.PtrString(""),
		ForceRestart: util.PtrInt32(int32(0)),
		Name:         util.PtrString(""),
		Value:        util.PtrString(""),
	}
	modifyParametersRequest := &ModifyParametersRequest{
		InstanceId: util.PtrString(""),
		Parameter:  Parameter,
	}
	err := SCS_CLIENT.ModifyParameters(modifyParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyReplicationZone(t *testing.T) {
	modifyReplicationZoneRequest := &ModifyReplicationZoneRequest{
		InstanceId:      util.PtrString(""),
		IsDefer:         util.PtrBool(false),
		ReplicationInfo: []*ReplicationItem{},
	}
	err := SCS_CLIENT.ModifyReplicationZone(modifyReplicationZoneRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifySyncGroupName(t *testing.T) {
	modifySyncGroupNameRequest := &ModifySyncGroupNameRequest{
		SyncGroupShowId: util.PtrString(""),
		GroupName:       util.PtrString(""),
	}
	err := SCS_CLIENT.ModifySyncGroupName(modifySyncGroupNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTimeWindow(t *testing.T) {
	modifyTimeWindowRequest := &ModifyTimeWindowRequest{
		InstanceId: util.PtrString(""),
		StartTime:  util.PtrString(""),
		Duration:   util.PtrInt32(int32(0)),
		Period:     []*int32{},
	}
	err := SCS_CLIENT.ModifyTimeWindow(modifyTimeWindowRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ParameterTemplateDeleteParameters(t *testing.T) {
	parameterTemplateDeleteParametersRequest := &ParameterTemplateDeleteParametersRequest{
		TemplateShowId: util.PtrString(""),
		Parameters:     []*string{},
	}
	err := SCS_CLIENT.ParameterTemplateDeleteParameters(parameterTemplateDeleteParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ParameterTemplateDetails(t *testing.T) {
	parameterTemplateDetailsRequest := &ParameterTemplateDetailsRequest{
		TemplateShowId: util.PtrString(""),
	}
	result := &ParameterTemplateDetailsResponse{}
	result, err := SCS_CLIENT.ParameterTemplateDetails(parameterTemplateDetailsRequest)
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
func TestClient_ParameterTemplateModifyParameters(t *testing.T) {
	parameterTemplateModifyParametersRequest := &ParameterTemplateModifyParametersRequest{
		TemplateShowId: util.PtrString(""),
		Parameters:     []*Parameters{},
	}
	err := SCS_CLIENT.ParameterTemplateModifyParameters(parameterTemplateModifyParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PostPaidToPrepaid(t *testing.T) {
	postPaidToPrepaidRequest := &PostPaidToPrepaidRequest{
		Duration:    util.PtrInt32(int32(0)),
		InstanceIds: []*string{},
	}
	result := &PostPaidToPrepaidResponse{}
	result, err := SCS_CLIENT.PostPaidToPrepaid(postPaidToPrepaidRequest)
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
func TestClient_PrepaidToPostpaid(t *testing.T) {
	prepaidToPostpaidRequest := &PrepaidToPostpaidRequest{
		InstanceIds: []*string{},
	}
	result := &PrepaidToPostpaidResponse{}
	result, err := SCS_CLIENT.PrepaidToPostpaid(prepaidToPostpaidRequest)
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
func TestClient_ProxyNodeReplace(t *testing.T) {
	proxyNodeReplaceRequest := &ProxyNodeReplaceRequest{
		InstanceId: util.PtrString(""),
		ProxyList:  []*string{},
		ScsDefer:   util.PtrBool(false),
	}
	err := SCS_CLIENT.ProxyNodeReplace(proxyNodeReplaceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ProxyVersionUpgradeOrRestart(t *testing.T) {
	proxyVersionUpgradeOrRestartRequest := &ProxyVersionUpgradeOrRestartRequest{
		InstanceId:  util.PtrString(""),
		ProxyList:   []*string{},
		UpgradeType: util.PtrString(""),
		IsDefer:     util.PtrBool(false),
	}
	err := SCS_CLIENT.ProxyVersionUpgradeOrRestart(proxyVersionUpgradeOrRestartRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryIpWhitelist(t *testing.T) {
	queryIpWhitelistRequest := &QueryIpWhitelistRequest{
		InstanceId: util.PtrString(""),
	}
	result := &QueryIpWhitelistResponse{}
	result, err := SCS_CLIENT.QueryIpWhitelist(queryIpWhitelistRequest)
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
func TestClient_QueryMemoryScalingConfig(t *testing.T) {
	queryMemoryScalingConfigRequest := &QueryMemoryScalingConfigRequest{
		InstanceId: util.PtrString(""),
	}
	result := &QueryMemoryScalingConfigResponse{}
	result, err := SCS_CLIENT.QueryMemoryScalingConfig(queryMemoryScalingConfigRequest)
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
func TestClient_ReleaseHotGroup(t *testing.T) {
	releaseHotGroupRequest := &ReleaseHotGroupRequest{
		GroupId: util.PtrString(""),
	}
	err := SCS_CLIENT.ReleaseHotGroup(releaseHotGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseInstance(t *testing.T) {
	releaseInstanceRequest := &ReleaseInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.ReleaseInstance(releaseInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenewInstance(t *testing.T) {
	renewInstanceRequest := &RenewInstanceRequest{
		InstanceIds: []*string{},
		Duration:    util.PtrInt32(int32(0)),
	}
	result := &RenewInstanceResponse{}
	result, err := SCS_CLIENT.RenewInstance(renewInstanceRequest)
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
func TestClient_RestartInstance(t *testing.T) {
	restartInstanceRequest := &RestartInstanceRequest{
		InstanceId: util.PtrString(""),
		IsDefer:    util.PtrBool(false),
	}
	err := SCS_CLIENT.RestartInstance(restartInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetBackupPolicy(t *testing.T) {
	setBackupPolicyRequest := &SetBackupPolicyRequest{
		InstanceId: util.PtrString(""),
		BackupTime: util.PtrString(""),
		BackupDays: util.PtrString(""),
		ExpireDay:  util.PtrInt32(int32(0)),
		IsEncrypt:  util.PtrString(""),
	}
	err := SCS_CLIENT.SetBackupPolicy(setBackupPolicyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetClusterAsMaster(t *testing.T) {
	setClusterAsMasterRequest := &SetClusterAsMasterRequest{
		InstanceId: util.PtrString(""),
	}
	err := SCS_CLIENT.SetClusterAsMaster(setClusterAsMasterRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetClusterAsSlave(t *testing.T) {
	setClusterAsSlaveRequest := &SetClusterAsSlaveRequest{
		InstanceId:   util.PtrString(""),
		MasterDomain: util.PtrString(""),
		MasterPort:   util.PtrInt32(int32(0)),
	}
	err := SCS_CLIENT.SetClusterAsSlave(setClusterAsSlaveRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetMemoryScalingConfig(t *testing.T) {
	MemSpec := &MemSpec{
		MemUsageUpperThreshold:        util.PtrInt32(int32(0)),
		MemUsageDownThreshold:         util.PtrInt32(int32(0)),
		MaxNodeType:                   util.PtrString(""),
		MinNodeType:                   util.PtrString(""),
		ObservationWindowSizeForUpper: util.PtrString(""),
		ObservationWindowSizeForDown:  util.PtrString(""),
	}
	setMemoryScalingConfigRequest := &SetMemoryScalingConfigRequest{
		InstanceId: util.PtrString(""),
		MemSpec:    MemSpec,
	}
	err := SCS_CLIENT.SetMemoryScalingConfig(setMemoryScalingConfigRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetPermissions(t *testing.T) {
	setPermissionsRequest := &SetPermissionsRequest{
		InstanceId: util.PtrString(""),
		UserName:   util.PtrString(""),
		UserType:   util.PtrInt32(int32(0)),
	}
	err := SCS_CLIENT.SetPermissions(setPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SyncGroupAddInstance(t *testing.T) {
	syncGroupAddInstanceRequest := &SyncGroupAddInstanceRequest{
		GroupId:  util.PtrString(""),
		MemberId: util.PtrString(""),
		Region:   util.PtrString(""),
	}
	err := SCS_CLIENT.SyncGroupAddInstance(syncGroupAddInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SyncGroupDelayInfo(t *testing.T) {
	syncGroupDelayInfoRequest := &SyncGroupDelayInfoRequest{
		GroupId: util.PtrString(""),
	}
	result := &SyncGroupDelayInfoResponse{}
	result, err := SCS_CLIENT.SyncGroupDelayInfo(syncGroupDelayInfoRequest)
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
func TestClient_SyncGroupDetail(t *testing.T) {
	syncGroupDetailRequest := &SyncGroupDetailRequest{
		SyncGroupShowId: util.PtrString(""),
	}
	result := &SyncGroupDetailResponse{}
	result, err := SCS_CLIENT.SyncGroupDetail(syncGroupDetailRequest)
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
func TestClient_SyncGroupList(t *testing.T) {
	syncGroupListRequest := &SyncGroupListRequest{
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &SyncGroupListResponse{}
	result, err := SCS_CLIENT.SyncGroupList(syncGroupListRequest)
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
func TestClient_SyncGroupModifyBnsgroup(t *testing.T) {
	syncGroupModifyBnsgroupRequest := &SyncGroupModifyBnsgroupRequest{
		GroupId:  util.PtrString(""),
		BnsGroup: util.PtrString(""),
	}
	err := SCS_CLIENT.SyncGroupModifyBnsgroup(syncGroupModifyBnsgroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SyncGroupPreCheck(t *testing.T) {
	syncGroupPreCheckRequest := &SyncGroupPreCheckRequest{
		SyncGroupShowId: util.PtrString(""),
		Members:         []*CheckSyncGroupRequestMember{},
	}
	result := &SyncGroupPreCheckResponse{}
	result, err := SCS_CLIENT.SyncGroupPreCheck(syncGroupPreCheckRequest)
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
func TestClient_SyncGroupRemoveInstance(t *testing.T) {
	syncGroupRemoveInstanceRequest := &SyncGroupRemoveInstanceRequest{
		GroupId:  util.PtrString(""),
		MemberId: util.PtrString(""),
		Region:   util.PtrString(""),
	}
	err := SCS_CLIENT.SyncGroupRemoveInstance(syncGroupRemoveInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TdeEncryption(t *testing.T) {
	tdeEncryptionRequest := &TdeEncryptionRequest{
		InstanceId: util.PtrString(""),
		Action:     util.PtrString(""),
	}
	result := &TdeEncryptionResponse{}
	result, err := SCS_CLIENT.TdeEncryption(tdeEncryptionRequest)
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
func TestClient_UnbindSecurityGroup(t *testing.T) {
	unbindSecurityGroupRequest := &UnbindSecurityGroupRequest{
		InstanceId:       util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	result := &UnbindSecurityGroupResponse{}
	result, err := SCS_CLIENT.UnbindSecurityGroup(unbindSecurityGroupRequest)
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
func TestClient_UnbindTags(t *testing.T) {
	unbindTagsRequest := &UnbindTagsRequest{
		InstanceId: util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := SCS_CLIENT.UnbindTags(unbindTagsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceWhiteGroup(t *testing.T) {
	updateInstanceWhiteGroupRequest := &UpdateInstanceWhiteGroupRequest{
		InstanceId:    util.PtrString(""),
		GroupName:     util.PtrString(""),
		NewGroupName:  util.PtrString(""),
		ClusterIpList: []*string{},
	}
	err := SCS_CLIENT.UpdateInstanceWhiteGroup(updateInstanceWhiteGroupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSecurityGroup(t *testing.T) {
	updateSecurityGroupRequest := &UpdateSecurityGroupRequest{
		InstanceId:       util.PtrString(""),
		SecurityGroupIds: []*string{},
	}
	result := &UpdateSecurityGroupResponse{}
	result, err := SCS_CLIENT.UpdateSecurityGroup(updateSecurityGroupRequest)
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
func TestClient_UpdateTlsEncryption(t *testing.T) {
	updateTlsEncryptionRequest := &UpdateTlsEncryptionRequest{
		InstanceId: util.PtrString(""),
		Action:     util.PtrString(""),
	}
	err := SCS_CLIENT.UpdateTlsEncryption(updateTlsEncryptionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ViewSecurityGroup(t *testing.T) {
	viewSecurityGroupRequest := &ViewSecurityGroupRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ViewSecurityGroupResponse{}
	result, err := SCS_CLIENT.ViewSecurityGroup(viewSecurityGroupRequest)
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
