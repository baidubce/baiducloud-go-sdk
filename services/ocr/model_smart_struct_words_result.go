package ocr

type SmartStructWordsResult struct {
	StructInfo *SmartStructStructInfo `json:"struct_info,omitempty"`
	Relations  *SmartStructRelations  `json:"relations,omitempty"`
	LineInfo   []*SmartStructLineInfo `json:"line_info,omitempty"`
}
