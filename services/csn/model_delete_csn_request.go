package csn

type DeleteCsnRequest struct {
	CsnId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
