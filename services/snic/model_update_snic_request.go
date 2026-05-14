package snic

type UpdateSnicRequest struct {
	EndpointId  *string `json:"-"`
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
