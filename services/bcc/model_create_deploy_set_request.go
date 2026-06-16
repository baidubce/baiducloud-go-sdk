package bcc

type CreateDeploySetRequest struct {
	Name        *string `json:"name,omitempty"`
	Desc        *string `json:"desc,omitempty"`
	Strategy    *string `json:"strategy,omitempty"`
	Concurrency *int32  `json:"concurrency,omitempty"`
}
