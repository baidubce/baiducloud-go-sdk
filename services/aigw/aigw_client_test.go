package aigw

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
    AIGW_CLIENT  *Client
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
    AIGW_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

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

func TestClient_CreateAIGateway(t *testing.T) {
    AihcArgs := &AihcArgs{
        AccountId : util.PtrString(""),
    SubnetId : util.PtrString(""),
    SecurityGroupIds : util.PtrString(""),
    VpcCidr : util.PtrString(""),
    DomainPrefix : util.PtrString(""),

    }
    createAIGatewayRequest := &CreateAIGatewayRequest{
        XRegion : util.PtrString(""),
        Name : util.PtrString(""),
        VpcId : util.PtrString(""),
        VpcCidr : util.PtrString(""),
        SubnetId : util.PtrString(""),
        GatewayType : util.PtrString(""),
        IsInternal : util.PtrString(""),
        NetworkTypes : []*string{},
        Replicas : util.PtrInt32(int32(0)),
        InstallMode : util.PtrString(""),
        Description : util.PtrString(""),
        DeleteProtection : util.PtrBool(false),
        SrcProduct : util.PtrString(""),
        AccountId : util.PtrString(""),
        WorkspaceId : util.PtrString(""),
        WorkspaceName : util.PtrString(""),
        BlbId : util.PtrString(""),
        BlbIp : util.PtrString(""),
        Clusters : []*ClusterInfo{},
        CpromInstanceId : util.PtrString(""),
        CpromBearerToken : util.PtrString(""),
        BlsEnabled : util.PtrBool(false),
        LogStoreName : util.PtrString(""),
        Version : util.PtrString(""),
        Tags : []*Tag{},
        ResourceGroupId : util.PtrString(""),
        AihcArgs : AihcArgs,
    }
    result := &CreateAIGatewayResponse{}
    result, err := AIGW_CLIENT.CreateAIGateway(createAIGatewayRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateConsumer(t *testing.T) {
    Credential := &ConsumerCredentialSpec{
        Name : util.PtrString(""),
    GenerateMode : util.PtrString(""),
    Value : util.PtrString(""),
    InHeader : util.PtrBool(false),
    InQuery : util.PtrBool(false),
    KeyNames : []*string{},
    Description : util.PtrString(""),

    }
    IamCredential := &IAMCredentialSpec{
        Name : util.PtrString(""),
    IamApiKeyId : util.PtrString(""),
    IamTokenIdMasked : util.PtrString(""),
    IamUserId : util.PtrString(""),
    IamDomainId : util.PtrString(""),
    ResourceIds : []*string{},
    InHeader : util.PtrBool(false),
    InQuery : util.PtrBool(false),
    KeyNames : []*string{},
    Status : util.PtrString(""),

    }
    createConsumerRequest := &CreateConsumerRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        ConsumerName : util.PtrString(""),
        Description : util.PtrString(""),
        AuthType : util.PtrString(""),
        CredentialType : util.PtrString(""),
        RouteNames : []*string{},
        Tags : []*Tag{},
        Credential : Credential,
        IamCredential : IamCredential,
    }
    result := &CreateConsumerResponse{}
    result, err := AIGW_CLIENT.CreateConsumer(createConsumerRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateRoute(t *testing.T) {
    MatchRules := &MatchRule{
        PathRule : &PathRule{
    MatchType : util.PtrString(""),
    Value : util.PtrString(""),
    CaseSensitive : util.PtrBool(false),

    },
    Methods : []*string{},
    Headers : []*MatcherRule{},
    QueryParams : []*MatcherRule{},

    }
    TargetService := &TargetService{
        ServiceSource : util.PtrString(""),
    ServiceName : util.PtrString(""),
    Namespace : util.PtrString(""),
    ServicePort : util.PtrInt32(int32(0)),
    LoadBalanceAlgorithm : util.PtrString(""),
    HashType : util.PtrString(""),
    HashKey : util.PtrString(""),
    RequestRatio : util.PtrInt32(int32(0)),
    ModelName : util.PtrString(""),
    WeightFactor : util.PtrInt32(int32(0)),
    ModelNameMode : util.PtrString(""),
    SpecifiedModelName : util.PtrString(""),

    }
    Rewrite := &Rewrite{
        Enabled : util.PtrBool(false),
    Path : util.PtrString(""),

    }
    RegexRewrite := &RegexRewrite{
        Match : util.PtrString(""),
    Rewrite : util.PtrString(""),

    }
    TokenRateLimit := &TokenRateLimit{
        RuleName : util.PtrString(""),
    Enabled : util.PtrBool(false),
    PreReserveRemainingRatio : util.PtrFloat32(float32(0)),
    PreReserveHistoryWindowSeconds : util.PtrInt32(int32(0)),
    PreReserveSafetyFactor : util.PtrFloat32(float32(0)),
    PreReserveEstimationMode : util.PtrString(""),
    PreReserveInitialTokens : nil,
    SlidingWindowBucketCount : util.PtrInt32(int32(0)),
    PreReserveAdmissionMode : util.PtrString(""),
    PreReserveAdmissionBurstSeconds : util.PtrInt32(int32(0)),
    PreReserveRetryJitterMs : util.PtrInt32(int32(0)),
    RuleItems : []*RuleItem{},

    }
    RequestRateLimit := &RequestRateLimit{
        RuleName : util.PtrString(""),
    Enabled : util.PtrBool(false),
    RuleItems : []*RuleItem{},

    }
    TimeoutPolicy := &TimeoutPolicy{
        Enabled : util.PtrBool(false),
    Timeout : util.PtrInt32(int32(0)),

    }
    RetryPolicy := &RetryPolicy{
        Enabled : util.PtrBool(false),
    RetryConditions : util.PtrString(""),
    NumRetries : util.PtrInt32(int32(0)),

    }
    CorsPolicy := &CorsPolicy{
        Enabled : util.PtrBool(false),
    AllowOrigins : []*OriginMatch[string]string{},
    AllowMethods : []*string{},
    AllowHeaders : []*string{},
    ExposeHeaders : []*string{},
    MaxAge : util.PtrInt32(int32(0)),
    AllowCredentials : util.PtrBool(false),

    }
    ResponseHeaders := &ResponseHeaders{
        Enabled : util.PtrBool(false),
    Headers : []*CustomHeader{},

    }
    FallbackConfig := &FallbackConfig{
        Enabled : util.PtrBool(false),
    ServiceName : util.PtrString(""),
    ModelNameMode : util.PtrString(""),
    SpecifiedModelName : util.PtrString(""),

    }
    createRouteRequest := &CreateRouteRequest{
        InstanceId : util.PtrString(""),
        ClusterId : util.PtrString(""),
        XRegion : util.PtrString(""),
        RouteName : util.PtrString(""),
        SrcProduct : util.PtrString(""),
        AccessMode : util.PtrString(""),
        WebSubdomain : util.PtrString(""),
        ServicePath : util.PtrString(""),
        Domains : []*string{},
        MatchRules : MatchRules,
        MultiService : util.PtrBool(false),
        TrafficDistributionStrategy : util.PtrString(""),
        EnableWeightAdjust : util.PtrBool(false),
        TargetService : TargetService,
        Rewrite : Rewrite,
        RegexRewrite : RegexRewrite,
        CustomHeaders : []*CustomHeader{},
        SkipSetHostHeader : util.PtrBool(false),
        AuthEnabled : util.PtrBool(false),
        AllowedConsumers : []*string{},
        TokenRateLimit : TokenRateLimit,
        RequestRateLimit : RequestRateLimit,
        TimeoutPolicy : TimeoutPolicy,
        RetryPolicy : RetryPolicy,
        CorsPolicy : CorsPolicy,
        ResponseHeaders : ResponseHeaders,
        FallbackConfig : FallbackConfig,
    }
    result := &CreateRouteResponse{}
    result, err := AIGW_CLIENT.CreateRoute(createRouteRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateService(t *testing.T) {
    createServiceRequest := &CreateServiceRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        ServiceSource : util.PtrString(""),
        Namespace : util.PtrString(""),
        ServiceName : util.PtrString(""),
        ClusterId : util.PtrString(""),
        ClusterIds : []*string{},
        ServiceList : []*ServiceItem{},
        RegistryId : util.PtrString(""),
        ServiceAddresses : []*string{},
        ServiceProtocol : util.PtrString(""),
        Provider : util.PtrString(""),
        Endpoint : util.PtrString(""),
        ApiKeys : []*string{},
        CredentialSource : util.PtrString(""),
        CredentialNames : []*string{},
        FailoverEnabled : util.PtrBool(false),
        FailoverModel : util.PtrString(""),
    }
    result := &CreateServiceResponse{}
    result, err := AIGW_CLIENT.CreateService(createServiceRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteAIGateway(t *testing.T) {
    deleteAIGatewayRequest := &DeleteAIGatewayRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        Force : util.PtrBool(false),
    }
    result := &DeleteAIGatewayResponse{}
    result, err := AIGW_CLIENT.DeleteAIGateway(deleteAIGatewayRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteConsumer(t *testing.T) {
    deleteConsumerRequest := &DeleteConsumerRequest{
        InstanceId : util.PtrString(""),
        ConsumerId : util.PtrString(""),
        XRegion : util.PtrString(""),
        KeyType : util.PtrString(""),
    }
    result := &DeleteConsumerResponse{}
    result, err := AIGW_CLIENT.DeleteConsumer(deleteConsumerRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteRoute(t *testing.T) {
    deleteRouteRequest := &DeleteRouteRequest{
        InstanceId : util.PtrString(""),
        RouteName : util.PtrString(""),
        XRegion : util.PtrString(""),
    }
    result := &DeleteRouteResponse{}
    result, err := AIGW_CLIENT.DeleteRoute(deleteRouteRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteService(t *testing.T) {
    deleteServiceRequest := &DeleteServiceRequest{
        InstanceId : util.PtrString(""),
        ServiceName : util.PtrString(""),
        Namespace : util.PtrString(""),
        XRegion : util.PtrString(""),
    }
    result := &DeleteServiceResponse{}
    result, err := AIGW_CLIENT.DeleteService(deleteServiceRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetAIGatewayDetail(t *testing.T) {
    getAIGatewayDetailRequest := &GetAIGatewayDetailRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        SrcProduct : util.PtrString(""),
    }
    result := &GetAIGatewayDetailResponse{}
    result, err := AIGW_CLIENT.GetAIGatewayDetail(getAIGatewayDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetConsumer(t *testing.T) {
    getConsumerRequest := &GetConsumerRequest{
        InstanceId : util.PtrString(""),
        ConsumerId : util.PtrString(""),
        XRegion : util.PtrString(""),
        KeyType : util.PtrString(""),
    }
    result := &GetConsumerResponse{}
    result, err := AIGW_CLIENT.GetConsumer(getConsumerRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetConsumerList(t *testing.T) {
    getConsumerListRequest := &GetConsumerListRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
        TagKey : util.PtrString(""),
        TagValue : util.PtrString(""),
    }
    result := &GetConsumerListResponse{}
    result, err := AIGW_CLIENT.GetConsumerList(getConsumerListRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetServiceDetail(t *testing.T) {
    getServiceDetailRequest := &GetServiceDetailRequest{
        InstanceId : util.PtrString(""),
        ServiceName : util.PtrString(""),
        XRegion : util.PtrString(""),
    }
    result := &GetServiceDetailResponse{}
    result, err := AIGW_CLIENT.GetServiceDetail(getServiceDetailRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetServiceList(t *testing.T) {
    getServiceListRequest := &GetServiceListRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        ServiceSource : util.PtrString(""),
    }
    result := &GetServiceListResponse{}
    result, err := AIGW_CLIENT.GetServiceList(getServiceListRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListAIGateways(t *testing.T) {
    listAIGatewaysRequest := &ListAIGatewaysRequest{
        XRegion : util.PtrString(""),
        Keyword : util.PtrString(""),
        KeywordType : util.PtrString(""),
        Status : util.PtrString(""),
        SrcProduct : util.PtrString(""),
        TagKey : util.PtrString(""),
        TagValue : util.PtrString(""),
        ResourceGroupId : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
        OrderBy : util.PtrString(""),
        Order : util.PtrString(""),
    }
    result := &ListAIGatewaysResponse{}
    result, err := AIGW_CLIENT.ListAIGateways(listAIGatewaysRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListServicesBySource(t *testing.T) {
    listServicesBySourceRequest := &ListServicesBySourceRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        ServiceSource : util.PtrString(""),
    }
    result := &ListServicesBySourceResponse{}
    result, err := AIGW_CLIENT.ListServicesBySource(listServicesBySourceRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryRoutingDetails(t *testing.T) {
    queryRoutingDetailsRequest := &QueryRoutingDetailsRequest{
        InstanceId : util.PtrString(""),
        RouteName : util.PtrString(""),
        XRegion : util.PtrString(""),
    }
    result := &QueryRoutingDetailsResponse{}
    result, err := AIGW_CLIENT.QueryRoutingDetails(queryRoutingDetailsRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryRoutingList(t *testing.T) {
    queryRoutingListRequest := &QueryRoutingListRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        RouteName : util.PtrString(""),
        PageNo : util.PtrInt32(int32(0)),
        PageSize : util.PtrInt32(int32(0)),
        OrderBy : util.PtrString(""),
        Order : util.PtrString(""),
    }
    result := &QueryRoutingListResponse{}
    result, err := AIGW_CLIENT.QueryRoutingList(queryRoutingListRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAIGateway(t *testing.T) {
    updateAIGatewayRequest := &UpdateAIGatewayRequest{
        InstanceId : util.PtrString(""),
        XRegion : util.PtrString(""),
        Name : util.PtrString(""),
        Description : util.PtrString(""),
        DeleteProtection : util.PtrBool(false),
        PublicAccessible : util.PtrBool(false),
        Replicas : util.PtrInt32(int32(0)),
        NetworkTypes : []*string{},
        Tags : []*Tag{},
    }
    result := &UpdateAIGatewayResponse{}
    result, err := AIGW_CLIENT.UpdateAIGateway(updateAIGatewayRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateConsumer(t *testing.T) {
    CredentialOp := &CredentialOp{
        Operation : util.PtrString(""),
    CredentialName : util.PtrString(""),
    Value : util.PtrString(""),

    }
    CredentialLocation := &ConsumerCredentialLocation{
        InHeader : util.PtrBool(false),
    InQuery : util.PtrBool(false),
    KeyNames : []*string{},

    }
    IamCredential := &IAMCredentialSpec{
        Name : util.PtrString(""),
    IamApiKeyId : util.PtrString(""),
    IamTokenIdMasked : util.PtrString(""),
    IamUserId : util.PtrString(""),
    IamDomainId : util.PtrString(""),
    ResourceIds : []*string{},
    InHeader : util.PtrBool(false),
    InQuery : util.PtrBool(false),
    KeyNames : []*string{},
    Status : util.PtrString(""),

    }
    updateConsumerRequest := &UpdateConsumerRequest{
        InstanceId : util.PtrString(""),
        ConsumerId : util.PtrString(""),
        XRegion : util.PtrString(""),
        KeyType : util.PtrString(""),
        Description : util.PtrString(""),
        RouteNames : []*string{},
        Tags : []*Tag{},
        CredentialOp : CredentialOp,
        CredentialLocation : CredentialLocation,
        IamCredential : IamCredential,
    }
    result := &UpdateConsumerResponse{}
    result, err := AIGW_CLIENT.UpdateConsumer(updateConsumerRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateRoute(t *testing.T) {
    MatchRules := &MatchRule{
        PathRule : &PathRule{
    MatchType : util.PtrString(""),
    Value : util.PtrString(""),
    CaseSensitive : util.PtrBool(false),

    },
    Methods : []*string{},
    Headers : []*MatcherRule{},
    QueryParams : []*MatcherRule{},

    }
    TargetService := &TargetService{
        ServiceSource : util.PtrString(""),
    ServiceName : util.PtrString(""),
    Namespace : util.PtrString(""),
    ServicePort : util.PtrInt32(int32(0)),
    LoadBalanceAlgorithm : util.PtrString(""),
    HashType : util.PtrString(""),
    HashKey : util.PtrString(""),
    RequestRatio : util.PtrInt32(int32(0)),
    ModelName : util.PtrString(""),
    WeightFactor : util.PtrInt32(int32(0)),
    ModelNameMode : util.PtrString(""),
    SpecifiedModelName : util.PtrString(""),

    }
    Rewrite := &Rewrite{
        Enabled : util.PtrBool(false),
    Path : util.PtrString(""),

    }
    RegexRewrite := &RegexRewrite{
        Match : util.PtrString(""),
    Rewrite : util.PtrString(""),

    }
    TokenRateLimit := &TokenRateLimit{
        RuleName : util.PtrString(""),
    Enabled : util.PtrBool(false),
    PreReserveRemainingRatio : util.PtrFloat32(float32(0)),
    PreReserveHistoryWindowSeconds : util.PtrInt32(int32(0)),
    PreReserveSafetyFactor : util.PtrFloat32(float32(0)),
    PreReserveEstimationMode : util.PtrString(""),
    PreReserveInitialTokens : nil,
    SlidingWindowBucketCount : util.PtrInt32(int32(0)),
    PreReserveAdmissionMode : util.PtrString(""),
    PreReserveAdmissionBurstSeconds : util.PtrInt32(int32(0)),
    PreReserveRetryJitterMs : util.PtrInt32(int32(0)),
    RuleItems : []*RuleItem{},

    }
    RequestRateLimit := &RequestRateLimit{
        RuleName : util.PtrString(""),
    Enabled : util.PtrBool(false),
    RuleItems : []*RuleItem{},

    }
    TimeoutPolicy := &TimeoutPolicy{
        Enabled : util.PtrBool(false),
    Timeout : util.PtrInt32(int32(0)),

    }
    RetryPolicy := &RetryPolicy{
        Enabled : util.PtrBool(false),
    RetryConditions : util.PtrString(""),
    NumRetries : util.PtrInt32(int32(0)),

    }
    CorsPolicy := &CorsPolicy{
        Enabled : util.PtrBool(false),
    AllowOrigins : []*OriginMatch[string]string{},
    AllowMethods : []*string{},
    AllowHeaders : []*string{},
    ExposeHeaders : []*string{},
    MaxAge : util.PtrInt32(int32(0)),
    AllowCredentials : util.PtrBool(false),

    }
    ResponseHeaders := &ResponseHeaders{
        Enabled : util.PtrBool(false),
    Headers : []*CustomHeader{},

    }
    FallbackConfig := &FallbackConfig{
        Enabled : util.PtrBool(false),
    ServiceName : util.PtrString(""),
    ModelNameMode : util.PtrString(""),
    SpecifiedModelName : util.PtrString(""),

    }
    updateRouteRequest := &UpdateRouteRequest{
        InstanceId : util.PtrString(""),
        RouteName : util.PtrString(""),
        XRegion : util.PtrString(""),
        SrcProduct : util.PtrString(""),
        AccessMode : util.PtrString(""),
        WebSubdomain : util.PtrString(""),
        ServicePath : util.PtrString(""),
        Domains : []*string{},
        MatchRules : MatchRules,
        MultiService : util.PtrBool(false),
        TrafficDistributionStrategy : util.PtrString(""),
        EnableWeightAdjust : util.PtrBool(false),
        TargetService : TargetService,
        Rewrite : Rewrite,
        RegexRewrite : RegexRewrite,
        CustomHeaders : []*CustomHeader{},
        SkipSetHostHeader : util.PtrBool(false),
        AuthEnabled : util.PtrBool(false),
        AllowedConsumers : []*string{},
        TokenRateLimit : TokenRateLimit,
        RequestRateLimit : RequestRateLimit,
        TimeoutPolicy : TimeoutPolicy,
        RetryPolicy : RetryPolicy,
        CorsPolicy : CorsPolicy,
        ResponseHeaders : ResponseHeaders,
        FallbackConfig : FallbackConfig,
    }
    result := &UpdateRouteResponse{}
    result, err := AIGW_CLIENT.UpdateRoute(updateRouteRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateService(t *testing.T) {
    updateServiceRequest := &UpdateServiceRequest{
        InstanceId : util.PtrString(""),
        ServiceNamePath : util.PtrString(""),
        XRegion : util.PtrString(""),
        ServiceName : util.PtrString(""),
        ServiceAddresses : []*string{},
        ServiceProtocol : util.PtrString(""),
        Provider : util.PtrString(""),
        Endpoint : util.PtrString(""),
        ApiKeys : []*string{},
        FailoverEnabled : util.PtrBool(false),
        FailoverModel : util.PtrString(""),
        CredentialSource : util.PtrString(""),
        CredentialNames : []*string{},
    }
    result := &UpdateServiceResponse{}
    result, err := AIGW_CLIENT.UpdateService(updateServiceRequest)
    if err != nil {
        fmt.Println("request failed:", err)
        return
    }
    data, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        fmt.Println("json marshalIndent failed:", err)
        return
    }
    fmt.Println(string(data))
    ExpectEqual(t.Errorf, nil, err)
}
