package pfs

type DescFilesetRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	FilesetId  *string `json:"filesetId,omitempty"`
}
