package aihcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aihc"
)

func DescribeJobMetrics() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := aihc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	describeJobMetricsRequest := &aihc.DescribeJobMetricsRequest{
		ResourcePoolId: util.PtrString(""),
		QueueID:        util.PtrString(""),
		JobId:          util.PtrString(""),
		StartTime:      util.PtrString(""),
		EndTime:        util.PtrString(""),
		TimeStep:       util.PtrString(""),
		MetricType:     util.PtrString(""),
		RateInterval:   util.PtrString(""),
	}
	result, err := client.DescribeJobMetrics(describeJobMetricsRequest)
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
