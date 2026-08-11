package ax

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// QuerySandboxes
//
// PARAMS:
//   - request: the arguments to QuerySandboxes
//
// RETURNS:
//   - QuerySandboxesResponse: The return type of the QuerySandboxes interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySandboxes(request *QuerySandboxesRequest) (*QuerySandboxesResponse, error) {
	result := &QuerySandboxesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQuerySandboxesUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
