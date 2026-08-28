package ocr

type HandwritingGetEssayResult struct {
	RecognizeGranularity *string                                 `json:"recognize_granularity,omitempty"`
	Grade                *string                                 `json:"grade,omitempty"`
	Ids                  *string                                 `json:"ids,omitempty"`
	Name                 *string                                 `json:"name,omitempty"`
	EssayOverall         *HandwritingGetEssayOverall             `json:"essayOverall,omitempty"`
	Title                *Title                                  `json:"title,omitempty"`
	Content              *HandwritingCompositionGetResultContent `json:"content,omitempty"`
}
