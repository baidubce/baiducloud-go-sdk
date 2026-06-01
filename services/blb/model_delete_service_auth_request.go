package blb

type DeleteServiceAuthRequest struct {
	Service     *string   `json:"-"`
	ClientToken *string   `json:"-"`
	Action      *string   `json:"-"`
	UidList     []*string `json:"uidList,omitempty"`
}
