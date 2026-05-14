package snic

type DeleteSnicRequest struct {
	EndpointId  *string `json:"-"`
	ClientToken *string `json:"-"`
}
