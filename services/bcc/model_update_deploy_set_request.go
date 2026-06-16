package bcc

type UpdateDeploySetRequest struct {
	DeployId *string `json:"-"`
	Name     *string `json:"name,omitempty"`
	Desc     *string `json:"desc,omitempty"`
}
