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
func TestClient_BindTagVolume(t *testing.T) {
	bindTagVolumeRequest := &BindTagVolumeRequest{
		VolumeId:   util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.BindTagVolume(bindTagVolumeRequest)
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
func TestClient_DetachVolume(t *testing.T) {
	detachVolumeRequest := &DetachVolumeRequest{
		VolumeId:   util.PtrString(""),
		InstanceId: util.PtrString(""),
	}
	err := BCC_CLIENT.DetachVolume(detachVolumeRequest)
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
func TestClient_RenameVolume(t *testing.T) {
	renameVolumeRequest := &RenameVolumeRequest{
		VolumeId: util.PtrString(""),
		Name:     util.PtrString(""),
	}
	err := BCC_CLIENT.RenameVolume(renameVolumeRequest)
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
func TestClient_UnbindTagVolume(t *testing.T) {
	unbindTagVolumeRequest := &UnbindTagVolumeRequest{
		VolumeId:   util.PtrString(""),
		ChangeTags: []*TagModel{},
	}
	err := BCC_CLIENT.UnbindTagVolume(unbindTagVolumeRequest)
	ExpectEqual(t.Errorf, nil, err)
}
