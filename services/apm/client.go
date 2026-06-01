package apm

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "apm." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V1 = "v1"

	CONSTANT_APM = "apm"

	CONSTANT_QUERY = "query"

	CONSTANT_ALARM = "alarm"
)

// Client of apm service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getApmCreateAlarmPolicyUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmDeleteAlarmPolicyUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmDescribeAlarmUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_ALARM
}
func getApmDescribeAlarmPoliciesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmDescribeAlarmPolicyUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmDescribeAlarmsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_ALARM
}
func getApmUpdateAlarmPolicyUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmUpdateAlarmPolicyActionUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getApmUpdateAlarmPolicyStateUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getBindServiceTagUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getDeleteServicesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getDescribeDbStatementUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeDefaultConfigUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getDescribeDimensionValuesUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeEnvUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeExceptionsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMDimensionValuesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMMetricDataUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMServicesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMSessionUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMSessionsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMSessionsStatisticsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMSpansUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMTraceUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMTracesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeLLMTracesStatisticsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeMetricDataUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeRetentionLimitUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getDescribeServiceConfigUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
func getDescribeServicesMetricsUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeServicesNamesUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeSpanFieldValuesUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeSpansUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeTopologyUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeTraceUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getDescribeTraceMetricDataUri(Action string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM + bce.URI_PREFIX + CONSTANT_QUERY
}
func getUpdateServiceConfigUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_APM
}
