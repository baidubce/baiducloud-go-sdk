package scs

type ModifySyncGroupNameRequest struct {
	SyncGroupShowId *string `json:"-"`
	GroupName       *string `json:"groupName,omitempty"`
}
