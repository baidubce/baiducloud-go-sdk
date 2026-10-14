package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// QueryTaskList
//
// PARAMS:
//   - request: the arguments to QueryTaskList
//
// RETURNS:
//   - QueryTaskListResponse: The return type of the QueryTaskList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTaskList(request *QueryTaskListRequest) (*QueryTaskListResponse, error) {
	result := &QueryTaskListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryTaskListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
