package dbsc

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
	DBSC_CLIENT *Client
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
	DBSC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_CheckMysqlRateLimitSupport(t *testing.T) {
	checkMysqlRateLimitSupportRequest := &CheckMysqlRateLimitSupportRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &CheckMysqlRateLimitSupportResponse{}
	result, err := DBSC_CLIENT.CheckMysqlRateLimitSupport(checkMysqlRateLimitSupportRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateMysqlRateLimitTask(t *testing.T) {
	createMysqlRateLimitTaskRequest := &CreateMysqlRateLimitTaskRequest{
		AppId:       util.PtrString(""),
		NodeId:      util.PtrString(""),
		FilterKey:   util.PtrString(""),
		FilterLimit: util.PtrInt32(int32(0)),
		FilterType:  util.PtrString(""),
	}
	err := DBSC_CLIENT.CreateMysqlRateLimitTask(createMysqlRateLimitTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateRedisBigKeyAnalysisTask(t *testing.T) {
	createRedisBigKeyAnalysisTaskRequest := &CreateRedisBigKeyAnalysisTaskRequest{
		AppId:      util.PtrString(""),
		ClusterId:  util.PtrString(""),
		BackupType: util.PtrInt32(int32(0)),
		BackupId:   util.PtrString(""),
	}
	result := &CreateRedisBigKeyAnalysisTaskResponse{}
	result, err := DBSC_CLIENT.CreateRedisBigKeyAnalysisTask(createRedisBigKeyAnalysisTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteMysqlRateLimitTask(t *testing.T) {
	deleteMysqlRateLimitTaskRequest := &DeleteMysqlRateLimitTaskRequest{
		FilterId: util.PtrInt32(int32(0)),
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
	}
	err := DBSC_CLIENT.DeleteMysqlRateLimitTask(deleteMysqlRateLimitTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRedisBigKeyAnalysisTask(t *testing.T) {
	deleteRedisBigKeyAnalysisTaskRequest := &DeleteRedisBigKeyAnalysisTaskRequest{
		Ids:   []*int64{},
		AppId: util.PtrString(""),
	}
	err := DBSC_CLIENT.DeleteRedisBigKeyAnalysisTask(deleteRedisBigKeyAnalysisTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbCollectionIndexes(t *testing.T) {
	getMongodbCollectionIndexesRequest := &GetMongodbCollectionIndexesRequest{
		Product:    util.PtrString(""),
		AppId:      util.PtrString(""),
		NodeId:     util.PtrString(""),
		Database:   util.PtrString(""),
		Collection: util.PtrString(""),
	}
	err := DBSC_CLIENT.GetMongodbCollectionIndexes(getMongodbCollectionIndexesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbCollectionSpace(t *testing.T) {
	getMongodbCollectionSpaceRequest := &GetMongodbCollectionSpaceRequest{
		AppId:      util.PtrString(""),
		NodeId:     util.PtrString(""),
		Database:   util.PtrString(""),
		Collection: util.PtrString(""),
		OrderBy:    util.PtrString(""),
		Order:      util.PtrString(""),
		Page:       util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result := &GetMongodbCollectionSpaceResponse{}
	result, err := DBSC_CLIENT.GetMongodbCollectionSpace(getMongodbCollectionSpaceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbCollectionSpaceTrend(t *testing.T) {
	getMongodbCollectionSpaceTrendRequest := &GetMongodbCollectionSpaceTrendRequest{
		AppId:      util.PtrString(""),
		Database:   util.PtrString(""),
		Collection: util.PtrString(""),
		Period:     util.PtrInt32(int32(0)),
		NodeId:     util.PtrString(""),
		Start:      util.PtrString(""),
		End:        util.PtrString(""),
		Metrics:    util.PtrString(""),
		Statistics: util.PtrString(""),
	}
	result := &GetMongodbCollectionSpaceTrendResponse{}
	result, err := DBSC_CLIENT.GetMongodbCollectionSpaceTrend(getMongodbCollectionSpaceTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbDatabaseSpace(t *testing.T) {
	getMongodbDatabaseSpaceRequest := &GetMongodbDatabaseSpaceRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Database: util.PtrString(""),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &GetMongodbDatabaseSpaceResponse{}
	result, err := DBSC_CLIENT.GetMongodbDatabaseSpace(getMongodbDatabaseSpaceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbDatabaseSpaceTrend(t *testing.T) {
	getMongodbDatabaseSpaceTrendRequest := &GetMongodbDatabaseSpaceTrendRequest{
		AppId:      util.PtrString(""),
		Database:   util.PtrString(""),
		Period:     util.PtrInt32(int32(0)),
		NodeId:     util.PtrString(""),
		Start:      util.PtrString(""),
		End:        util.PtrString(""),
		Metrics:    util.PtrString(""),
		Statistics: util.PtrString(""),
	}
	result := &GetMongodbDatabaseSpaceTrendResponse{}
	result, err := DBSC_CLIENT.GetMongodbDatabaseSpaceTrend(getMongodbDatabaseSpaceTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbSlowLogTimeDistribution(t *testing.T) {
	getMongodbSlowLogTimeDistributionRequest := &GetMongodbSlowLogTimeDistributionRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		DbNames:        util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
	}
	result := &GetMongodbSlowLogTimeDistributionResponse{}
	result, err := DBSC_CLIENT.GetMongodbSlowLogTimeDistribution(getMongodbSlowLogTimeDistributionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbSlowLogTrend(t *testing.T) {
	getMongodbSlowLogTrendRequest := &GetMongodbSlowLogTrendRequest{
		AppId:  util.PtrString(""),
		Start:  util.PtrString(""),
		End:    util.PtrString(""),
		Period: util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMongodbSlowLogTrendResponse{}
	result, err := DBSC_CLIENT.GetMongodbSlowLogTrend(getMongodbSlowLogTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbSlowQueryTemplate(t *testing.T) {
	getMongodbSlowQueryTemplateRequest := &GetMongodbSlowQueryTemplateRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		Namespace:      util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &GetMongodbSlowQueryTemplateResponse{}
	result, err := DBSC_CLIENT.GetMongodbSlowQueryTemplate(getMongodbSlowQueryTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMongodbSpaceSummary(t *testing.T) {
	getMongodbSpaceSummaryRequest := &GetMongodbSpaceSummaryRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMongodbSpaceSummaryResponse{}
	result, err := DBSC_CLIENT.GetMongodbSpaceSummary(getMongodbSpaceSummaryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlActiveSessions(t *testing.T) {
	getMysqlActiveSessionsRequest := &GetMysqlActiveSessionsRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMysqlActiveSessionsResponse{}
	result, err := DBSC_CLIENT.GetMysqlActiveSessions(getMysqlActiveSessionsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlDatabaseSpace(t *testing.T) {
	getMysqlDatabaseSpaceRequest := &GetMysqlDatabaseSpaceRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Database: util.PtrString(""),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &GetMysqlDatabaseSpaceResponse{}
	result, err := DBSC_CLIENT.GetMysqlDatabaseSpace(getMysqlDatabaseSpaceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlDeadlockInfo(t *testing.T) {
	getMysqlDeadlockInfoRequest := &GetMysqlDeadlockInfoRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMysqlDeadlockInfoResponse{}
	result, err := DBSC_CLIENT.GetMysqlDeadlockInfo(getMysqlDeadlockInfoRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlKillSessionHistory(t *testing.T) {
	getMysqlKillSessionHistoryRequest := &GetMysqlKillSessionHistoryRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
		Start:  util.PtrString(""),
		End:    util.PtrString(""),
	}
	result := &GetMysqlKillSessionHistoryResponse{}
	result, err := DBSC_CLIENT.GetMysqlKillSessionHistory(getMysqlKillSessionHistoryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlRateLimitTaskDetail(t *testing.T) {
	getMysqlRateLimitTaskDetailRequest := &GetMysqlRateLimitTaskDetailRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		FilterId: util.PtrInt32(int32(0)),
	}
	result := &GetMysqlRateLimitTaskDetailResponse{}
	result, err := DBSC_CLIENT.GetMysqlRateLimitTaskDetail(getMysqlRateLimitTaskDetailRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlSlowLogTemplate(t *testing.T) {
	getMysqlSlowLogTemplateRequest := &GetMysqlSlowLogTemplateRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &GetMysqlSlowLogTemplateResponse{}
	result, err := DBSC_CLIENT.GetMysqlSlowLogTemplate(getMysqlSlowLogTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlSlowLogTimeDistribution(t *testing.T) {
	getMysqlSlowLogTimeDistributionRequest := &GetMysqlSlowLogTimeDistributionRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		DbNames:        util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
	}
	result := &GetMysqlSlowLogTimeDistributionResponse{}
	result, err := DBSC_CLIENT.GetMysqlSlowLogTimeDistribution(getMysqlSlowLogTimeDistributionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlSlowLogTrend(t *testing.T) {
	getMysqlSlowLogTrendRequest := &GetMysqlSlowLogTrendRequest{
		AppId:  util.PtrString(""),
		Start:  util.PtrString(""),
		End:    util.PtrString(""),
		Period: util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMysqlSlowLogTrendResponse{}
	result, err := DBSC_CLIENT.GetMysqlSlowLogTrend(getMysqlSlowLogTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlSpaceSummary(t *testing.T) {
	getMysqlSpaceSummaryRequest := &GetMysqlSpaceSummaryRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &GetMysqlSpaceSummaryResponse{}
	result, err := DBSC_CLIENT.GetMysqlSpaceSummary(getMysqlSpaceSummaryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlTableIndexes(t *testing.T) {
	getMysqlTableIndexesRequest := &GetMysqlTableIndexesRequest{
		Product:  util.PtrString(""),
		AppId:    util.PtrString(""),
		Database: util.PtrString(""),
		Table:    util.PtrString(""),
	}
	result := &GetMysqlTableIndexesResponse{}
	result, err := DBSC_CLIENT.GetMysqlTableIndexes(getMysqlTableIndexesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetMysqlTableSpace(t *testing.T) {
	getMysqlTableSpaceRequest := &GetMysqlTableSpaceRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Database: util.PtrString(""),
		Table:    util.PtrString(""),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &GetMysqlTableSpaceResponse{}
	result, err := DBSC_CLIENT.GetMysqlTableSpace(getMysqlTableSpaceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPegadbSlowLogTemplate(t *testing.T) {
	getPegadbSlowLogTemplateRequest := &GetPegadbSlowLogTemplateRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
	}
	result := &GetPegadbSlowLogTemplateResponse{}
	result, err := DBSC_CLIENT.GetPegadbSlowLogTemplate(getPegadbSlowLogTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPegadbSlowLogTimeDistribution(t *testing.T) {
	getPegadbSlowLogTimeDistributionRequest := &GetPegadbSlowLogTimeDistributionRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
	}
	result := &GetPegadbSlowLogTimeDistributionResponse{}
	result, err := DBSC_CLIENT.GetPegadbSlowLogTimeDistribution(getPegadbSlowLogTimeDistributionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPegadbSlowLogTrend(t *testing.T) {
	getPegadbSlowLogTrendRequest := &GetPegadbSlowLogTrendRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Period:   util.PtrInt32(int32(0)),
	}
	result := &GetPegadbSlowLogTrendResponse{}
	result, err := DBSC_CLIENT.GetPegadbSlowLogTrend(getPegadbSlowLogTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPostgresqlSlowLogTemplate(t *testing.T) {
	getPostgresqlSlowLogTemplateRequest := &GetPostgresqlSlowLogTemplateRequest{
		AppId:          util.PtrString(""),
		NodeId:         util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
	}
	result := &GetPostgresqlSlowLogTemplateResponse{}
	result, err := DBSC_CLIENT.GetPostgresqlSlowLogTemplate(getPostgresqlSlowLogTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPostgresqlSlowLogTimeDistribution(t *testing.T) {
	getPostgresqlSlowLogTimeDistributionRequest := &GetPostgresqlSlowLogTimeDistributionRequest{
		Product:   util.PtrString(""),
		AppId:     util.PtrString(""),
		NodeId:    util.PtrString(""),
		Start:     util.PtrString(""),
		End:       util.PtrString(""),
		DbNames:   []*string{},
		Users:     []*string{},
		ClientIPs: []*string{},
	}
	result := &GetPostgresqlSlowLogTimeDistributionResponse{}
	result, err := DBSC_CLIENT.GetPostgresqlSlowLogTimeDistribution(getPostgresqlSlowLogTimeDistributionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPostgresqlSlowLogTrend(t *testing.T) {
	getPostgresqlSlowLogTrendRequest := &GetPostgresqlSlowLogTrendRequest{
		Product: util.PtrString(""),
		AppId:   util.PtrString(""),
		NodeId:  util.PtrString(""),
		Start:   util.PtrString(""),
		End:     util.PtrString(""),
		Period:  util.PtrInt32(int32(0)),
	}
	result := &GetPostgresqlSlowLogTrendResponse{}
	result, err := DBSC_CLIENT.GetPostgresqlSlowLogTrend(getPostgresqlSlowLogTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetRedisBigKeyAnalysisResult(t *testing.T) {
	getRedisBigKeyAnalysisResultRequest := &GetRedisBigKeyAnalysisResultRequest{
		AppId:    util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
	}
	result := &GetRedisBigKeyAnalysisResultResponse{}
	result, err := DBSC_CLIENT.GetRedisBigKeyAnalysisResult(getRedisBigKeyAnalysisResultRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetRedisSlowLogTemplate(t *testing.T) {
	getRedisSlowLogTemplateRequest := &GetRedisSlowLogTemplateRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
	}
	result := &GetRedisSlowLogTemplateResponse{}
	result, err := DBSC_CLIENT.GetRedisSlowLogTemplate(getRedisSlowLogTemplateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetRedisSlowLogTimeDistribution(t *testing.T) {
	getRedisSlowLogTimeDistributionRequest := &GetRedisSlowLogTimeDistributionRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
	}
	result := &GetRedisSlowLogTimeDistributionResponse{}
	result, err := DBSC_CLIENT.GetRedisSlowLogTimeDistribution(getRedisSlowLogTimeDistributionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetRedisSlowLogTrend(t *testing.T) {
	getRedisSlowLogTrendRequest := &GetRedisSlowLogTrendRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Period:   util.PtrInt32(int32(0)),
	}
	result := &GetRedisSlowLogTrendResponse{}
	result, err := DBSC_CLIENT.GetRedisSlowLogTrend(getRedisSlowLogTrendRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_KillMysqlSession(t *testing.T) {
	killMysqlSessionRequest := &KillMysqlSessionRequest{
		AppId:   util.PtrString(""),
		NodeId:  util.PtrString(""),
		IdItems: []*int32{},
	}
	result := &KillMysqlSessionResponse{}
	result, err := DBSC_CLIENT.KillMysqlSession(killMysqlSessionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListMongodbSlowLogs(t *testing.T) {
	listMongodbSlowLogsRequest := &ListMongodbSlowLogsRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		Namespace:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &ListMongodbSlowLogsResponse{}
	result, err := DBSC_CLIENT.ListMongodbSlowLogs(listMongodbSlowLogsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListMysqlRateLimitTasks(t *testing.T) {
	listMysqlRateLimitTasksRequest := &ListMysqlRateLimitTasksRequest{
		AppId:  util.PtrString(""),
		NodeId: util.PtrString(""),
	}
	result := &ListMysqlRateLimitTasksResponse{}
	result, err := DBSC_CLIENT.ListMysqlRateLimitTasks(listMysqlRateLimitTasksRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListMysqlSlowLogs(t *testing.T) {
	listMysqlSlowLogsRequest := &ListMysqlSlowLogsRequest{
		AppId:          util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		NodeId:         util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
	}
	result := &ListMysqlSlowLogsResponse{}
	result, err := DBSC_CLIENT.ListMysqlSlowLogs(listMysqlSlowLogsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListPegadbSlowLogs(t *testing.T) {
	listPegadbSlowLogsRequest := &ListPegadbSlowLogsRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
	}
	result := &ListPegadbSlowLogsResponse{}
	result, err := DBSC_CLIENT.ListPegadbSlowLogs(listPegadbSlowLogsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListPostgresqlSlowLogs(t *testing.T) {
	listPostgresqlSlowLogsRequest := &ListPostgresqlSlowLogsRequest{
		AppId:     util.PtrString(""),
		NodeId:    util.PtrString(""),
		Start:     util.PtrString(""),
		End:       util.PtrString(""),
		Page:      util.PtrInt32(int32(0)),
		PageSize:  util.PtrInt32(int32(0)),
		DbNames:   []*string{},
		ClientIPs: []*string{},
		Users:     []*string{},
		OrderBy:   util.PtrString(""),
		Order:     util.PtrString(""),
	}
	result := &ListPostgresqlSlowLogsResponse{}
	result, err := DBSC_CLIENT.ListPostgresqlSlowLogs(listPostgresqlSlowLogsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListRedisBigKeyAnalysisTasks(t *testing.T) {
	listRedisBigKeyAnalysisTasksRequest := &ListRedisBigKeyAnalysisTasksRequest{
		Id:        util.PtrInt32(int32(0)),
		AppId:     util.PtrString(""),
		ClusterId: util.PtrString(""),
		DataType:  util.PtrString(""),
		Order:     util.PtrString(""),
		OrderBy:   util.PtrString(""),
	}
	result := &ListRedisBigKeyAnalysisTasksResponse{}
	result, err := DBSC_CLIENT.ListRedisBigKeyAnalysisTasks(listRedisBigKeyAnalysisTasksRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListRedisSlowLogs(t *testing.T) {
	listRedisSlowLogsRequest := &ListRedisSlowLogsRequest{
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
		Start:    util.PtrString(""),
		End:      util.PtrString(""),
		DbEngine: util.PtrString(""),
		Page:     util.PtrInt32(int32(0)),
		PageSize: util.PtrInt32(int32(0)),
		OrderBy:  util.PtrString(""),
		Order:    util.PtrString(""),
	}
	result := &ListRedisSlowLogsResponse{}
	result, err := DBSC_CLIENT.ListRedisSlowLogs(listRedisSlowLogsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StartStopMysqlInstanceFlowLimitingTask(t *testing.T) {
	startStopMysqlInstanceFlowLimitingTaskRequest := &StartStopMysqlInstanceFlowLimitingTaskRequest{
		FilterId: util.PtrInt32(int32(0)),
		Action:   util.PtrString(""),
		AppId:    util.PtrString(""),
		NodeId:   util.PtrString(""),
	}
	err := DBSC_CLIENT.StartStopMysqlInstanceFlowLimitingTask(startStopMysqlInstanceFlowLimitingTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateMysqlRateLimitTask(t *testing.T) {
	updateMysqlRateLimitTaskRequest := &UpdateMysqlRateLimitTaskRequest{
		FilterId:    util.PtrInt32(int32(0)),
		AppId:       util.PtrString(""),
		NodeId:      util.PtrString(""),
		FilterKey:   util.PtrString(""),
		FilterLimit: util.PtrInt32(int32(0)),
		FilterType:  util.PtrString(""),
	}
	err := DBSC_CLIENT.UpdateMysqlRateLimitTask(updateMysqlRateLimitTaskRequest)
	ExpectEqual(t.Errorf, nil, err)
}
