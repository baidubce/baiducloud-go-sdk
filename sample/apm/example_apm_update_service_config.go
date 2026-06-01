package apmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/services/apm"
)

func UpdateServiceConfig() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := apm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	SampleConfig := &apm.SampleConfig{
		Enabled:    util.PtrBool(false),
		Processors: []*apm.SampleProcessor{},
	}
	LoggingConfig := &apm.LoggingConfig{
		Enabled:      util.PtrBool(false),
		Region:       util.PtrString(""),
		Project:      util.PtrString(""),
		LogStoreName: util.PtrString(""),
		TraceIdIndex: util.PtrString(""),
		TraceIdKey:   util.PtrString(""),
		SpanIdIndex:  util.PtrString(""),
		SpanIdKey:    util.PtrString(""),
	}
	RequestConfig := &apm.ServiceRequestConfig{
		Global:                            util.PtrBool(false),
		ServerSlowRequestThresholdSeconds: util.PtrFloat64(float64(0)),
		DbSlowRequestThresholdSeconds:     util.PtrFloat64(float64(0)),
		OkHttpStatus:                      []*int32{},
	}
	TopoConfig := &apm.ServiceTopoConfig{
		Global:                  util.PtrBool(false),
		RequestSecondsThreshold: util.PtrFloat64(float64(0)),
		ErrorRateThreshold:      util.PtrFloat64(float64(0)),
	}
	MllmResourceDumpConfig := &apm.MllmResourceDumpConfig{
		RetentionDays: util.PtrInt32(int32(0)),
		Bucket:        util.PtrString(""),
	}
	updateServiceConfigRequest := &apm.UpdateServiceConfigRequest{
		ServiceNames:           []*string{},
		SampleConfig:           SampleConfig,
		LoggingConfig:          LoggingConfig,
		RequestConfig:          RequestConfig,
		TopoConfig:             TopoConfig,
		MllmResourceDumpConfig: MllmResourceDumpConfig,
	}
	result, err := client.UpdateServiceConfig(updateServiceConfigRequest)
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
