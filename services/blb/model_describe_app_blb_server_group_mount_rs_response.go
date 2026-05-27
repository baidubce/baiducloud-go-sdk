package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAppBlbServerGroupMountRsResponse struct {
	bce.BaseResponse
	BackendServerList []*AppBackendServer `json:"backendServerList,omitempty"`
}
