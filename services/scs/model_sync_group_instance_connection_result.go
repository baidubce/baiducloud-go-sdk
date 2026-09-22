package scs

type SyncGroupInstanceConnectionResult struct {
	SourceId    *string `json:"sourceId,omitempty"`
	SourceRole  *string `json:"sourceRole,omitempty"`
	TargetId    *string `json:"targetId,omitempty"`
	TargetRole  *string `json:"targetRole,omitempty"`
	Connectable *bool   `json:"connectable,omitempty"`
}
