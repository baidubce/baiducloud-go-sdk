package oos

type EventExecution struct {
	Namespace        *string      `json:"namespace,omitempty"`
	Name             *string      `json:"name,omitempty"`
	State            *string      `json:"state,omitempty"`
	Description      *string      `json:"description,omitempty"`
	Template         *Template    `json:"template,omitempty"`
	Properties       *interface{} `json:"properties,omitempty"`
	Event            *Event       `json:"event,omitempty"`
	SilentCycleMilli *int32       `json:"silentCycleMilli,omitempty"`
	Tags             []*Tag       `json:"tags,omitempty"`
	TemplateDeleted  *bool        `json:"templateDeleted,omitempty"`
	Locale           *string      `json:"locale,omitempty"`
}
