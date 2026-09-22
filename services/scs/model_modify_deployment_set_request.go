package scs

type ModifyDeploymentSetRequest struct {
	DeploySetId *string `json:"-"`
	Desc        *string `json:"desc,omitempty"`
	Name        *string `json:"name,omitempty"`
	Concurrency *int32  `json:"concurrency,omitempty"`
}
