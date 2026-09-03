package image

type SearchResult struct {
	ContSign *string  `json:"cont_sign,omitempty"`
	Score    *float32 `json:"score,omitempty"`
	Brief    *string  `json:"brief,omitempty"`
}
