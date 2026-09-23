package nlp

import "github.com/baidubce/baiducloud-go-sdk/bce"

type SimnetResponse struct {
	bce.BaseResponse
	ErrorCode *int32   `json:"error_code,omitempty"`
	ErrorMsg  *string  `json:"error_msg,omitempty"`
	LogId     *int64   `json:"log_id,omitempty"`
	Score     *float32 `json:"score,omitempty"`
	Texts     *Texts   `json:"texts,omitempty"`
}
