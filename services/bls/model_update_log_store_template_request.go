package bls

type UpdateLogStoreTemplateRequest struct {
	Name             *string   `json:"name,omitempty"`
	ProjectPatterns  []*string `json:"projectPatterns,omitempty"`
	LogstorePatterns []*string `json:"logstorePatterns,omitempty"`
	Priority         *int32    `json:"priority,omitempty"`
	Template         *Template `json:"template,omitempty"`
}
