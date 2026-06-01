package apm

type DescribeLLMSpansRequest struct {
	BeginDatetime       *string   `json:"beginDatetime,omitempty"`
	EndDatetime         *string   `json:"endDatetime,omitempty"`
	ParseLLMInputOutput *bool     `json:"parseLLMInputOutput,omitempty"`
	Filters             []*Filter `json:"filters,omitempty"`
	OrderBy             *string   `json:"orderBy,omitempty"`
	Order               *string   `json:"order,omitempty"`
	Marker              *string   `json:"marker,omitempty"`
}
