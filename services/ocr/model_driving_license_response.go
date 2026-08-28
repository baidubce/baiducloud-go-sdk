package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DrivingLicenseResponse struct {
	bce.BaseResponse
	ErrorCode          *int32                            `json:"error_code,omitempty"`
	ErrorMsg           *string                           `json:"error_msg,omitempty"`
	LogId              *int64                            `json:"log_id,omitempty"`
	Direction          *int32                            `json:"direction,omitempty"`
	WordsResultNum     *int32                            `json:"words_result_num,omitempty"`
	WordsResult        *interface{}                      `json:"words_result,omitempty"`
	WarnInfos          []*string                         `json:"warn_infos,omitempty"`
	QualityPropobility *DrivingLicenseQualityPropobility `json:"quality_propobility,omitempty"`
	RiskType           *string                           `json:"risk_type,omitempty"`
	EditTool           *string                           `json:"edit_tool,omitempty"`
}
