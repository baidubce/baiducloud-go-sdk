package bls

type SearchInfo struct {
	QueryType *string `json:"queryType,omitempty"`
	Took      *int32  `json:"took,omitempty"`
	Hits      *int32  `json:"hits,omitempty"`
}
