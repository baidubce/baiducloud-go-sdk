package apm

type DeleteServicesRequest struct {
	ServiceNames []*string `json:"serviceNames,omitempty"`
}
