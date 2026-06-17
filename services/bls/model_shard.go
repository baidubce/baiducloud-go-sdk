package bls

type Shard struct {
	Total      *int32 `json:"total,omitempty"`
	Successful *int32 `json:"successful,omitempty"`
	Skipped    *int32 `json:"skipped,omitempty"`
	Failed     *int32 `json:"failed,omitempty"`
}
