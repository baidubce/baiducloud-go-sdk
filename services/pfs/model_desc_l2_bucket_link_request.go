package pfs

type DescL2BucketLinkRequest struct {
	Action       *string `json:"-"`
	InstanceId   *string `json:"instanceId,omitempty"`
	BucketLinkId *string `json:"bucketLinkId,omitempty"`
}
