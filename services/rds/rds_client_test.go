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

func TestClient_AddAccountApiAtTableGranularity(t *testing.T) {
	addAccountApiAtTableGranularityRequest := &AddAccountApiAtTableGranularityRequest{
		InstanceId: util.PtrString(""),
	}
	err := RDS_CLIENT.AddAccountApiAtTableGranularity(addAccountApiAtTableGranularityRequest)
	ExpectEqual(t.Errorf, nil, err)
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
func TestClient_ChangeAccountPassword(t *testing.T) {
	changeAccountPasswordRequest := &ChangeAccountPasswordRequest{
		InstanceId:  util.PtrString(""),
		AccountName: util.PtrString(""),
		Password:    util.PtrString(""),
	}
	err := RDS_CLIENT.ChangeAccountPassword(changeAccountPasswordRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CheckTheWhitelist(t *testing.T) {
	checkTheWhitelistRequest := &CheckTheWhitelistRequest{
		InstanceId: util.PtrString(""),
	}
	result := &CheckTheWhitelistResponse{}
	result, err := RDS_CLIENT.CheckTheWhitelist(checkTheWhitelistRequest)
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
func TestClient_CopyParameterTemplate(t *testing.T) {
	copyParameterTemplateRequest := &CopyParameterTemplateRequest{
		Name:       util.PtrString(""),
		Desc:       util.PtrString(""),
		TemplateId: util.PtrString(""),
	}
	err := RDS_CLIENT.CopyParameterTemplate(copyParameterTemplateRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateADatabase(t *testing.T) {
	createADatabaseRequest := &CreateADatabaseRequest{
		InstanceId:        util.PtrString(""),
		CharacterSetName:  util.PtrString(""),
		DbName:            util.PtrString(""),
		Remark:            util.PtrString(""),
		AccountPrivileges: []*AccountPrivilege{},
		Ctype:             util.PtrString(""),
		Collate:           util.PtrString(""),
		Owner:             util.PtrString(""),
	}
	err := RDS_CLIENT.CreateADatabase(createADatabaseRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAPrimaryInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
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
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
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
func TestClient_CreateAccount(t *testing.T) {
	createAccountRequest := &CreateAccountRequest{
		InstanceId:         util.PtrString(""),
		AccountName:        util.PtrString(""),
		Password:           util.PtrString(""),
		AccountType:        util.PtrString(""),
		DatabasePrivileges: []*DatabasePrivilege{},
		Desc:               util.PtrString(""),
		RdsType:            util.PtrString(""),
	}
	err := RDS_CLIENT.CreateAccount(createAccountRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateReadOnlyInstance(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
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
func TestClient_DeleteAccount(t *testing.T) {
	deleteAccountRequest := &DeleteAccountRequest{
		InstanceId:  util.PtrString(""),
		AccountName: util.PtrString(""),
	}
	err := RDS_CLIENT.DeleteAccount(deleteAccountRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTheDatabase(t *testing.T) {
	deleteTheDatabaseRequest := &DeleteTheDatabaseRequest{
		InstanceId: util.PtrString(""),
		DbName:     util.PtrString(""),
	}
	err := RDS_CLIENT.DeleteTheDatabase(deleteTheDatabaseRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTheSpecifiedBackupSet(t *testing.T) {
	deleteTheSpecifiedBackupSetRequest := &DeleteTheSpecifiedBackupSetRequest{
		InstanceId: util.PtrString(""),
		SnapshotId: util.PtrString(""),
	}
	err := RDS_CLIENT.DeleteTheSpecifiedBackupSet(deleteTheSpecifiedBackupSetRequest)
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
func TestClient_ErrorLogList(t *testing.T) {
	errorLogListRequest := &ErrorLogListRequest{
		InstanceId: util.PtrString(""),
		Datetime:   util.PtrString(""),
	}
	result := &ErrorLogListResponse{}
	result, err := RDS_CLIENT.ErrorLogList(errorLogListRequest)
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
func TestClient_GetParameterTemplateDetails(t *testing.T) {
	getParameterTemplateDetailsRequest := &GetParameterTemplateDetailsRequest{
		TemplateId: util.PtrString(""),
	}
	result := &GetParameterTemplateDetailsResponse{}
	result, err := RDS_CLIENT.GetParameterTemplateDetails(getParameterTemplateDetailsRequest)
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
		PageNo:    util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
		Type:      util.PtrString(""),
		DbType:    util.PtrString(""),
		DbVersion: util.PtrString(""),
	}
	result := &GetParameterTemplateListResponse{}
	result, err := RDS_CLIENT.GetParameterTemplateList(getParameterTemplateListRequest)
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
func TestClient_GetPgLogDetails(t *testing.T) {
	getPgLogDetailsRequest := &GetPgLogDetailsRequest{
		InstanceId:             util.PtrString(""),
		PglogId:                util.PtrString(""),
		DownloadValidTimeInSec: util.PtrInt32(int32(0)),
	}
	result := &GetPgLogDetailsResponse{}
	result, err := RDS_CLIENT.GetPgLogDetails(getPgLogDetailsRequest)
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
func TestClient_GetPgLogList(t *testing.T) {
	getPgLogListRequest := &GetPgLogListRequest{
		InstanceId: util.PtrString(""),
		Date:       util.PtrString(""),
	}
	result := &GetPgLogListResponse{}
	result, err := RDS_CLIENT.GetPgLogList(getPgLogListRequest)
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
func TestClient_InstanceEnablesAutomaticRenewal(t *testing.T) {
	instanceEnablesAutomaticRenewalRequest := &InstanceEnablesAutomaticRenewalRequest{
		InstanceIds:       []*string{},
		AutoRenewTimeUnit: util.PtrString(""),
		AutoRenewTime:     util.PtrInt32(int32(0)),
	}
	err := RDS_CLIENT.InstanceEnablesAutomaticRenewal(instanceEnablesAutomaticRenewalRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ManuallyInitiateFullPhysicalBackup(t *testing.T) {
	manuallyInitiateFullPhysicalBackupRequest := &ManuallyInitiateFullPhysicalBackupRequest{
		InstanceId:     util.PtrString(""),
		EffectiveTime:  util.PtrString(""),
		DataBackupType: util.PtrString(""),
	}
	err := RDS_CLIENT.ManuallyInitiateFullPhysicalBackup(manuallyInitiateFullPhysicalBackupRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyAccountNotes(t *testing.T) {
	modifyAccountNotesRequest := &ModifyAccountNotesRequest{
		InstanceId:  util.PtrString(""),
		AccountName: util.PtrString(""),
		Remark:      util.PtrString(""),
	}
	err := RDS_CLIENT.ModifyAccountNotes(modifyAccountNotesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyAccountPermissions(t *testing.T) {
	modifyAccountPermissionsRequest := &ModifyAccountPermissionsRequest{
		InstanceId:         util.PtrString(""),
		AccountName:        util.PtrString(""),
		DatabasePrivileges: []*DatabasePrivilege{},
	}
	err := RDS_CLIENT.ModifyAccountPermissions(modifyAccountPermissionsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyBackupStrategy(t *testing.T) {
	modifyBackupStrategyRequest := &ModifyBackupStrategyRequest{
		InstanceId:                    util.PtrString(""),
		BackupDays:                    util.PtrString(""),
		BackupTime:                    util.PtrString(""),
		Persistent:                    util.PtrBool(false),
		ExpireInDays:                  util.PtrInt32(int32(0)),
		LogBackupRetainDays:           util.PtrInt32(int32(0)),
		DataBackupType:                util.PtrString(""),
		IncrementalDataBackupEnable:   util.PtrBool(false),
		IncrementalDataBackupInterval: util.PtrInt32(int32(0)),
		LatestDataBackupRetainDays:    util.PtrInt32(int32(0)),
		RetentionPeriod:               util.PtrInt32(int32(0)),
		TimeUnit:                      util.PtrString(""),
		DataBackupCopyEnable:          util.PtrBool(false),
		DataBackupCopyStoragesRegions: []*string{},
		DataBackupCopyRetainDays:      util.PtrInt32(int32(0)),
	}
	err := RDS_CLIENT.ModifyBackupStrategy(modifyBackupStrategyRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyConfigurationParameters(t *testing.T) {
	modifyConfigurationParametersRequest := &ModifyConfigurationParametersRequest{
		InstanceId:    util.PtrString(""),
		EffectiveTime: util.PtrString(""),
		Parameters:    []*Parameter{},
	}
	err := RDS_CLIENT.ModifyConfigurationParameters(modifyConfigurationParametersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyDatabasePort(t *testing.T) {
	modifyDatabasePortRequest := &ModifyDatabasePortRequest{
		InstanceId: util.PtrString(""),
		EntryPort:  util.PtrInt32(int32(0)),
	}
	err := RDS_CLIENT.ModifyDatabasePort(modifyDatabasePortRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyDescriptionInformation(t *testing.T) {
	modifyDescriptionInformationRequest := &ModifyDescriptionInformationRequest{
		InstanceId: util.PtrString(""),
		DbName:     util.PtrString(""),
		Remark:     util.PtrString(""),
	}
	err := RDS_CLIENT.ModifyDescriptionInformation(modifyDescriptionInformationRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ObtainBinlogInformation(t *testing.T) {
	obtainBinlogInformationRequest := &ObtainBinlogInformationRequest{
		InstanceId:             util.PtrString(""),
		LogId:                  util.PtrString(""),
		DownloadValidTimeInSec: util.PtrString(""),
	}
	result := &ObtainBinlogInformationResponse{}
	result, err := RDS_CLIENT.ObtainBinlogInformation(obtainBinlogInformationRequest)
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
func TestClient_ObtainSslEncryptionCertificate(t *testing.T) {
	result := &ObtainSslEncryptionCertificateResponse{}
	result, err := RDS_CLIENT.ObtainSslEncryptionCertificate()
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
func TestClient_ObtainSslEncryptionInformation(t *testing.T) {
	obtainSslEncryptionInformationRequest := &ObtainSslEncryptionInformationRequest{
		InstanceId: util.PtrString(""),
	}
	result := &ObtainSslEncryptionInformationResponse{}
	result, err := RDS_CLIENT.ObtainSslEncryptionInformation(obtainSslEncryptionInformationRequest)
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
func TestClient_QueryAccountDetails(t *testing.T) {
	queryAccountDetailsRequest := &QueryAccountDetailsRequest{
		InstanceId:  util.PtrString(""),
		AccountName: util.PtrString(""),
	}
	result := &QueryAccountDetailsResponse{}
	result, err := RDS_CLIENT.QueryAccountDetails(queryAccountDetailsRequest)
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
func TestClient_QueryAccountList(t *testing.T) {
	queryAccountListRequest := &QueryAccountListRequest{
		InstanceId: util.PtrString(""),
	}
	result := &QueryAccountListResponse{}
	result, err := RDS_CLIENT.QueryAccountList(queryAccountListRequest)
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
func TestClient_QueryAndModifyHistory(t *testing.T) {
	queryAndModifyHistoryRequest := &QueryAndModifyHistoryRequest{
		InstanceId: util.PtrString(""),
	}
	result := &QueryAndModifyHistoryResponse{}
	result, err := RDS_CLIENT.QueryAndModifyHistory(queryAndModifyHistoryRequest)
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
func TestClient_QueryBackupDetails(t *testing.T) {
	queryBackupDetailsRequest := &QueryBackupDetailsRequest{
		InstanceId: util.PtrString(""),
		BackupId:   util.PtrString(""),
	}
	result := &QueryBackupDetailsResponse{}
	result, err := RDS_CLIENT.QueryBackupDetails(queryBackupDetailsRequest)
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
func TestClient_QueryBackupList(t *testing.T) {
	queryBackupListRequest := &QueryBackupListRequest{
		InstanceId: util.PtrString(""),
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
	}
	result := &QueryBackupListResponse{}
	result, err := RDS_CLIENT.QueryBackupList(queryBackupListRequest)
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
func TestClient_QueryParameterList(t *testing.T) {
	queryParameterListRequest := &QueryParameterListRequest{
		InstanceId: util.PtrString(""),
		Keyword:    util.PtrString(""),
	}
	result := &QueryParameterListResponse{}
	result, err := RDS_CLIENT.QueryParameterList(queryParameterListRequest)
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
func TestClient_RenewalExample(t *testing.T) {
	renewalExampleRequest := &RenewalExampleRequest{
		Duration:    util.PtrInt32(int32(0)),
		InstanceIds: []*string{},
	}
	result := &RenewalExampleResponse{}
	result, err := RDS_CLIENT.RenewalExample(renewalExampleRequest)
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
func TestClient_RestoreLibraryTablesBasedOnBackupSets(t *testing.T) {
	restoreLibraryTablesBasedOnBackupSetsRequest := &RestoreLibraryTablesBasedOnBackupSetsRequest{
		InstanceId:       util.PtrString(""),
		SourceInstanceId: util.PtrString(""),
		TargetInstanceId: util.PtrString(""),
		SnapshotId:       util.PtrString(""),
		Data:             []*RecoveryToSourceInstanceModel{},
		DbName:           util.PtrString(""),
		NewDbname:        util.PtrString(""),
		RestoreMode:      util.PtrString(""),
		Tables:           []*RecoveryToSourceInstanceTableModel{},
		TableName:        util.PtrString(""),
		NewTablename:     util.PtrString(""),
	}
	err := RDS_CLIENT.RestoreLibraryTablesBasedOnBackupSets(restoreLibraryTablesBasedOnBackupSetsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RestoreTheDatabaseTableByTimePoint(t *testing.T) {
	restoreTheDatabaseTableByTimePointRequest := &RestoreTheDatabaseTableByTimePointRequest{
		InstanceId:       util.PtrString(""),
		SourceInstanceId: util.PtrString(""),
		TargetInstanceId: util.PtrString(""),
		Datetime:         util.PtrString(""),
		Data:             []*RecoveryToSourceInstanceModel{},
		DbName:           util.PtrString(""),
		NewDbname:        util.PtrString(""),
		RestoreMode:      util.PtrString(""),
		Tables:           []*RecoveryToSourceInstanceTableModel{},
		TableName:        util.PtrString(""),
		NewTablename:     util.PtrString(""),
	}
	err := RDS_CLIENT.RestoreTheDatabaseTableByTimePoint(restoreTheDatabaseTableByTimePointRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RetrieveDatabaseList(t *testing.T) {
	retrieveDatabaseListRequest := &RetrieveDatabaseListRequest{
		InstanceId: util.PtrString(""),
	}
	result := &RetrieveDatabaseListResponse{}
	result, err := RDS_CLIENT.RetrieveDatabaseList(retrieveDatabaseListRequest)
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
func TestClient_RetrieveTheBinlogList(t *testing.T) {
	retrieveTheBinlogListRequest := &RetrieveTheBinlogListRequest{
		InstanceId: util.PtrString(""),
		Datetime:   util.PtrString(""),
	}
	result := &RetrieveTheBinlogListResponse{}
	result, err := RDS_CLIENT.RetrieveTheBinlogList(retrieveTheBinlogListRequest)
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
func TestClient_SetEncryptionPolicyApi(t *testing.T) {
	setEncryptionPolicyApiRequest := &SetEncryptionPolicyApiRequest{
		InstanceId:   util.PtrString(""),
		BackupPolicy: []*UpdateEncryptPolicyRequest{},
	}
	err := RDS_CLIENT.SetEncryptionPolicyApi(setEncryptionPolicyApiRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SlowLogList(t *testing.T) {
	slowLogListRequest := &SlowLogListRequest{
		InstanceId: util.PtrString(""),
		Datetime:   util.PtrString(""),
	}
	result := &SlowLogListResponse{}
	result, err := RDS_CLIENT.SlowLogList(slowLogListRequest)
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
func TestClient_TableGranularityQueryUserDetailsApi(t *testing.T) {
	tableGranularityQueryUserDetailsApiRequest := &TableGranularityQueryUserDetailsApiRequest{
		InstanceId: util.PtrString(""),
		User:       util.PtrString(""),
		Host:       util.PtrString(""),
	}
	result := &TableGranularityQueryUserDetailsApiResponse{}
	result, err := RDS_CLIENT.TableGranularityQueryUserDetailsApi(tableGranularityQueryUserDetailsApiRequest)
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
func TestClient_TableGranularityQueryUserListApi(t *testing.T) {
	tableGranularityQueryUserListApiRequest := &TableGranularityQueryUserListApiRequest{
		InstanceId: util.PtrString(""),
	}
	result := &TableGranularityQueryUserListApiResponse{}
	result, err := RDS_CLIENT.TableGranularityQueryUserListApi(tableGranularityQueryUserListApiRequest)
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
func TestClient_TableGranularityUpdateUserDescriptionApi(t *testing.T) {
	tableGranularityUpdateUserDescriptionApiRequest := &TableGranularityUpdateUserDescriptionApiRequest{
		InstanceId: util.PtrString(""),
		User:       util.PtrString(""),
		Host:       util.PtrString(""),
		Remark:     util.PtrString(""),
	}
	err := RDS_CLIENT.TableGranularityUpdateUserDescriptionApi(tableGranularityUpdateUserDescriptionApiRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TargetRegionMappingListApi(t *testing.T) {
	targetRegionMappingListApiRequest := &TargetRegionMappingListApiRequest{
		InstanceId: util.PtrString(""),
	}
	result := &TargetRegionMappingListApiResponse{}
	result, err := RDS_CLIENT.TargetRegionMappingListApi(targetRegionMappingListApiRequest)
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
func TestClient_UpdateWhitelist(t *testing.T) {
	updateWhitelistRequest := &UpdateWhitelistRequest{
		InstanceId:  util.PtrString(""),
		Action:      util.PtrString(""),
		SecurityIps: []*string{},
		Tag:         util.PtrString(""),
	}
	err := RDS_CLIENT.UpdateWhitelist(updateWhitelistRequest)
	ExpectEqual(t.Errorf, nil, err)
}
