package scs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// InstanceList
//
// PARAMS:
//   - request: the arguments to InstanceList
//
// RETURNS:
//   - InstanceListResponse: The return type of the InstanceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceList(request *InstanceListRequest) (*InstanceListResponse, error) {
	result := &InstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceListUri(util.StringValue(request.Marker), util.StringValue(request.MaxKeys), util.StringValue(request.InstanceIds), util.StringValue(request.VnetIp))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
