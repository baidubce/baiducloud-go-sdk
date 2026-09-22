package scs

type SyncGroupRemoveInstanceRequest struct {
	GroupId  *string `json:"-"`
	MemberId *string `json:"memberId,omitempty"`
	Region   *string `json:"region,omitempty"`
}
