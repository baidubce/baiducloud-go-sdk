package ocr

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AccurateResponse struct {
	bce.BaseResponse
	ErrorCode           *int32                      `json:"error_code,omitempty"`
	ErrorMsg            *string                     `json:"error_msg,omitempty"`
	LogId               *int64                      `json:"log_id,omitempty"`
	Direction           *int32                      `json:"direction,omitempty"`
	WordsResultNum      *int32                      `json:"words_result_num,omitempty"`
	WordsResult         []*AccurateWordsResult      `json:"words_result,omitempty"`
	ParagraphsResult    []*AccurateParagraphsResult `json:"paragraphs_result,omitempty"`
	ParagraphsResultNum *int32                      `json:"paragraphs_result_num,omitempty"`
	PdfFileSize         *int32                      `json:"pdf_file_size,omitempty"`
	OfdFileSize         *string                     `json:"ofd_file_size,omitempty"`
}
