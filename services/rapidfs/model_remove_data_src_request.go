package rapidfs

type RemoveDataSrcRequest struct {
	Action      *string `json:"-"`
	ClientToken *string `json:"-"`
	DataSrcId   *string `json:"dataSrcId,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty"`
	Token       *string `json:"token,omitempty"`
}
