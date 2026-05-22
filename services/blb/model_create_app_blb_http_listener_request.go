package blb

type CreateAppBlbHttpListenerRequest struct {
	BlbId                 *string                    `json:"-"`
	ClientToken           *string                    `json:"-"`
	ListenerPort          *int32                     `json:"listenerPort,omitempty"`
	Scheduler             *string                    `json:"scheduler,omitempty"`
	KeepSession           *bool                      `json:"keepSession,omitempty"`
	KeepSessionType       *string                    `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    *int32                     `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName *string                    `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool                      `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool                      `json:"xForwardedProto,omitempty"`
	AdditionalAttributes  *AdditionalAttributesModel `json:"additionalAttributes,omitempty"`
	ServerTimeout         *int32                     `json:"serverTimeout,omitempty"`
	RedirectPort          *int32                     `json:"redirectPort,omitempty"`
	Description           *string                    `json:"description,omitempty"`
}
