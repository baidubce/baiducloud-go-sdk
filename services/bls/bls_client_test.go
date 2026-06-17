package bls

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
	BLS_CLIENT *Client
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

	BLS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AsyncSearch(t *testing.T) {
	Highlight := &Highlight{
		PreTags:  []*string{},
		PostTags: []*string{},
	}
	Sort := make(map[string]map[string]string)
	asyncSearchRequest := &AsyncSearchRequest{
		Name:        util.PtrString(""),
		Query:       nil,
		Aggs:        nil,
		Fields:      []*string{},
		Sort:        &Sort,
		SearchAfter: []*string{},
		Highlight:   Highlight,
		Size:        util.PtrInt32(int32(0)),
	}
	result := &AsyncSearchResponse{}
	result, err := BLS_CLIENT.AsyncSearch(asyncSearchRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchGetLogStore(t *testing.T) {
	batchGetLogStoreRequest := &BatchGetLogStoreRequest{
		LogStores: []*LogStoreBatchRequest{},
	}
	result := &BatchGetLogStoreResponse{}
	result, err := BLS_CLIENT.BatchGetLogStore(batchGetLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BulkDeleteLogShipper(t *testing.T) {
	bulkDeleteLogShipperRequest := &BulkDeleteLogShipperRequest{
		LogShipperIDs: []*string{},
	}
	result := &BulkDeleteLogShipperResponse{}
	result, err := BLS_CLIENT.BulkDeleteLogShipper(bulkDeleteLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BulkSetLogShipperStatus(t *testing.T) {
	bulkSetLogShipperStatusRequest := &BulkSetLogShipperStatusRequest{
		DesiredStatus: util.PtrString(""),
		LogShipperIDs: []*string{},
	}
	result := &BulkSetLogShipperStatusResponse{}
	result, err := BLS_CLIENT.BulkSetLogShipperStatus(bulkSetLogShipperStatusRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAlarmPolicy(t *testing.T) {
	Schedule := &Schedule{
		IntervalMinute: util.PtrInt32(int32(0)),
		FixTimeMinute:  util.PtrInt32(int32(0)),
		DayOfWeek:      util.PtrInt32(int32(0)),
	}
	createAlarmPolicyRequest := &CreateAlarmPolicyRequest{
		Name:                 util.PtrString(""),
		Objects:              []*LogStore{},
		Targets:              []*Target{},
		TriggerConditions:    []*TriggerCondition{},
		Groups:               []*string{},
		Schedule:             Schedule,
		PendingCount:         util.PtrInt32(int32(0)),
		RepeatIntervalMinute: util.PtrInt32(int32(0)),
		RecoverWithoutNotice: util.PtrBool(false),
		State:                util.PtrString(""),
		NoticeState:          util.PtrString(""),
		Notices:              []*Notice{},
		NoticeRawLogs:        []*NoticeRawLog{},
	}
	result := &CreateAlarmPolicyResponse{}
	result, err := BLS_CLIENT.CreateAlarmPolicy(createAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateDownloadTask(t *testing.T) {
	createDownloadTaskRequest := &CreateDownloadTaskRequest{
		Name:           util.PtrString(""),
		Project:        util.PtrString(""),
		LogStoreName:   util.PtrString(""),
		LogStreamName:  util.PtrString(""),
		Query:          util.PtrString(""),
		QueryStartTime: util.PtrString(""),
		QueryEndTime:   util.PtrString(""),
		Format:         util.PtrString(""),
		Limit:          util.PtrInt32(int32(0)),
		Order:          util.PtrString(""),
		FileDir:        util.PtrString(""),
	}
	result := &CreateDownloadTaskResponse{}
	result, err := BLS_CLIENT.CreateDownloadTask(createDownloadTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateFastQuery(t *testing.T) {
	createFastQueryRequest := &CreateFastQueryRequest{
		FastQueryName: util.PtrString(""),
		Query:         util.PtrString(""),
		Description:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStoreName:  util.PtrString(""),
		LogStreamName: util.PtrString(""),
	}
	result := &CreateFastQueryResponse{}
	result, err := BLS_CLIENT.CreateFastQuery(createFastQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateIndex(t *testing.T) {
	Fields := make(map[string]*IndexField)
	createIndexRequest := &CreateIndexRequest{
		LogStoreName:   util.PtrString(""),
		Project:        util.PtrString(""),
		Fulltext:       util.PtrBool(false),
		CaseSensitive:  util.PtrBool(false),
		IncludeChinese: util.PtrBool(false),
		Separators:     util.PtrString(""),
		Fields:         &Fields,
	}
	result := &CreateIndexResponse{}
	result, err := BLS_CLIENT.CreateIndex(createIndexRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateLogShipper(t *testing.T) {
	DestConfig := &DestConfig{
		BOSPath:                  util.PtrString(""),
		PartitionFormatTS:        util.PtrString(""),
		PartitionFormatLogStream: util.PtrBool(false),
		MaxObjectSize:            util.PtrInt32(int32(0)),
		CompressType:             util.PtrString(""),
		DeliverInterval:          util.PtrInt32(int32(0)),
		StorageFormat:            util.PtrString(""),
		CsvHeadline:              util.PtrBool(false),
		CsvDelimiter:             util.PtrString(""),
		CsvQuote:                 util.PtrString(""),
		NullIdentifier:           util.PtrString(""),
		SelectedColumnName:       util.PtrString(""),
		SelectedColumnType:       util.PtrString(""),
		FieldsName:               []*string{},
		FieldsType:               []*string{},
		ShipperType:              util.PtrString(""),
		KafkaConfig:              util.PtrString(""),
		DestType:                 util.PtrString(""),
		LogStore:                 util.PtrString(""),
		RateLimit:                util.PtrInt32(int32(0)),
		ClientCount:              util.PtrInt32(int32(0)),
	}
	createLogShipperRequest := &CreateLogShipperRequest{
		Project:        util.PtrString(""),
		LogStoreName:   util.PtrString(""),
		LogShipperName: util.PtrString(""),
		StartTime:      util.PtrString(""),
		DestType:       util.PtrString(""),
		DestConfig:     DestConfig,
	}
	result := &CreateLogShipperResponse{}
	result, err := BLS_CLIENT.CreateLogShipper(createLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateLogStore(t *testing.T) {
	Index := &Index{
		Fulltext:       util.PtrBool(false),
		CaseSensitive:  util.PtrBool(false),
		IncludeChinese: util.PtrBool(false),
		Separators:     util.PtrString(""),
		Fields:         make(map[string]*Field),
	}
	createLogStoreRequest := &CreateLogStoreRequest{
		Project:               util.PtrString(""),
		LogStoreName:          util.PtrString(""),
		Retention:             util.PtrInt32(int32(0)),
		Tags:                  []*Tag{},
		Index:                 Index,
		ShardCount:            util.PtrInt32(int32(0)),
		MaxShardCount:         util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		IndexEnabled:          util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
		EnableHotRetention:    util.PtrBool(false),
	}
	result := &CreateLogStoreResponse{}
	result, err := BLS_CLIENT.CreateLogStore(createLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateLogStoreTemplate(t *testing.T) {
	Template := &Template{
		Retention:             util.PtrInt32(int32(0)),
		ShardCount:            util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		MaxShardCount:         util.PtrInt32(int32(0)),
		EnableHotRetention:    util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
		Index: &Index{
			Fulltext:       util.PtrBool(false),
			CaseSensitive:  util.PtrBool(false),
			IncludeChinese: util.PtrBool(false),
			Separators:     util.PtrString(""),
			Fields:         make(map[string]*Field),
		},
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		CreatedTimestamp: util.PtrString(""),
		UpdatedTimestamp: util.PtrString(""),
	}
	createLogStoreTemplateRequest := &CreateLogStoreTemplateRequest{
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		Template:         Template,
	}
	result := &CreateLogStoreTemplateResponse{}
	result, err := BLS_CLIENT.CreateLogStoreTemplate(createLogStoreTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateLogStoreView(t *testing.T) {
	createLogStoreViewRequest := &CreateLogStoreViewRequest{
		Project:   util.PtrString(""),
		Name:      util.PtrString(""),
		Logstores: []*LogStore{},
	}
	err := BLS_CLIENT.CreateLogStoreView(createLogStoreViewRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateProject(t *testing.T) {
	createProjectRequest := &CreateProjectRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateProjectResponse{}
	result, err := BLS_CLIENT.CreateProject(createProjectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateTask(t *testing.T) {
	Config := &TaskConfig{
		SrcConfig: &SrcConfig{
			SrcType:        util.PtrString(""),
			LogType:        util.PtrString(""),
			SrcDir:         util.PtrString(""),
			MatchedPattern: util.PtrString(""),
			IgnorePattern:  util.PtrString(""),
			TimeFormat:     util.PtrString(""),
			Ttl:            util.PtrInt32(int32(0)),
			UseMultiline:   util.PtrBool(false),
			MultilineRegex: util.PtrString(""),
			RecursiveDir:   util.PtrBool(false),
			ProcessType:    util.PtrString(""),
			ProcessConfig: &ProcessConfig{
				Regex:            util.PtrString(""),
				Separator:        util.PtrString(""),
				CustomSeparator:  util.PtrString(""),
				Quote:            util.PtrString(""),
				KvKeyIndex:       util.PtrInt32(int32(0)),
				KvValueIndex:     util.PtrInt32(int32(0)),
				SampleLog:        util.PtrString(""),
				Keys:             util.PtrString(""),
				DataType:         util.PtrString(""),
				DiscardOnFailure: util.PtrBool(false),
				KeepOriginal:     util.PtrBool(false),
			},
			LogTime:        util.PtrString(""),
			TimestampKey:   util.PtrString(""),
			DateFormat:     util.PtrString(""),
			FilterExpr:     util.PtrString(""),
			AdditionConfig: nil,
			MetaEnv:        []*string{},
			MetaLabel:      []*string{},
			MetaContainer:  []*string{},
			MetaToFields:   util.PtrBool(false),
			HarvesterLimit: util.PtrInt32(int32(0)),
		},
		DestConfig: &DestConfig{
			BOSPath:                  util.PtrString(""),
			PartitionFormatTS:        util.PtrString(""),
			PartitionFormatLogStream: util.PtrBool(false),
			MaxObjectSize:            util.PtrInt32(int32(0)),
			CompressType:             util.PtrString(""),
			DeliverInterval:          util.PtrInt32(int32(0)),
			StorageFormat:            util.PtrString(""),
			CsvHeadline:              util.PtrBool(false),
			CsvDelimiter:             util.PtrString(""),
			CsvQuote:                 util.PtrString(""),
			NullIdentifier:           util.PtrString(""),
			SelectedColumnName:       util.PtrString(""),
			SelectedColumnType:       util.PtrString(""),
			FieldsName:               []*string{},
			FieldsType:               []*string{},
			ShipperType:              util.PtrString(""),
			KafkaConfig:              util.PtrString(""),
			DestType:                 util.PtrString(""),
			LogStore:                 util.PtrString(""),
			RateLimit:                util.PtrInt32(int32(0)),
			ClientCount:              util.PtrInt32(int32(0)),
		},
	}
	createTaskRequest := &CreateTaskRequest{
		Name:   util.PtrString(""),
		Config: Config,
		Hosts:  []*Host{},
		Tags:   []*Tag{},
	}
	result := &CreateTaskResponse{}
	result, err := BLS_CLIENT.CreateTask(createTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAlarmPolicy(t *testing.T) {
	deleteAlarmPolicyRequest := &DeleteAlarmPolicyRequest{
		Name: []*string{},
	}
	result := &DeleteAlarmPolicyResponse{}
	result, err := BLS_CLIENT.DeleteAlarmPolicy(deleteAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteDownloadTask(t *testing.T) {
	deleteDownloadTaskRequest := &DeleteDownloadTaskRequest{
		Uuid: util.PtrString(""),
	}
	result := &DeleteDownloadTaskResponse{}
	result, err := BLS_CLIENT.DeleteDownloadTask(deleteDownloadTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteFastQuery(t *testing.T) {
	deleteFastQueryRequest := &DeleteFastQueryRequest{
		FastQueryName: util.PtrString(""),
	}
	result := &DeleteFastQueryResponse{}
	result, err := BLS_CLIENT.DeleteFastQuery(deleteFastQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIndex(t *testing.T) {
	deleteIndexRequest := &DeleteIndexRequest{
		LogStoreName: util.PtrString(""),
		Project:      util.PtrString(""),
	}
	result := &DeleteIndexResponse{}
	result, err := BLS_CLIENT.DeleteIndex(deleteIndexRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteLogStore(t *testing.T) {
	deleteLogStoreRequest := &DeleteLogStoreRequest{
		LogStoreName: util.PtrString(""),
		Project:      util.PtrString(""),
	}
	result := &DeleteLogStoreResponse{}
	result, err := BLS_CLIENT.DeleteLogStore(deleteLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteLogStoreTemplates(t *testing.T) {
	deleteLogStoreTemplatesRequest := &DeleteLogStoreTemplatesRequest{
		Names: []*string{},
	}
	result := &DeleteLogStoreTemplatesResponse{}
	result, err := BLS_CLIENT.DeleteLogStoreTemplates(deleteLogStoreTemplatesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteLogStoreView(t *testing.T) {
	deleteLogStoreViewRequest := &DeleteLogStoreViewRequest{
		Name:    util.PtrString(""),
		Project: util.PtrString(""),
	}
	result := &DeleteLogStoreViewResponse{}
	result, err := BLS_CLIENT.DeleteLogStoreView(deleteLogStoreViewRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteProject(t *testing.T) {
	deleteProjectRequest := &DeleteProjectRequest{
		Uuid: util.PtrString(""),
	}
	result := &DeleteProjectResponse{}
	result, err := BLS_CLIENT.DeleteProject(deleteProjectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSingleLogShipper(t *testing.T) {
	deleteSingleLogShipperRequest := &DeleteSingleLogShipperRequest{
		LogShipperID: util.PtrString(""),
	}
	result := &DeleteSingleLogShipperResponse{}
	result, err := BLS_CLIENT.DeleteSingleLogShipper(deleteSingleLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeAlarmPolicy(t *testing.T) {
	describeAlarmPolicyRequest := &DescribeAlarmPolicyRequest{
		PolicyName: util.PtrString(""),
	}
	result := &DescribeAlarmPolicyResponse{}
	result, err := BLS_CLIENT.DescribeAlarmPolicy(describeAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeAlarmRecord(t *testing.T) {
	describeAlarmRecordRequest := &DescribeAlarmRecordRequest{
		AlarmId: util.PtrString(""),
	}
	result := &DescribeAlarmRecordResponse{}
	result, err := BLS_CLIENT.DescribeAlarmRecord(describeAlarmRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeDownloadTask(t *testing.T) {
	describeDownloadTaskRequest := &DescribeDownloadTaskRequest{
		Uuid: util.PtrString(""),
	}
	result := &DescribeDownloadTaskResponse{}
	result, err := BLS_CLIENT.DescribeDownloadTask(describeDownloadTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeFastQuery(t *testing.T) {
	describeFastQueryRequest := &DescribeFastQueryRequest{
		FastQueryName: util.PtrString(""),
	}
	result := &DescribeFastQueryResponse{}
	result, err := BLS_CLIENT.DescribeFastQuery(describeFastQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeIndex(t *testing.T) {
	describeIndexRequest := &DescribeIndexRequest{
		LogStoreName: util.PtrString(""),
		Project:      util.PtrString(""),
	}
	result := &DescribeIndexResponse{}
	result, err := BLS_CLIENT.DescribeIndex(describeIndexRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeLogStore(t *testing.T) {
	describeLogStoreRequest := &DescribeLogStoreRequest{
		LogStoreName: util.PtrString(""),
		Project:      util.PtrString(""),
	}
	result := &DescribeLogStoreResponse{}
	result, err := BLS_CLIENT.DescribeLogStore(describeLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeLogStoreTemplate(t *testing.T) {
	describeLogStoreTemplateRequest := &DescribeLogStoreTemplateRequest{
		Name: util.PtrString(""),
	}
	result := &DescribeLogStoreTemplateResponse{}
	result, err := BLS_CLIENT.DescribeLogStoreTemplate(describeLogStoreTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeLogStoreTemplates(t *testing.T) {
	describeLogStoreTemplatesRequest := &DescribeLogStoreTemplatesRequest{
		Name:     util.PtrString(""),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &DescribeLogStoreTemplatesResponse{}
	result, err := BLS_CLIENT.DescribeLogStoreTemplates(describeLogStoreTemplatesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeLogStoreView(t *testing.T) {
	describeLogStoreViewRequest := &DescribeLogStoreViewRequest{
		Name:    util.PtrString(""),
		Project: util.PtrString(""),
	}
	result := &DescribeLogStoreViewResponse{}
	result, err := BLS_CLIENT.DescribeLogStoreView(describeLogStoreViewRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DescribeProject(t *testing.T) {
	describeProjectRequest := &DescribeProjectRequest{
		Uuid: util.PtrString(""),
	}
	result := &DescribeProjectResponse{}
	result, err := BLS_CLIENT.DescribeProject(describeProjectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DisableAlarmPolicy(t *testing.T) {
	disableAlarmPolicyRequest := &DisableAlarmPolicyRequest{
		Name: []*string{},
	}
	result := &DisableAlarmPolicyResponse{}
	result, err := BLS_CLIENT.DisableAlarmPolicy(disableAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableAlarmPolicy(t *testing.T) {
	enableAlarmPolicyRequest := &EnableAlarmPolicyRequest{
		Name: []*string{},
	}
	result := &EnableAlarmPolicyResponse{}
	result, err := BLS_CLIENT.EnableAlarmPolicy(enableAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FieldCaps(t *testing.T) {
	fieldCapsRequest := &FieldCapsRequest{
		Name:   util.PtrString(""),
		Fields: []*string{},
	}
	result := &FieldCapsResponse{}
	result, err := BLS_CLIENT.FieldCaps(fieldCapsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetDownloadTaskLink(t *testing.T) {
	getDownloadTaskLinkRequest := &GetDownloadTaskLinkRequest{
		Uuid: util.PtrString(""),
	}
	result := &GetDownloadTaskLinkResponse{}
	result, err := BLS_CLIENT.GetDownloadTaskLink(getDownloadTaskLinkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetLogShipper(t *testing.T) {
	getLogShipperRequest := &GetLogShipperRequest{
		LogShipperID: util.PtrString(""),
	}
	result := &GetLogShipperResponse{}
	result, err := BLS_CLIENT.GetLogShipper(getLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAlarmExecutionStats(t *testing.T) {
	listAlarmExecutionStatsRequest := &ListAlarmExecutionStatsRequest{
		PolicyId:      util.PtrString(""),
		PolicyName:    util.PtrString(""),
		LogStoreName:  util.PtrString(""),
		States:        []*string{},
		NoticeStates:  []*string{},
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &ListAlarmExecutionStatsResponse{}
	result, err := BLS_CLIENT.ListAlarmExecutionStats(listAlarmExecutionStatsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAlarmExecutions(t *testing.T) {
	listAlarmExecutionsRequest := &ListAlarmExecutionsRequest{
		PolicyId:      util.PtrString(""),
		LogStoreName:  util.PtrString(""),
		State:         util.PtrString(""),
		NoticeStates:  []*string{},
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		OrderBy:       util.PtrString(""),
		Order:         util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &ListAlarmExecutionsResponse{}
	result, err := BLS_CLIENT.ListAlarmExecutions(listAlarmExecutionsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAlarmPolicy(t *testing.T) {
	listAlarmPolicyRequest := &ListAlarmPolicyRequest{
		PolicyNamePattern:   util.PtrString(""),
		PolicyIdPattern:     util.PtrString(""),
		LogStoreNamePattern: util.PtrString(""),
		State:               util.PtrString(""),
		NoticeState:         util.PtrString(""),
		OrderBy:             util.PtrString(""),
		Order:               util.PtrString(""),
		PageNo:              util.PtrInt32(int32(0)),
		PageSize:            util.PtrInt32(int32(0)),
	}
	result := &ListAlarmPolicyResponse{}
	result, err := BLS_CLIENT.ListAlarmPolicy(listAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAlarmRecord(t *testing.T) {
	listAlarmRecordRequest := &ListAlarmRecordRequest{
		PolicyNamePattern:   util.PtrString(""),
		PolicyIdPattern:     util.PtrString(""),
		LogStoreNamePattern: util.PtrString(""),
		Level:               util.PtrString(""),
		State:               util.PtrString(""),
		StartDateTime:       util.PtrString(""),
		EndDateTime:         util.PtrString(""),
		OrderBy:             util.PtrString(""),
		Order:               util.PtrString(""),
		PageNo:              util.PtrInt32(int32(0)),
		PageSize:            util.PtrInt32(int32(0)),
	}
	result := &ListAlarmRecordResponse{}
	result, err := BLS_CLIENT.ListAlarmRecord(listAlarmRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListDownloadTask(t *testing.T) {
	listDownloadTaskRequest := &ListDownloadTaskRequest{
		Project:      util.PtrString(""),
		LogStoreName: util.PtrString(""),
		OrderBy:      util.PtrString(""),
		Order:        util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &ListDownloadTaskResponse{}
	result, err := BLS_CLIENT.ListDownloadTask(listDownloadTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListFastQuery(t *testing.T) {
	listFastQueryRequest := &ListFastQueryRequest{
		Project:      util.PtrString(""),
		LogStoreName: util.PtrString(""),
		NamePattern:  util.PtrString(""),
		Order:        util.PtrString(""),
		OrderBy:      util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &ListFastQueryResponse{}
	result, err := BLS_CLIENT.ListFastQuery(listFastQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLogShipper(t *testing.T) {
	listLogShipperRequest := &ListLogShipperRequest{
		LogShipperID:   util.PtrString(""),
		LogShipperName: util.PtrString(""),
		Project:        util.PtrString(""),
		LogStoreName:   util.PtrString(""),
		DestType:       util.PtrString(""),
		Status:         util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		PageNo:         util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &ListLogShipperResponse{}
	result, err := BLS_CLIENT.ListLogShipper(listLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLogShipperRecord(t *testing.T) {
	listLogShipperRecordRequest := &ListLogShipperRecordRequest{
		LogShipperID: util.PtrString(""),
		SinceHours:   util.PtrInt32(int32(0)),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &ListLogShipperRecordResponse{}
	result, err := BLS_CLIENT.ListLogShipperRecord(listLogShipperRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLogStore(t *testing.T) {
	listLogStoreRequest := &ListLogStoreRequest{
		Project:     util.PtrString(""),
		NamePattern: util.PtrString(""),
		Order:       util.PtrString(""),
		OrderBy:     util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &ListLogStoreResponse{}
	result, err := BLS_CLIENT.ListLogStore(listLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLogStoreView(t *testing.T) {
	listLogStoreViewRequest := &ListLogStoreViewRequest{
		Project:  util.PtrString(""),
		Name:     util.PtrString(""),
		Order:    util.PtrString(""),
		OrderBy:  util.PtrString(""),
		PageNo:   util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &ListLogStoreViewResponse{}
	result, err := BLS_CLIENT.ListLogStoreView(listLogStoreViewRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListLogStream(t *testing.T) {
	listLogStreamRequest := &ListLogStreamRequest{
		Project:      util.PtrString(""),
		LogStoreName: util.PtrString(""),
		NamePattern:  util.PtrString(""),
		Order:        util.PtrString(""),
		OrderBy:      util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &ListLogStreamResponse{}
	result, err := BLS_CLIENT.ListLogStream(listLogStreamRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListProject(t *testing.T) {
	listProjectRequest := &ListProjectRequest{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		OrderBy:     util.PtrString(""),
		Order:       util.PtrString(""),
		PageNo:      util.PtrInt32(int32(0)),
		PageSize:    util.PtrInt32(int32(0)),
	}
	result := &ListProjectResponse{}
	result, err := BLS_CLIENT.ListProject(listProjectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PullLogRecord(t *testing.T) {
	pullLogRecordRequest := &PullLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		LogStreamName: util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		Limit:         util.PtrInt32(int32(0)),
		Marker:        util.PtrString(""),
	}
	result := &PullLogRecordResponse{}
	result, err := BLS_CLIENT.PullLogRecord(pullLogRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PushLogRecord(t *testing.T) {
	pushLogRecordRequest := &PushLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
		BlsType:       util.PtrString(""),
		LogRecords:    []*LogRecord{},
		Tags:          []*LogTag{},
	}
	result := &PushLogRecordResponse{}
	result, err := BLS_CLIENT.PushLogRecord(pushLogRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryLogHistogram(t *testing.T) {
	queryLogHistogramRequest := &QueryLogHistogramRequest{
		LogStoreName:  util.PtrString(""),
		Query:         util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
	}
	result := &QueryLogHistogramResponse{}
	result, err := BLS_CLIENT.QueryLogHistogram(queryLogHistogramRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryLogRecord(t *testing.T) {
	queryLogRecordRequest := &QueryLogRecordRequest{
		LogStoreName:  util.PtrString(""),
		Query:         util.PtrString(""),
		StartDateTime: util.PtrString(""),
		EndDateTime:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStreamName: util.PtrString(""),
		Marker:        util.PtrString(""),
		Limit:         util.PtrInt32(int32(0)),
		Sort:          util.PtrString(""),
		PageNo:        util.PtrInt32(int32(0)),
		PageSize:      util.PtrInt32(int32(0)),
	}
	result := &QueryLogRecordResponse{}
	result, err := BLS_CLIENT.QueryLogRecord(queryLogRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResolveIndex(t *testing.T) {
	resolveIndexRequest := &ResolveIndexRequest{
		Name: util.PtrString(""),
	}
	result := &ResolveIndexResponse{}
	result, err := BLS_CLIENT.ResolveIndex(resolveIndexRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SetSingleLogShipperStatus(t *testing.T) {
	setSingleLogShipperStatusRequest := &SetSingleLogShipperStatusRequest{
		LogShipperID:  util.PtrString(""),
		DesiredStatus: util.PtrString(""),
	}
	result := &SetSingleLogShipperStatusResponse{}
	result, err := BLS_CLIENT.SetSingleLogShipperStatus(setSingleLogShipperStatusRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TermsEnum(t *testing.T) {
	termsEnumRequest := &TermsEnumRequest{
		Name:        util.PtrString(""),
		Field:       util.PtrString(""),
		BlsString:   util.PtrString(""),
		Size:        util.PtrInt32(int32(0)),
		IndexFilter: nil,
	}
	result := &TermsEnumResponse{}
	result, err := BLS_CLIENT.TermsEnum(termsEnumRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAlarmPolicy(t *testing.T) {
	Schedule := &Schedule{
		IntervalMinute: util.PtrInt32(int32(0)),
		FixTimeMinute:  util.PtrInt32(int32(0)),
		DayOfWeek:      util.PtrInt32(int32(0)),
	}
	updateAlarmPolicyRequest := &UpdateAlarmPolicyRequest{
		Name:                 util.PtrString(""),
		Objects:              []*LogStore{},
		Targets:              []*Target{},
		TriggerConditions:    []*TriggerCondition{},
		Groups:               []*string{},
		Schedule:             Schedule,
		PendingCount:         util.PtrInt32(int32(0)),
		RepeatIntervalMinute: util.PtrInt32(int32(0)),
		RecoverWithoutNotice: util.PtrBool(false),
		State:                util.PtrString(""),
		NoticeState:          util.PtrString(""),
		Notices:              []*Notice{},
		NoticeRawLogs:        []*NoticeRawLog{},
	}
	result := &UpdateAlarmPolicyResponse{}
	result, err := BLS_CLIENT.UpdateAlarmPolicy(updateAlarmPolicyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateFastQuery(t *testing.T) {
	updateFastQueryRequest := &UpdateFastQueryRequest{
		Name:          util.PtrString(""),
		FastQueryName: util.PtrString(""),
		Query:         util.PtrString(""),
		Description:   util.PtrString(""),
		Project:       util.PtrString(""),
		LogStoreName:  util.PtrString(""),
		LogStreamName: util.PtrString(""),
	}
	result := &UpdateFastQueryResponse{}
	result, err := BLS_CLIENT.UpdateFastQuery(updateFastQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateIndex(t *testing.T) {
	Fields := make(map[string]*IndexField)
	updateIndexRequest := &UpdateIndexRequest{
		LogStoreName:   util.PtrString(""),
		Project:        util.PtrString(""),
		Fulltext:       util.PtrBool(false),
		CaseSensitive:  util.PtrBool(false),
		IncludeChinese: util.PtrBool(false),
		Separators:     util.PtrString(""),
		Fields:         &Fields,
	}
	result := &UpdateIndexResponse{}
	result, err := BLS_CLIENT.UpdateIndex(updateIndexRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLogShipper(t *testing.T) {
	DestConfig := &DestConfig{
		BOSPath:                  util.PtrString(""),
		PartitionFormatTS:        util.PtrString(""),
		PartitionFormatLogStream: util.PtrBool(false),
		MaxObjectSize:            util.PtrInt32(int32(0)),
		CompressType:             util.PtrString(""),
		DeliverInterval:          util.PtrInt32(int32(0)),
		StorageFormat:            util.PtrString(""),
		CsvHeadline:              util.PtrBool(false),
		CsvDelimiter:             util.PtrString(""),
		CsvQuote:                 util.PtrString(""),
		NullIdentifier:           util.PtrString(""),
		SelectedColumnName:       util.PtrString(""),
		SelectedColumnType:       util.PtrString(""),
		FieldsName:               []*string{},
		FieldsType:               []*string{},
		ShipperType:              util.PtrString(""),
		KafkaConfig:              util.PtrString(""),
		DestType:                 util.PtrString(""),
		LogStore:                 util.PtrString(""),
		RateLimit:                util.PtrInt32(int32(0)),
		ClientCount:              util.PtrInt32(int32(0)),
	}
	updateLogShipperRequest := &UpdateLogShipperRequest{
		LogShipperID:   util.PtrString(""),
		Project:        util.PtrString(""),
		LogShipperName: util.PtrString(""),
		DestConfig:     DestConfig,
	}
	result := &UpdateLogShipperResponse{}
	result, err := BLS_CLIENT.UpdateLogShipper(updateLogShipperRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLogStore(t *testing.T) {
	updateLogStoreRequest := &UpdateLogStoreRequest{
		LogStoreName:          util.PtrString(""),
		Project:               util.PtrString(""),
		Retention:             util.PtrInt32(int32(0)),
		Tags:                  []*Tag{},
		ShardCount:            util.PtrInt32(int32(0)),
		MaxShardCount:         util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		IndexEnabled:          util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
	}
	result := &UpdateLogStoreResponse{}
	result, err := BLS_CLIENT.UpdateLogStore(updateLogStoreRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLogStoreTemplate(t *testing.T) {
	Template := &Template{
		Retention:             util.PtrInt32(int32(0)),
		ShardCount:            util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		MaxShardCount:         util.PtrInt32(int32(0)),
		EnableHotRetention:    util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
		Index: &Index{
			Fulltext:       util.PtrBool(false),
			CaseSensitive:  util.PtrBool(false),
			IncludeChinese: util.PtrBool(false),
			Separators:     util.PtrString(""),
			Fields:         make(map[string]*Field),
		},
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		CreatedTimestamp: util.PtrString(""),
		UpdatedTimestamp: util.PtrString(""),
	}
	updateLogStoreTemplateRequest := &UpdateLogStoreTemplateRequest{
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		Template:         Template,
	}
	result := &UpdateLogStoreTemplateResponse{}
	result, err := BLS_CLIENT.UpdateLogStoreTemplate(updateLogStoreTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateLogStoreView(t *testing.T) {
	updateLogStoreViewRequest := &UpdateLogStoreViewRequest{
		Project:   util.PtrString(""),
		Name:      util.PtrString(""),
		Logstores: []*LogStore{},
	}
	err := BLS_CLIENT.UpdateLogStoreView(updateLogStoreViewRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateProject(t *testing.T) {
	updateProjectRequest := &UpdateProjectRequest{
		Uuid:        util.PtrString(""),
		Description: util.PtrString(""),
		Top:         util.PtrBool(false),
	}
	result := &UpdateProjectResponse{}
	result, err := BLS_CLIENT.UpdateProject(updateProjectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateTask(t *testing.T) {
	Config := &TaskConfig{
		SrcConfig: &SrcConfig{
			SrcType:        util.PtrString(""),
			LogType:        util.PtrString(""),
			SrcDir:         util.PtrString(""),
			MatchedPattern: util.PtrString(""),
			IgnorePattern:  util.PtrString(""),
			TimeFormat:     util.PtrString(""),
			Ttl:            util.PtrInt32(int32(0)),
			UseMultiline:   util.PtrBool(false),
			MultilineRegex: util.PtrString(""),
			RecursiveDir:   util.PtrBool(false),
			ProcessType:    util.PtrString(""),
			ProcessConfig: &ProcessConfig{
				Regex:            util.PtrString(""),
				Separator:        util.PtrString(""),
				CustomSeparator:  util.PtrString(""),
				Quote:            util.PtrString(""),
				KvKeyIndex:       util.PtrInt32(int32(0)),
				KvValueIndex:     util.PtrInt32(int32(0)),
				SampleLog:        util.PtrString(""),
				Keys:             util.PtrString(""),
				DataType:         util.PtrString(""),
				DiscardOnFailure: util.PtrBool(false),
				KeepOriginal:     util.PtrBool(false),
			},
			LogTime:        util.PtrString(""),
			TimestampKey:   util.PtrString(""),
			DateFormat:     util.PtrString(""),
			FilterExpr:     util.PtrString(""),
			AdditionConfig: nil,
			MetaEnv:        []*string{},
			MetaLabel:      []*string{},
			MetaContainer:  []*string{},
			MetaToFields:   util.PtrBool(false),
			HarvesterLimit: util.PtrInt32(int32(0)),
		},
		DestConfig: &DestConfig{
			BOSPath:                  util.PtrString(""),
			PartitionFormatTS:        util.PtrString(""),
			PartitionFormatLogStream: util.PtrBool(false),
			MaxObjectSize:            util.PtrInt32(int32(0)),
			CompressType:             util.PtrString(""),
			DeliverInterval:          util.PtrInt32(int32(0)),
			StorageFormat:            util.PtrString(""),
			CsvHeadline:              util.PtrBool(false),
			CsvDelimiter:             util.PtrString(""),
			CsvQuote:                 util.PtrString(""),
			NullIdentifier:           util.PtrString(""),
			SelectedColumnName:       util.PtrString(""),
			SelectedColumnType:       util.PtrString(""),
			FieldsName:               []*string{},
			FieldsType:               []*string{},
			ShipperType:              util.PtrString(""),
			KafkaConfig:              util.PtrString(""),
			DestType:                 util.PtrString(""),
			LogStore:                 util.PtrString(""),
			RateLimit:                util.PtrInt32(int32(0)),
			ClientCount:              util.PtrInt32(int32(0)),
		},
	}
	updateTaskRequest := &UpdateTaskRequest{
		TaskId: util.PtrString(""),
		Name:   util.PtrString(""),
		Config: Config,
		Hosts:  []*Host{},
		Tags:   []*Tag{},
	}
	err := BLS_CLIENT.UpdateTask(updateTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ValidateAlarmCondition(t *testing.T) {
	validateAlarmConditionRequest := &ValidateAlarmConditionRequest{
		FieldTypes: []*string{},
		Conditions: []*string{},
	}
	result := &ValidateAlarmConditionResponse{}
	result, err := BLS_CLIENT.ValidateAlarmCondition(validateAlarmConditionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ValidateAlarmPolicySQL(t *testing.T) {
	validateAlarmPolicySQLRequest := &ValidateAlarmPolicySQLRequest{
		LogStores: []*LogStore{},
		Query:     util.PtrString(""),
	}
	result := &ValidateAlarmPolicySQLResponse{}
	result, err := BLS_CLIENT.ValidateAlarmPolicySQL(validateAlarmPolicySQLRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
