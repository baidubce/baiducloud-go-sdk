package eip

type RestoreEipFromRecycleRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
