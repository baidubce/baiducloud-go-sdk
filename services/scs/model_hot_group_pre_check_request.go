package scs

type HotGroupPreCheckRequest struct {
	Leader    *Leader          `json:"leader,omitempty"`
	Followers []*FollowersItem `json:"followers,omitempty"`
}
