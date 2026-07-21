package iam

type ACLEntry struct {
	Service    *string   `json:"service,omitempty"`
	Region     *string   `json:"region,omitempty"`
	Resource   []*string `json:"resource,omitempty"`
	Permission []*string `json:"permission,omitempty"`
	Effect     *string   `json:"effect,omitempty"`
}
