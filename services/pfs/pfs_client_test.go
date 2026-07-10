package pfs

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
	PFS_CLIENT *Client
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

	PFS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CancelL2BucketLink(t *testing.T) {
	cancelL2BucketLinkRequest := &CancelL2BucketLinkRequest{
		Action:       util.PtrString(""),
		BucketLinkId: util.PtrString(""),
		InstanceId:   util.PtrString(""),
	}
	result := &CancelL2BucketLinkResponse{}
	result, err := PFS_CLIENT.CancelL2BucketLink(cancelL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateFileset(t *testing.T) {
	createFilesetRequest := &CreateFilesetRequest{
		Action:           util.PtrString(""),
		InstanceId:       util.PtrString(""),
		FilesetName:      util.PtrString(""),
		FilesetPath:      util.PtrString(""),
		BlockQuota:       util.PtrInt32(int32(0)),
		FilesQuota:       util.PtrInt64(int64(0)),
		QpsLimit:         util.PtrInt32(int32(0)),
		BandwidthLimitMb: util.PtrInt32(int32(0)),
	}
	result := &CreateFilesetResponse{}
	result, err := PFS_CLIENT.CreateFileset(createFilesetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateL2BucketLink(t *testing.T) {
	createL2BucketLinkRequest := &CreateL2BucketLinkRequest{
		Action:               util.PtrString(""),
		InstanceId:           util.PtrString(""),
		ConflictPolicy:       util.PtrString(""),
		BucketName:           util.PtrString(""),
		BucketPrefix:         util.PtrString(""),
		ThroughputLimitBytes: util.PtrString(""),
		ReportObjectName:     util.PtrString(""),
		BucketLinkName:       util.PtrString(""),
		TransferType:         util.PtrInt32(int32(0)),
		PfsPath:              util.PtrString(""),
		Cron:                 util.PtrString(""),
		BucketBelongUserId:   util.PtrString(""),
		LccId:                util.PtrString(""),
		Scope:                util.PtrInt32(int32(0)),
	}
	result := &CreateL2BucketLinkResponse{}
	result, err := PFS_CLIENT.CreateL2BucketLink(createL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateL2Policy(t *testing.T) {
	createL2PolicyRequest := &CreateL2PolicyRequest{
		Action:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		PolicyName:   util.PtrString(""),
		Path:         util.PtrString(""),
		ExpiredTime:  util.PtrInt32(int32(0)),
		PfsType:      util.PtrInt32(int32(0)),
		ExecuteTime:  util.PtrInt32(int32(0)),
		BucketName:   util.PtrString(""),
		BucketPrefix: util.PtrString(""),
	}
	result := &CreateL2PolicyResponse{}
	result, err := PFS_CLIENT.CreateL2Policy(createL2PolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreatePfs(t *testing.T) {
	createPfsRequest := &CreatePfsRequest{
		Name:         util.PtrString(""),
		InstanceType: util.PtrString(""),
		Capacity:     util.PtrInt32(int32(0)),
		Zone:         util.PtrString(""),
		SubnetId:     util.PtrString(""),
		Description:  util.PtrString(""),
		Tags:         []*Tag{},
	}
	result := &CreatePfsResponse{}
	result, err := PFS_CLIENT.CreatePfs(createPfsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteFileset(t *testing.T) {
	deleteFilesetRequest := &DeleteFilesetRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		FilesetId:  util.PtrString(""),
	}
	result := &DeleteFilesetResponse{}
	result, err := PFS_CLIENT.DeleteFileset(deleteFilesetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteL2BucketLink(t *testing.T) {
	deleteL2BucketLinkRequest := &DeleteL2BucketLinkRequest{
		InstanceId:   util.PtrString(""),
		BucketLinkId: util.PtrString(""),
	}
	result := &DeleteL2BucketLinkResponse{}
	result, err := PFS_CLIENT.DeleteL2BucketLink(deleteL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteL2Policy(t *testing.T) {
	deleteL2PolicyRequest := &DeleteL2PolicyRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		PolicyId:   util.PtrString(""),
	}
	result := &DeleteL2PolicyResponse{}
	result, err := PFS_CLIENT.DeleteL2Policy(deleteL2PolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletePfs(t *testing.T) {
	deletePfsRequest := &DeletePfsRequest{
		InstanceId: util.PtrString(""),
	}
	err := PFS_CLIENT.DeletePfs(deletePfsRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescFileset(t *testing.T) {
	descFilesetRequest := &DescFilesetRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		FilesetId:  util.PtrString(""),
	}
	result := &DescFilesetResponse{}
	result, err := PFS_CLIENT.DescFileset(descFilesetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescL2BucketLink(t *testing.T) {
	descL2BucketLinkRequest := &DescL2BucketLinkRequest{
		Action:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		BucketLinkId: util.PtrString(""),
	}
	result := &DescL2BucketLinkResponse{}
	result, err := PFS_CLIENT.DescL2BucketLink(descL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescL2Policy(t *testing.T) {
	descL2PolicyRequest := &DescL2PolicyRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		PolicyId:   util.PtrString(""),
	}
	result := &DescL2PolicyResponse{}
	result, err := PFS_CLIENT.DescL2Policy(descL2PolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescPfs(t *testing.T) {
	descPfsRequest := &DescPfsRequest{
		InstanceId: util.PtrString(""),
	}
	result := &DescPfsResponse{}
	result, err := PFS_CLIENT.DescPfs(descPfsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_InstanceListClients(t *testing.T) {
	instanceListClientsRequest := &InstanceListClientsRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Manner:     util.PtrString(""),
		Marker:     util.PtrString(""),
	}
	result := &InstanceListClientsResponse{}
	result, err := PFS_CLIENT.InstanceListClients(instanceListClientsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListFileset(t *testing.T) {
	listFilesetRequest := &ListFilesetRequest{
		Action:      util.PtrString(""),
		InstanceId:  util.PtrString(""),
		FilesetId:   util.PtrString(""),
		FilesetName: util.PtrString(""),
		Manner:      util.PtrString(""),
		Order:       util.PtrString(""),
		OrderBy:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &ListFilesetResponse{}
	result, err := PFS_CLIENT.ListFileset(listFilesetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListL2BucketLink(t *testing.T) {
	listL2BucketLinkRequest := &ListL2BucketLinkRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Manner:     util.PtrString(""),
		Marker:     util.PtrString(""),
	}
	result := &ListL2BucketLinkResponse{}
	result, err := PFS_CLIENT.ListL2BucketLink(listL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListL2Policy(t *testing.T) {
	listL2PolicyRequest := &ListL2PolicyRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		Marker:     util.PtrString(""),
	}
	result := &ListL2PolicyResponse{}
	result, err := PFS_CLIENT.ListL2Policy(listL2PolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListPfs(t *testing.T) {
	listPfsRequest := &ListPfsRequest{
		Manner:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		Marker:    util.PtrString(""),
		FilterTag: util.PtrString(""),
	}
	result := &ListPfsResponse{}
	result, err := PFS_CLIENT.ListPfs(listPfsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_LstPerL2BktLnkExecLog(t *testing.T) {
	lstPerL2BktLnkExecLogRequest := &LstPerL2BktLnkExecLogRequest{
		Action:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		BucketLinkId: util.PtrString(""),
		StartTime:    util.PtrInt32(int32(0)),
		EndTime:      util.PtrInt32(int32(0)),
	}
	result := &LstPerL2BktLnkExecLogResponse{}
	result, err := PFS_CLIENT.LstPerL2BktLnkExecLog(lstPerL2BktLnkExecLogRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MountTargetListClients(t *testing.T) {
	mountTargetListClientsRequest := &MountTargetListClientsRequest{
		Action:        util.PtrString(""),
		MountTargetId: util.PtrString(""),
		MaxKeys:       util.PtrInt32(int32(0)),
		Manner:        util.PtrString(""),
		Marker:        util.PtrString(""),
	}
	result := &MountTargetListClientsResponse{}
	result, err := PFS_CLIENT.MountTargetListClients(mountTargetListClientsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PauseL2BucketLink(t *testing.T) {
	pauseL2BucketLinkRequest := &PauseL2BucketLinkRequest{
		Action:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		BucketLinkId: util.PtrString(""),
	}
	result := &PauseL2BucketLinkResponse{}
	result, err := PFS_CLIENT.PauseL2BucketLink(pauseL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QryL2PolExecDetail(t *testing.T) {
	qryL2PolExecDetailRequest := &QryL2PolExecDetailRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		PolicyId:   util.PtrString(""),
		JobId:      util.PtrString(""),
	}
	result := &QryL2PolExecDetailResponse{}
	result, err := PFS_CLIENT.QryL2PolExecDetail(qryL2PolExecDetailRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QryL2PolExecLog(t *testing.T) {
	qryL2PolExecLogRequest := &QryL2PolExecLogRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		PolicyId:   util.PtrString(""),
		StartTime:  util.PtrInt32(int32(0)),
		EndTime:    util.PtrInt32(int32(0)),
	}
	result := &QryL2PolExecLogResponse{}
	result, err := PFS_CLIENT.QryL2PolExecLog(qryL2PolExecLogRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResumeL2BucketLink(t *testing.T) {
	resumeL2BucketLinkRequest := &ResumeL2BucketLinkRequest{
		Action:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		BucketLinkId: util.PtrString(""),
	}
	result := &ResumeL2BucketLinkResponse{}
	result, err := PFS_CLIENT.ResumeL2BucketLink(resumeL2BucketLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdPerL2BktLnkInfo(t *testing.T) {
	updPerL2BktLnkInfoRequest := &UpdPerL2BktLnkInfoRequest{
		Action:                  util.PtrString(""),
		InstanceId:              util.PtrString(""),
		BucketLinkId:            util.PtrString(""),
		NewCron:                 util.PtrString(""),
		NewBucketLinkName:       util.PtrString(""),
		NewConflictPolicy:       util.PtrInt32(int32(0)),
		NewThroughputLimitBytes: util.PtrInt32(int32(0)),
		NewScope:                util.PtrInt32(int32(0)),
	}
	result := &UpdPerL2BktLnkInfoResponse{}
	result, err := PFS_CLIENT.UpdPerL2BktLnkInfo(updPerL2BktLnkInfoRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateFileset(t *testing.T) {
	updateFilesetRequest := &UpdateFilesetRequest{
		Action:           util.PtrString(""),
		InstanceId:       util.PtrString(""),
		FilesetId:        util.PtrString(""),
		FilesetName:      util.PtrString(""),
		BlockQuota:       util.PtrInt32(int32(0)),
		FilesQuota:       util.PtrInt64(int64(0)),
		QpsLimit:         util.PtrInt32(int32(0)),
		BandwidthLimitMb: util.PtrInt32(int32(0)),
	}
	result := &UpdateFilesetResponse{}
	result, err := PFS_CLIENT.UpdateFileset(updateFilesetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateL2Policy(t *testing.T) {
	updateL2PolicyRequest := &UpdateL2PolicyRequest{
		Action:        util.PtrString(""),
		InstanceId:    util.PtrString(""),
		PolicyId:      util.PtrString(""),
		NewPolicyName: util.PtrString(""),
		ExpiredTime:   util.PtrInt32(int32(0)),
		ExecuteTime:   util.PtrInt32(int32(0)),
		BucketName:    util.PtrString(""),
		BucketPrefix:  util.PtrString(""),
	}
	result := &UpdateL2PolicyResponse{}
	result, err := PFS_CLIENT.UpdateL2Policy(updateL2PolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePFSTag(t *testing.T) {
	updatePFSTagRequest := &UpdatePFSTagRequest{
		InstanceId: []*string{},
		Tags:       []*Tag{},
	}
	err := PFS_CLIENT.UpdatePFSTag(updatePFSTagRequest)
	ExpectEqual(t.Errorf, nil, err)
}
