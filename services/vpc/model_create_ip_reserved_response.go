

package vpc


import "github.com/baidubce/baiducloud-go-sdk/bce"



type CreateIpReservedResponse struct {
	bce.BaseResponse
	IpReserveId *string `json:"ipReserveId,omitempty"`
}

