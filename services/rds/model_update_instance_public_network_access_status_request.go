package rds

type UpdateInstancePublicNetworkAccessStatusRequest struct {
	InstanceId   *string `json:"-"`
	PublicAccess *bool   `json:"publicAccess,omitempty"`
}
