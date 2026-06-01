package apm

type BindServiceTagRequest struct {
	ServiceName *string `json:"serviceName,omitempty"`
	ServiceId   *string `json:"serviceId,omitempty"`
	Tags        []*Tag  `json:"tags,omitempty"`
}
