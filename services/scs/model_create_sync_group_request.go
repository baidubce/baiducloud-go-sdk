package scs

type CreateSyncGroupRequest struct {
	SyncGroupName *string   `json:"syncGroupName,omitempty"`
	Members       []*Member `json:"members,omitempty"`
}
