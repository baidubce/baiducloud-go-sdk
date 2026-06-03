package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeZonesResponse struct {
	bce.BaseResponse
	Zones []*string `json:"zones,omitempty"`
}
