package nlp

type SentimentClassifyItem struct {
	Sentiment    *int32   `json:"sentiment,omitempty"`
	Confidence   *float32 `json:"confidence,omitempty"`
	PositiveProb *float32 `json:"positive_prob,omitempty"`
	NegativeProb *float32 `json:"negative_prob,omitempty"`
}
