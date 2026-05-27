package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddRecordResponse struct {
	bce.BaseResponse
	RecordId *string `json:"recordId,omitempty"`
}
