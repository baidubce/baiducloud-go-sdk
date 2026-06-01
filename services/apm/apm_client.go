package apm

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// ApmCreateAlarmPolicy
//
// PARAMS:
//   - request: the arguments to ApmCreateAlarmPolicy
//
// RETURNS:
//   - ApmCreateAlarmPolicyResponse: The return type of the ApmCreateAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmCreateAlarmPolicy(request *ApmCreateAlarmPolicyRequest) (*ApmCreateAlarmPolicyResponse, error) {
	result := &ApmCreateAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmCreateAlarmPolicyUri()).
		WithQueryParamFilter("action", "CreateAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmDeleteAlarmPolicy
//
// PARAMS:
//   - request: the arguments to ApmDeleteAlarmPolicy
//
// RETURNS:
//   - ApmDeleteAlarmPolicyResponse: The return type of the ApmDeleteAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmDeleteAlarmPolicy(request *ApmDeleteAlarmPolicyRequest) (*ApmDeleteAlarmPolicyResponse, error) {
	result := &ApmDeleteAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmDeleteAlarmPolicyUri()).
		WithQueryParamFilter("action", "DeleteAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmDescribeAlarm
//
// PARAMS:
//   - request: the arguments to ApmDescribeAlarm
//
// RETURNS:
//   - ApmDescribeAlarmResponse: The return type of the ApmDescribeAlarm interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmDescribeAlarm(request *ApmDescribeAlarmRequest) (*ApmDescribeAlarmResponse, error) {
	result := &ApmDescribeAlarmResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmDescribeAlarmUri()).
		WithQueryParamFilter("action", "DescribeAlarm").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmDescribeAlarmPolicies
//
// PARAMS:
//   - request: the arguments to ApmDescribeAlarmPolicies
//
// RETURNS:
//   - ApmDescribeAlarmPoliciesResponse: The return type of the ApmDescribeAlarmPolicies interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmDescribeAlarmPolicies(request *ApmDescribeAlarmPoliciesRequest) (*ApmDescribeAlarmPoliciesResponse, error) {
	result := &ApmDescribeAlarmPoliciesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmDescribeAlarmPoliciesUri()).
		WithQueryParamFilter("action", "DescribeAlarmPolicies").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmDescribeAlarmPolicy
//
// PARAMS:
//   - request: the arguments to ApmDescribeAlarmPolicy
//
// RETURNS:
//   - ApmDescribeAlarmPolicyResponse: The return type of the ApmDescribeAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmDescribeAlarmPolicy(request *ApmDescribeAlarmPolicyRequest) (*ApmDescribeAlarmPolicyResponse, error) {
	result := &ApmDescribeAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmDescribeAlarmPolicyUri()).
		WithQueryParamFilter("action", "DescribeAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmDescribeAlarms
//
// PARAMS:
//   - request: the arguments to ApmDescribeAlarms
//
// RETURNS:
//   - ApmDescribeAlarmsResponse: The return type of the ApmDescribeAlarms interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmDescribeAlarms(request *ApmDescribeAlarmsRequest) (*ApmDescribeAlarmsResponse, error) {
	result := &ApmDescribeAlarmsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmDescribeAlarmsUri()).
		WithQueryParamFilter("action", "DescribeAlarms").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmUpdateAlarmPolicy
//
// PARAMS:
//   - request: the arguments to ApmUpdateAlarmPolicy
//
// RETURNS:
//   - ApmUpdateAlarmPolicyResponse: The return type of the ApmUpdateAlarmPolicy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmUpdateAlarmPolicy(request *ApmUpdateAlarmPolicyRequest) (*ApmUpdateAlarmPolicyResponse, error) {
	result := &ApmUpdateAlarmPolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmUpdateAlarmPolicyUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicy").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmUpdateAlarmPolicyAction
//
// PARAMS:
//   - request: the arguments to ApmUpdateAlarmPolicyAction
//
// RETURNS:
//   - ApmUpdateAlarmPolicyActionResponse: The return type of the ApmUpdateAlarmPolicyAction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmUpdateAlarmPolicyAction(request *ApmUpdateAlarmPolicyActionRequest) (*ApmUpdateAlarmPolicyActionResponse, error) {
	result := &ApmUpdateAlarmPolicyActionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmUpdateAlarmPolicyActionUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicyAction").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ApmUpdateAlarmPolicyState
//
// PARAMS:
//   - request: the arguments to ApmUpdateAlarmPolicyState
//
// RETURNS:
//   - ApmUpdateAlarmPolicyStateResponse: The return type of the ApmUpdateAlarmPolicyState interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ApmUpdateAlarmPolicyState(request *ApmUpdateAlarmPolicyStateRequest) (*ApmUpdateAlarmPolicyStateResponse, error) {
	result := &ApmUpdateAlarmPolicyStateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getApmUpdateAlarmPolicyStateUri()).
		WithQueryParamFilter("action", "UpdateAlarmPolicyState").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindServiceTag
//
// PARAMS:
//   - request: the arguments to BindServiceTag
//
// RETURNS:
//   - BindServiceTagResponse: The return type of the BindServiceTag interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BindServiceTag(request *BindServiceTagRequest) (*BindServiceTagResponse, error) {
	result := &BindServiceTagResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindServiceTagUri()).
		WithQueryParamFilter("action", "BindServiceTag").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteServices
//
// PARAMS:
//   - request: the arguments to DeleteServices
//
// RETURNS:
//   - DeleteServicesResponse: The return type of the DeleteServices interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteServices(request *DeleteServicesRequest) (*DeleteServicesResponse, error) {
	result := &DeleteServicesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteServicesUri()).
		WithQueryParamFilter("action", "DeleteServices").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDbStatement
//
// PARAMS:
//   - request: the arguments to DescribeDbStatement
//
// RETURNS:
//   - DescribeDbStatementResponse: The return type of the DescribeDbStatement interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDbStatement(request *DescribeDbStatementRequest) (*DescribeDbStatementResponse, error) {
	result := &DescribeDbStatementResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDbStatementUri()).
		WithQueryParamFilter("action", "DescribeDbStatement").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDefaultConfig
//
// PARAMS:
//   - request: the arguments to DescribeDefaultConfig
//
// RETURNS:
//   - DescribeDefaultConfigResponse: The return type of the DescribeDefaultConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDefaultConfig() (*DescribeDefaultConfigResponse, error) {
	result := &DescribeDefaultConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDefaultConfigUri()).
		WithQueryParamFilter("action", "DescribeDefaultConfig").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeDimensionValues
//
// PARAMS:
//   - request: the arguments to DescribeDimensionValues
//
// RETURNS:
//   - DescribeDimensionValuesResponse: The return type of the DescribeDimensionValues interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeDimensionValues(request *DescribeDimensionValuesRequest) (*DescribeDimensionValuesResponse, error) {
	result := &DescribeDimensionValuesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeDimensionValuesUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeDimensionValues").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeEnv
//
// PARAMS:
//   - request: the arguments to DescribeEnv
//
// RETURNS:
//   - DescribeEnvResponse: The return type of the DescribeEnv interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeEnv() (*DescribeEnvResponse, error) {
	result := &DescribeEnvResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeEnvUri()).
		WithQueryParamFilter("action", "DescribeEnv").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeExceptions
//
// PARAMS:
//   - request: the arguments to DescribeExceptions
//
// RETURNS:
//   - DescribeExceptionsResponse: The return type of the DescribeExceptions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeExceptions(request *DescribeExceptionsRequest) (*DescribeExceptionsResponse, error) {
	result := &DescribeExceptionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeExceptionsUri()).
		WithQueryParamFilter("action", "DescribeExceptions").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMDimensionValues
//
// PARAMS:
//   - request: the arguments to DescribeLLMDimensionValues
//
// RETURNS:
//   - DescribeLLMDimensionValuesResponse: The return type of the DescribeLLMDimensionValues interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMDimensionValues(request *DescribeLLMDimensionValuesRequest) (*DescribeLLMDimensionValuesResponse, error) {
	result := &DescribeLLMDimensionValuesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMDimensionValuesUri()).
		WithQueryParamFilter("action", "DescribeLLMDimensionValues").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMMetricData
//
// PARAMS:
//   - request: the arguments to DescribeLLMMetricData
//
// RETURNS:
//   - DescribeLLMMetricDataResponse: The return type of the DescribeLLMMetricData interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMMetricData(request *DescribeLLMMetricDataRequest) (*DescribeLLMMetricDataResponse, error) {
	result := &DescribeLLMMetricDataResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMMetricDataUri()).
		WithQueryParamFilter("action", "DescribeLLMMetricData").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMServices
//
// PARAMS:
//   - request: the arguments to DescribeLLMServices
//
// RETURNS:
//   - DescribeLLMServicesResponse: The return type of the DescribeLLMServices interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMServices(request *DescribeLLMServicesRequest) (*DescribeLLMServicesResponse, error) {
	result := &DescribeLLMServicesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMServicesUri()).
		WithQueryParamFilter("action", "DescribeLLMServices").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMSession
//
// PARAMS:
//   - request: the arguments to DescribeLLMSession
//
// RETURNS:
//   - DescribeLLMSessionResponse: The return type of the DescribeLLMSession interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMSession(request *DescribeLLMSessionRequest) (*DescribeLLMSessionResponse, error) {
	result := &DescribeLLMSessionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMSessionUri()).
		WithQueryParamFilter("action", "DescribeLLMSession").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMSessions
//
// PARAMS:
//   - request: the arguments to DescribeLLMSessions
//
// RETURNS:
//   - DescribeLLMSessionsResponse: The return type of the DescribeLLMSessions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMSessions(request *DescribeLLMSessionsRequest) (*DescribeLLMSessionsResponse, error) {
	result := &DescribeLLMSessionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMSessionsUri()).
		WithQueryParamFilter("action", "DescribeLLMSessions").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMSessionsStatistics
//
// PARAMS:
//   - request: the arguments to DescribeLLMSessionsStatistics
//
// RETURNS:
//   - DescribeLLMSessionsStatisticsResponse: The return type of the DescribeLLMSessionsStatistics interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMSessionsStatistics(request *DescribeLLMSessionsStatisticsRequest) (*DescribeLLMSessionsStatisticsResponse, error) {
	result := &DescribeLLMSessionsStatisticsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMSessionsStatisticsUri()).
		WithQueryParamFilter("action", "DescribeLLMSessionsStatistics").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMSpans
//
// PARAMS:
//   - request: the arguments to DescribeLLMSpans
//
// RETURNS:
//   - DescribeLLMSpansResponse: The return type of the DescribeLLMSpans interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMSpans(request *DescribeLLMSpansRequest) (*DescribeLLMSpansResponse, error) {
	result := &DescribeLLMSpansResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMSpansUri()).
		WithQueryParamFilter("action", "DescribeLLMSpans").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMTrace
//
// PARAMS:
//   - request: the arguments to DescribeLLMTrace
//
// RETURNS:
//   - DescribeLLMTraceResponse: The return type of the DescribeLLMTrace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMTrace(request *DescribeLLMTraceRequest) (*DescribeLLMTraceResponse, error) {
	result := &DescribeLLMTraceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMTraceUri()).
		WithQueryParamFilter("action", "DescribeLLMTrace").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMTraces
//
// PARAMS:
//   - request: the arguments to DescribeLLMTraces
//
// RETURNS:
//   - DescribeLLMTracesResponse: The return type of the DescribeLLMTraces interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMTraces(request *DescribeLLMTracesRequest) (*DescribeLLMTracesResponse, error) {
	result := &DescribeLLMTracesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMTracesUri()).
		WithQueryParamFilter("action", "DescribeLLMTraces").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeLLMTracesStatistics
//
// PARAMS:
//   - request: the arguments to DescribeLLMTracesStatistics
//
// RETURNS:
//   - DescribeLLMTracesStatisticsResponse: The return type of the DescribeLLMTracesStatistics interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeLLMTracesStatistics(request *DescribeLLMTracesStatisticsRequest) (*DescribeLLMTracesStatisticsResponse, error) {
	result := &DescribeLLMTracesStatisticsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeLLMTracesStatisticsUri()).
		WithQueryParamFilter("action", "DescribeLLMTracesStatistics").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeMetricData
//
// PARAMS:
//   - request: the arguments to DescribeMetricData
//
// RETURNS:
//   - DescribeMetricDataResponse: The return type of the DescribeMetricData interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeMetricData(request *DescribeMetricDataRequest) (*DescribeMetricDataResponse, error) {
	result := &DescribeMetricDataResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeMetricDataUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeMetricData").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeRetentionLimit
//
// PARAMS:
//   - request: the arguments to DescribeRetentionLimit
//
// RETURNS:
//   - DescribeRetentionLimitResponse: The return type of the DescribeRetentionLimit interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeRetentionLimit() (*DescribeRetentionLimitResponse, error) {
	result := &DescribeRetentionLimitResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeRetentionLimitUri()).
		WithQueryParamFilter("action", "DescribeRetentionLimit").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeServiceConfig
//
// PARAMS:
//   - request: the arguments to DescribeServiceConfig
//
// RETURNS:
//   - DescribeServiceConfigResponse: The return type of the DescribeServiceConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeServiceConfig(request *DescribeServiceConfigRequest) (*DescribeServiceConfigResponse, error) {
	result := &DescribeServiceConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeServiceConfigUri()).
		WithQueryParamFilter("action", "DescribeServiceConfig").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeServicesMetrics
//
// PARAMS:
//   - request: the arguments to DescribeServicesMetrics
//
// RETURNS:
//   - DescribeServicesMetricsResponse: The return type of the DescribeServicesMetrics interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeServicesMetrics(request *DescribeServicesMetricsRequest) (*DescribeServicesMetricsResponse, error) {
	result := &DescribeServicesMetricsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeServicesMetricsUri()).
		WithQueryParamFilter("action", "DescribeServicesMetrics").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeServicesNames
//
// PARAMS:
//   - request: the arguments to DescribeServicesNames
//
// RETURNS:
//   - DescribeServicesNamesResponse: The return type of the DescribeServicesNames interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeServicesNames(request *DescribeServicesNamesRequest) (*DescribeServicesNamesResponse, error) {
	result := &DescribeServicesNamesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeServicesNamesUri()).
		WithQueryParamFilter("action", "DescribeServicesNames").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeSpanFieldValues
//
// PARAMS:
//   - request: the arguments to DescribeSpanFieldValues
//
// RETURNS:
//   - DescribeSpanFieldValuesResponse: The return type of the DescribeSpanFieldValues interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeSpanFieldValues(request *DescribeSpanFieldValuesRequest) (*DescribeSpanFieldValuesResponse, error) {
	result := &DescribeSpanFieldValuesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeSpanFieldValuesUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeSpanFieldValues").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeSpans
//
// PARAMS:
//   - request: the arguments to DescribeSpans
//
// RETURNS:
//   - DescribeSpansResponse: The return type of the DescribeSpans interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeSpans(request *DescribeSpansRequest) (*DescribeSpansResponse, error) {
	result := &DescribeSpansResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeSpansUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeSpans").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeTopology
//
// PARAMS:
//   - request: the arguments to DescribeTopology
//
// RETURNS:
//   - DescribeTopologyResponse: The return type of the DescribeTopology interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeTopology(request *DescribeTopologyRequest) (*DescribeTopologyResponse, error) {
	result := &DescribeTopologyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeTopologyUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeTopology").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeTrace
//
// PARAMS:
//   - request: the arguments to DescribeTrace
//
// RETURNS:
//   - DescribeTraceResponse: The return type of the DescribeTrace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeTrace(request *DescribeTraceRequest) (*DescribeTraceResponse, error) {
	result := &DescribeTraceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeTraceUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeTrace").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescribeTraceMetricData
//
// PARAMS:
//   - request: the arguments to DescribeTraceMetricData
//
// RETURNS:
//   - DescribeTraceMetricDataResponse: The return type of the DescribeTraceMetricData interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescribeTraceMetricData(request *DescribeTraceMetricDataRequest) (*DescribeTraceMetricDataResponse, error) {
	result := &DescribeTraceMetricDataResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescribeTraceMetricDataUri(util.StringValue(request.Action))).
		WithQueryParamFilter("action", "DescribeTraceMetricData").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateServiceConfig
//
// PARAMS:
//   - request: the arguments to UpdateServiceConfig
//
// RETURNS:
//   - UpdateServiceConfigResponse: The return type of the UpdateServiceConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateServiceConfig(request *UpdateServiceConfigRequest) (*UpdateServiceConfigResponse, error) {
	result := &UpdateServiceConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateServiceConfigUri()).
		WithQueryParamFilter("action", "UpdateServiceConfig").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
