package scs

type ModifyInstanceDomainNameRequest struct {
	InstanceId *string `json:"-"`
	Domain     *string `json:"domain,omitempty"`
}
