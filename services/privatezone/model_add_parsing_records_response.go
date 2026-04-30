package privatezone

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddParsingRecordsResponse struct {
	bce.BaseResponse
	RecordId *string `json:"recordId,omitempty"`
}
