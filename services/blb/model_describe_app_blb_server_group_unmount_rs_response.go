package blb

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeAppBlbServerGroupUnmountRsResponse struct {
	bce.BaseResponse
	BackendServerList []*AppBackendServer `json:"backendServerList,omitempty"`
}
