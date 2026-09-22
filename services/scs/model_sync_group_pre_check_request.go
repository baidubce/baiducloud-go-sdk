package scs

type SyncGroupPreCheckRequest struct {
	SyncGroupShowId *string                        `json:"syncGroupShowId,omitempty"`
	Members         []*CheckSyncGroupRequestMember `json:"members,omitempty"`
}
