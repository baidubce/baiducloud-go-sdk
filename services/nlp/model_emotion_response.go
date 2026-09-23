package nlp

import "github.com/baidubce/baiducloud-go-sdk/bce"

type EmotionResponse struct {
	bce.BaseResponse
	ErrorCode *int32         `json:"error_code,omitempty"`
	ErrorMsg  *string        `json:"error_msg,omitempty"`
	LogId     *int64         `json:"log_id,omitempty"`
	Text      *string        `json:"text,omitempty"`
	Items     []*EmotionItem `json:"items,omitempty"`
}
