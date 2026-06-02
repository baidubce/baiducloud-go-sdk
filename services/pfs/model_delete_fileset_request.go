package pfs

type DeleteFilesetRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	FilesetId  *string `json:"filesetId,omitempty"`
}
