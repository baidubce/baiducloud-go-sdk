package scs

type ModifyParametersRequest struct {
	InstanceId *string    `json:"-"`
	Parameter  *Parameter `json:"parameter,omitempty"`
}
