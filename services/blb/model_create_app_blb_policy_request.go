package blb

type CreateAppBlbPolicyRequest struct {
	BlbId        *string            `json:"-"`
	ClientToken  *string            `json:"-"`
	ListenerPort *int32             `json:"listenerPort,omitempty"`
	AppPolicyVos []*CreateAppPolicy `json:"appPolicyVos,omitempty"`
	BlbType      *string            `json:"type,omitempty"`
}
