package rds

type TableGranularityUpdateUserDescriptionApiRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	User       *string `json:"user,omitempty"`
	Host       *string `json:"host,omitempty"`
	Remark     *string `json:"remark,omitempty"`
}
