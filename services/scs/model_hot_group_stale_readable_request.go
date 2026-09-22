package scs

type HotGroupStaleReadableRequest struct {
	GroupId       *string `json:"-"`
	FollowerId    *string `json:"followerId,omitempty"`
	StaleReadable *bool   `json:"staleReadable,omitempty"`
}
