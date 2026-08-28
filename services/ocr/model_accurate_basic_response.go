package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AccurateBasicResponse struct {
	bce.BaseResponse
	ErrorCode           *int32                           `json:"error_code,omitempty"`
	ErrorMsg            *string                          `json:"error_msg,omitempty"`
	Direction           *int32                           `json:"direction,omitempty"`
	LogId               *int64                           `json:"log_id,omitempty"`
	WordsResultNum      *int32                           `json:"words_result_num,omitempty"`
	WordsResult         []*AccurateBasicWordsResult      `json:"words_result,omitempty"`
	ParagraphsResult    []*AccurateBasicParagraphsResult `json:"paragraphs_result,omitempty"`
	ParagraphsResultNum *int32                           `json:"paragraphs_result_num,omitempty"`
	PdfFileSize         *int32                           `json:"pdf_file_size,omitempty"`
	OfdFileSize         *string                          `json:"ofd_file_size,omitempty"`
}
