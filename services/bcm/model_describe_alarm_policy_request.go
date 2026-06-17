package bcm

type DescribeAlarmPolicyRequest struct {
	Id                     *string `json:"id,omitempty"`
	RequireSubResourceType *bool   `json:"requireSubResourceType,omitempty"`
}
