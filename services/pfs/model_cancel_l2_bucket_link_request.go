package pfs

type CancelL2BucketLinkRequest struct {
	Action       *string `json:"-"`
	BucketLinkId *string `json:"bucketLinkId,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty"`
}
