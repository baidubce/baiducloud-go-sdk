package scs

type Deploy struct {
	ScsInstanceCount *int32    `json:"scsInstanceCount,omitempty"`
	ScsInstancelds   []*string `json:"scsInstancelds,omitempty"`
	Name             *string   `json:"name,omitempty"`
	Strategy         *string   `json:"strategy,omitempty"`
	Concurrency      *int32    `json:"concurrency,omitempty"`
	DeploySetId      *string   `json:"deploySetId,omitempty"`
	Desc             *string   `json:"desc,omitempty"`
}
