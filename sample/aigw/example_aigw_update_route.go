package aigwsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aigw"
)

func UpdateRoute() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := aigw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	MatchRules := &aigw.MatchRule{
		PathRule: &aigw.PathRule{
			MatchType:     util.PtrString(""),
			Value:         util.PtrString(""),
			CaseSensitive: util.PtrBool(false),
		},
		Methods:     []*string{},
		Headers:     []*aigw.MatcherRule{},
		QueryParams: []*aigw.MatcherRule{},
	}
	TargetService := &aigw.TargetService{
		ServiceSource:        util.PtrString(""),
		ServiceName:          util.PtrString(""),
		Namespace:            util.PtrString(""),
		ServicePort:          util.PtrInt32(int32(0)),
		LoadBalanceAlgorithm: util.PtrString(""),
		HashType:             util.PtrString(""),
		HashKey:              util.PtrString(""),
		RequestRatio:         util.PtrInt32(int32(0)),
		ModelName:            util.PtrString(""),
		WeightFactor:         util.PtrInt32(int32(0)),
		ModelNameMode:        util.PtrString(""),
		SpecifiedModelName:   util.PtrString(""),
	}
	Rewrite := &aigw.Rewrite{
		Enabled: util.PtrBool(false),
		Path:    util.PtrString(""),
	}
	RegexRewrite := &aigw.RegexRewrite{
		Match:   util.PtrString(""),
		Rewrite: util.PtrString(""),
	}
	TokenRateLimit := &aigw.TokenRateLimit{
		RuleName:                        util.PtrString(""),
		Enabled:                         util.PtrBool(false),
		PreReserveRemainingRatio:        util.PtrFloat32(float32(0)),
		PreReserveHistoryWindowSeconds:  util.PtrInt32(int32(0)),
		PreReserveSafetyFactor:          util.PtrFloat32(float32(0)),
		PreReserveEstimationMode:        util.PtrString(""),
		PreReserveInitialTokens:         nil,
		SlidingWindowBucketCount:        util.PtrInt32(int32(0)),
		PreReserveAdmissionMode:         util.PtrString(""),
		PreReserveAdmissionBurstSeconds: util.PtrInt32(int32(0)),
		PreReserveRetryJitterMs:         util.PtrInt32(int32(0)),
		RuleItems:                       []*aigw.RuleItem{},
	}
	RequestRateLimit := &aigw.RequestRateLimit{
		RuleName:  util.PtrString(""),
		Enabled:   util.PtrBool(false),
		RuleItems: []*aigw.RuleItem{},
	}
	TimeoutPolicy := &aigw.TimeoutPolicy{
		Enabled: util.PtrBool(false),
		Timeout: util.PtrInt32(int32(0)),
	}
	RetryPolicy := &aigw.RetryPolicy{
		Enabled:         util.PtrBool(false),
		RetryConditions: util.PtrString(""),
		NumRetries:      util.PtrInt32(int32(0)),
	}
	CorsPolicy := &aigw.CorsPolicy{
		Enabled:          util.PtrBool(false),
		AllowOrigins:     []*aigw.OriginMatch{},
		AllowMethods:     []*string{},
		AllowHeaders:     []*string{},
		ExposeHeaders:    []*string{},
		MaxAge:           util.PtrInt32(int32(0)),
		AllowCredentials: util.PtrBool(false),
	}
	ResponseHeaders := &aigw.ResponseHeaders{
		Enabled: util.PtrBool(false),
		Headers: []*aigw.CustomHeader{},
	}
	FallbackConfig := &aigw.FallbackConfig{
		Enabled:            util.PtrBool(false),
		ServiceName:        util.PtrString(""),
		ModelNameMode:      util.PtrString(""),
		SpecifiedModelName: util.PtrString(""),
	}
	updateRouteRequest := &aigw.UpdateRouteRequest{
		InstanceId:                  util.PtrString(""),
		RouteName:                   util.PtrString(""),
		XRegion:                     util.PtrString(""),
		SrcProduct:                  util.PtrString(""),
		AccessMode:                  util.PtrString(""),
		WebSubdomain:                util.PtrString(""),
		ServicePath:                 util.PtrString(""),
		Domains:                     []*string{},
		MatchRules:                  MatchRules,
		MultiService:                util.PtrBool(false),
		TrafficDistributionStrategy: util.PtrString(""),
		EnableWeightAdjust:          util.PtrBool(false),
		TargetService:               TargetService,
		Rewrite:                     Rewrite,
		RegexRewrite:                RegexRewrite,
		CustomHeaders:               []*aigw.CustomHeader{},
		SkipSetHostHeader:           util.PtrBool(false),
		AuthEnabled:                 util.PtrBool(false),
		AllowedConsumers:            []*string{},
		TokenRateLimit:              TokenRateLimit,
		RequestRateLimit:            RequestRateLimit,
		TimeoutPolicy:               TimeoutPolicy,
		RetryPolicy:                 RetryPolicy,
		CorsPolicy:                  CorsPolicy,
		ResponseHeaders:             ResponseHeaders,
		FallbackConfig:              FallbackConfig,
	}
	result, err := client.UpdateRoute(updateRouteRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
}
