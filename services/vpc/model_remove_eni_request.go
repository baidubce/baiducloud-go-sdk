package vpc

type RemoveEniRequest struct {
	EniId       *string `json:"-"`
	ClientToken *string `json:"-"`
}
