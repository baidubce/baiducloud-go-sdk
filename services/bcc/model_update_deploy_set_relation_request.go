package bcc

type UpdateDeploySetRelationRequest struct {
	InstanceId      *string   `json:"instanceId,omitempty"`
	DeploysetIdList []*string `json:"deploysetIdList,omitempty"`
}
