package apm

type AlarmTarget struct {
	ApmType  *string   `json:"type,omitempty"`
	Tags     []*Tag    `json:"tags,omitempty"`
	Services []*string `json:"services,omitempty"`
}
