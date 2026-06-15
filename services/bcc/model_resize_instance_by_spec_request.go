package bcc

type ResizeInstanceBySpecRequest struct {
	InstanceId       *string `json:"-"`
	Action           *string `json:"-"`
	Spec             *string `json:"spec,omitempty"`
	EnableJumboFrame *bool   `json:"enableJumboFrame,omitempty"`
}
