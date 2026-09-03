package image

type MultiObjectItem struct {
	Name     *string              `json:"name,omitempty"`
	Score    *float32             `json:"score,omitempty"`
	Location *MultiObjectLocation `json:"location,omitempty"`
}
