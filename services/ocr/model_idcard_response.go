package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type IdcardResponse struct {
	bce.BaseResponse
	ErrorCode        *int32          `json:"error_code,omitempty"`
	ErrorMsg         *string         `json:"error_msg,omitempty"`
	LogId            *int64          `json:"log_id,omitempty"`
	WordsResultNum   *int32          `json:"words_result_num,omitempty"`
	WordsResult      *interface{}    `json:"words_result,omitempty"`
	Direction        *int32          `json:"direction,omitempty"`
	ImageStatus      *string         `json:"image_status,omitempty"`
	RiskType         *string         `json:"risk_type,omitempty"`
	CardQuality      *IdCardQuality  `json:"card_quality,omitempty"`
	Photo            *string         `json:"photo,omitempty"`
	PhotoLocation    *IdcardLocation `json:"photo_location,omitempty"`
	CardImage        *string         `json:"card_image,omitempty"`
	CardLocation     *IdcardLocation `json:"card_location,omitempty"`
	IdcardNumberType *int32          `json:"idcard_number_type,omitempty"`
	CardPs           *int32          `json:"card_ps,omitempty"`
	EditTool         *string         `json:"edit_tool,omitempty"`
}
