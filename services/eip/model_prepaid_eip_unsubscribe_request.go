package eip

type PrepaidEipUnsubscribeRequest struct {
	Eip         *string `json:"-"`
	ClientToken *string `json:"-"`
}
