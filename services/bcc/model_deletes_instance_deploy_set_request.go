package bcc

type DeletesInstanceDeploySetRequest struct {
	DeployId       *string   `json:"deployId,omitempty"`
	InstanceIdList []*string `json:"instanceIdList,omitempty"`
}
