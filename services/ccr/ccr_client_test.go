package ccr

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "reflect"
    "runtime"
    "testing"

    "github.com/baidubce/baiducloud-go-sdk/core/util/log"
    "github.com/baidubce/baiducloud-go-sdk/core/util"
)

var (
    CCR_CLIENT  *Client
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

    CCR_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_AddPublicNetworkWhitelist(t *testing.T) {
    addPublicNetworkWhitelistRequest := &AddPublicNetworkWhitelistRequest{
        InstanceId : util.PtrString(""),
        IpAddr : util.PtrString(""),
        Description : util.PtrString(""),
    }
    err := CCR_CLIENT.AddPublicNetworkWhitelist(addPublicNetworkWhitelistRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AddVpcLink(t *testing.T) {
    addVpcLinkRequest := &AddVpcLinkRequest{
        InstanceId : util.PtrString(""),
        VpcID : util.PtrString(""),
        SubnetID : util.PtrString(""),
        IpType : util.PtrString(""),
        IpAddress : util.PtrString(""),
        AutoDnsResolve : util.PtrBool(false),
    }
    err := CCR_CLIENT.AddVpcLink(addVpcLinkRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAcceleratorFilter(t *testing.T) {
    createAcceleratorFilterRequest := &CreateAcceleratorFilterRequest{
        InstanceId : util.PtrString(""),
        Description : util.PtrString(""),
        Filters : []*AcceleratorFilter{},
        Name : util.PtrString(""),
    }
    err := CCR_CLIENT.CreateAcceleratorFilter(createAcceleratorFilterRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateImageMigrationRule(t *testing.T) {
    SrcRegistry := &ReplicationRegistryRequest{
        Id : util.PtrInt32(int32(0)),

    }
    Trigger := &ReplicationTriggerRequest{
        CcrType : util.PtrString(""),

    }
    createImageMigrationRuleRequest := &CreateImageMigrationRuleRequest{
        InstanceId : util.PtrString(""),
        Description : util.PtrString(""),
        DestProjectName : util.PtrString(""),
        Filters : []*ReplicationFilterRequest{},
        Name : util.PtrString(""),
        Override : util.PtrBool(false),
        SrcRegistry : SrcRegistry,
        Trigger : Trigger,
    }
    err := CCR_CLIENT.CreateImageMigrationRule(createImageMigrationRuleRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateInstanceSync(t *testing.T) {
    Trigger := &ReplicationSyncTriggerRequest{
        CcrType : util.PtrString(""),

    }
    createInstanceSyncRequest := &CreateInstanceSyncRequest{
        InstanceId : util.PtrString(""),
        Description : util.PtrString(""),
        DestInstanceId : util.PtrString(""),
        DestProjectName : util.PtrString(""),
        Name : util.PtrString(""),
        Override : util.PtrBool(false),
        SrcProjectName : util.PtrString(""),
        SrcRepository : util.PtrString(""),
        SrcTag : util.PtrString(""),
        SyncType : util.PtrString(""),
        Trigger : Trigger,
    }
    err := CCR_CLIENT.CreateInstanceSync(createInstanceSyncRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateProject(t *testing.T) {
    createProjectRequest := &CreateProjectRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        Public : util.PtrString(""),
    }
    result := &CreateProjectResponse{}
    result, err := CCR_CLIENT.CreateProject(createProjectRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateRobotAccount(t *testing.T) {
    createRobotAccountRequest := &CreateRobotAccountRequest{
        InstanceId : util.PtrString(""),
        Name : util.PtrString(""),
        Secret : util.PtrString(""),
        Disable : util.PtrBool(false),
        Duration : util.PtrInt32(int32(0)),
        Description : util.PtrString(""),
        Permissions : []*RobotPermission{},
    }
    result := &CreateRobotAccountResponse{}
    result, err := CCR_CLIENT.CreateRobotAccount(createRobotAccountRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateTemporaryPassword(t *testing.T) {
    createTemporaryPasswordRequest := &CreateTemporaryPasswordRequest{
        InstanceId : util.PtrString(""),
        Duration : util.PtrInt32(int32(0)),
    }
    result := &CreateTemporaryPasswordResponse{}
    result, err := CCR_CLIENT.CreateTemporaryPassword(createTemporaryPasswordRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateTrigger(t *testing.T) {
    createTriggerRequest := &CreateTriggerRequest{
        InstanceId : util.PtrString(""),
        Description : util.PtrString(""),
        EventTypes : []*string{},
        Filters : []*TriggerFilter{},
        Name : util.PtrString(""),
        Targets : []*TriggerTarget{},
    }
    err := CCR_CLIENT.CreateTrigger(createTriggerRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAcceleratorFilter(t *testing.T) {
    deleteAcceleratorFilterRequest := &DeleteAcceleratorFilterRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteAcceleratorFilter(deleteAcceleratorFilterRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAcceleratorFilters(t *testing.T) {
    deleteAcceleratorFiltersRequest := &DeleteAcceleratorFiltersRequest{
        InstanceId : util.PtrString(""),
        Items : []*int32{},
    }
    err := CCR_CLIENT.DeleteAcceleratorFilters(deleteAcceleratorFiltersRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteChart(t *testing.T) {
    deleteChartRequest := &DeleteChartRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        ChartName : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteChart(deleteChartRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteChartVersion(t *testing.T) {
    deleteChartVersionRequest := &DeleteChartVersionRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        ChartName : util.PtrString(""),
        ChartVersion : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteChartVersion(deleteChartVersionRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteChartVersions(t *testing.T) {
    deleteChartVersionsRequest := &DeleteChartVersionsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        ChartName : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeleteChartVersions(deleteChartVersionsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteCharts(t *testing.T) {
    deleteChartsRequest := &DeleteChartsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeleteCharts(deleteChartsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteImageMigrationRule(t *testing.T) {
    deleteImageMigrationRuleRequest := &DeleteImageMigrationRuleRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteImageMigrationRule(deleteImageMigrationRuleRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteInstanceSync(t *testing.T) {
    deleteInstanceSyncRequest := &DeleteInstanceSyncRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteInstanceSync(deleteInstanceSyncRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteProject(t *testing.T) {
    deleteProjectRequest := &DeleteProjectRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteProject(deleteProjectRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteProjects(t *testing.T) {
    deleteProjectsRequest := &DeleteProjectsRequest{
        InstanceId : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeleteProjects(deleteProjectsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeletePublicNetworkWhitelist(t *testing.T) {
    deletePublicNetworkWhitelistRequest := &DeletePublicNetworkWhitelistRequest{
        InstanceId : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeletePublicNetworkWhitelist(deletePublicNetworkWhitelistRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRepositories(t *testing.T) {
    deleteRepositoriesRequest := &DeleteRepositoriesRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeleteRepositories(deleteRepositoriesRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRepository(t *testing.T) {
    deleteRepositoryRequest := &DeleteRepositoryRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteRepository(deleteRepositoryRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRobotAccount(t *testing.T) {
    deleteRobotAccountRequest := &DeleteRobotAccountRequest{
        InstanceId : util.PtrString(""),
        RobotID : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteRobotAccount(deleteRobotAccountRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTag(t *testing.T) {
    deleteTagRequest := &DeleteTagRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteTag(deleteTagRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTags(t *testing.T) {
    deleteTagsRequest := &DeleteTagsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        Items : []*string{},
    }
    err := CCR_CLIENT.DeleteTags(deleteTagsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTrigger(t *testing.T) {
    deleteTriggerRequest := &DeleteTriggerRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteTrigger(deleteTriggerRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteTriggers(t *testing.T) {
    deleteTriggersRequest := &DeleteTriggersRequest{
        InstanceId : util.PtrString(""),
        Items : []*int32{},
    }
    err := CCR_CLIENT.DeleteTriggers(deleteTriggersRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteVpcLink(t *testing.T) {
    deleteVpcLinkRequest := &DeleteVpcLinkRequest{
        InstanceId : util.PtrString(""),
        VpcID : util.PtrString(""),
        SubnetID : util.PtrString(""),
    }
    err := CCR_CLIENT.DeleteVpcLink(deleteVpcLinkRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DownloadChart(t *testing.T) {
    downloadChartRequest := &DownloadChartRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        Filename : util.PtrString(""),
    }
    err := CCR_CLIENT.DownloadChart(downloadChartRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ExecuteImageMigration(t *testing.T) {
    executeImageMigrationRequest := &ExecuteImageMigrationRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrInt32(int32(0)),
    }
    err := CCR_CLIENT.ExecuteImageMigration(executeImageMigrationRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ExecuteInstanceSync(t *testing.T) {
    executeInstanceSyncRequest := &ExecuteInstanceSyncRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrInt32(int32(0)),
    }
    err := CCR_CLIENT.ExecuteInstanceSync(executeInstanceSyncRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetAcceleratorFilterDetail(t *testing.T) {
    getAcceleratorFilterDetailRequest := &GetAcceleratorFilterDetailRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    result := &GetAcceleratorFilterDetailResponse{}
    result, err := CCR_CLIENT.GetAcceleratorFilterDetail(getAcceleratorFilterDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetImageMigrationExecutionRecordDetail(t *testing.T) {
    getImageMigrationExecutionRecordDetailRequest := &GetImageMigrationExecutionRecordDetailRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
    }
    result := &GetImageMigrationExecutionRecordDetailResponse{}
    result, err := CCR_CLIENT.GetImageMigrationExecutionRecordDetail(getImageMigrationExecutionRecordDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetImageMigrationRuleDetail(t *testing.T) {
    getImageMigrationRuleDetailRequest := &GetImageMigrationRuleDetailRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    result := &GetImageMigrationRuleDetailResponse{}
    result, err := CCR_CLIENT.GetImageMigrationRuleDetail(getImageMigrationRuleDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetImageMigrationTaskLogs(t *testing.T) {
    getImageMigrationTaskLogsRequest := &GetImageMigrationTaskLogsRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
        TaskId : util.PtrString(""),
    }
    err := CCR_CLIENT.GetImageMigrationTaskLogs(getImageMigrationTaskLogsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetInstanceDetail(t *testing.T) {
    getInstanceDetailRequest := &GetInstanceDetailRequest{
        InstanceId : util.PtrString(""),
    }
    result := &GetInstanceDetailResponse{}
    result, err := CCR_CLIENT.GetInstanceDetail(getInstanceDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetInstanceSyncDetail(t *testing.T) {
    getInstanceSyncDetailRequest := &GetInstanceSyncDetailRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    result := &GetInstanceSyncDetailResponse{}
    result, err := CCR_CLIENT.GetInstanceSyncDetail(getInstanceSyncDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetInstanceSyncExecutionDetail(t *testing.T) {
    getInstanceSyncExecutionDetailRequest := &GetInstanceSyncExecutionDetailRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
    }
    result := &GetInstanceSyncExecutionDetailResponse{}
    result, err := CCR_CLIENT.GetInstanceSyncExecutionDetail(getInstanceSyncExecutionDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetInstanceSyncTaskLogs(t *testing.T) {
    getInstanceSyncTaskLogsRequest := &GetInstanceSyncTaskLogsRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
        TaskId : util.PtrString(""),
    }
    err := CCR_CLIENT.GetInstanceSyncTaskLogs(getInstanceSyncTaskLogsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetProjectDetail(t *testing.T) {
    getProjectDetailRequest := &GetProjectDetailRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
    }
    result := &GetProjectDetailResponse{}
    result, err := CCR_CLIENT.GetProjectDetail(getProjectDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetPublicNetworkConfig(t *testing.T) {
    getPublicNetworkConfigRequest := &GetPublicNetworkConfigRequest{
        InstanceId : util.PtrString(""),
    }
    result := &GetPublicNetworkConfigResponse{}
    result, err := CCR_CLIENT.GetPublicNetworkConfig(getPublicNetworkConfigRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetRepository(t *testing.T) {
    getRepositoryRequest := &GetRepositoryRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
    }
    result := &GetRepositoryResponse{}
    result, err := CCR_CLIENT.GetRepository(getRepositoryRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTagBuildHistory(t *testing.T) {
    getTagBuildHistoryRequest := &GetTagBuildHistoryRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
    }
    result := &GetTagBuildHistoryResponse{}
    result, err := CCR_CLIENT.GetTagBuildHistory(getTagBuildHistoryRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTagDetail(t *testing.T) {
    getTagDetailRequest := &GetTagDetailRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
    }
    result := &GetTagDetailResponse{}
    result, err := CCR_CLIENT.GetTagDetail(getTagDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTagsScanOverview(t *testing.T) {
    getTagsScanOverviewRequest := &GetTagsScanOverviewRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &GetTagsScanOverviewResponse{}
    result, err := CCR_CLIENT.GetTagsScanOverview(getTagsScanOverviewRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetTriggerDetail(t *testing.T) {
    getTriggerDetailRequest := &GetTriggerDetailRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
    }
    result := &GetTriggerDetailResponse{}
    result, err := CCR_CLIENT.GetTriggerDetail(getTriggerDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetUserDetail(t *testing.T) {
    getUserDetailRequest := &GetUserDetailRequest{
        UserId : util.PtrString(""),
    }
    result := &GetUserDetailResponse{}
    result, err := CCR_CLIENT.GetUserDetail(getUserDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAcceleratorFilters(t *testing.T) {
    listAcceleratorFiltersRequest := &ListAcceleratorFiltersRequest{
        InstanceId : util.PtrString(""),
        PolicyName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListAcceleratorFiltersResponse{}
    result, err := CCR_CLIENT.ListAcceleratorFilters(listAcceleratorFiltersRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListChartVersions(t *testing.T) {
    listChartVersionsRequest := &ListChartVersionsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        ChartName : util.PtrString(""),
        ChartVersion : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListChartVersionsResponse{}
    result, err := CCR_CLIENT.ListChartVersions(listChartVersionsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListCharts(t *testing.T) {
    listChartsRequest := &ListChartsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        ChartName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListChartsResponse{}
    result, err := CCR_CLIENT.ListCharts(listChartsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListImageMigrationRecords(t *testing.T) {
    listImageMigrationRecordsRequest := &ListImageMigrationRecordsRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListImageMigrationRecordsResponse{}
    result, err := CCR_CLIENT.ListImageMigrationRecords(listImageMigrationRecordsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListImageMigrationRules(t *testing.T) {
    listImageMigrationRulesRequest := &ListImageMigrationRulesRequest{
        InstanceId : util.PtrString(""),
        PolicyName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListImageMigrationRulesResponse{}
    result, err := CCR_CLIENT.ListImageMigrationRules(listImageMigrationRulesRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListImageMigrationTaskRecords(t *testing.T) {
    listImageMigrationTaskRecordsRequest := &ListImageMigrationTaskRecordsRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListImageMigrationTaskRecordsResponse{}
    result, err := CCR_CLIENT.ListImageMigrationTaskRecords(listImageMigrationTaskRecordsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListInstanceSyncRecords(t *testing.T) {
    listInstanceSyncRecordsRequest := &ListInstanceSyncRecordsRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListInstanceSyncRecordsResponse{}
    result, err := CCR_CLIENT.ListInstanceSyncRecords(listInstanceSyncRecordsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListInstanceSyncTaskRecords(t *testing.T) {
    listInstanceSyncTaskRecordsRequest := &ListInstanceSyncTaskRecordsRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListInstanceSyncTaskRecordsResponse{}
    result, err := CCR_CLIENT.ListInstanceSyncTaskRecords(listInstanceSyncTaskRecordsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListInstanceSyncs(t *testing.T) {
    listInstanceSyncsRequest := &ListInstanceSyncsRequest{
        InstanceId : util.PtrString(""),
        PolicyName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListInstanceSyncsResponse{}
    result, err := CCR_CLIENT.ListInstanceSyncs(listInstanceSyncsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
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
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
        KeywordType : util.PtrString(""),
        Keyword : util.PtrString(""),
        Acrossregion : util.PtrString(""),
    }
    result := &ListInstancesResponse{}
    result, err := CCR_CLIENT.ListInstances(listInstancesRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListProjects(t *testing.T) {
    listProjectsRequest := &ListProjectsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListProjectsResponse{}
    result, err := CCR_CLIENT.ListProjects(listProjectsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListRepositories(t *testing.T) {
    listRepositoriesRequest := &ListRepositoriesRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListRepositoriesResponse{}
    result, err := CCR_CLIENT.ListRepositories(listRepositoriesRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListRobotAccounts(t *testing.T) {
    listRobotAccountsRequest := &ListRobotAccountsRequest{
        InstanceId : util.PtrString(""),
        Status : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListRobotAccountsResponse{}
    result, err := CCR_CLIENT.ListRobotAccounts(listRobotAccountsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTags(t *testing.T) {
    listTagsRequest := &ListTagsRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListTagsResponse{}
    result, err := CCR_CLIENT.ListTags(listTagsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTriggerTasks(t *testing.T) {
    listTriggerTasksRequest := &ListTriggerTasksRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListTriggerTasksResponse{}
    result, err := CCR_CLIENT.ListTriggerTasks(listTriggerTasksRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListTriggers(t *testing.T) {
    listTriggersRequest := &ListTriggersRequest{
        InstanceId : util.PtrString(""),
        PolicyName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
    }
    result := &ListTriggersResponse{}
    result, err := CCR_CLIENT.ListTriggers(listTriggersRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListVpcLinks(t *testing.T) {
    listVpcLinksRequest := &ListVpcLinksRequest{
        InstanceId : util.PtrString(""),
    }
    result := &ListVpcLinksResponse{}
    result, err := CCR_CLIENT.ListVpcLinks(listVpcLinksRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReExecuteTriggerTask(t *testing.T) {
    reExecuteTriggerTaskRequest := &ReExecuteTriggerTaskRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        JobId : util.PtrString(""),
    }
    err := CCR_CLIENT.ReExecuteTriggerTask(reExecuteTriggerTaskRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RefreshRobotAccountKey(t *testing.T) {
    refreshRobotAccountKeyRequest := &RefreshRobotAccountKeyRequest{
        InstanceId : util.PtrString(""),
        RobotID : util.PtrString(""),
        Secret : util.PtrString(""),
    }
    result := &RefreshRobotAccountKeyResponse{}
    result, err := CCR_CLIENT.RefreshRobotAccountKey(refreshRobotAccountKeyRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ResetPassword(t *testing.T) {
    resetPasswordRequest := &ResetPasswordRequest{
        InstanceId : util.PtrString(""),
        Password : util.PtrString(""),
    }
    err := CCR_CLIENT.ResetPassword(resetPasswordRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StopImageMigration(t *testing.T) {
    stopImageMigrationRequest := &StopImageMigrationRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
    }
    err := CCR_CLIENT.StopImageMigration(stopImageMigrationRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_StopInstanceSync(t *testing.T) {
    stopInstanceSyncRequest := &StopInstanceSyncRequest{
        InstanceId : util.PtrString(""),
        ExecutionId : util.PtrString(""),
    }
    err := CCR_CLIENT.StopInstanceSync(stopInstanceSyncRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TestAcceleratorFilter(t *testing.T) {
    testAcceleratorFilterRequest := &TestAcceleratorFilterRequest{
        Filters : []*AcceleratorFilter{},
        Repository : util.PtrString(""),
    }
    result := &TestAcceleratorFilterResponse{}
    result, err := CCR_CLIENT.TestAcceleratorFilter(testAcceleratorFilterRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TestTriggerTargetAddress(t *testing.T) {
    Headers := make(map[string]string)
    testTriggerTargetAddressRequest := &TestTriggerTargetAddressRequest{
        InstanceId : util.PtrString(""),
        Address : util.PtrString(""),
        Headers : ,
    }
    err := CCR_CLIENT.TestTriggerTargetAddress(testTriggerTargetAddressRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ToggleAcceleratorFilter(t *testing.T) {
    toggleAcceleratorFilterRequest := &ToggleAcceleratorFilterRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Enabled : util.PtrString(""),
    }
    err := CCR_CLIENT.ToggleAcceleratorFilter(toggleAcceleratorFilterRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ToggleTrigger(t *testing.T) {
    toggleTriggerRequest := &ToggleTriggerRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Enabled : util.PtrString(""),
    }
    err := CCR_CLIENT.ToggleTrigger(toggleTriggerRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TriggerTagScan(t *testing.T) {
    triggerTagScanRequest := &TriggerTagScanRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        TagName : util.PtrString(""),
    }
    err := CCR_CLIENT.TriggerTagScan(triggerTagScanRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAcceleratorFilter(t *testing.T) {
    updateAcceleratorFilterRequest := &UpdateAcceleratorFilterRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Description : util.PtrString(""),
        Filters : []*AcceleratorFilter{},
        Name : util.PtrString(""),
    }
    err := CCR_CLIENT.UpdateAcceleratorFilter(updateAcceleratorFilterRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateImageMigrationRule(t *testing.T) {
    SrcRegistry := &ReplicationRegistryRequest{
        Id : util.PtrInt32(int32(0)),

    }
    Trigger := &ReplicationTriggerRequest{
        CcrType : util.PtrString(""),

    }
    updateImageMigrationRuleRequest := &UpdateImageMigrationRuleRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Description : util.PtrString(""),
        DestProjectName : util.PtrString(""),
        Filters : []*ReplicationFilterRequest{},
        Name : util.PtrString(""),
        Override : util.PtrBool(false),
        SrcRegistry : SrcRegistry,
        Trigger : Trigger,
    }
    err := CCR_CLIENT.UpdateImageMigrationRule(updateImageMigrationRuleRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceName(t *testing.T) {
    updateInstanceNameRequest := &UpdateInstanceNameRequest{
        InstanceId : util.PtrString(""),
        Name : util.PtrString(""),
    }
    result := &UpdateInstanceNameResponse{}
    result, err := CCR_CLIENT.UpdateInstanceName(updateInstanceNameRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceSync(t *testing.T) {
    Trigger := &ReplicationSyncTriggerRequest{
        CcrType : util.PtrString(""),

    }
    updateInstanceSyncRequest := &UpdateInstanceSyncRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Description : util.PtrString(""),
        DestInstanceId : util.PtrString(""),
        DestProjectName : util.PtrString(""),
        Name : util.PtrString(""),
        Override : util.PtrBool(false),
        SrcProjectName : util.PtrString(""),
        SrcRepository : util.PtrString(""),
        SrcTag : util.PtrString(""),
        SyncType : util.PtrString(""),
        Trigger : Trigger,
    }
    err := CCR_CLIENT.UpdateInstanceSync(updateInstanceSyncRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateInstanceTags(t *testing.T) {
    updateInstanceTagsRequest := &UpdateInstanceTagsRequest{
        InstanceId : util.PtrString(""),
        Tags : []*LogicalTag{},
    }
    err := CCR_CLIENT.UpdateInstanceTags(updateInstanceTagsRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateProject(t *testing.T) {
    updateProjectRequest := &UpdateProjectRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        AutoScan : util.PtrString(""),
        Public : util.PtrString(""),
    }
    result := &UpdateProjectResponse{}
    result, err := CCR_CLIENT.UpdateProject(updateProjectRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdatePublicNetwork(t *testing.T) {
    updatePublicNetworkRequest := &UpdatePublicNetworkRequest{
        InstanceId : util.PtrString(""),
        Action : util.PtrString(""),
    }
    err := CCR_CLIENT.UpdatePublicNetwork(updatePublicNetworkRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRepository(t *testing.T) {
    updateRepositoryRequest := &UpdateRepositoryRequest{
        InstanceId : util.PtrString(""),
        ProjectName : util.PtrString(""),
        RepositoryName : util.PtrString(""),
        Description : util.PtrString(""),
    }
    result := &UpdateRepositoryResponse{}
    result, err := CCR_CLIENT.UpdateRepository(updateRepositoryRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRobotAccount(t *testing.T) {
    updateRobotAccountRequest := &UpdateRobotAccountRequest{
        InstanceId : util.PtrString(""),
        RobotID : util.PtrString(""),
        Disable : util.PtrBool(false),
        Duration : util.PtrInt32(int32(0)),
        Description : util.PtrString(""),
        Permissions : []*RobotPermission{},
    }
    err := CCR_CLIENT.UpdateRobotAccount(updateRobotAccountRequest)
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateTrigger(t *testing.T) {
    updateTriggerRequest := &UpdateTriggerRequest{
        InstanceId : util.PtrString(""),
        PolicyId : util.PtrString(""),
        Description : util.PtrString(""),
        EventTypes : []*string{},
        Filters : []*TriggerFilter{},
        Name : util.PtrString(""),
        Targets : []*TriggerTarget{},
    }
    err := CCR_CLIENT.UpdateTrigger(updateTriggerRequest)
    ExpectEqual(t.Errorf, nil, err)
}
