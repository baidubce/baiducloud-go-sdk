package eip

type ReleaseEipFromRecycleRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
