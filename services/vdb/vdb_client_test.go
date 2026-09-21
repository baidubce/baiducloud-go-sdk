package vdb

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
	VDB_CLIENT *Client
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
	VDB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_AccountListUsingGET(t *testing.T) {
	accountListUsingGETRequest := &AccountListUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &AccountListUsingGETResponse{}
	result, err := VDB_CLIENT.AccountListUsingGET(accountListUsingGETRequest)
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
func TestClient_BindEipUsingPOST(t *testing.T) {
	bindEipUsingPOSTRequest := &BindEipUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		Eip:        util.PtrString(""),
	}
	err := VDB_CLIENT.BindEipUsingPOST(bindEipUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateInstanceUsingPOST(t *testing.T) {
	InstanceParam := &InstanceParam{
		AvailabilityZone:     util.PtrString(""),
		AzInfos:              []*AzInfo{},
		CloneDataAppBackupId: util.PtrString(""),
		CloneDataAppId:       util.PtrString(""),
		Components:           []*MilvusComponent{},
		DataNodeNum:          util.PtrInt32(int32(0)),
		DiskFlavor:           util.PtrInt32(int32(0)),
		DiskType:             util.PtrString(""),
		EnableEmbedding:      util.PtrBool(false),
		EnableEncryption:     util.PtrBool(false),
		EngineVersion:        util.PtrString(""),
		From:                 util.PtrString(""),
		InstanceName:         util.PtrString(""),
		InstanceNum:          util.PtrInt32(int32(0)),
		InstanceType:         util.PtrString(""),
		MasterNodeSpec:       util.PtrString(""),
		MasterNum:            util.PtrInt32(int32(0)),
		NodeSpec:             util.PtrString(""),
		NodeType:             util.PtrString(""),
		OrderId:              util.PtrString(""),
		Password:             util.PtrString(""),
		Port:                 util.PtrInt32(int32(0)),
		ProxyNodeSpec:        util.PtrString(""),
		ProxyNum:             util.PtrInt32(int32(0)),
		ReqSource:            util.PtrString(""),
		SubnetId:             util.PtrString(""),
		SwitchEntrance:       util.PtrString(""),
		VpcId:                util.PtrString(""),
	}
	createInstanceUsingPOSTRequest := &CreateInstanceUsingPOSTRequest{
		EngineType:        util.PtrString(""),
		AutoRenew:         util.PtrBool(false),
		AutoRenewTime:     util.PtrInt32(int32(0)),
		AutoRenewTimeUnit: util.PtrString(""),
		Components:        []*MilvusComponent{},
		Duration:          util.PtrInt32(int32(0)),
		Env:               util.PtrString(""),
		InstanceParam:     InstanceParam,
		ProductType:       util.PtrString(""),
		TimeUnit:          util.PtrString(""),
	}
	result := &CreateInstanceUsingPOSTResponse{}
	result, err := VDB_CLIENT.CreateInstanceUsingPOST(createInstanceUsingPOSTRequest)
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
func TestClient_DeleteInstanceUsingDELETE(t *testing.T) {
	deleteInstanceUsingDELETERequest := &DeleteInstanceUsingDELETERequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.DeleteInstanceUsingDELETE(deleteInstanceUsingDELETERequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRecordUsingDELETE(t *testing.T) {
	deleteRecordUsingDELETERequest := &DeleteRecordUsingDELETERequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		BatchId:    util.PtrString(""),
		BackupId:   util.PtrString(""),
	}
	err := VDB_CLIENT.DeleteRecordUsingDELETE(deleteRecordUsingDELETERequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRecyclerInstance(t *testing.T) {
	deleteRecyclerInstanceRequest := &DeleteRecyclerInstanceRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.DeleteRecyclerInstance(deleteRecyclerInstanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Deleteinstanceusingdelete1(t *testing.T) {
	deleteinstanceusingdelete1Request := &Deleteinstanceusingdelete1Request{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.Deleteinstanceusingdelete1(deleteinstanceusingdelete1Request)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeInstanceConfigs(t *testing.T) {
	describeInstanceConfigsRequest := &DescribeInstanceConfigsRequest{
		InstanceId: util.PtrString(""),
	}
	err := VDB_CLIENT.DescribeInstanceConfigs(describeInstanceConfigsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeInstanceConfigsUsingGET(t *testing.T) {
	describeInstanceConfigsUsingGETRequest := &DescribeInstanceConfigsUsingGETRequest{
		InstanceId: util.PtrString(""),
	}
	err := VDB_CLIENT.DescribeInstanceConfigsUsingGET(describeInstanceConfigsUsingGETRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetConfigUsingGET(t *testing.T) {
	getConfigUsingGETRequest := &GetConfigUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &GetConfigUsingGETResponse{}
	result, err := VDB_CLIENT.GetConfigUsingGET(getConfigUsingGETRequest)
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
func TestClient_GetFreeInstanceQuota(t *testing.T) {
	result := &GetFreeInstanceQuotaResponse{}
	result, err := VDB_CLIENT.GetFreeInstanceQuota()
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
func TestClient_GetFreeInstanceQuotaUsingGET(t *testing.T) {
	result := &GetFreeInstanceQuotaUsingGETResponse{}
	result, err := VDB_CLIENT.GetFreeInstanceQuotaUsingGET()
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
func TestClient_GetInstanceListUsingGET(t *testing.T) {
	getInstanceListUsingGETRequest := &GetInstanceListUsingGETRequest{
		EngineType:   util.PtrString(""),
		InstanceType: util.PtrString(""),
	}
	result := &GetInstanceListUsingGETResponse{}
	result, err := VDB_CLIENT.GetInstanceListUsingGET(getInstanceListUsingGETRequest)
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
func TestClient_GetNodeSpecListUsingGET(t *testing.T) {
	getNodeSpecListUsingGETRequest := &GetNodeSpecListUsingGETRequest{
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.GetNodeSpecListUsingGET(getNodeSpecListUsingGETRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPriceUsingPOST(t *testing.T) {
	InstanceParam := &InstanceParam{
		AvailabilityZone:     util.PtrString(""),
		AzInfos:              []*AzInfo{},
		CloneDataAppBackupId: util.PtrString(""),
		CloneDataAppId:       util.PtrString(""),
		Components:           []*MilvusComponent{},
		DataNodeNum:          util.PtrInt32(int32(0)),
		DiskFlavor:           util.PtrInt32(int32(0)),
		DiskType:             util.PtrString(""),
		EnableEmbedding:      util.PtrBool(false),
		EnableEncryption:     util.PtrBool(false),
		EngineVersion:        util.PtrString(""),
		From:                 util.PtrString(""),
		InstanceName:         util.PtrString(""),
		InstanceNum:          util.PtrInt32(int32(0)),
		InstanceType:         util.PtrString(""),
		MasterNodeSpec:       util.PtrString(""),
		MasterNum:            util.PtrInt32(int32(0)),
		NodeSpec:             util.PtrString(""),
		NodeType:             util.PtrString(""),
		OrderId:              util.PtrString(""),
		Password:             util.PtrString(""),
		Port:                 util.PtrInt32(int32(0)),
		ProxyNodeSpec:        util.PtrString(""),
		ProxyNum:             util.PtrInt32(int32(0)),
		ReqSource:            util.PtrString(""),
		SubnetId:             util.PtrString(""),
		SwitchEntrance:       util.PtrString(""),
		VpcId:                util.PtrString(""),
	}
	getPriceUsingPOSTRequest := &GetPriceUsingPOSTRequest{
		EngineType:        util.PtrString(""),
		AutoRenew:         util.PtrBool(false),
		AutoRenewTime:     util.PtrInt32(int32(0)),
		AutoRenewTimeUnit: util.PtrString(""),
		Components:        []*MilvusComponent{},
		Duration:          util.PtrInt32(int32(0)),
		Env:               util.PtrString(""),
		InstanceParam:     InstanceParam,
		ProductType:       util.PtrString(""),
		TimeUnit:          util.PtrString(""),
	}
	result := &GetPriceUsingPOSTResponse{}
	result, err := VDB_CLIENT.GetPriceUsingPOST(getPriceUsingPOSTRequest)
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
func TestClient_GetQuotaUsingGET(t *testing.T) {
	getQuotaUsingGETRequest := &GetQuotaUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &GetQuotaUsingGETResponse{}
	result, err := VDB_CLIENT.GetQuotaUsingGET(getQuotaUsingGETRequest)
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
func TestClient_GetTLSCertificateUsingGET(t *testing.T) {
	getTLSCertificateUsingGETRequest := &GetTLSCertificateUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &GetTLSCertificateUsingGETResponse{}
	result, err := VDB_CLIENT.GetTLSCertificateUsingGET(getTLSCertificateUsingGETRequest)
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
func TestClient_GetTLSInfoUsingGET(t *testing.T) {
	getTLSInfoUsingGETRequest := &GetTLSInfoUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &GetTLSInfoUsingGETResponse{}
	result, err := VDB_CLIENT.GetTLSInfoUsingGET(getTLSInfoUsingGETRequest)
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
func TestClient_Getinstancelistusingget1(t *testing.T) {
	getinstancelistusingget1Request := &Getinstancelistusingget1Request{
		EngineType: util.PtrString(""),
	}
	result := &Getinstancelistusingget1Response{}
	result, err := VDB_CLIENT.Getinstancelistusingget1(getinstancelistusingget1Request)
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
func TestClient_InstanceDetailUsingGET(t *testing.T) {
	instanceDetailUsingGETRequest := &InstanceDetailUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &InstanceDetailUsingGETResponse{}
	result, err := VDB_CLIENT.InstanceDetailUsingGET(instanceDetailUsingGETRequest)
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
func TestClient_ListRecordsUsingGET(t *testing.T) {
	listRecordsUsingGETRequest := &ListRecordsUsingGETRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		ListOrder:  util.PtrString(""),
		Page:       util.PtrString(""),
		PageSize:   util.PtrString(""),
	}
	result := &ListRecordsUsingGETResponse{}
	result, err := VDB_CLIENT.ListRecordsUsingGET(listRecordsUsingGETRequest)
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
func TestClient_ManualBackupUsingPOST(t *testing.T) {
	manualBackupUsingPOSTRequest := &ManualBackupUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		Comment:    util.PtrString(""),
	}
	err := VDB_CLIENT.ManualBackupUsingPOST(manualBackupUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceConfig(t *testing.T) {
	modifyInstanceConfigRequest := &ModifyInstanceConfigRequest{
		InstanceId:  util.PtrString(""),
		Reason:      util.PtrString(""),
		UserConfigs: []*InstanceConfigUserConfig{},
	}
	err := VDB_CLIENT.ModifyInstanceConfig(modifyInstanceConfigRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyInstanceConfigUsingPOST(t *testing.T) {
	modifyInstanceConfigUsingPOSTRequest := &ModifyInstanceConfigUsingPOSTRequest{
		InstanceId:  util.PtrString(""),
		Reason:      util.PtrString(""),
		UserConfigs: []*InstanceConfigUserConfig{},
	}
	err := VDB_CLIENT.ModifyInstanceConfigUsingPOST(modifyInstanceConfigUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyPasswordUsingPOST(t *testing.T) {
	modifyPasswordUsingPOSTRequest := &ModifyPasswordUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		From:       util.PtrString(""),
		Password:   util.PtrString(""),
		Username:   util.PtrString(""),
	}
	err := VDB_CLIENT.ModifyPasswordUsingPOST(modifyPasswordUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyPublicAccess(t *testing.T) {
	modifyPublicAccessRequest := &ModifyPublicAccessRequest{
		InstanceId:   util.PtrString(""),
		PublicAccess: util.PtrBool(false),
	}
	err := VDB_CLIENT.ModifyPublicAccess(modifyPublicAccessRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyPublicAccessUsingPUT(t *testing.T) {
	modifyPublicAccessUsingPUTRequest := &ModifyPublicAccessUsingPUTRequest{
		InstanceId:   util.PtrString(""),
		PublicAccess: util.PtrBool(false),
	}
	err := VDB_CLIENT.ModifyPublicAccessUsingPUT(modifyPublicAccessUsingPUTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyTLSUsingPUT(t *testing.T) {
	modifyTLSUsingPUTRequest := &ModifyTLSUsingPUTRequest{
		EngineType: util.PtrString(""),
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	err := VDB_CLIENT.ModifyTLSUsingPUT(modifyTLSUsingPUTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PasswordUsingGET(t *testing.T) {
	passwordUsingGETRequest := &PasswordUsingGETRequest{
		InstanceId: util.PtrString(""),
		Username:   util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &PasswordUsingGETResponse{}
	result, err := VDB_CLIENT.PasswordUsingGET(passwordUsingGETRequest)
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
func TestClient_RecoverInstanceUsingPOST(t *testing.T) {
	recoverInstanceUsingPOSTRequest := &RecoverInstanceUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.RecoverInstanceUsingPOST(recoverInstanceUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RecoverUsingPOST(t *testing.T) {
	recoverUsingPOSTRequest := &RecoverUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		BackupId:   util.PtrString(""),
		BatchId:    util.PtrString(""),
		Confirmed:  util.PtrBool(false),
	}
	err := VDB_CLIENT.RecoverUsingPOST(recoverUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResizeInstanceUsingPOST(t *testing.T) {
	resizeInstanceUsingPOSTRequest := &ResizeInstanceUsingPOSTRequest{
		EngineType:     util.PtrString(""),
		Components:     []*MilvusComponent{},
		DataNodeNum:    util.PtrInt32(int32(0)),
		DiskFlavor:     util.PtrInt32(int32(0)),
		DiskType:       util.PtrString(""),
		Env:            util.PtrString(""),
		InstanceId:     util.PtrString(""),
		MasterNodeSpec: util.PtrString(""),
		MasterNum:      util.PtrInt32(int32(0)),
		NodeSpec:       util.PtrString(""),
		NodeType:       util.PtrString(""),
		OrderId:        util.PtrString(""),
		ProxyNodeSpec:  util.PtrString(""),
		ProxyNum:       util.PtrInt32(int32(0)),
	}
	result := &ResizeInstanceUsingPOSTResponse{}
	result, err := VDB_CLIENT.ResizeInstanceUsingPOST(resizeInstanceUsingPOSTRequest)
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
func TestClient_SetCommentUsingPOST(t *testing.T) {
	setCommentUsingPOSTRequest := &SetCommentUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		BackupId:   util.PtrString(""),
		BatchId:    util.PtrString(""),
		Comment:    util.PtrString(""),
	}
	err := VDB_CLIENT.SetCommentUsingPOST(setCommentUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetConfigUsingPOST(t *testing.T) {
	setConfigUsingPOSTRequest := &SetConfigUsingPOSTRequest{
		InstanceId:        util.PtrString(""),
		EngineType:        util.PtrString(""),
		AutoBackupConfig:  util.PtrString(""),
		AutoBackupEnabled: util.PtrBool(false),
		IsEncrypt:         util.PtrString(""),
	}
	err := VDB_CLIENT.SetConfigUsingPOST(setConfigUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindEipUsingPOST(t *testing.T) {
	unbindEipUsingPOSTRequest := &UnbindEipUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	err := VDB_CLIENT.UnbindEipUsingPOST(unbindEipUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceDomain(t *testing.T) {
	updateInstanceDomainRequest := &UpdateInstanceDomainRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		Domain:     util.PtrString(""),
	}
	err := VDB_CLIENT.UpdateInstanceDomain(updateInstanceDomainRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceDomainUsingPOST(t *testing.T) {
	updateInstanceDomainUsingPOSTRequest := &UpdateInstanceDomainUsingPOSTRequest{
		InstanceId: util.PtrString(""),
		EngineType: util.PtrString(""),
		Domain:     util.PtrString(""),
	}
	err := VDB_CLIENT.UpdateInstanceDomainUsingPOST(updateInstanceDomainUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceName(t *testing.T) {
	updateInstanceNameRequest := &UpdateInstanceNameRequest{
		InstanceId:   util.PtrString(""),
		EngineType:   util.PtrString(""),
		InstanceName: util.PtrString(""),
	}
	err := VDB_CLIENT.UpdateInstanceName(updateInstanceNameRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceNameUsingPOST(t *testing.T) {
	updateInstanceNameUsingPOSTRequest := &UpdateInstanceNameUsingPOSTRequest{
		InstanceId:   util.PtrString(""),
		EngineType:   util.PtrString(""),
		InstanceName: util.PtrString(""),
	}
	err := VDB_CLIENT.UpdateInstanceNameUsingPOST(updateInstanceNameUsingPOSTRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ZoneListUsingGET(t *testing.T) {
	zoneListUsingGETRequest := &ZoneListUsingGETRequest{
		From:       util.PtrString(""),
		EngineType: util.PtrString(""),
	}
	result := &ZoneListUsingGETResponse{}
	result, err := VDB_CLIENT.ZoneListUsingGET(zoneListUsingGETRequest)
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
