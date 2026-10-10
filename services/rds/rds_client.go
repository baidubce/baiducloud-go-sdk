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

// - error: nil if success otherwise the specific error
func (c *Client) QueryTaskList(request *QueryTaskListRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryTaskListUri()).
		WithBody(request).
		Do()
}
