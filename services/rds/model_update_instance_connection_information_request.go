package rds

type UpdateInstanceConnectionInformationRequest struct {
	InstanceId *string `json:"-"`
	Address    *string `json:"address,omitempty"`
}
