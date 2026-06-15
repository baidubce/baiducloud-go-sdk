package bcc

type DeleteInstanceDeploySetRequest struct {
	DeployId       *string   `json:"deployId,omitempty"`
	InstanceIdList []*string `json:"instanceIdList,omitempty"`
}
