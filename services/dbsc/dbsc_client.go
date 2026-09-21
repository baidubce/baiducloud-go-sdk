package dbsc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"strings"
)

const (
	VERSION_V1 = "v1"
)

// CheckMysqlRateLimitSupport
//
// PARAMS:
//   - request: the arguments to CheckMysqlRateLimitSupport
//
// RETURNS:
//   - CheckMysqlRateLimitSupportResponse: The return type of the CheckMysqlRateLimitSupport interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckMysqlRateLimitSupport(request *CheckMysqlRateLimitSupportRequest) (*CheckMysqlRateLimitSupportResponse, error) {
	result := &CheckMysqlRateLimitSupportResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCheckMysqlRateLimitSupportUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateMysqlRateLimitTask
//
// PARAMS:
//   - request: the arguments to CreateMysqlRateLimitTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateMysqlRateLimitTask(request *CreateMysqlRateLimitTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateMysqlRateLimitTaskUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateRedisBigKeyAnalysisTask
//
// PARAMS:
//   - request: the arguments to CreateRedisBigKeyAnalysisTask
//
// RETURNS:
//   - CreateRedisBigKeyAnalysisTaskResponse: The return type of the CreateRedisBigKeyAnalysisTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateRedisBigKeyAnalysisTask(request *CreateRedisBigKeyAnalysisTaskRequest) (*CreateRedisBigKeyAnalysisTaskResponse, error) {
	result := &CreateRedisBigKeyAnalysisTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateRedisBigKeyAnalysisTaskUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteMysqlRateLimitTask
//
// PARAMS:
//   - request: the arguments to DeleteMysqlRateLimitTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteMysqlRateLimitTask(request *DeleteMysqlRateLimitTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteMysqlRateLimitTaskUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("filterId", "filterId").
		WithQueryParamFilter("filterId", util.Int32Value(request.FilterId)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		Do()
}

// DeleteRedisBigKeyAnalysisTask
//
// PARAMS:
//   - request: the arguments to DeleteRedisBigKeyAnalysisTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRedisBigKeyAnalysisTask(request *DeleteRedisBigKeyAnalysisTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRedisBigKeyAnalysisTaskUri(VERSION_V1)).
		WithQueryParamFilter("ids", strings.Join(util.PtrSliceToStringSlice(request.Ids), ",")).
		WithBody(request).
		Do()
}

// GetMongodbCollectionIndexes
//
// PARAMS:
//   - request: the arguments to GetMongodbCollectionIndexes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetMongodbCollectionIndexes(request *GetMongodbCollectionIndexesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbCollectionIndexesUri(VERSION_V1)).
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("nodeId", "string").
		WithQueryParamFilter("database", "string").
		WithQueryParamFilter("collection", "string").
		WithQueryParamFilter("product", util.StringValue(request.Product)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("collection", util.StringValue(request.Collection)).
		Do()
}

// GetMongodbCollectionSpace
//
// PARAMS:
//   - request: the arguments to GetMongodbCollectionSpace
//
// RETURNS:
//   - GetMongodbCollectionSpaceResponse: The return type of the GetMongodbCollectionSpace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbCollectionSpace(request *GetMongodbCollectionSpaceRequest) (*GetMongodbCollectionSpaceResponse, error) {
	result := &GetMongodbCollectionSpaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbCollectionSpaceUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("collection", util.StringValue(request.Collection)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbCollectionSpaceTrend
//
// PARAMS:
//   - request: the arguments to GetMongodbCollectionSpaceTrend
//
// RETURNS:
//   - GetMongodbCollectionSpaceTrendResponse: The return type of the GetMongodbCollectionSpaceTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbCollectionSpaceTrend(request *GetMongodbCollectionSpaceTrendRequest) (*GetMongodbCollectionSpaceTrendResponse, error) {
	result := &GetMongodbCollectionSpaceTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbCollectionSpaceTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("collection", util.StringValue(request.Collection)).
		WithQueryParamFilter("period", util.Int32Value(request.Period)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("metrics", util.StringValue(request.Metrics)).
		WithQueryParamFilter("statistics", util.StringValue(request.Statistics)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbDatabaseSpace
//
// PARAMS:
//   - request: the arguments to GetMongodbDatabaseSpace
//
// RETURNS:
//   - GetMongodbDatabaseSpaceResponse: The return type of the GetMongodbDatabaseSpace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbDatabaseSpace(request *GetMongodbDatabaseSpaceRequest) (*GetMongodbDatabaseSpaceResponse, error) {
	result := &GetMongodbDatabaseSpaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbDatabaseSpaceUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbDatabaseSpaceTrend
//
// PARAMS:
//   - request: the arguments to GetMongodbDatabaseSpaceTrend
//
// RETURNS:
//   - GetMongodbDatabaseSpaceTrendResponse: The return type of the GetMongodbDatabaseSpaceTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbDatabaseSpaceTrend(request *GetMongodbDatabaseSpaceTrendRequest) (*GetMongodbDatabaseSpaceTrendResponse, error) {
	result := &GetMongodbDatabaseSpaceTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbDatabaseSpaceTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("period", util.Int32Value(request.Period)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("metrics", util.StringValue(request.Metrics)).
		WithQueryParamFilter("statistics", util.StringValue(request.Statistics)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbSlowLogTimeDistribution
//
// PARAMS:
//   - request: the arguments to GetMongodbSlowLogTimeDistribution
//
// RETURNS:
//   - GetMongodbSlowLogTimeDistributionResponse: The return type of the GetMongodbSlowLogTimeDistribution interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbSlowLogTimeDistribution(request *GetMongodbSlowLogTimeDistributionRequest) (*GetMongodbSlowLogTimeDistributionResponse, error) {
	result := &GetMongodbSlowLogTimeDistributionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbSlowLogTimeDistributionUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbSlowLogTrend
//
// PARAMS:
//   - request: the arguments to GetMongodbSlowLogTrend
//
// RETURNS:
//   - GetMongodbSlowLogTrendResponse: The return type of the GetMongodbSlowLogTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbSlowLogTrend(request *GetMongodbSlowLogTrendRequest) (*GetMongodbSlowLogTrendResponse, error) {
	result := &GetMongodbSlowLogTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbSlowLogTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("period", util.StringValue(request.Period)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbSlowQueryTemplate
//
// PARAMS:
//   - request: the arguments to GetMongodbSlowQueryTemplate
//
// RETURNS:
//   - GetMongodbSlowQueryTemplateResponse: The return type of the GetMongodbSlowQueryTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbSlowQueryTemplate(request *GetMongodbSlowQueryTemplateRequest) (*GetMongodbSlowQueryTemplateResponse, error) {
	result := &GetMongodbSlowQueryTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbSlowQueryTemplateUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("users", util.StringValue(request.Users)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("clientIps", util.StringValue(request.ClientIps)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMongodbSpaceSummary
//
// PARAMS:
//   - request: the arguments to GetMongodbSpaceSummary
//
// RETURNS:
//   - GetMongodbSpaceSummaryResponse: The return type of the GetMongodbSpaceSummary interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMongodbSpaceSummary(request *GetMongodbSpaceSummaryRequest) (*GetMongodbSpaceSummaryResponse, error) {
	result := &GetMongodbSpaceSummaryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMongodbSpaceSummaryUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlActiveSessions
//
// PARAMS:
//   - request: the arguments to GetMysqlActiveSessions
//
// RETURNS:
//   - GetMysqlActiveSessionsResponse: The return type of the GetMysqlActiveSessions interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlActiveSessions(request *GetMysqlActiveSessionsRequest) (*GetMysqlActiveSessionsResponse, error) {
	result := &GetMysqlActiveSessionsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlActiveSessionsUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlDatabaseSpace
//
// PARAMS:
//   - request: the arguments to GetMysqlDatabaseSpace
//
// RETURNS:
//   - GetMysqlDatabaseSpaceResponse: The return type of the GetMysqlDatabaseSpace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlDatabaseSpace(request *GetMysqlDatabaseSpaceRequest) (*GetMysqlDatabaseSpaceResponse, error) {
	result := &GetMysqlDatabaseSpaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlDatabaseSpaceUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlDeadlockInfo
//
// PARAMS:
//   - request: the arguments to GetMysqlDeadlockInfo
//
// RETURNS:
//   - GetMysqlDeadlockInfoResponse: The return type of the GetMysqlDeadlockInfo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlDeadlockInfo(request *GetMysqlDeadlockInfoRequest) (*GetMysqlDeadlockInfoResponse, error) {
	result := &GetMysqlDeadlockInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlDeadlockInfoUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlKillSessionHistory
//
// PARAMS:
//   - request: the arguments to GetMysqlKillSessionHistory
//
// RETURNS:
//   - GetMysqlKillSessionHistoryResponse: The return type of the GetMysqlKillSessionHistory interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlKillSessionHistory(request *GetMysqlKillSessionHistoryRequest) (*GetMysqlKillSessionHistoryResponse, error) {
	result := &GetMysqlKillSessionHistoryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetMysqlKillSessionHistoryUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("start", "start").
		WithQueryParamFilter("end", "end").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlRateLimitTaskDetail
//
// PARAMS:
//   - request: the arguments to GetMysqlRateLimitTaskDetail
//
// RETURNS:
//   - GetMysqlRateLimitTaskDetailResponse: The return type of the GetMysqlRateLimitTaskDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlRateLimitTaskDetail(request *GetMysqlRateLimitTaskDetailRequest) (*GetMysqlRateLimitTaskDetailResponse, error) {
	result := &GetMysqlRateLimitTaskDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlRateLimitTaskDetailUri(VERSION_V1)).
		WithQueryParam("filterId", "").
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("filterId", util.Int32Value(request.FilterId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlSlowLogTemplate
//
// PARAMS:
//   - request: the arguments to GetMysqlSlowLogTemplate
//
// RETURNS:
//   - GetMysqlSlowLogTemplateResponse: The return type of the GetMysqlSlowLogTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlSlowLogTemplate(request *GetMysqlSlowLogTemplateRequest) (*GetMysqlSlowLogTemplateResponse, error) {
	result := &GetMysqlSlowLogTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlSlowLogTemplateUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("users", util.StringValue(request.Users)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("clientIps", util.StringValue(request.ClientIps)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlSlowLogTimeDistribution
//
// PARAMS:
//   - request: the arguments to GetMysqlSlowLogTimeDistribution
//
// RETURNS:
//   - GetMysqlSlowLogTimeDistributionResponse: The return type of the GetMysqlSlowLogTimeDistribution interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlSlowLogTimeDistribution(request *GetMysqlSlowLogTimeDistributionRequest) (*GetMysqlSlowLogTimeDistributionResponse, error) {
	result := &GetMysqlSlowLogTimeDistributionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlSlowLogTimeDistributionUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlSlowLogTrend
//
// PARAMS:
//   - request: the arguments to GetMysqlSlowLogTrend
//
// RETURNS:
//   - GetMysqlSlowLogTrendResponse: The return type of the GetMysqlSlowLogTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlSlowLogTrend(request *GetMysqlSlowLogTrendRequest) (*GetMysqlSlowLogTrendResponse, error) {
	result := &GetMysqlSlowLogTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlSlowLogTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("period", util.StringValue(request.Period)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlSpaceSummary
//
// PARAMS:
//   - request: the arguments to GetMysqlSpaceSummary
//
// RETURNS:
//   - GetMysqlSpaceSummaryResponse: The return type of the GetMysqlSpaceSummary interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlSpaceSummary(request *GetMysqlSpaceSummaryRequest) (*GetMysqlSpaceSummaryResponse, error) {
	result := &GetMysqlSpaceSummaryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlSpaceSummaryUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlTableIndexes
//
// PARAMS:
//   - request: the arguments to GetMysqlTableIndexes
//
// RETURNS:
//   - GetMysqlTableIndexesResponse: The return type of the GetMysqlTableIndexes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlTableIndexes(request *GetMysqlTableIndexesRequest) (*GetMysqlTableIndexesResponse, error) {
	result := &GetMysqlTableIndexesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlTableIndexesUri(VERSION_V1)).
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("database", "string").
		WithQueryParamFilter("table", "string").
		WithQueryParamFilter("product", util.StringValue(request.Product)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("table", util.StringValue(request.Table)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMysqlTableSpace
//
// PARAMS:
//   - request: the arguments to GetMysqlTableSpace
//
// RETURNS:
//   - GetMysqlTableSpaceResponse: The return type of the GetMysqlTableSpace interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetMysqlTableSpace(request *GetMysqlTableSpaceRequest) (*GetMysqlTableSpaceResponse, error) {
	result := &GetMysqlTableSpaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetMysqlTableSpaceUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("database", util.StringValue(request.Database)).
		WithQueryParamFilter("table", util.StringValue(request.Table)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPegadbSlowLogTemplate
//
// PARAMS:
//   - request: the arguments to GetPegadbSlowLogTemplate
//
// RETURNS:
//   - GetPegadbSlowLogTemplateResponse: The return type of the GetPegadbSlowLogTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPegadbSlowLogTemplate(request *GetPegadbSlowLogTemplateRequest) (*GetPegadbSlowLogTemplateResponse, error) {
	result := &GetPegadbSlowLogTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPegadbSlowLogTemplateUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPegadbSlowLogTimeDistribution
//
// PARAMS:
//   - request: the arguments to GetPegadbSlowLogTimeDistribution
//
// RETURNS:
//   - GetPegadbSlowLogTimeDistributionResponse: The return type of the GetPegadbSlowLogTimeDistribution interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPegadbSlowLogTimeDistribution(request *GetPegadbSlowLogTimeDistributionRequest) (*GetPegadbSlowLogTimeDistributionResponse, error) {
	result := &GetPegadbSlowLogTimeDistributionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPegadbSlowLogTimeDistributionUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPegadbSlowLogTrend
//
// PARAMS:
//   - request: the arguments to GetPegadbSlowLogTrend
//
// RETURNS:
//   - GetPegadbSlowLogTrendResponse: The return type of the GetPegadbSlowLogTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPegadbSlowLogTrend(request *GetPegadbSlowLogTrendRequest) (*GetPegadbSlowLogTrendResponse, error) {
	result := &GetPegadbSlowLogTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPegadbSlowLogTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("period", util.Int32Value(request.Period)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPostgresqlSlowLogTemplate
//
// PARAMS:
//   - request: the arguments to GetPostgresqlSlowLogTemplate
//
// RETURNS:
//   - GetPostgresqlSlowLogTemplateResponse: The return type of the GetPostgresqlSlowLogTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPostgresqlSlowLogTemplate(request *GetPostgresqlSlowLogTemplateRequest) (*GetPostgresqlSlowLogTemplateResponse, error) {
	result := &GetPostgresqlSlowLogTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPostgresqlSlowLogTemplateUri(VERSION_V1)).
		WithQueryParam("", "").
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("nodeId", "string").
		WithQueryParamFilter("start", "string").
		WithQueryParamFilter("end", "string").
		WithQueryParamFilter("dbNames", "dbNames").
		WithQueryParamFilter("clientIps", "clientIps").
		WithQueryParamFilter("fingerprintMd5", "fingerprintMd5page").
		WithQueryParamFilter("pageSize", "int").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("users", util.StringValue(request.Users)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("clientIps", util.StringValue(request.ClientIps)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPostgresqlSlowLogTimeDistribution
//
// PARAMS:
//   - request: the arguments to GetPostgresqlSlowLogTimeDistribution
//
// RETURNS:
//   - GetPostgresqlSlowLogTimeDistributionResponse: The return type of the GetPostgresqlSlowLogTimeDistribution interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPostgresqlSlowLogTimeDistribution(request *GetPostgresqlSlowLogTimeDistributionRequest) (*GetPostgresqlSlowLogTimeDistributionResponse, error) {
	result := &GetPostgresqlSlowLogTimeDistributionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPostgresqlSlowLogTimeDistributionUri(VERSION_V1)).
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("product", util.StringValue(request.Product)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbNames", strings.Join(util.PtrSliceToStringSlice(request.DbNames), ",")).
		WithQueryParamFilter("users", strings.Join(util.PtrSliceToStringSlice(request.Users), ",")).
		WithQueryParamFilter("clientIPs", strings.Join(util.PtrSliceToStringSlice(request.ClientIPs), ",")).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPostgresqlSlowLogTrend
//
// PARAMS:
//   - request: the arguments to GetPostgresqlSlowLogTrend
//
// RETURNS:
//   - GetPostgresqlSlowLogTrendResponse: The return type of the GetPostgresqlSlowLogTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPostgresqlSlowLogTrend(request *GetPostgresqlSlowLogTrendRequest) (*GetPostgresqlSlowLogTrendResponse, error) {
	result := &GetPostgresqlSlowLogTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPostgresqlSlowLogTrendUri(VERSION_V1)).
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("nodeId", "string").
		WithQueryParamFilter("start", "string").
		WithQueryParamFilter("end", "string").
		WithQueryParamFilter("period", "int").
		WithQueryParamFilter("product", util.StringValue(request.Product)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("period", util.Int32Value(request.Period)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRedisBigKeyAnalysisResult
//
// PARAMS:
//   - request: the arguments to GetRedisBigKeyAnalysisResult
//
// RETURNS:
//   - GetRedisBigKeyAnalysisResultResponse: The return type of the GetRedisBigKeyAnalysisResult interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRedisBigKeyAnalysisResult(request *GetRedisBigKeyAnalysisResultRequest) (*GetRedisBigKeyAnalysisResultResponse, error) {
	result := &GetRedisBigKeyAnalysisResultResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRedisBigKeyAnalysisResultUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("page", "page").
		WithQueryParamFilter("pageSize", "pageSize").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRedisSlowLogTemplate
//
// PARAMS:
//   - request: the arguments to GetRedisSlowLogTemplate
//
// RETURNS:
//   - GetRedisSlowLogTemplateResponse: The return type of the GetRedisSlowLogTemplate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRedisSlowLogTemplate(request *GetRedisSlowLogTemplateRequest) (*GetRedisSlowLogTemplateResponse, error) {
	result := &GetRedisSlowLogTemplateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRedisSlowLogTemplateUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRedisSlowLogTimeDistribution
//
// PARAMS:
//   - request: the arguments to GetRedisSlowLogTimeDistribution
//
// RETURNS:
//   - GetRedisSlowLogTimeDistributionResponse: The return type of the GetRedisSlowLogTimeDistribution interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRedisSlowLogTimeDistribution(request *GetRedisSlowLogTimeDistributionRequest) (*GetRedisSlowLogTimeDistributionResponse, error) {
	result := &GetRedisSlowLogTimeDistributionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRedisSlowLogTimeDistributionUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRedisSlowLogTrend
//
// PARAMS:
//   - request: the arguments to GetRedisSlowLogTrend
//
// RETURNS:
//   - GetRedisSlowLogTrendResponse: The return type of the GetRedisSlowLogTrend interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetRedisSlowLogTrend(request *GetRedisSlowLogTrendRequest) (*GetRedisSlowLogTrendResponse, error) {
	result := &GetRedisSlowLogTrendResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetRedisSlowLogTrendUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("period", util.Int32Value(request.Period)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// KillMysqlSession
//
// PARAMS:
//   - request: the arguments to KillMysqlSession
//
// RETURNS:
//   - KillMysqlSessionResponse: The return type of the KillMysqlSession interface.
//   - error: nil if success otherwise the specific error
func (c *Client) KillMysqlSession(request *KillMysqlSessionRequest) (*KillMysqlSessionResponse, error) {
	result := &KillMysqlSessionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getKillMysqlSessionUri(VERSION_V1)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListMongodbSlowLogs
//
// PARAMS:
//   - request: the arguments to ListMongodbSlowLogs
//
// RETURNS:
//   - ListMongodbSlowLogsResponse: The return type of the ListMongodbSlowLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListMongodbSlowLogs(request *ListMongodbSlowLogsRequest) (*ListMongodbSlowLogsResponse, error) {
	result := &ListMongodbSlowLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListMongodbSlowLogsUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("users", util.StringValue(request.Users)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("clientIps", util.StringValue(request.ClientIps)).
		WithQueryParamFilter("namespace", util.StringValue(request.Namespace)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListMysqlRateLimitTasks
//
// PARAMS:
//   - request: the arguments to ListMysqlRateLimitTasks
//
// RETURNS:
//   - ListMysqlRateLimitTasksResponse: The return type of the ListMysqlRateLimitTasks interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListMysqlRateLimitTasks(request *ListMysqlRateLimitTasksRequest) (*ListMysqlRateLimitTasksResponse, error) {
	result := &ListMysqlRateLimitTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListMysqlRateLimitTasksUri(VERSION_V1)).
		WithQueryParamFilter("appId", "rds-OEEsaajh").
		WithQueryParamFilter("nodeId", "rds-OEEsaajh").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListMysqlSlowLogs
//
// PARAMS:
//   - request: the arguments to ListMysqlSlowLogs
//
// RETURNS:
//   - ListMysqlSlowLogsResponse: The return type of the ListMysqlSlowLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListMysqlSlowLogs(request *ListMysqlSlowLogsRequest) (*ListMysqlSlowLogsResponse, error) {
	result := &ListMysqlSlowLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListMysqlSlowLogsUri(VERSION_V1)).
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("users", util.StringValue(request.Users)).
		WithQueryParamFilter("dbNames", util.StringValue(request.DbNames)).
		WithQueryParamFilter("clientIps", util.StringValue(request.ClientIps)).
		WithQueryParamFilter("fingerprintMd5", util.StringValue(request.FingerprintMd5)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPegadbSlowLogs
//
// PARAMS:
//   - request: the arguments to ListPegadbSlowLogs
//
// RETURNS:
//   - ListPegadbSlowLogsResponse: The return type of the ListPegadbSlowLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPegadbSlowLogs(request *ListPegadbSlowLogsRequest) (*ListPegadbSlowLogsResponse, error) {
	result := &ListPegadbSlowLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPegadbSlowLogsUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPostgresqlSlowLogs
//
// PARAMS:
//   - request: the arguments to ListPostgresqlSlowLogs
//
// RETURNS:
//   - ListPostgresqlSlowLogsResponse: The return type of the ListPostgresqlSlowLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPostgresqlSlowLogs(request *ListPostgresqlSlowLogsRequest) (*ListPostgresqlSlowLogsResponse, error) {
	result := &ListPostgresqlSlowLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPostgresqlSlowLogsUri(VERSION_V1)).
		WithQueryParamFilter("product", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("dbNames", strings.Join(util.PtrSliceToStringSlice(request.DbNames), ",")).
		WithQueryParamFilter("clientIPs", strings.Join(util.PtrSliceToStringSlice(request.ClientIPs), ",")).
		WithQueryParamFilter("users", strings.Join(util.PtrSliceToStringSlice(request.Users), ",")).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRedisBigKeyAnalysisTasks
//
// PARAMS:
//   - request: the arguments to ListRedisBigKeyAnalysisTasks
//
// RETURNS:
//   - ListRedisBigKeyAnalysisTasksResponse: The return type of the ListRedisBigKeyAnalysisTasks interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRedisBigKeyAnalysisTasks(request *ListRedisBigKeyAnalysisTasksRequest) (*ListRedisBigKeyAnalysisTasksResponse, error) {
	result := &ListRedisBigKeyAnalysisTasksResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRedisBigKeyAnalysisTasksUri(VERSION_V1)).
		WithQueryParamFilter("id", "id").
		WithQueryParamFilter("appId", "appId").
		WithQueryParamFilter("nodeId", "nodeId").
		WithQueryParamFilter("dataType", "dataType").
		WithQueryParamFilter("orderBy", "orderBy").
		WithQueryParamFilter("order", "order").
		WithQueryParamFilter("id", util.Int32Value(request.Id)).
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("clusterId", util.StringValue(request.ClusterId)).
		WithQueryParamFilter("dataType", util.StringValue(request.DataType)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRedisSlowLogs
//
// PARAMS:
//   - request: the arguments to ListRedisSlowLogs
//
// RETURNS:
//   - ListRedisSlowLogsResponse: The return type of the ListRedisSlowLogs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRedisSlowLogs(request *ListRedisSlowLogsRequest) (*ListRedisSlowLogsResponse, error) {
	result := &ListRedisSlowLogsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRedisSlowLogsUri(VERSION_V1)).
		WithQueryParamFilter("appId", "string").
		WithQueryParamFilter("appId", util.StringValue(request.AppId)).
		WithQueryParamFilter("nodeId", util.StringValue(request.NodeId)).
		WithQueryParamFilter("start", util.StringValue(request.Start)).
		WithQueryParamFilter("end", util.StringValue(request.End)).
		WithQueryParamFilter("dbEngine", util.StringValue(request.DbEngine)).
		WithQueryParamFilter("page", util.Int32Value(request.Page)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("orderBy", util.StringValue(request.OrderBy)).
		WithQueryParamFilter("order", util.StringValue(request.Order)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StartStopMysqlInstanceFlowLimitingTask
//
// PARAMS:
//   - request: the arguments to StartStopMysqlInstanceFlowLimitingTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) StartStopMysqlInstanceFlowLimitingTask(request *StartStopMysqlInstanceFlowLimitingTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getStartStopMysqlInstanceFlowLimitingTaskUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// UpdateMysqlRateLimitTask
//
// PARAMS:
//   - request: the arguments to UpdateMysqlRateLimitTask
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateMysqlRateLimitTask(request *UpdateMysqlRateLimitTaskRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateMysqlRateLimitTaskUri(VERSION_V1)).
		WithBody(request).
		Do()
}
