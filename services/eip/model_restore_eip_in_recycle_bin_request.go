package eip

type RestoreEipInRecycleBinRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
