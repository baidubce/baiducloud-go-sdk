package bcmsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcm"
)

func DescribeMetricDataLatest() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcm.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	describeMetricDataLatestRequest := &bcm.DescribeMetricDataLatestRequest{
		Action:              util.PtrString(""),
		Scope:               util.PtrString(""),
		ResourceType:        util.PtrString(""),
		Region:              util.PtrString(""),
		EndDatetime:         util.PtrString(""),
		MetricName:          util.PtrString(""),
		Filters:             []*bcm.Filter{},
		Limit:               util.PtrInt32(int32(0)),
		Offset:              util.PtrInt32(int32(0)),
		PeriodSeconds:       util.PtrInt32(int32(0)),
		AggregationOverTime: []*string{},
	}
	result, err := client.DescribeMetricDataLatest(describeMetricDataLatestRequest)
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
