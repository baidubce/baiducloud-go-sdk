package vpc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type QueryTheIpAddressesOccupiedByProductsWithinVpcResponse struct {
	bce.BaseResponse
	PageNo     *int32        `json:"pageNo,omitempty"`
	PageSize   *int32        `json:"pageSize,omitempty"`
	TotalCount *int32        `json:"totalCount,omitempty"`
	Result     []*ResourceIp `json:"result,omitempty"`
}
