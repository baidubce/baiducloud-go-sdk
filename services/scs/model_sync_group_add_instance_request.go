package scs

type SyncGroupAddInstanceRequest struct {
	GroupId  *string `json:"-"`
	MemberId *string `json:"memberId,omitempty"`
	Region   *string `json:"region,omitempty"`
}
