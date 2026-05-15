package cfw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// QueryCfwList
//
// PARAMS:
//   - request: the arguments to QueryCfwList
//
// RETURNS:
//   - QueryCfwListResponse: The return type of the QueryCfwList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryCfwList(request *QueryCfwListRequest) (*QueryCfwListResponse, error) {
	result := &QueryCfwListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryCfwListUri(VERSION_V1)).
		WithQueryParamFilter("marker", "cfw-egx34bzjj43k").
		WithQueryParamFilter("maxKeys", "1").
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
