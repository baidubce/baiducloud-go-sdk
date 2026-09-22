package scs

type HotGroupAddClusterRequest struct {
	GroupId        *string `json:"-"`
	FollowerId     *string `json:"followerId,omitempty"`
	FollowerRegion *string `json:"followerRegion,omitempty"`
	SyncMaster     *string `json:"syncMaster,omitempty"`
}
